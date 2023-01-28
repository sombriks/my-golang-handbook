package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type todo struct {
	description string
	done        bool
}

var fileName = "todo.txt"
var todos = make([]todo, 0)

func loadFile() {

	var todoFile *os.File
	_, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		todoFile, _ = os.Create(fileName)
	} else {
		todoFile, _ = os.Open(fileName)
	}
	defer todoFile.Close()

	todoFileScanner := bufio.NewScanner(todoFile)
	todoFileScanner.Split(bufio.ScanLines)

	for todoFileScanner.Scan() {
		todoLine := todoFileScanner.Text()
		todoSplit := strings.Split(todoLine, ":::")
		t := todo{description: todoSplit[0]}
		t.done, _ = strconv.ParseBool(todoSplit[1])
		todos = append(todos, t)
	}

}

func list() {
	for i, e := range todos {
		// https://stackoverflow.com/a/31483763/420096
		fmt.Printf("%d) [%s] %s\n", i, map[bool]string{true: "✓", false: " "}[e.done], e.description)
	}
}

func add() {
	fmt.Println("provide the new todo description")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	todos = append(todos, todo{input, false})
}

func resolve() {
	list()
	fmt.Println("choose a number to resolve")
	var i int
	fmt.Scanln(&i)
	todos[i].done = true
}

func save() {

	todoFile, _ := os.OpenFile(fileName, os.O_WRONLY, 666)
	defer todoFile.Close()

	for _, e := range todos {
		line := fmt.Sprintf("%s:::%t", e.description, e.done)
		fmt.Fprintln(todoFile, line)
	}
	todoFile.Sync()
}

func main() {
	loadFile()
	var op string
	for op != "quit" {
		fmt.Println("choose an option:")
		fmt.Println("list")
		fmt.Println("add")
		fmt.Println("resolve")
		fmt.Println("quit")
		fmt.Println()
		fmt.Scanln(&op)
		fmt.Println()
		switch op {
		case "list":
			list()
		case "add":
			add()
		case "resolve":
			resolve()
		case "quit":
			op = "quit"
			save()
		default:
			fmt.Println("invalid option")
		}
		fmt.Println()
	}
}
