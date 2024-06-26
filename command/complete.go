package command

import (
	"fmt"
	"github.com/ronveen/godoit/internal"
	"github.com/teris-io/cli"
	"strconv"
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

	idStr := args[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID is not a number")
	}

	todos, highValue := internal.Load()
	found := false
	for i, todo := range todos {
		if todo.Id == id {
			todos[i].Done = true
			found = true
			fmt.Printf("Completed %d (%s)\n", todo.Id, todo.Text)
		}
	}

	internal.Store(todos, highValue)

	if !found {
		fmt.Printf("Could not find Todo %d!\n", id)
	}
	return 0
}
