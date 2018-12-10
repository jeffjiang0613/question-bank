package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jeffjiang0613/question-bank/routes"
)

func main() {
	app := gin.Default()
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.Use(cors.Default())
	app.Use()
	routes.Route(app)
	app.Run(":8081")
}
