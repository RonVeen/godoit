package command

import (
	"fmt"
	"github.com/teris-io/cli"
)

func CompleteCommand() cli.Command {
	return cli.NewCommand("complete", "complete a reminder").
		WithOption(cli.NewOption("id", "ID of the reminder to complete").
			WithChar('n').
			WithType(cli.TypeInt)).
		WithAction(completeAction)
}

func completeAction(args []string, options map[string]string) int {
	if options["id"] == "" {
		fmt.Println("ID is required for option Complete")
		return 1
	}

	fmt.Println("Completing " + options["id"])
	return 0
}
