module github.com/giantswarm/e2esetup/v2

go 1.14

require (
	github.com/giantswarm/apprclient/v2 v2.0.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/helmclient/v2 v2.0.0
	github.com/giantswarm/k8sclient/v4 v4.0.0
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.3
	github.com/giantswarm/valuemodifier v0.2.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/opencontainers/runc v1.0.0-rc2.0.20190611121236-6cc515888830 // indirect
	github.com/spf13/afero v1.3.4
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	k8s.io/api v0.18.5
	k8s.io/apiextensions-apiserver v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v0.18.5
	sigs.k8s.io/yaml v1.2.0
)
