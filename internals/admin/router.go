package admin

import (
	"log"
	"net/http"
	"strconv"

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

func DeleteBlogPostHandler(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "an error occured ",
		})
		log.Panic(err)
		return
	}
	if err = DeleteBlogPost(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "an error occured ",
		})
		log.Panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted blog post",
	})

}
