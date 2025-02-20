package handlers

import (
	"context"
	"net/http"

	"yourmodule/ent"

	"github.com/aditilondhe07/theschool-api/ent/class"
	"github.com/gin-gonic/gin"
)

// CreateStudent creates a new student and assigns them to a class.
func CreateStudent(c *gin.Context, client *ent.Client) {
	type CreateStudentInput struct {
		Name    string `json:"name" binding:"required"`
		ClassID int    `json:"class_id" binding:"required"`
	}
	var input CreateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Retrieve the class by ID.
	cls, err := client.Class.
		Query().
		Where(class.IDEQ(input.ClassID)).
		Only(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "class not found"})
		return
	}
	// Create the student and set its class edge.
	stu, err := client.Student.
		Create().
		SetName(input.Name).
		SetClass(cls).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, stu)
}

// ListStudents returns all students along with their associated class.
func ListStudents(c *gin.Context, client *ent.Client) {
	students, err := client.Student.
		Query().
		WithClass().
		All(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list students"})
		return
	}
	c.JSON(http.StatusOK, students)
}
