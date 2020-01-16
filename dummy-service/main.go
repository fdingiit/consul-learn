package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func health(c *gin.Context) {
	c.JSON(200, struct{}{})
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("should have 2 args")
		os.Exit(1)
	}

	host := os.Args[1]

	r := gin.Default()

	r.GET("/health", health)

	if err := http.ListenAndServe(host, r); err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}
