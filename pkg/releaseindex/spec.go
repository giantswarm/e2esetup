package releaseindex

type Authority string

func (c Authority) String() string {
	return string(c)
}

const (
	AuthorityAWSOperator     Authority = "aws-operator"
	AuthorityAzureOperator   Authority = "azure-operator"
	AuthorityCertOperator    Authority = "cert-operator"
	AuthorityChartOperator   Authority = "chart-operator"
	AuthorityClusterOperator Authority = "cluster-operator"
	AuthorityFlannelOperator Authority = "flannel-operator"
	AuthorityKVMOperator     Authority = "kvm-operator"
)

type Provider string

func (p Provider) String() string {
	return string(p)
}

const (
	ProviderAWS   Provider = "aws"
	ProviderAzure Provider = "azure"
	ProviderKVM   Provider = "kvm"
)

type VersionType string

func (v VersionType) String() string {
	return string(v)
}

const (
	VersionTypeLatest  VersionType = "latest"
	VersionTypeCurrent VersionType = "current"
)
