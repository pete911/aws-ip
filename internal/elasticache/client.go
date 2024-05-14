package elasticache

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/pete911/aws-ip/internal"
	"log/slog"
)

type Client struct {
	logger *slog.Logger
	svc    *elasticache.Client
}

func NewClient(logger *slog.Logger, config internal.Config) Client {
	return Client{
		logger: logger.With("component", "elasticache"),
		svc:    elasticache.NewFromConfig(config.Config),
	}
}

func (c Client) DescribeDBClusters(ctx context.Context) ([]Cluster, error) {
	subnetGroupsByName, err := c.describeCacheSubnetGroupsByName(ctx)
	if err != nil {
		return nil, err
	}

	in := &elasticache.DescribeCacheClustersInput{ShowCacheNodeInfo: aws.Bool(true)}

	var clusters []Cluster
	for {
		out, err := c.svc.DescribeCacheClusters(ctx, in)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, toClusters(out.CacheClusters, subnetGroupsByName)...)
		if aws.ToString(out.Marker) == "" {
			break
		}
		in.Marker = out.Marker
	}
	return clusters, nil
}

type subnetGroup struct {
	vpcId   string
	subnets []string
}

func toSubnetGroup(in types.CacheSubnetGroup) subnetGroup {
	var subnets []string
	for _, v := range in.Subnets {
		subnets = append(subnets, aws.ToString(v.SubnetIdentifier))
	}

	return subnetGroup{
		vpcId:   aws.ToString(in.VpcId),
		subnets: subnets,
	}
}

func (c Client) describeCacheSubnetGroupsByName(ctx context.Context) (map[string]subnetGroup, error) {
	in := &elasticache.DescribeCacheSubnetGroupsInput{}

	subnetsGroups := make(map[string]subnetGroup)
	for {
		out, err := c.svc.DescribeCacheSubnetGroups(ctx, in)
		if err != nil {
			return nil, err
		}
		for _, group := range out.CacheSubnetGroups {
			subnetsGroups[aws.ToString(group.CacheSubnetGroupName)] = toSubnetGroup(group)
		}
		if aws.ToString(out.Marker) == "" {
			break
		}
	}
	return subnetsGroups, nil
}
