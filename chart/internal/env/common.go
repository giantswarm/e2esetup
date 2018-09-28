package env

import (
	"os"
)

const (
	envVarCircleCI      = "CIRCLECI"
	envVarKeepResources = "KEEP_RESOURCES"
)

var (
	circleCI      string
	keepResources string
)

func init() {
	circleCI = os.Getenv(envVarCircleCI)
	keepResources = os.Getenv(envVarKeepResources)
}

func CircleCI() string {
	return circleCI
}

func KeepResources() string {
	return keepResources
}
