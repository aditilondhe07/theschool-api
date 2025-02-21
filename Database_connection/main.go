package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Database connection
	var err error
	db, err = sql.Open("mysql", "aditilondhe07:Venkatesh@777@tcp(127.0.0.1:3306)/theschool-api")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")

	// Routes
	http.HandleFunc("/teachers", teachersHandler)
	http.HandleFunc("/classes", classesHandler)
	http.HandleFunc("/students", studentsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Teacher struct
type Teacher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}

// Class struct
type Class struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacher_id"`
}

// Student struct
type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	ClassID int    `json:"class_id"`
}

// CRUD Handlers

// teachersHandler handles requests for the /teachers endpoint
func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Retrieve list of teachers
	case "POST":
		// Add a new teacher
	case "PUT":
		// Update a teacher
	case "DELETE":
		// Delete a teacher
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// classesHandler handles requests for the /classes endpoint
func classesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Retrieve list of classes
	case "POST":
		// Add a new class
	case "PUT":
		// Update a class
	case "DELETE":
		// Delete a class
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// studentsHandler handles requests for the /students endpoint
func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Retrieve list of students
	case "POST":
		// Add a new student
	case "PUT":
		// Update a student
	case "DELETE":
		// Delete a student
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
