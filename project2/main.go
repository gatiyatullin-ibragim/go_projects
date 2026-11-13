package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	id        int
	Title     string
	Completed bool
}

var tasks []Task
var nextId = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		fmt.Println("Choose action: ")

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			listTasks()

		case "2":
			fmt.Println("Enter name of a task: ")
			if scanner.Scan() {
				title := strings.TrimSpace(scanner.Text())
				if title != "" {
					addTask(title)
				}
			}

		case "3":
			listTasks()
			fmt.Println("Enter id of task you want to mark as done: ")
			if scanner.Scan() {
				var id int
				_, err := fmt.Sscan(scanner.Text(), &id)
				if err == nil {
					completeTask(id)
				}
			}

		case "4":
			listTasks()
			fmt.Println("Enter id of task you want to delete: ")
			if scanner.Scan() {
				var id int
				_, err := fmt.Sscan(scanner.Text(), &id)
				if err == nil {
					deleteTask(id)
				}
			}

		case "5":
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("Wrong choice. Try again")
		}
	}
}

func printMenu() {
	fmt.Println("=======MENU=======")
	fmt.Println("1.Show the list of tasks")
	fmt.Println("2.Add task to list")
	fmt.Println("3.Mark as done")
	fmt.Println("4.Remove task")
	fmt.Println("5.Exit")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("List of tasks is empty")
		return
	}
	fmt.Println("Your tasks: ")

	for _, tsk := range tasks {
		status := "[]"
		if tsk.Completed {
			status = "[x]"
		}
		fmt.Println("%d, %s %s\n", tsk.id, status, tsk.Title)
	}
}

func addTask(title string) {
	tasks = append(tasks, Task{
		id:        nextId,
		Title:     title,
		Completed: false,
	})
	fmt.Printf("Task added with id: %d\n", nextId)
	nextId++
}

func completeTask(id int) {
	for i := range tasks {
		if tasks[i].id == id {
			tasks[i].Completed = true
			fmt.Println("Task marked as completed")
			return
		}
	}
	fmt.Println("Task with this id is not found")
}

func deleteTask(id int) {
	for i, t := range tasks {
		if t.id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Tasks is deleted")
			return
		}
	}
	fmt.Println("Tasks with this Id is not found")
}
