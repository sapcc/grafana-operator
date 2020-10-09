package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/integr8ly/grafana-operator/v3/pkg/apis"
	"github.com/integr8ly/grafana-operator/v3/pkg/controller"
	"github.com/integr8ly/grafana-operator/v3/pkg/controller/common"
	config2 "github.com/integr8ly/grafana-operator/v3/pkg/controller/config"
	"github.com/integr8ly/grafana-operator/v3/pkg/controller/grafanadashboard"
	"github.com/integr8ly/grafana-operator/v3/version"
	routev1 "github.com/openshift/api/route/v1"
	"github.com/operator-framework/operator-sdk/pkg/leader"
	"github.com/operator-framework/operator-sdk/pkg/metrics"
	"github.com/operator-framework/operator-sdk/pkg/ready"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/runtime/signals"
)

var log = logf.Log.WithName("cmd")
var flagImage string
var flagImageTag string
var flagPluginsInitContainerImage string
var flagPluginsInitContainerTag string
var flagProxyContainerImage string
var flagProxyContainerImageTag string
var flagNamespaces string
var scanAll bool
var syncPeriod int

var (
	metricsHost       = "0.0.0.0"
	metricsPort int32 = 8080
)

func printVersion() {
	log.Info(fmt.Sprintf("Go Version: %s", runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH))
	log.Info(fmt.Sprintf("operator-sdk Version: %v", sdkVersion.Version))
	log.Info(fmt.Sprintf("operator Version: %v", version.Version))
}

func init() {
	flagset := flag.CommandLine
	flagset.StringVar(&flagImage, "grafana-image", "", "Overrides the default Grafana image")
	flagset.StringVar(&flagImageTag, "grafana-image-tag", "", "Overrides the default Grafana image tag")
	flagset.StringVar(&flagPluginsInitContainerImage, "grafana-plugins-init-container-image", "", "Overrides the default Grafana Plugins Init Container image")
	flagset.StringVar(&flagPluginsInitContainerTag, "grafana-plugins-init-container-tag", "", "Overrides the default Grafana Plugins Init Container tag")
	flagset.StringVar(&flagProxyContainerImage, "grafana-proxy-container-image", "", "Overrides the default Grafana Proxy Container image")
	flagset.StringVar(&flagProxyContainerImageTag, "grafana-proxy-container-image-tag", "", "Overrides the default Grafana Proxy Container image tag")
	flagset.StringVar(&flagNamespaces, "namespaces", "", "Namespaces to scope the interaction of the Grafana operator. Mutually exclusive with --scan-all")
	flagset.BoolVar(&scanAll, "scan-all", true, "Scans all namespaces for dashboards")
	flagset.IntVar(&syncPeriod, "sync-period", 10, "SyncPeriod determines the minimum frequency at which watched resources are reconciled.")
	flagset.Parse(os.Args[1:])
}

// Starts a separate controller for the dashboard reconciliation in the background
func startDashboardController(ns string, cfg *rest.Config, signalHandler <-chan struct{}, autodetectChannel chan schema.GroupVersionKind) {
	// Create a new Cmd to provide shared dependencies and start components
	dashboardMgr, err := manager.New(cfg, manager.Options{
		MetricsBindAddress: "0",
		Namespace:          ns,
	})
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup Scheme for the dashboard resource
	if err := apis.AddToScheme(dashboardMgr.GetScheme()); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Use a separate manager for the dashboard controller
	grafanadashboard.Add(dashboardMgr, autodetectChannel, "")

	go func() {
		if err := dashboardMgr.Start(signalHandler); err != nil {
			log.Error(err, "dashboard manager exited non-zero")
			os.Exit(1)
		}
	}()
}

// Get the trimmed and sanitized list of namespaces (if --namespaces was provided)
func getSanitizedNamespaceList() []string {
	provided := strings.Split(flagNamespaces, ",")
	var selected []string

	for _, v := range provided {
		v = strings.TrimSpace(v)

		if v != "" {
			selected = append(selected, v)
		}
	}

	return selected
}

