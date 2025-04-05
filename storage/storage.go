package storage 

import (
	"os"
	"encoding/json"
	"models.tasks"
)

const  fileName = "tasks.json"

func loadTasks()([]Task,error){
	var tasks [] Task
	data,err := os.ReadFile(fileName)
	if err != nil{
		if os.IsNotExist(err){
			return []Task{}, nil 
		}
		return nil,err
	}
	err = json.Unmarshal(data,&tasks)
	return tasks,err 
}

func saveTasks(taks)

