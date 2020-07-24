module github.com/giantswarm/e2esetup

go 1.14

require (
	github.com/giantswarm/apprclient v0.2.1-0.20200724085653-63c7eb430dcf
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/helmclient v1.0.6-0.20200724085934-291f74d2499a
	github.com/giantswarm/k8sclient/v3 v3.1.3-0.20200724085258-345602646ea8
	github.com/giantswarm/microerror v0.2.0
	github.com/giantswarm/micrologger v0.3.1
	github.com/giantswarm/valuemodifier v0.2.0
	github.com/google/go-github v17.0.0+incompatible
	github.com/opencontainers/runc v1.0.0-rc2.0.20190611121236-6cc515888830 // indirect
	github.com/spf13/afero v1.3.2
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	k8s.io/api v0.18.5
	k8s.io/apiextensions-apiserver v0.18.5
	k8s.io/apimachinery v0.18.5
	k8s.io/client-go v0.18.5
	sigs.k8s.io/yaml v1.2.0
)
