package router

import (
	"net/http"
	"strconv"

	"github.com/ernestngugi/sil-devops/internal/controller"
	"github.com/ernestngugi/sil-devops/internal/db"
	"github.com/ernestngugi/sil-devops/internal/form"
	"github.com/ernestngugi/sil-devops/internal/repos"
	"github.com/gin-gonic/gin"
)

type AppRouter struct {
	*gin.Engine
}

func BuildRouter(
	dB db.DB,
) *AppRouter {

	router := gin.Default()
	appRouter := router.Group("/v1")

	blogRepository := repos.NewBlogRepository()

	blogController := controller.NewBlogRepository(blogRepository)

	appRouter.POST("/blog", createBlog(dB, blogController))
	appRouter.GET("/blogs/:id", getBlogByID(dB, blogController))
	appRouter.GET("/blogs", listBlogs(dB, blogController))
	appRouter.PUT("/blogs/:id", updateBlog(dB, blogController))
	appRouter.DELETE("/blogs/:id", deleteBlog(dB, blogController))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error_message": "Endpoint not found"})
	})

	return &AppRouter{router}
}

func createBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		var form form.Blog

		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false})
			return
		}

		blog, err := blogController.CreateBlog(c.Request.Context(), dB, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, blog)
	}
}

func updateBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		var form form.UpdateBlog

		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false})
			return
		}

		blogID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		blog, err := blogController.UpdateBlog(c.Request.Context(), dB, blogID, &form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, blog)
	}
}

func getBlogByID(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		blogID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		blog, err := blogController.BlogByID(c.Request.Context(), dB, blogID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, blog)
	}
}

func listBlogs(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		blogs, err := blogController.AllBlogs(c.Request.Context(), dB)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, blogs)
	}
}

func deleteBlog(
	dB db.DB,
	blogController controller.BlogController,
) func(c *gin.Context) {
	return func(c *gin.Context) {

		blogID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		err = blogController.DeleteBlog(c.Request.Context(), dB, blogID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true})
	}
}
