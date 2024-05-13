package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"time"
)

type Config struct {
	Account string
	Region  string
	Config  aws.Config
}

// NewConfigWithProfile returns config for specific AWS Profile. Used in exploratory tests and debugging
func NewConfigWithProfile(region, profile string) (Config, error) {
	cfg, err := newAWSConfig(profile)
	if err != nil {
		return Config{}, fmt.Errorf("load aws config: %w", err)
	}
	return newConfig(cfg, region)
}

func NewConfig(region string) (Config, error) {
	cfg, err := newAWSConfig("")
	if err != nil {
		return Config{}, fmt.Errorf("load aws config: %w", err)
	}
	return newConfig(cfg, region)
}

func newConfig(cfg aws.Config, region string) (Config, error) {
	if region == "" && cfg.Region == "" {
		return Config{}, errors.New("missing aws region")
	}

	if region != "" {
		cfg.Region = region
	}

	account, err := getCurrentAWSAccount(cfg)
	if err != nil {
		return Config{}, fmt.Errorf("get current aws account: %w", err)
	}

	return Config{
		Account: account,
		Region:  cfg.Region,
		Config:  cfg,
	}, nil
}

func newAWSConfig(profile string) (aws.Config, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if profile == "" {
		return config.LoadDefaultConfig(ctx)
	}
	return config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
}

func getCurrentAWSAccount(cfg aws.Config) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	svc := sts.NewFromConfig(cfg)
	resp, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}
	return aws.ToString(resp.Account), nil
}