func main() {
	// The logger instantiated here can be changed to any logger
	// implementing the logr.Logger interface. This logger will
	// be propagated through the whole operator, generating
	// uniform and structured logs.
	logf.SetLogger(logf.ZapLogger(false))

	printVersion()

	// Controller configuration
	controllerConfig := config2.GetControllerConfig()
	controllerConfig.AddConfigItem(config2.ConfigGrafanaImage, flagImage)
	controllerConfig.AddConfigItem(config2.ConfigGrafanaImageTag, flagImageTag)
	controllerConfig.AddConfigItem(config2.ConfigPluginsInitContainerImage, flagPluginsInitContainerImage)
	controllerConfig.AddConfigItem(config2.ConfigPluginsInitContainerTag, flagPluginsInitContainerTag)
	controllerConfig.AddConfigItem(config2.ConfigGrafanaProxyImage, flagProxyContainerImage)
	controllerConfig.AddConfigItem(config2.ConfigGrafanaProxyImageTag, flagProxyContainerImageTag)
	controllerConfig.AddConfigItem(config2.ConfigOperatorNamespace, "grafana-operator")
	controllerConfig.AddConfigItem(config2.ConfigDashboardLabelSelector, "")

	// Get the namespaces to scan for dashboards
	// It's either the same namespace as the controller's or it's all namespaces if the
	// --scan-all flag has been passed
	var dashboardNamespaces = []string{}
	if scanAll {
		dashboardNamespaces = []string{""}
		log.Info("Scanning for dashboards in all namespaces")
	}

	if flagNamespaces != "" {
		dashboardNamespaces = getSanitizedNamespaceList()
		if len(dashboardNamespaces) == 0 {
			fmt.Fprint(os.Stderr, "--namespaces provided but no valid namespaces in list")
			os.Exit(1)
		}
		log.Info(fmt.Sprintf("Scanning for dashboards in the following namespaces: [%s]", strings.Join(dashboardNamespaces, ",")))
	}

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Become the leader before proceeding
	leader.Become(context.TODO(), "grafana-operator-lock")

	r := ready.NewFileReady()
	err = r.Set()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}
	defer r.Unset()

	sync := time.Duration(syncPeriod) * time.Hour
	// Create a new Cmd to provide shared dependencies and start components
	mgr, err := manager.New(cfg, manager.Options{
		SyncPeriod:         &sync,
		MetricsBindAddress: fmt.Sprintf("%s:%d", metricsHost, metricsPort),
	})
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	log.Info("Registering Components.")

	// Starting the resource auto-detection for the grafana controller
	autodetect, err := common.NewAutoDetect(mgr)
	if err != nil {
		log.Error(err, "failed to start the background process to auto-detect the operator capabilities")
	} else {
		autodetect.Start()
		defer autodetect.Stop()
	}

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup Scheme for OpenShift routes
	if err := routev1.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup all Controllers
	if err := controller.AddToManager(mgr, autodetect.SubscriptionChannel, ""); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	servicePorts := []v1.ServicePort{
		{
			Name:       metrics.OperatorPortName,
			Protocol:   v1.ProtocolTCP,
			Port:       metricsPort,
			TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: metricsPort},
		},
	}
	_, err = metrics.CreateMetricsService(context.TODO(), cfg, servicePorts)
	if err != nil {
		log.Error(err, "error starting metrics service")
	}

	log.Info("Starting the Cmd.")

	signalHandler := signals.SetupSignalHandler()

	// Start one dashboard controller per watch namespace
	//for _, ns := range dashboardNamespaces {
	//startDashboardController(ns, cfg, signalHandler, autodetect.SubscriptionChannel)
	//}

	if err := mgr.Start(signalHandler); err != nil {
		log.Error(err, "manager exited non-zero")
		os.Exit(1)
	}
}
