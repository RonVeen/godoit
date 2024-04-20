package command

import (
	"fmt"
	"github.com/teris-io/cli"
)

func CompleteCommand() cli.Command {
	return cli.NewCommand("complete", "complete a reminder").
		WithShortcut("co").
		WithArg(cli.NewArg("id", "ID of the reminder to complete")).
		WithAction(completeAction)
}

func completeAction(args []string, options map[string]string) int {
	if len(args) == 0 {
		fmt.Println("ID is required for option Complete")
		return 1
	}

	fmt.Println("Completing " + args[0])
	return 0
}
