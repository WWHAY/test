package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api", Hello)
	router.Run(":8080")
}

//Hello .
func Hello(c *gin.Context) {
	fmt.Println("hello world！")
}
