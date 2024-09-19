package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Task        string
	IsCompleted bool
}

var todos []Todo

func addTodo(task string) {
	todos = append(todos, Todo{Task: task, IsCompleted: false})
}

func completeTodo(index int) {
	if index >= 0 && index < len(todos) {
		todos[index].IsCompleted = true
	}
}

func displayTodos() {
	if len(todos) == 0 {
		fmt.Println("No todos yet!")
		return
	}
	for i, todo := range todos {
		status := " "
		if todo.IsCompleted {
			status = "X"
		}
		fmt.Printf("[%s] %d. %s\n", status, i+1, todo.Task)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo List ---")
		displayTodos()
		fmt.Println("\nCommands: add <task>, complete <number>, quit")
		fmt.Print("Enter a command: ")

		scanner.Scan()
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)

		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Please provide a task to add.")
				continue
			}
			addTodo(parts[1])
		case "complete":
			if len(parts) < 2 {
				fmt.Println("Please provide the number of the task to complete.")
				continue
			}
			var index int
			fmt.Sscanf(parts[1], "%d", &index)
			completeTodo(index - 1)
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}
