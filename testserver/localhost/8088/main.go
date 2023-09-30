package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/health")
	})
	router.Run(":8088")
}
