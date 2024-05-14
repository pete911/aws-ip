package flag

import "github.com/spf13/cobra"

type Elasticache struct {
	Status   bool
	Vpc      bool
	Subnet   bool
	VpcId    bool
	SubnetId bool
}

func InitElasticacheFlags(cmd *cobra.Command, flags *Elasticache) {
	cmd.Flags().BoolVar(
		&flags.Status,
		"status",
		false,
		"print status",
	)
	cmd.Flags().BoolVar(
		&flags.Vpc,
		"vpc",
		false,
		"print vpc name",
	)
	cmd.Flags().BoolVar(
		&flags.Subnet,
		"subnet",
		false,
		"print subnet names",
	)
	cmd.Flags().BoolVar(
		&flags.VpcId,
		"vpc-id",
		false,
		"print vpc id",
	)
	cmd.Flags().BoolVar(
		&flags.SubnetId,
		"subnet-id",
		false,
		"print subnet ids",
	)
}
