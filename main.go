package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.LoadHTMLFiles("./web/login.html")

	//r.Static("/sdk", "./web/sdk")
	r.Static("/web", "./web")
	r.Static("/pages", "./web/pages")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "")
	})
	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "",
			"data":   "",
		})
	})
	r.Run(":8099")
}
