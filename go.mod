module github.com/integr8ly/grafana-operator/v3

go 1.13

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/databus23/goslo.policy v0.0.0-20170317131957-3ae74dd07ebf
	github.com/databus23/keystone v0.0.0-20180111110916-350fd0e663cd
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.1.0
	github.com/go-openapi/errors v0.19.2
	github.com/go-openapi/loads v0.19.4
	github.com/go-openapi/runtime v0.19.4
	github.com/go-openapi/spec v0.19.4
	github.com/go-openapi/strfmt v0.19.3
	github.com/go-openapi/swag v0.19.5
	github.com/go-openapi/validate v0.19.5
	github.com/google/go-jsonnet v0.16.0 // indirect
	github.com/gophercloud/gophercloud v0.6.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/namsral/flag v1.7.4-pre
	github.com/openshift/api v3.9.0+incompatible
	github.com/operator-framework/operator-sdk v0.18.2
	github.com/pkg/errors v0.9.1
	github.com/pmylund/go-cache v2.1.0+incompatible // indirect
	github.com/prometheus/client_golang v1.5.1
	github.com/sapcc/kubernikus v1.5.0
	github.com/spf13/pflag v1.0.5
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.18.2
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200121204235-bf4fb3bd569c
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.2+incompatible // Required by OLM
	github.com/mattn/go-sqlite3 => github.com/mattn/go-sqlite3 v1.10.0
	k8s.io/client-go => k8s.io/client-go v0.18.2
)
