package main

import (
	"github.com/ronveen/godoit/command"
	cli "github.com/teris-io/cli"
	"os"
)

func main() {
	app := cli.New("Manage reminders").
		WithCommand(command.AddCommand()).
		WithCommand(command.ListCommand()).
		WithCommand(command.CompleteCommand()).
		WithCommand(command.UndoCommand()).
		WithCommand(command.PurgeCommand())

	os.Exit(app.Run(os.Args, os.Stdout))

}
