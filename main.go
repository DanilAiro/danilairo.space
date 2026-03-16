package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", handleMain)

	pr := router.Group("/projects", projectsHeader) // projectsHeader не совсем хедер, но просто грузится на всех страницах группы
	pr.GET("/", handleProjects)

	router.Run()
}

func handleMain(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func projectsHeader(c *gin.Context) {
	c.HTML(http.StatusOK, "projects-group.html", gin.H{})
}

func handleProjects(c *gin.Context) {
	c.HTML(http.StatusOK, "projects.html", gin.H{})
}
