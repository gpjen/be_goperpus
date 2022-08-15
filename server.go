package main

import (
	"be_goperpus/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "server oke",
		})
	})

	V1 := router.Group("/v1")
	V1.GET("/books", handler.GetBooks)
	V1.GET("/book/:id", handler.GetBookById)
	V1.POST("/book", handler.NewBook)
	V1.PATCH("/book/:id", handler.UpdateBook)
	V1.DELETE("/book/:id", handler.DeleteBook)

	router.Run(":8888")
}
