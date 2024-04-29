package command

import (
	"github.com/ronveen/godoit/internal"
	"github.com/teris-io/cli"
	"log"
	"os"
	"text/template"
	"time"
)

func ListCommand() cli.Command {
	return cli.NewCommand("list", "list reminders").
		WithShortcut("ls").
		WithOption(cli.NewOption("complete", "Show completed todos").
			WithChar('c').
			WithType(cli.TypeBool)).
		WithOption(cli.NewOption("all", "Show all, regardless of date").
			WithChar('c').
			WithType(cli.TypeBool)).
		WithOption(cli.NewOption("due", "alternative due date (format yyyymmmdd)").
			WithChar('d').
			WithType(cli.TypeString)).
		WithAction(listAction)
}

func listAction(args []string, options map[string]string) int {
	type listData struct {
		Due   string
		Todos []internal.Todo
	}

	var todos, _ = internal.Load()
	var filterDate = truncateToStartOfDay(time.Now())

	showCompletes := options["complete"] == "true"
	showAll := options["all"] == "true"
	dateStr, found := options["due"]

	if showAll && found {
		log.Fatal("Cannot use both --due and --all")
	}

	if found {
		date, err := time.Parse("20060102", dateStr)
		if err == nil {
			filterDate = date
		}
	}

	result := make([]internal.Todo, 0)
	for _, t := range todos {
		if showCompletes || !t.Done {
			if showAll || filterDate.Equal(t.Due) {
				result = append(result, t)
			}
		}
	}

	var templateName = "command/list.tmpl"
	tmpl := template.Must(template.ParseFiles(templateName))
	err := tmpl.Execute(os.Stdout, listData{
		Due:   filterDate.Format("2006-01-02"),
		Todos: result,
	})
	if err != nil {
		log.Fatal(err)
	}

	return 0
}

func truncateToStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}
