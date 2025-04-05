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

```
### 5 🧠 Concepts Covered
✅ Structs and JSON marshaling

✅ CLI argument parsing using flag

✅ File I/O with error handling

✅ Modular project structure in Go

✅ Beginner-friendly and extensible architecture

### 6 📈 Lines of Code (Approx)
File	LOC
main.go	~50
models/tasks.go	~10
storage/storage.go	~40
Total	~100
Small but powerful codebase — easy to read, easy to scale.

### 7 🛠️ Future Improvements
 Delete task by ID

 Edit task title

 Switch from JSON to BoltDB or SQLite

 Add interactive UI using Bubble Tea

 Cobra CLI for subcommands (todo add, todo list, etc.)

 Dockerized CLI tool

 Homebrew distribution for macOS

### 🤝 Contributing
Pull requests are welcome! If you’d like to improve functionality or add new features, feel free to fork and open a PR.

📄 License
MIT License — feel free to use, modify, and share!

### ✨ Author
Made with ❤️ by Rishav-R03

---

Let me know if you want to add:
- Badges (Go version, license, build status)
- Docker instructions
- A GIF demo of the CLI

Want to turn this into a Cobra-powered CLI app next?


