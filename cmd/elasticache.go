package cmd

import (
	"context"
	"fmt"
	"github.com/pete911/aws-ip/internal/elasticache"
	"github.com/pete911/aws-ip/internal/out"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	elasticacheCmd = &cobra.Command{
		Use:     "elasticache",
		Aliases: []string{"elasticcache", "elastic-cache"},
		Short:   "view elastic cache IPs",
		Long:    "",
		Run:     runElasticache,
	}
)

func init() {
	Root.AddCommand(elasticacheCmd)
}

func runElasticache(cmd *cobra.Command, _ []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logger := GlobalFlags.Logger().With("cmd", cmd.Use)
	elasticacheRunner{
		logger: logger,
		client: elasticache.NewClient(logger, GlobalFlags.Config()),
	}.Run(ctx)
}

type elasticacheRunner struct {
	logger *slog.Logger
	client elasticache.Client
}

func (r elasticacheRunner) Run(ctx context.Context) {
	clusters, err := r.client.DescribeDBClusters(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	r.print(clusters)
}

func (r elasticacheRunner) print(clusters []elasticache.Cluster) {
	table := out.NewTable(r.logger, os.Stdout)
	table.AddRow("ID", "STATUS", "ENGINE", "VERSION", "NODE", "NODE STATUS", "NODE ENDPOINT", "NODE PORT", "IP")
	for _, c := range clusters {
		for _, n := range c.CacheNodes {
			ips, err := n.LookupIp()
			if err != nil {
				r.logger.Error(fmt.Sprintf("lookup %s ip: %v", n.Endpoint.Address, err))
			}
			table.AddRow(
				c.CacheClusterId,
				c.CacheClusterStatus,
				c.Engine,
				c.EngineVersion,
				n.CacheNodeId,
				n.CacheNodeStatus,
				n.Endpoint.Address,
				fmt.Sprintf("%d", n.Endpoint.Port),
				strings.Join(ips, ", "),
			)
		}
	}
	table.Print()
}
