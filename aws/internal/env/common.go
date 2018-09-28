package env

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/giantswarm/e2e-harness/pkg/framework"
)

const (
	component = "aws-operator"
	provider  = "aws"
)

const (
	envVarCircleCI             = "CIRCLECI"
	envVarCircleSHA1           = "CIRCLE_SHA1"
	envVarGithubBotToken       = "GITHUB_BOT_TOKEN"
	envVarKeepResources        = "KEEP_RESOURCES"
	envVarRegistryPullSecret   = "REGISTRY_PULL_SECRET"
	envVarTestedVersion        = "TESTED_VERSION"
	envVarTestDir              = "TEST_DIR"
	envVarVersionBundleVersion = "VERSION_BUNDLE_VERSION"
)

var (
	circleCI             string
	circleSHA1           string
	clusterID            string
	registryPullSecret   string
	githubToken          string
	testDir              string
	testedVersion        string
	keepResources        string
	versionBundleVersion string
)

func init() {
	var err error

	circleCI = os.Getenv(envVarCircleCI)
	keepResources = os.Getenv(envVarKeepResources)

	circleSHA1 = os.Getenv(envVarCircleSHA1)
	if circleSHA1 == "" {
		panic(fmt.Sprintf("env var %#q must not be empty", envVarCircleSHA1))
	}

	githubToken = os.Getenv(envVarGithubBotToken)
	if githubToken == "" {
		panic(fmt.Sprintf("env var %#q must not be empty", envVarGithubBotToken))
	}

	testedVersion = os.Getenv(envVarTestedVersion)
	if testedVersion == "" {
		panic(fmt.Sprintf("env var %#q must not be empty", envVarTestedVersion))
	}

	registryPullSecret = os.Getenv(envVarRegistryPullSecret)
	if registryPullSecret == "" {
		panic(fmt.Sprintf("env var %#q must not be empty", envVarRegistryPullSecret))
	}

	testDir = os.Getenv(envVarTestDir)
	if testDir == "" {
		panic(fmt.Sprintf("env var %#q must not be empty", envVarTestDir))
	}

	params := &framework.VBVParams{
		Component: component,
		Provider:  provider,
		Token:     githubToken,
		VType:     testedVersion,
	}
	versionBundleVersion, err = framework.GetVersionBundleVersion(params)
	if err != nil {
		panic(err.Error())
	}
	if VersionBundleVersion() == "" {
		if strings.ToLower(testedVersion) == "wip" {
			log.Println("WIP version bundle version not present, exiting.")
			os.Exit(0)
		}
		panic("version bundle version  must not be empty")
	}
}

func CircleCI() bool {
	return circleCI == "true"
}

// ClusterID returns a cluster ID unique to a run integration test. It might
// look like ci-wip-3cc75-5e958.
//
//     ci is a static identifier stating a CI run of the aws-operator.
//     wip is a version reference which can also be cur for the current version.
//     3cc75 is the Git SHA.
//     5e958 is a hash of the integration test dir, if any.
//
func ClusterID() string {
	var hash string
	{
		h := sha1.New()
		h.Write([]byte(testDir))
		hash = fmt.Sprintf("%x", h.Sum(nil))
	}

	parts := []string{
		"ci",
		testedVersion[0:3],
		circleSHA1[0:5],
		hash[0:5],
	}

	return strings.Join(parts, "-")
}

func KeepResources() bool {
	return keepResources == "true"
}

func RegistryPullSecret() string {
	return registryPullSecret
}

func VersionBundleVersion() string {
	return versionBundleVersion
}
