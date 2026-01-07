package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readTasks() []string {
	data, err := os.ReadFile("tasks.txt")
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

		file, _ := os.OpenFile(
			"tasks.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		defer file.Close()

		file.WriteString("[ ] " + task + "\n")
		fmt.Println("Task added ✅")

	case "list":
		tasks := readTasks()

		for i, task := range tasks {
			if task != "" {
				fmt.Printf("%d. %s\n", i+1, task)
			}
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task number")
			return
		}

		taskNum, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Println("Invalid task number")
			return
		}

		tasks := readTasks()

		if taskNum <=0 || taskNum > len(tasks){
			fmt.Println("Task not found")
			return 
		}

		tasks[taskNum-1] = strings.Replace(
			tasks[taskNum-1],
			"[ ]",
			"[x]",
			1,
		)

		output := strings.Join(tasks,"\n")
		os.WriteFile("tasks.txt",[]byte(output),0644)

		fmt.Println("Task marked as done ✅")

	default:
		fmt.Println("Unknown command")
	}

}
