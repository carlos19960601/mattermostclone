package commands

import (
	"github.com/spf13/cobra"
)

func getConfigDSN(command *cobra.Command) string {
	configDSN, _ := command.Flags().GetString("config")
	return configDSN
}
