package chart

import (
	"github.com/giantswarm/helmclient"
	"github.com/giantswarm/k8sclient/v3/pkg/k8sclient"
)

type Config struct {
	HelmClient *helmclient.Client
	Setup      *k8sclient.Setup
}
