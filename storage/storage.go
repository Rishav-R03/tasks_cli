package storage

import (
	"encoding/json"
	"os"

	"github.com/Rishav-R03/tasks_cli/models"
	// "models.Task"
)

const fileName = "tasks.json"

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
