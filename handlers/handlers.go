package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home is a handler for the home page
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title":   "Welcome to Gin",
		"Message": "This is the home page. RAJVEERRRR",
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"Title": "Welcome to Gin ABOUT PAGE RAJVEER",
		"Para":  "this is a paragrtaph from the about page :)))))",
	})
}
