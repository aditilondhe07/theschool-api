package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/school-api/ent"
)

// createTeacher creates a new teacher.
func createTeacher(c *gin.Context, client *ent.Client) {
	type Request struct {
		Name string `json:"name" binding:"required"`
	}
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t, err := client.Teacher.
		Create().
		SetName(req.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create teacher"})
		return
	}
	c.JSON(http.StatusCreated, t)
}

// listTeachers returns a list of teachers.
func listTeachers(c *gin.Context, client *ent.Client) {
	teachers, err := client.Teacher.
		Query().
		All(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query teachers"})
		return
	}
	c.JSON(http.StatusOK, teachers)
}
