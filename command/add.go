package command

import (
	"fmt"
	"github.com/teris-io/cli"
)

func AddCommand() cli.Command {
	return cli.NewCommand("add", "Add a reminder").
		WithArg(cli.NewArg("text", "Text of the new reminder")).
		WithOption(cli.NewOption("due", "Due date for the reminder").
			WithChar('d').
			WithType(cli.TypeString)).
		WithAction(addAction)
}

func addAction(args []string, options map[string]string) int {
	fmt.Println("Added reminder: " + args[0] + " for date " + options["due"])
	return 0
}
