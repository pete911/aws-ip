package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pete911/aws-ip/internal"
	"log/slog"
)

type Client struct {
	logger *slog.Logger
	svc    *ec2.Client
}

func NewClient(logger *slog.Logger, config internal.Config) Client {
	return Client{
		logger: logger.With("component", "ec2"),
		svc:    ec2.NewFromConfig(config.Config),
	}
}

func (c Client) DescribeVPCs(ctx context.Context) (Vpcs, error) {
	in := &ec2.DescribeVpcsInput{}

	var vpcs Vpcs
	for {
		out, err := c.svc.DescribeVpcs(ctx, in)
		if err != nil {
			return nil, err
		}
		vpcs = append(vpcs, toVpcs(out.Vpcs)...)
		if aws.ToString(out.NextToken) == "" {
			break
		}
		in.NextToken = out.NextToken
	}
	return vpcs, nil
}

func (c Client) DescribeSubnets(ctx context.Context) (Subnets, error) {
	in := &ec2.DescribeSubnetsInput{}

	var subnets Subnets
	for {
		out, err := c.svc.DescribeSubnets(ctx, in)
		if err != nil {
			return nil, err
		}
		subnets = append(subnets, toSubnets(out.Subnets)...)
		if aws.ToString(out.NextToken) == "" {
			break
		}
		in.NextToken = out.NextToken
	}
	return subnets, nil
}
