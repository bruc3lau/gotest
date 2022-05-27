package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/weapon", getWeapon)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getWeapon(c *gin.Context) {
	keys := c.Keys
	a := keys["type"]
	fmt.Println(a)
}
