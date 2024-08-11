package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
)

//#! /bin/zsh
//aws sso login --profile sc-development
//
//aws eks update-kubeconfig --name "eks01-ap-southeast-2-development" --alias "eks01-ap-southeast-2-development" --profile "sc-development" --region "ap-southeast-2"
//
//kubectl config set-context --current --namespace reports
//~

const (
	flagProfileKey   = "profile"
	flagProfileUsage = "the aws sso profile to use for authentication"
)

func main() {
	profile := readProfileFlag()

	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	if err != nil {
		fatalf("failed to load config for profile '%s': %v", profile, err)
	}

}

func readProfileFlag() string {
	var profile string
	flag.StringVar(&profile, flagProfileKey, "", flagProfileUsage)

	return profile
}

func fatalf(msg string, args ...any) {
	fmt.Printf(msg, args)
	os.Exit(1)
}
