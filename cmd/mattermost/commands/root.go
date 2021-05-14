package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "mattermost",
}

func Run(args []string) error {
	RootCmd.SetArgs(args)
	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringP("config", "c", "", "配置文件路径")
}
