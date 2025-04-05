# 📝 To-Do List CLI in Go

A simple and efficient command-line To-Do application built using **Golang**, following best practices in code organization and modular structure. This project is perfect for learning Go fundamentals such as file I/O, structs, flags, and working with JSON.

---

## 📌 Features

- Add new tasks via CLI
- List all tasks
- Mark tasks as completed
- Data persistence using a JSON file
- Modular folder structure (`models`, `storage`)
- Simple, fast, and extensible

---

## 🏗️ Project Structure

tasks_cli/ │ 
  ├── main.go # Entry point with CLI flags 
  ├── go.mod # Go module file │ 
  ├── models/ │ 
    └── tasks.go # Task struct definition │ 
  └── storage/   
    └── storage.go # Load/save logic using JSON


---

## 📦 Tech Stack

- Language: **Golang**
- Modules:
  - `encoding/json` – for marshaling/unmarshaling tasks
  - `os` – for file operations
  - `flag` – for command-line flag parsing

---

## 🚀 How to Use

### 1. Clone the repository

```bash
git clone https://github.com/Rishav-R03/tasks_cli.git
cd tasks_cli
```
### 2. Build the App
```
go build -o todo
```
### 3. Running the App
```
# Add a task
./todo --add "Buy groceries"

# List tasks
./todo --list

# Mark a task as done
./todo --done 1
```
### 4. Example of JSON file
```
[
  {
    "id": 1,
    "title": "Buy groceries",
    "done": true
  },
  {
    "id": 2,
    "title": "Finish Go project",
    "done": false
  }
]


