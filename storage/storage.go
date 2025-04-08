package storage

import (
	"encoding/json"
	"os"

	"github.com/Rishav-R03/tasks_cli/models"
	// "models.Task"
)

const fileName = "tasks.json"

// LoadTasks reads tasks from the JSON file and returns them as a slice.
// If the file doesn't exist, it returns an empty slice without an error

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task            // create empty slice to store the tasks of type struct
	data, err := os.ReadFile(fileName) // tries to read all the content from tasks.json file
	// data will be a byte slice ([]byte)
	if err != nil { // handling the error if file is not found
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &tasks) // parsing the json to Go structs
	// converts json into the tasks slice (uses Go's encoding/json package)
	// if json is invalid error will be not nil
	return tasks, err //return the result
	// if all went well err will be nil and tasks list will be returned
	// if decoding failed it will return the empty tasks and the error
}

func SaveTasks(tasks []models.Task) error {
	// Convert the tasks slice into nicely formatted JSON (indented with spaces)
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		// If marshalling fails, return the error to the caller
		return err
	}

	// Write the JSON data to the file with read/write permissions for the owner
	return os.WriteFile(fileName, data, 0644)
}

