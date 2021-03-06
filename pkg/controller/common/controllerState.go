package common

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var ControllerEvents = make(chan ControllerState)

type ControllerState struct {
	DashboardSelectors []*v1.LabelSelector
	AdminUsername      string
	AdminPassword      string
	AdminUrl           string
	ProxyUrl           string
	GrafanaReady       bool
	GrafanaProxyReady  bool
	ClientTimeout      int
	FixAnnotations     bool
	FixHeights         bool
}
