package services

import (
	"github.com/gin-gonic/gin"
	"github.com/msg2324/go/CRUD_Backend/repositories"
	"net/http"
)

func GetFeed(c *gin.Context) {
	feed := repositories.DB
	c.JSON(http.StatusOK, feed)
}
