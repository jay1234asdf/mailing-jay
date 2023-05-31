package router

import (
	"net/http"
	"os"

	"github.com/EmeraldLS/MailService/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.Use(cors.Default())
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.POST("/sendmail", controller.SendMail)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"response": "success",
			"message":  "This is the homepage",
		})
	})
	r.Run("0.0.0.0:" + port)
}
