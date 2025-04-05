package main

import (
	"fmt"

	"flag"

	"github.com/Rishav-R03/tasks_cli/models"
	"github.com/Rishav-R03/tasks_cli/storage"
)

func main() {
	add := flag.String("add", "", "Add a new task")
	list := flag.Bool("list", false, "List all tasks")
	done := flag.Int("done", -1, "Mark as done (provide ID)")
	flag.Parse()
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks : ", err)
		return
	}
	switch {
	case *add != "":
		tasks = append(tasks, models.Task{
			ID:    len(tasks) + 1,
			Title: *add,
			Done:  false,
		})
		storage.SaveTasks(tasks)
		fmt.Println("Task added!")

	case *list:
		if len(tasks) == 0 {
			fmt.Println("No tasks yet!")
		}
		for _, t := range tasks {
			status := "[ ]"
			if t.Done {
				status = "[✓]"
			}
			fmt.Printf("%d. %s %s\n", t.ID, status, t.Title)
		}
	case *done > 0:
		for i, t := range tasks {
			if t.ID == *done {
				tasks[i].Done = true
				storage.SaveTasks(tasks)
				fmt.Println("Marked as done!")
				return
			}
		}
		fmt.Println("Tasks ID not found")
	default:
		fmt.Println("Usage: ")
		flag.PrintDefaults()
	}

}
