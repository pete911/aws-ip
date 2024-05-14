package cmd

import (
	"context"
	"fmt"
	"github.com/pete911/aws-ip/cmd/flag"
	"github.com/pete911/aws-ip/internal/ec2"
	"github.com/pete911/aws-ip/internal/elasticache"
	"github.com/pete911/aws-ip/internal/out"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	elasticacheFlags flag.Elasticache
	elasticacheCmd   = &cobra.Command{
		Use:     "elasticache",
		Aliases: []string{"elasticcache", "elastic-cache"},
		Short:   "view elastic cache IPs",
		Long:    "",
		Run:     runElasticache,
	}
)

func init() {
	flag.InitElasticacheFlags(elasticacheCmd, &elasticacheFlags)
	Root.AddCommand(elasticacheCmd)
}

func runElasticache(cmd *cobra.Command, _ []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logger := GlobalFlags.Logger().With("cmd", cmd.Use)
	runner := ElasticacheRunner{
		logger: logger,
		flags:  elasticacheFlags,
		client: elasticache.NewClient(logger, GlobalFlags.Config()),
	}
	// load vpc and subnet names if needed
	ec2client := ec2.NewClient(logger, GlobalFlags.Config())
	if elasticacheFlags.Vpc {
		runner.vpcs = loadVpcs(ctx, ec2client)
	}
	if elasticacheFlags.Subnet {
		runner.subnets = loadSubnets(ctx, ec2client)
	}
	runner.Run(ctx)
}

type ElasticacheRunner struct {
	logger  *slog.Logger
	flags   flag.Elasticache
	vpcs    ec2.Vpcs
	subnets ec2.Subnets
	client  elasticache.Client
}

func (r ElasticacheRunner) Run(ctx context.Context) {
	clusters, err := r.client.DescribeDBClusters(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	r.print(clusters)
}

func (r ElasticacheRunner) print(clusters []elasticache.Cluster) {
	table := out.NewTable(r.logger, os.Stdout)
	table.AddRow(r.tableHeader()...)
	for _, c := range clusters {
		for _, n := range c.CacheNodes {
			table.AddRow(r.tableRow(c, n)...)
		}
	}
	table.Print()
}

func (r ElasticacheRunner) tableHeader() []string {
	header := []string{"ID", "ENGINE", "VERSION", "NODE", "NODE ENDPOINT", "NODE PORT", "IP"}
	if r.flags.Status {
		header = append(header, "NODE STATUS")
	}
	if r.flags.VpcId {
		header = append(header, "VPC ID")
	}
	if r.flags.Vpc {
		header = append(header, "VPC")
	}
	if r.flags.SubnetId {
		header = append(header, "SUBNET ID")
	}
	if r.flags.Subnet {
		header = append(header, "SUBNET")
	}
	return header
}

func (r ElasticacheRunner) tableRow(cluster elasticache.Cluster, node elasticache.Node) []string {
	ips, err := node.LookupIp()
	if err != nil {
		r.logger.Error(fmt.Sprintf("lookup %s ip: %v", node.Endpoint.Address, err))
	}

	row := []string{
		node.CacheNodeId,
		cluster.Engine,
		cluster.EngineVersion,
		node.CacheNodeId,
		node.Endpoint.Address,
		fmt.Sprintf("%d", node.Endpoint.Port),
		strings.Join(ips, ", "),
	}

	if r.flags.Status {
		row = append(row, node.CacheNodeStatus)
	}
	if r.flags.VpcId {
		row = append(row, cluster.VpcId)
	}
	if r.flags.Vpc {
		row = append(row, r.vpcs.ById(cluster.VpcId).Tags["Name"])
	}
	if r.flags.SubnetId {
		row = append(row, strings.Join(cluster.Subnets, ", "))
	}
	if r.flags.Subnet {
		var names []string
		for _, n := range r.subnets.BySubnetIds(cluster.Subnets) {
			names = append(names, n.Tags["Name"])
		}
		row = append(row, strings.Join(names, ", "))
	}
	return row
}
