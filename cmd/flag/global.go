package flag

import (
	"fmt"
	"github.com/pete911/aws-ip/internal"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

type Global struct {
	Region   string
	logLevel string
}

func (f Global) Logger() *slog.Logger {
	l, err := internal.NewLogger(f.logLevel)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return l
}

func (f Global) Config() internal.Config {
	cfg, err := internal.NewConfig(f.Region)
	if err != nil {
		fmt.Printf("new aws config: %v\n", err)
		os.Exit(1)
	}
	return cfg
}

func InitPersistentFlags(cmd *cobra.Command, flags *Global) {
	cmd.PersistentFlags().StringVar(
		&flags.Region,
		"region",
		getStringEnv("REGION", ""),
		"aws region",
	)
	cmd.PersistentFlags().StringVar(
		&flags.logLevel,
		"log-level",
		"warn",
		"log level - debug, info, warn, error",
	)
}

func getStringEnv(envName string, defaultValue string) string {
	env, ok := os.LookupEnv(fmt.Sprintf("AWS_IP_%s", envName))
	if !ok {
		return defaultValue
	}
	return env
}
