package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/v1shn3vsk7/Everlinq/internal/models"
	"net/http"
)

func Index() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", models.Page{
			Title: "Main page",
			Url:   "",
		})
	}
}
