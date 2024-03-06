package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/msg2324/go/CRUD_Backend/services"
)

func NewController(port string) {
	r := gin.New()

	r.Use(gin.Logger())

	r.GET("/get-feed", func(c *gin.Context) {
		services.GetFeed(c)
	})

	r.POST("/post-feed", func(c *gin.Context) {
		
	})
	err := r.Run(port)
	if err != nil {
		panic(err.Error())

	}
}
