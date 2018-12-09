package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/routes"
)

func main() {
	app := gin.Default()
	routes.Route(app)
	app.Run(":8081")
}
