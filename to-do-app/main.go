package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Todo struct {
	Task        string
	IsCompleted bool
}

type TodoList struct {
	Todos []Todo
}

func (tl *TodoList) addTodo(task string) {
	tl.Todos = append(tl.Todos, Todo{Task: task, IsCompleted: false})
}

func (tl *TodoList) completeTodo(index int) {
	if index >= 0 && index < len(tl.Todos) {
		tl.Todos[index].IsCompleted = true
	}
}

func (tl *TodoList) removeTodo(index int) {
	if index >= 0 && index < len(tl.Todos) {
		tl.Todos = append(tl.Todos[:index], tl.Todos[index+1:]...)
	}
}

func (tl *TodoList) displayTodos() {
	for i, todo := range tl.Todos {
		status := " "
		if todo.IsCompleted {
			status = "X"
		}
		fmt.Printf("[%s] %d. %s\n", status, i+1, todo.Task)
	}
}

func (tl *TodoList) saveToFile(filename string) error {
	data, err := json.Marshal(tl)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func loadFromFile(filename string) (*TodoList, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return &TodoList{}, nil
	}
	var tl TodoList
	err = json.Unmarshal(data, &tl)
	if err != nil {
		return nil, err
	}
	return &tl, nil
}

func main() {
	filename := "todos.json"
	todoList, err := loadFromFile(filename)
	if err != nil {
		fmt.Println("Error loading todo list:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Todo List ---")
		todoList.displayTodos()
		fmt.Println("\nCommands: add, complete, remove, quit")
		fmt.Print("Enter a command: ")

		scanner.Scan()
		input := scanner.Text()
		command := strings.Fields(input)

		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case "add":
			if len(command) < 2 {
				fmt.Println("Please provide a task to add.")
				continue
			}
			task := strings.Join(command[1:], " ")
			todoList.addTodo(task)
		case "complete":
			if len(command) != 2 {
				fmt.Println("Please provide the number of the task to complete.")
				continue
			}
			var index int
			fmt.Sscanf(command[1], "%d", &index)
			todoList.completeTodo(index - 1)
		case "remove":
			if len(command) != 2 {
				fmt.Println("Please provide the number of the task to remove.")
				continue
			}
			var index int
			fmt.Sscanf(command[1], "%d", &index)
			todoList.removeTodo(index - 1)
		case "quit":
			err := todoList.saveToFile(filename)
			if err != nil {
				fmt.Println("Error saving todo list:", err)
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command. Please try again.")
		}

		err := todoList.saveToFile(filename)
		if err != nil {
			fmt.Println("Error saving todo list:", err)
		}
	}
}
