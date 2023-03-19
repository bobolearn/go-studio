package routes

import (
	controller "go-studio/controller"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QuestionRouter(r *gin.Engine, db *gorm.DB) {
	r.GET("/questions", func(c *gin.Context) {
		questions, err := controller.QueryAllQuestions(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"questions": questions})
	})
	r.GET("/question/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		question, err := controller.QueryOneQuestion(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"question": question})
	})
	r.POST("/questions", func(c *gin.Context) {
		var postRequest struct {
			UserID  uint   `json:"user_id"`
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.BindJSON(&postRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		question, err := controller.CreateQuestion(db, postRequest.UserID, postRequest.Title, postRequest.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, question)
	})
	r.PUT("/question/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		var postRequest struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}
		if err := c.BindJSON(&postRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := controller.UpdateQuestion(db, id, postRequest.Title, postRequest.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusCreated)
	})
	r.DELETE("/question/:id", func(c *gin.Context) {
		id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		_, err := controller.QueryOneQuestion(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = controller.DeleteQuestion(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	})
}
