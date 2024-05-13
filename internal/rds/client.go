package rds

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/pete911/aws-ip/internal"
	"log/slog"
)

type Client struct {
	logger *slog.Logger
	svc    *rds.Client
}

func NewClient(logger *slog.Logger, config internal.Config) Client {
	return Client{
		logger: logger.With("component", "rds"),
		svc:    rds.NewFromConfig(config.Config),
	}
}

func (c Client) ListInstances(ctx context.Context) ([]Instance, error) {
	in := &rds.DescribeDBInstancesInput{}

	var instances []Instance
	for {
		out, err := c.svc.DescribeDBInstances(ctx, in)
		if err != nil {
			return nil, err
		}
		instances = append(instances, toInstances(out.DBInstances)...)
		if aws.ToString(out.Marker) == "" {
			break
		}
		in.Marker = out.Marker
	}
	return instances, nil
}
