package main

import (
	"os"

	"github.com/zengqiang96/mattermostclone/cmd/mattermost/commands"
)

func main() {
	if err := commands.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
