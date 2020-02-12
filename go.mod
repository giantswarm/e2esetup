module github.com/giantswarm/e2esetup

go 1.13

require (
	github.com/DATA-DOG/go-sqlmock v1.4.1 // indirect
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/chai2010/gettext-go v0.0.0-20170215093142-bf70f2a70fb1 // indirect
	github.com/docker/docker v1.13.1 // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/emicklei/go-restful v2.11.1+incompatible // indirect
	github.com/fortytw2/leaktest v1.3.0 // indirect
	github.com/giantswarm/apiextensions v0.0.0-20200208202837-856d5ad2e680 // indirect
	github.com/giantswarm/apprclient v0.0.0-20191209123802-955b7e89e6e2
	github.com/giantswarm/backoff v0.0.0-20190913091243-4dd491125192
	github.com/giantswarm/errors v0.0.0-20200205145547-c12c94d9a110 // indirect
	github.com/giantswarm/helmclient v0.0.0-20191209124624-d3c47349776d
	github.com/giantswarm/k8sclient v0.0.0-20191209120459-6cb127468cd6
	github.com/giantswarm/k8sportforward v0.0.0-20191209104600-676e7106283c // indirect
	github.com/giantswarm/microerror v0.1.1-0.20200205143715-01b76f66cae6
	github.com/giantswarm/micrologger v0.0.0-20200205144836-079154bcae45
	github.com/giantswarm/valuemodifier v0.0.0-20191004091153-c3b5fdecbd5c
	github.com/go-openapi/spec v0.19.4 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9 // indirect
	github.com/google/go-github v17.0.0+incompatible
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/huandu/xstrings v1.2.1 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.1 // indirect
	github.com/prometheus/common v0.7.0 // indirect
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20191121092708-da1cb182f00e // indirect
	github.com/spf13/afero v1.2.2
	golang.org/x/crypto v0.0.0-20191122220453-ac88ee75c92c // indirect
	golang.org/x/net v0.0.0-20191126235420-ef20fe5d7933 // indirect
	golang.org/x/oauth2 v0.0.0-20191122200657-5d9234df094c
	golang.org/x/sys v0.0.0-20191127021746-63cb32ae39b2 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	google.golang.org/genproto v0.0.0-20191115221424-83cc0476cb11 // indirect
	google.golang.org/grpc v1.25.1 // indirect
	gopkg.in/square/go-jose.v2 v2.4.0 // indirect
	k8s.io/api v0.16.6
	k8s.io/apiextensions-apiserver v0.16.6
	k8s.io/apimachinery v0.16.6
	k8s.io/client-go v0.16.6
	k8s.io/helm v2.16.1+incompatible
	k8s.io/kubernetes v1.16.3 // indirect
	k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6 // indirect
	sigs.k8s.io/controller-runtime v0.4.0 // indirect
	vbom.ml/util v0.0.0-20180919145318-efcd4e0f9787 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.1+incompatible
	// All of that is because helm has an import to k8s.io/kubernetes which
	// uses relative paths to those.
	k8s.io/api v0.0.0 => k8s.io/api v0.16.6
	k8s.io/apiextensions-apiserver v0.0.0 => k8s.io/apiextensions-apiserver v0.16.6
	k8s.io/apimachinery v0.0.0 => k8s.io/apimachinery v0.16.6
	k8s.io/apiserver v0.0.0 => k8s.io/apiserver v0.16.6
	k8s.io/cli-runtime v0.0.0 => k8s.io/cli-runtime v0.16.6
	k8s.io/client-go v0.0.0 => k8s.io/client-go v0.16.6
	k8s.io/cloud-provider v0.0.0 => k8s.io/cloud-provider v0.16.6
	k8s.io/cluster-bootstrap v0.0.0 => k8s.io/cluster-bootstrap v0.16.6
	k8s.io/code-generator v0.0.0 => k8s.io/code-generator v0.16.6
	k8s.io/component-base v0.0.0 => k8s.io/component-base v0.16.6
	k8s.io/cri-api v0.0.0 => k8s.io/cri-api v0.16.6
	k8s.io/csi-translation-lib v0.0.0 => k8s.io/csi-translation-lib v0.16.6
	k8s.io/kube-aggregator v0.0.0 => k8s.io/kube-aggregator v0.16.6
	k8s.io/kube-controller-manager v0.0.0 => k8s.io/kube-controller-manager v0.16.6
	k8s.io/kube-proxy v0.0.0 => k8s.io/kube-proxy v0.16.6
	k8s.io/kube-scheduler v0.0.0 => k8s.io/kube-scheduler v0.16.6
	k8s.io/kubectl v0.0.0 => k8s.io/kubectl v0.16.6
	k8s.io/kubelet v0.0.0 => k8s.io/kubelet v0.16.6
	k8s.io/legacy-cloud-providers v0.0.0 => k8s.io/legacy-cloud-providers v0.16.6
	k8s.io/metrics v0.0.0 => k8s.io/metrics v0.16.6
	k8s.io/sample-apiserver v0.0.0 => k8s.io/sample-apiserver v0.16.6
	k8s.io/utils v0.0.0 => k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6
)
