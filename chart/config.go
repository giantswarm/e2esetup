package chart

import (
	"github.com/giantswarm/helmclient/v2/pkg/helmclient"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
)

type Config struct {
	HelmClient *helmclient.Client
	Setup      *k8sclient.Setup
}
