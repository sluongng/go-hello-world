package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/someGet", func(c *gin.Context) {
		c.String(http.StatusOK, "<html><body><div>HELLO WORLD</div></body>")
	})

	router.POST("/somePost", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	router.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello users")
	})

	router.Run(":3000")
}
