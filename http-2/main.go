package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	g := gin.New()
	g.StaticFile("/static", "main.go")
	g.GET("/", func(c *gin.Context) {
		p := c.Writer.Pusher()
		if p == nil {
			fmt.Println("Nil pusher")

		} else {
			otps := new(http.PushOptions)
			otps.Method = "GET"
			otps.Header.Set("push-custom-header", "roshan patil")
			err := p.Push("main.go", otps)
			if err != nil {
				fmt.Println(err)
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})
	g.Run(":8080")
}
