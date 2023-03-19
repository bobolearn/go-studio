package main

import (
	database "go-studio/db"
	route "go-studio/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	route.QuestionRouter(r, db)
	r.Run(":3000")
}
