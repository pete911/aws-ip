package cmd

import (
	"context"
	"fmt"
	"github.com/pete911/aws-ip/internal/out"
	"github.com/pete911/aws-ip/internal/rds"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	rdsCmd = &cobra.Command{
		Use:   "rds",
		Short: "view rds IPs",
		Long:  "",
		Run:   runRds,
	}
)

func init() {
	Root.AddCommand(rdsCmd)
}

func runRds(cmd *cobra.Command, _ []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logger := GlobalFlags.Logger().With("cmd", cmd.Use)
	RdsRunner{
		logger: logger,
		client: rds.NewClient(logger, GlobalFlags.Config()),
	}.Run(ctx)
}

type RdsRunner struct {
	logger *slog.Logger
	client rds.Client
}

func (r RdsRunner) Run(ctx context.Context) {
	instances, err := r.client.ListInstances(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	r.print(instances)
}

func (r RdsRunner) print(instances []rds.Instance) {
	table := out.NewTable(r.logger, os.Stdout)
	table.AddRow("NAME", "ENDPOINT", "PORT", "IP", "VPC", "SUBNETS")
	for _, i := range instances {
		ips, err := i.LookupIp()
		if err != nil {
			r.logger.Error(fmt.Sprintf("lookup %s ip: %v", i.Endpoint.Address, err))
		}
		table.AddRow(
			i.DBName,
			i.Endpoint.Address,
			fmt.Sprintf("%d", i.Endpoint.Port),
			strings.Join(ips, ", "),
			i.VpcId,
			strings.Join(i.Subnets, ", "),
		)
	}
	table.Print()
}
