package main

import (
	"fmt"
	"os"
	"strings"
)

func readTasks() []string {
	data, err := os.ReadFile("tasks,txt")
	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(data),"\n")
	return lines
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : task add|list|done")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task")
			return
		}
		task := os.Args[2]

		file, _ := os.OpenFile("tasks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()

		file.WriteString(task + "\n")
		fmt.Println("Task added ✅")

	case "list":

		data, _ := os.ReadFile("tasks.txt")
		fmt.Println(string(data))
	case "done":

		fmt.Println("Done command selected")
	default:
		fmt.Println("Unknown command")
	}

}
