package controllers

import (
	ArticleService "blog/internal/modules/article/services"
	"blog/pkg/html"
	"strconv"

	// "blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
} 
	
// func (controller *Controller) Show(c *gin.Context) {
// 	// html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
// 	// 	"title":    "Home page",
// 	// 	"featured": controller.articleService.GetFeaturedArticles(),
// 	// 	"stories":  controller.articleService.GetStoriesArticles(),
// 	// })
// 	// get the article 
// 	id, err := strconv.Atoi(c.Param("id"));
 
// 	if err != nil{
// 		xz
 
// 		return
//  	}

// 	article, err := controller.articleService.Find(id);
// 	if err != nil{
// 		html.Render(c,http.StatusNotFound,"templates/errors/html/404", gin.H{
// 			"message":"Error getting article with id: " + c.Param("id"),
// 			"error": err.Error()} )
// 		return
//  	}

//  	c.JSON(http.StatusOK, gin.H{
// 		"article": article,
// 	})


// }

func (controller *Controller) Show(c *gin.Context) {
	// Get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error converting the id"})
		return
	}

	// Find the article from the database
	article, err := controller.articleService.Find(id)

	// If the article is not found, show error page
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// if article found, render article template
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})
}
