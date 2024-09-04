package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home is a handler for the home page

func NotFound404(c *gin.Context){
	c.HTML(http.StatusNotFound, "404.html", gin.H{
		"title": "Page Not Found",
	})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{
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

func Login(c *gin.Context){
	c.HTML(http.StatusOK, "loginForm.html", gin.H{})
}

func SignUp(c *gin.Context){
	c.HTML(http.StatusOK, "signupForm.html", gin.H{})
}
