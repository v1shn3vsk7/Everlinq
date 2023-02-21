package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/v1shn3vsk7/Everlinq/internal/models"
	"net/http"
)

func Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", models.Page{
			Title: "Log in to your account",
			Url:   "",
		})
	}
}
