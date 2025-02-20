package handlers

import (
	"context"
	"net/http"

	"yourmodule/ent"
	"yourmodule/ent/teacher"

	"github.com/gin-gonic/gin"
)

// CreateClass creates a new class and associates it with a teacher.
func CreateClass(c *gin.Context, client *ent.Client) {
	type CreateClassInput struct {
		Name      string `json:"name" binding:"required"`
		TeacherID int    `json:"teacher_id" binding:"required"`
	}
	var input CreateClassInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Retrieve teacher by ID.
	t, err := client.Teacher.
		Query().
		Where(teacher.IDEQ(input.TeacherID)).
		Only(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "teacher not found"})
		return
	}
	// Create the class and set its teacher edge.
	cls, err := client.Class.
		Create().
		SetName(input.Name).
		SetTeacher(t).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create class"})
		return
	}
	c.JSON(http.StatusCreated, cls)
}

// ListClasses returns a list of all classes along with their associated teacher.
func ListClasses(c *gin.Context, client *ent.Client) {
	classes, err := client.Class.
		Query().
		WithTeacher().
		All(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list classes"})
		return
	}
	c.JSON(http.StatusOK, classes)
}
