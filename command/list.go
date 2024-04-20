package command

import (
	"github.com/ronveen/godoit/internal"
	"github.com/teris-io/cli"
	"time"
)

func ListCommand() cli.Command {
	return cli.NewCommand("list", "list reminders").
		WithShortcut("ls").
		WithOption(cli.NewOption("id", "ID of the reminder to show").
			WithChar('n').
			WithType(cli.TypeInt)).
		WithOption(cli.NewOption("due", "alternative due date (format yyyymmmdd)").
			WithChar('d').
			WithType(cli.TypeString)).
		WithAction(listAction)
}

func listAction(args []string, options map[string]string) int {
	var todos, _ = internal.Load()
	if len(todos) == 0 {
		println("No todos found")
	}

	var filterDate = truncateToStartOfDay(time.Now())

	dateStr, found := options["due"]
	if found {
		date, err := time.Parse("20060102", dateStr)
		if err == nil {
			filterDate = date
		}
	}

	println("Items due on " + filterDate.Format("2006-01-02"))

	for _, t := range todos {

		if filterDate.Equal(t.Due) {
			printTodo(t)
		}
	}

	return 0
}

func printTodo(t internal.Todo) {
	if t.Done {
		print("\u2714 ")
	} else {
		print("  ")
	}
	print(t.Id)
	println(" " + t.Text)
}

func truncateToStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
