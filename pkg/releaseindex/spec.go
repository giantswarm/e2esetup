package releaseindex

type Component string

func (c Component) String() string {
	return string(c)
}

const (
	ComponentAWSOperator     Component = "aws-operator"
	ComponentAzureOperator   Component = "azure-operator"
	ComponentCertOperator    Component = "cert-operator"
	ComponentChartOperator   Component = "chart-operator"
	ComponentClusterOperator Component = "cluster-operator"
	ComponentFlannelOperator Component = "flannel-operator"
	ComponentKVMOperator     Component = "kvm-operator"
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
