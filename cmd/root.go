package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/pete911/aws-ip/cmd/flag"
	"github.com/pete911/aws-ip/internal/ec2"
	"github.com/spf13/cobra"
)

var (
	GlobalFlags flag.Global
	Root        = &cobra.Command{Short: "view public IPs of AWS services"}

	Version string
)

func init() {
	flag.InitPersistentFlags(Root, &GlobalFlags)
}

func loadVpcs(ctx context.Context, client ec2.Client) ec2.Vpcs {
	vpcs, err := client.DescribeVPCs(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return vpcs
}

func loadSubnets(ctx context.Context, client ec2.Client) ec2.Subnets {
	subnets, err := client.DescribeSubnets(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return subnets
}
