package main

import (
	"haki/blog/internals/admin"
	"haki/blog/internals/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	err := config.Loadenv()
	if err != nil {
		log.Panic(err)
	}
	err = config.Conndb()

	if err != nil {
		log.Panic(err)
	}

	adminr := app.Group("/admin")
	{
		adminr.POST("/blog/create", admin.CreateBlogPostHandler)
	}
	app.Run(":4001")
}
