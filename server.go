package main

import (
	"be_goperpus/books"
	"be_goperpus/config"
	"be_goperpus/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnDb()
	defer config.CloseDb(db)

	bookRepositoriy := books.NewRepository(db)
	bookServices := books.NewServices(bookRepositoriy)
	bookHandler := handler.NewBookHandler(bookServices)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "server oke",
		})
	})

	V1 := router.Group("/v1")
	V1.GET("/books", bookHandler.GetBooks)
	V1.GET("/book/:id", bookHandler.GetBookById)
	V1.POST("/book", bookHandler.NewBook)
	V1.PATCH("/book/:id", bookHandler.UpdateBook)
	V1.DELETE("/book/:id", bookHandler.DeleteBook)

	router.Run(":8888")
}
