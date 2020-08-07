package helmclient

import (
	helmclientlib "github.com/giantswarm/helmclient/v2"
	"github.com/giantswarm/k8sclient/v4/pkg/k8sclient"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
)

type ClientConfig struct {
	K8sClient k8sclient.Interface
	Logger    micrologger.Logger
}

type Client struct {
	logger micrologger.Logger

	helmClient helmclientlib.Interface
}

func NewClient(config ClientConfig) (*Client, error) {
	if config.K8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	var err error

	var helmClient helmclientlib.Interface
	{
		c := helmclientlib.Config{
			K8sClient: config.K8sClient,
			Logger:    config.Logger,
		}

		helmClient, err = helmclientlib.New(c)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	c := &Client{
		logger: config.Logger,

		helmClient: helmClient,
	}

	return c, nil
}

func (c *Client) HelmClient() helmclientlib.Interface {
	return c.helmClient
}
