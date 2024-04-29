package command

import (
	"fmt"
	"github.com/ronveen/godoit/internal"
	"github.com/teris-io/cli"
	"strconv"
)

func UndoCommand() cli.Command {
	return cli.NewCommand("undo", "Undo completing of a reminder").
		WithArg(cli.NewArg("id", "ID of the reminder to complete")).
		WithAction(undoAction)
}

func undoAction(args []string, options map[string]string) int {
	if len(args) == 0 {
		fmt.Println("ID is required for option Undo")
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
			todos[i].Done = false
			found = true
			fmt.Printf("Todo %d (%s) is reverted back to undone\n", todo.Id, todo.Text)
		}
	}

	internal.Store(todos, highValue)

	if !found {
		fmt.Printf("Could not find Todo %d!\n", id)
	}
	return 0
}
