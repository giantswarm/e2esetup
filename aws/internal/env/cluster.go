package env

import (
	"fmt"
	"os"
)

const (
	envVarCommonDomain = "COMMON_DOMAIN"
	envVarVaultToken   = "VAULT_TOKEN"
)

var (
	commonDomain string
	vaultToken   string
)

func init() {
	commonDomain = os.Getenv(envVarCommonDomain)
	if commonDomain == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarCommonDomain))
	}

	vaultToken = os.Getenv(envVarVaultToken)
	if vaultToken == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarVaultToken))
	}
}

func CommonDomain() string {
	return commonDomain
}

func VaultToken() string {
	return vaultToken
}
