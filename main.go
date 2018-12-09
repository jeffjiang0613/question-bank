package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	route(app)
	app.Run(":8081")
}
