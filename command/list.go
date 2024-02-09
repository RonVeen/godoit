package command

import (
	"fmt"
	"github.com/teris-io/cli"
)

func ListCommand() cli.Command {
	return cli.NewCommand("list", "list reminders").
		WithOption(cli.NewOption("id", "ID of the reminder to show").
			WithChar('n').
			WithType(cli.TypeInt)).
		WithOption(cli.NewOption("due", "alternative duedate (format yyyymmmdd)").
			WithChar('d').
			WithType(cli.TypeString)).
		WithAction(listAction)
}

func listAction(args []string, options map[string]string) int {
	if options["id"] != "" {
		fmt.Println("Listing details for " + options["id"])
	} else {
		if options["due"] == "" {
			fmt.Println("Listing options")
		} else {
			fmt.Println("Listing options for " + options["due"])
		}
	}
	return 0
}
