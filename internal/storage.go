package internal

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const File = "./todo.txt"

type Todo struct {
	Text string
	Done bool
	Id   int
	Due  time.Time
}

func Load() ([]Todo, int) {
	t, v := textToTodo(readFile(File))
	return t, v
}

func Store(todos []Todo, highValue int) {
	writeFile(todos, highValue)
}

func textToTodo(data []string) ([]Todo, int) {
	todos := make([]Todo, 0, 0)

	//  First entry contains the current highest value
	//  Retrieve it and remove the entry from the slice
	firstElement := data[0]
	data = data[1:]

	highValue, err := strconv.Atoi(firstElement)
	if err != nil {
		log.Fatal("Data file is corrupt (high value missing or invalid")
	}

	for _, s := range data {
		elem := strings.Split(s, ";")
		if len(elem) == 4 {
			id, _ := strconv.Atoi(elem[0])
			date, _ := time.Parse("20060102", elem[3])
			todo := Todo{Id: id,
				Text: elem[1],
				Done: elem[2] == "1",
				Due:  date}
			todos = append(todos, todo)
		}
	}
	return todos, highValue
}

func readFile(name string) []string {
	var data = make([]string, 0, 25)
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func writeFile(todos []Todo, highValue int) {
	f, err := os.Create(File)
	if err != nil {
		log.Fatal("Error writing data to disk")
	}
	defer f.Close()

	//  First, write the highest used value
	f.WriteString(writeLine(strconv.Itoa(highValue)))

	//  Write all details
	for _, t := range todos {
		line := writeLine(strconv.Itoa(t.Id), t.Text, booleanToNumeric(t.Done), t.Due.Format("20060102"))
		f.WriteString(line)
	}
}

func booleanToNumeric(value bool) string {
	if value {
		return "1"
	} else {
		return "0"
	}
}

func writeLine(parts ...string) string {
	sb := strings.Builder{}
	for i, s := range parts {
		sb.WriteString(s)
		if i == len(parts)-1 {
			sb.WriteString("\n")
		} else {
			sb.WriteString(";")
		}
	}
	return sb.String()
}
