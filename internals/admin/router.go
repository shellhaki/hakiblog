package admin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBlogPostHandler(c *gin.Context) {
	var req CreateBlogPostRequestType
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	if err := CreateBlogPost(req.Title, req.Body, req.Tags, req.Cover); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "An error occured while processing request",
		})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Blog post created successfully",
	})
}
