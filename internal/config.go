package internal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type Config struct {
	Region string
	Config aws.Config
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

	return Config{
		Region: cfg.Region,
		Config: cfg,
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
