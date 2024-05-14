package cmd

import (
	"context"
	"fmt"
	"github.com/pete911/aws-ip/cmd/flag"
	"github.com/pete911/aws-ip/internal/ec2"
	"github.com/pete911/aws-ip/internal/out"
	"github.com/pete911/aws-ip/internal/rds"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	rdsFlags flag.Rds
	rdsCmd   = &cobra.Command{
		Use:   "rds",
		Short: "view rds IPs",
		Long:  "",
		Run:   runRds,
	}
)

func init() {
	flag.InitRdsFlags(rdsCmd, &rdsFlags)
	Root.AddCommand(rdsCmd)
}

func runRds(cmd *cobra.Command, _ []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logger := GlobalFlags.Logger().With("cmd", cmd.Use)
	runner := RdsRunner{
		logger: logger,
		flags:  rdsFlags,
		client: rds.NewClient(logger, GlobalFlags.Config()),
	}

	// load vpc and subnet names if needed
	ec2client := ec2.NewClient(logger, GlobalFlags.Config())
	if rdsFlags.Vpc {
		runner.vpcs = loadVpcs(ctx, ec2client)
	}
	if rdsFlags.Subnet {
		runner.subnets = loadSubnets(ctx, ec2client)
	}
	runner.Run(ctx)
}

type RdsRunner struct {
	logger  *slog.Logger
	flags   flag.Rds
	vpcs    ec2.Vpcs
	subnets ec2.Subnets
	client  rds.Client
}

func (r RdsRunner) Run(ctx context.Context) {
	instances, err := r.client.DescribeDBInstances(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	r.print(instances)
}

func (r RdsRunner) print(instances []rds.Instance) {
	table := out.NewTable(r.logger, os.Stdout)
	table.AddRow(r.tableHeader()...)
	for _, i := range instances {
		table.AddRow(r.tableRow(i)...)
	}
	table.Print()
}

func (r RdsRunner) tableHeader() []string {
	header := []string{"NAME", "ENDPOINT", "PORT", "IP"}
	if r.flags.Status {
		header = append(header, "STATUS")
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

func (r RdsRunner) tableRow(i rds.Instance) []string {
	ips, err := i.LookupIp()
	if err != nil {
		r.logger.Error(fmt.Sprintf("lookup %s ip: %v", i.Endpoint.Address, err))
	}

	row := []string{
		i.DBName,
		i.Endpoint.Address,
		fmt.Sprintf("%d", i.Endpoint.Port),
		strings.Join(ips, ", "),
	}

	if r.flags.Status {
		row = append(row, i.DBInstanceStatus)
	}
	if r.flags.VpcId {
		row = append(row, i.VpcId)
	}
	if r.flags.Vpc {
		row = append(row, r.vpcs.ById(i.VpcId).Tags["Name"])
	}
	if r.flags.SubnetId {
		row = append(row, strings.Join(i.Subnets, ", "))
	}
	if r.flags.Subnet {
		var names []string
		for _, n := range r.subnets.BySubnetIds(i.Subnets) {
			names = append(names, n.Tags["Name"])
		}
		row = append(row, strings.Join(names, ", "))
	}
	return row
}
