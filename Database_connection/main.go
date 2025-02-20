package main

import (
	"context"
	"log"
	"time"

	"github.com/aditilondhe07/theschool-api/ent"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

func main() {
	// Replace these with your MySQL credentials and database name.
	dsn := "root:Venkatesh@777@tcp(127.0.0.1:3306)/theschool-apiparseTime=True"
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Set a connection timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Run the auto migration tool to create/update database schema.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize your router using Gin.
	router := gin.Default()

	// Define API endpoints.
	router.POST("/teachers", func(c *gin.Context) { createTeacher(c, client) })
	router.GET("/teachers", func(c *gin.Context) { listTeachers(c, client) })
	// Define similar endpoints for classes and students.

	// Start the server.
	router.Run(":8080")
}
