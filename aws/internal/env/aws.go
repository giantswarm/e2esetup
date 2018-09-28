package env

import (
	"fmt"
	"os"
)

const (
	envVarAWSAPIHostedZoneGuest     = "AWS_API_HOSTED_ZONE_GUEST"
	envVarAWSIngressHostedZoneGuest = "AWS_INGRESS_HOSTED_ZONE_GUEST"
	envVarAWSRegion                 = "AWS_REGION"
	envVarGuestAWSARN               = "GUEST_AWS_ARN"
	envVarGuestAWSAccessKeyID       = "GUEST_AWS_ACCESS_KEY_ID"
	envVarGuestAWSAccessKeySecret   = "GUEST_AWS_SECRET_ACCESS_KEY"
	envVarGuestAWSAccessKeyToken    = "GUEST_AWS_SESSION_TOKEN"
	envVarHostAWSAccessKeyID        = "HOST_AWS_ACCESS_KEY_ID"
	envVarHostAWSAccessKeySecret    = "HOST_AWS_SECRET_ACCESS_KEY"
	envVarHostAWSAccessKeyToken     = "HOST_AWS_SESSION_TOKEN"
	envVarIDRSAPub                  = "IDRSA_PUB"
)

var (
	awsAPIHostedZoneGuest     string
	awsIngressHostedZoneGuest string
	awsRegion                 string
	guestAWSARN               string
	guestAWSAccessKeyID       string
	guestAWSAccessKeySecret   string
	guestAWSAccessKeyToken    string
	hostAWSAccessKeyID        string
	hostAWSAccessKeySecret    string
	hostAWSAccessKeyToken     string
	idRSAPub                  string
)

func init() {
	awsAPIHostedZoneGuest = os.Getenv(envVarAWSAPIHostedZoneGuest)
	if awsAPIHostedZoneGuest == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarAWSAPIHostedZoneGuest))
	}

	awsIngressHostedZoneGuest = os.Getenv(envVarAWSIngressHostedZoneGuest)
	if awsIngressHostedZoneGuest == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarAWSIngressHostedZoneGuest))
	}

	awsRegion = os.Getenv(envVarAWSRegion)
	if awsRegion == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarAWSRegion))
	}

	guestAWSARN = os.Getenv(envVarGuestAWSARN)
	if guestAWSARN == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarGuestAWSARN))
	}

	guestAWSAccessKeyID = os.Getenv(envVarGuestAWSAccessKeyID)
	if guestAWSAccessKeyID == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarGuestAWSAccessKeyID))
	}

	guestAWSAccessKeySecret = os.Getenv(envVarGuestAWSAccessKeySecret)
	if guestAWSAccessKeySecret == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarGuestAWSAccessKeySecret))
	}

	guestAWSAccessKeyToken = os.Getenv(envVarGuestAWSAccessKeyToken)

	hostAWSAccessKeyID = os.Getenv(envVarHostAWSAccessKeyID)
	if hostAWSAccessKeyID == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarHostAWSAccessKeyID))
	}

	hostAWSAccessKeySecret = os.Getenv(envVarHostAWSAccessKeySecret)
	if hostAWSAccessKeySecret == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarHostAWSAccessKeySecret))
	}

	hostAWSAccessKeyToken = os.Getenv(envVarHostAWSAccessKeyToken)

	idRSAPub = os.Getenv(envVarIDRSAPub)
	if idRSAPub == "" {
		panic(fmt.Sprintf("env var '%s' must not be empty", envVarIDRSAPub))
	}
}

func AWSAPIHostedZoneGuest() string {
	return awsAPIHostedZoneGuest
}

func AWSIngressHostedZoneGuest() string {
	return awsIngressHostedZoneGuest
}

func AWSRegion() string {
	return awsRegion
}

func AWSRouteTable0() string {
	return ClusterID() + "_0"
}

func AWSRouteTable1() string {
	return ClusterID() + "_1"
}

func GuestAWSARN() string {
	return guestAWSARN
}

func GuestAWSAccessKeyID() string {
	return guestAWSAccessKeyID
}

func GuestAWSAccessKeySecret() string {
	return guestAWSAccessKeySecret
}

func GuestAWSAccessKeyToken() string {
	return guestAWSAccessKeyToken
}

func HostAWSAccessKeyID() string {
	return hostAWSAccessKeyID
}

func HostAWSAccessKeySecret() string {
	return hostAWSAccessKeySecret
}

func HostAWSAccessKeyToken() string {
	return hostAWSAccessKeyToken
}

func IDRSAPub() string {
	return idRSAPub
}
