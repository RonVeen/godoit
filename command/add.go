package command

import (
	"fmt"
	"github.com/ronveen/godoit/internal"
	"github.com/teris-io/cli"
	"log"
	"time"
)

const File = "todo.txt"

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
	var todos, highValue = internal.Load()
	highValue++
	time := checkTime(options)

	todo := internal.Todo{Id: highValue, Text: args[0], Done: false, Due: time}

	todos = append(todos, todo)
	internal.Store(todos, highValue)
	return 0
}

func checkTime(options map[string]string) time.Time {
	dateStr, found := options["due"]
	if found {
		date, err := time.Parse("20060102", dateStr)
		if err != nil {
			log.Fatal("Invalid date specified: " + dateStr)
		}
		return date
	}
	return time.Now()
}
