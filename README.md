# 📚 Student Management System (Go)

A simple Command Line Interface (CLI) application built with Go to manage student records using JSON as local storage.

This project was created as my first Go project to practice core Go concepts such as file handling, structs, slices, JSON encoding/decoding, and modular programming.

---

## ✨ Features

- 📄 Automatically create `data.json` if it doesn't exist
- ➕ Add a new student
- 🔍 Search student by ID
- ✏️ Update student information
- ❌ Delete student by ID
- 📋 View all students
- 💾 Save data into JSON file
- 📂 Load data from JSON file

---

## 📁 Project Structure

```
.
├── main.go              # Application entry point
├── student.go           # Student struct
├── file.go              # Create, Load, Save file functions
├── menu.go              # Add, Search, Update, Delete features
├── data.json            # Local database
└── README.md
```


## 🛠️ Technologies

- Go
- Standard Library
    - os
    - bufio
    - encoding/json
    - strconv
    - strings
    - fmt

No third-party packages are used.

---

## 📦 JSON Storage

Student data is stored locally inside `data.json`.

Example:

```json
[
  {
    "id": 1,
    "name": "Jouvan Augusto",
    "age": 20,
    "address": "Bandung"
  },
  {
    "id": 2,
    "name": "Riyan",
    "age": 21,
    "address": "Jakarta"
  }
]
```

---

## 🚀 How to Run

Clone the repository:

```bash
git clone https://github.com/yourusername/student-management-go.git
```

Go to the project directory:

```bash
cd student-management-go
```

Run the application:

```bash
go run .
```

---

## 📖 What I Learned

Through this project I practiced:

- Working with structs
- Creating and using slices
- Reading and writing files
- JSON Marshaling & Unmarshaling
- Error handling
- Building CRUD applications
- Using `bufio.Reader` for user input
- Organizing code into multiple Go files
- Using Go's standard library effectively

---

## 🏗️ Application Architecture

The application follows a simple file-based CRUD architecture.

```text
                User
                  │
                  ▼
             CLI Menu (main.go)
                  │
      ┌───────────┼────────────┐
      │           │            │
      ▼           ▼            ▼
 Add Student  Update Student  Delete Student
      │           │            │
      └───────────┼────────────┘
                  │
                  ▼
           LoadStudent()
                  │
                  ▼
      Read data.json from disk
                  │
                  ▼
        []Student (Slice)
                  │
      Add / Update / Delete
                  │
                  ▼
          SaveStudent()
                  │
                  ▼
      Write updated JSON
                  │
                  ▼
             data.json
```

Every CRUD operation follows the same workflow:

1. Load all student data from `data.json`.
2. Convert JSON into `[]Student`.
3. Modify the slice in memory.
4. Save the updated slice back to `data.json`.

This approach ensures that the JSON file always remains synchronized with the latest data.