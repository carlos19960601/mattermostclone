package commands

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/zengqiang96/mattermostclone/api4"
	"github.com/zengqiang96/mattermostclone/app"
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
	options := []app.Option{
		app.ConfigStore(configStore),
	}

	server, err := app.NewServer(options...)
	if err != nil {
		return err
	}
	defer server.Shutdown()

	a := app.New(app.ServerConnector(server))
	api4.Init(a, server.Router)

	serverErr := server.Start()
	if serverErr != nil {
		return serverErr
	}

	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan
	fmt.Println("关闭服务...")
	return nil
}
