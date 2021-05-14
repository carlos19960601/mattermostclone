package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zengqiang96/mattermostclone/config"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the Mattermost server",
	RunE:  serverCmdF,
}

func init() {
	RootCmd.AddCommand(serverCmd)
	RootCmd.RunE = serverCmdF
}

func serverCmdF(command *cobra.Command, args []string) error {
	interruptChan := make(chan os.Signal, 1)

	configStore, err := config.NewStore(getConfigDSN(command))
	if err != nil {
		return fmt.Errorf("加载配置失败 err: %w", err)
	}

	defer configStore.Close()

	return runServer(configStore, interruptChan)
}

func runServer(configStore *config.Store, interruptChan chan os.Signal) error {
	return nil
}
