package command

import (
	"fmt"
	"github.com/ronveen/godoit/internal"
	"github.com/teris-io/cli"
)

func PurgeCommand() cli.Command {
	return cli.NewCommand("purge", "Purges completed reminders").
		WithAction(purgeAction)
}

func purgeAction(args []string, options map[string]string) int {
	fmt.Println("Purging reminders...")
	var todos, highValue = internal.Load()

	remaining := make([]internal.Todo, 0)
	count := 0
	for _, todo := range todos {
		if !todo.Done {
			remaining = append(remaining, todo)
			count++
		}
	}

	internal.Store(remaining, highValue)
	fmt.Printf("Purged %d todos\n", count)
	return 0
}
