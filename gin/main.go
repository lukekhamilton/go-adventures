package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Logger HandlerFunc")
		t := time.Now()
		c.Set("example", "12346")
		c.Next()
		latency := time.Since(t)
		log.Println(latency)
		status := c.Writer.Status()
		log.Println(status)
		log.Println("Logger HandlerFunc END")
	}
}

func main() {

	r := gin.Default()
	r.Use(Logger())

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
