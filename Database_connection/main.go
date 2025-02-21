package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Database connection
	var err error
	db, err = sql.Open("mysql", "aditilondhe07:Venkatesh@777@tcp(192.168.0.144)/theschool_api")
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
		getTeachers(w, r)
	case "POST":
		createTeacher(w, r)
	case "PUT":
		updateTeacher(w, r)
	case "DELETE":
		deleteTeacher(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// createTeacher handles POST requests to add a new teacher
func createTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO teachers (name, subject) VALUES (?, ?)"
	_, err = db.Exec(query, teacher.Name, teacher.Subject)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

// getTeachers handles GET requests to retrieve all teachers
func getTeachers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, subject FROM teachers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var teachers []Teacher
	for rows.Next() {
		var teacher Teacher
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Subject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		teachers = append(teachers, teacher)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

// updateTeacher handles PUT requests to update an existing teacher
func updateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE teachers SET name = ?, subject = ? WHERE id = ?"
	_, err = db.Exec(query, teacher.Name, teacher.Subject, teacher.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teacher)
}

// deleteTeacher handles DELETE requests to remove a teacher
func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := "DELETE FROM teachers WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// classesHandler handles requests for the /classes endpoint
func classesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getClasses(w, r)
	case "POST":
		createClass(w, r)
	case "PUT":
		updateClass(w, r)
	case "DELETE":
		deleteClass(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// createClass handles POST requests to add a new class
func createClass(w http.ResponseWriter, r *http.Request) {
	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO classes (name, teacher_id) VALUES (?, ?)"
	_, err = db.Exec(query, class.Name, class.TeacherID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(class)
}

// getClasses handles GET requests to retrieve all classes
func getClasses(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, teacher_id FROM classes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var classes []Class
	for rows.Next() {
		var class Class
		err := rows.Scan(&class.ID, &class.Name, &class.TeacherID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		classes = append(classes, class)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

// updateClass handles PUT requests to update an existing class
func updateClass(w http.ResponseWriter, r *http.Request) {
	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE classes SET name = ?, teacher_id = ? WHERE id = ?"
	_, err = db.Exec(query, class.Name, class.TeacherID, class.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(class)
}

// deleteClass handles DELETE requests to remove a class
func deleteClass(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := "DELETE FROM classes WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// studentsHandler handles requests for the /students endpoint
func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getStudents(w, r)
	case "POST":
		createStudent(w, r)
	case "PUT":
		updateStudent(w, r)
	case "DELETE":
		deleteStudent(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getStudents handles GET requests to retrieve all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, class_id FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.Name, &student.ClassID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, student)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// createStudent handles POST requests to add a new student
func createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO students (name, class_id) VALUES (?, ?)"
	_, err = db.Exec(query, student.Name, student.ClassID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

// updateStudent handles PUT requests to update an existing student
func updateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE students SET name = ?, class_id = ? WHERE id = ?"
	_, err = db.Exec(query, student.Name, student.ClassID, student.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(student)
}

// deleteStudent handles DELETE requests to remove a student
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	query := "DELETE FROM students WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
