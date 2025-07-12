package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TitleUpdate struct {
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: 1, Title: "Освоить HTTP", Done: false},
	{ID: 2, Title: "Написать самому код. Понять", Done: false},
	{ID: 3, Title: "Показать Кириллу", Done: true},
}

func main() {
	r := gin.Default()

	r.POST("/todos", func(c *gin.Context) {
		var NewTask Todo
		if err := c.BindJSON(&NewTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": "non-correct JSON"})
			return
		}
		NewID := len(todos) + 1
		NewTask.ID = NewID
		todos = append(todos, NewTask)
		c.JSON(http.StatusCreated, NewTask)
	})
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	r.Run()
}
