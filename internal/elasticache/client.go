package elasticache

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
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
	in := &elasticache.DescribeCacheClustersInput{ShowCacheNodeInfo: aws.Bool(true)}

	var clusters []Cluster
	for {
		out, err := c.svc.DescribeCacheClusters(ctx, in)
		if err != nil {
			return nil, err
		}
		clusters = append(clusters, toClusters(out.CacheClusters)...)
		if aws.ToString(out.Marker) == "" {
			break
		}
		in.Marker = out.Marker
	}
	return clusters, nil
}
