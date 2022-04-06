package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func findProject(c *gin.Context) {
	id := c.Param("id")

	for _, a := range projects {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Project not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/projects", getProjects)
	router.GET("/projects/:id", findProject)

	router.POST("/users", createUser)

	router.Run("localhost:8080")
}
