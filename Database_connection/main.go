
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/aditilondhe07/theschool-api/ent"
	"context"
	"github.com/go-sql-driver/mysql"
)

var client *ent.Client

func main() {
	// Set up the connection to the database using Ent
	var err error
	client, err = ent.Open("mysql", "aditilondhe07:Venkatesh@777@tcp(192.168.0.144)/theschool_api")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Ping the database to check the connection
	if err := client.Ping(); err != nil {
		log.Fatalf("failed pinging database: %v", err)
	}

	log.Println("Successfully connected to the database!")

	// Routes
	http.HandleFunc("/teachers", teachersHandler)
	http.HandleFunc("/classes", classesHandler)
	http.HandleFunc("/students", studentsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Teacher struct
type Teacher struct {
	ID      int    json:"id"
	Name    string json:"name"
	Subject string json:"subject"
}

// Class struct
type Class struct {
	ID        int    json:"id"
	Name      string json:"name"
	TeacherID int    json:"teacher_id"
}

// Student struct
type Student struct {
	ID      int    json:"id"
	Name    string json:"name"
	ClassID int    json:"class_id"
}

// CRUD Handlers for Teachers

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

func createTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTeacher, err := client.Teacher.
		Create().
		SetName(teacher.Name).
		SetSubject(teacher.Subject).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTeacher)
}

func getTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := client.Teacher.
		Query().
		All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teachers)
}

func updateTeacher(w http.ResponseWriter, r *http.Request) {
	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTeacher, err := client.Teacher.
		UpdateOneID(teacher.ID).
		SetName(teacher.Name).
		SetSubject(teacher.Subject).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTeacher)
}

func deleteTeacher(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := client.Teacher.
		Delete().
		Where(ent.ID(id)).
		Exec(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CRUD Handlers for Classes

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

func createClass(w http.ResponseWriter, r *http.Request) {
	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdClass, err := client.Class.
		Create().
		SetName(class.Name).
		SetTeacherID(class.TeacherID).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdClass)
}

func getClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := client.Class.
		Query().
		All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classes)
}

func updateClass(w http.ResponseWriter, r *http.Request) {
	var class Class
	err := json.NewDecoder(r.Body).Decode(&class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedClass, err := client.Class.
		UpdateOneID(class.ID).
		SetName(class.Name).
		SetTeacherID(class.TeacherID).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedClass)
}

func deleteClass(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := client.Class.
		Delete().
		Where(ent.ID(id)).
		Exec(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CRUD Handlers for Students

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

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdStudent, err := client.Student.
		Create().
		SetName(student.Name).
		SetClassID(student.ClassID).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdStudent)
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	students, err := client.Student.
		Query().
		All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedStudent, err := client.Student.
		UpdateOneID(student.ID).
		SetName(student.Name).
		SetClassID(student.ClassID).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedStudent)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := client.Student.
		Delete().
		Where(ent.ID(id)).
		Exec(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
} 