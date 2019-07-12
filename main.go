package main

import "github.com/gin-gonic/gin"

func main() {

	g := gin.New()
	g.Static("/", "./WebBuild")
	// g.StaticFile("/", "./WebBuild/index.html")
	g.Run(":8080")
}
