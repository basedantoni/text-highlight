package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	e := gin.Default()
	e.LoadHTMLFiles("templates/index.html", "templates/habit.html")

	e.GET("/", func(c *gin.Context) {
		habits := ReadHabits()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"habits": habits,
		})
	})

	e.POST("/habits", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := CreateHabit(title, status)

		c.HTML(http.StatusOK, "habit.html", gin.H{
			"Title": title,
			"Status": status,
			"Id": id,
		})
	})

	e.DELETE("/habits/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteHabit(id)

		c.HTML(http.StatusOK, "habit.html", gin.H{})
	})

	e.Run(":8080")
}