package main

import (
	"be_goperpus/books"
	"be_goperpus/config"
	"be_goperpus/handler"

	"github.com/gin-gonic/gin"
)

var (
	db = config.ConnDb()

	//books
	bookRepositoriy = books.NewRepository(db)
	bookServices    = books.NewServices(bookRepositoriy)

	//handler
	bookHandler = handler.NewBookHandler(bookServices)
)

func main() {
	defer config.CloseDb(db)

	router := gin.Default()

	V1 := router.Group("/v1")
	V1.GET("/books", bookHandler.GetBooks)
	V1.GET("/book/:id", bookHandler.GetBookById)
	V1.POST("/book", bookHandler.NewBook)
	V1.PATCH("/book/:id", bookHandler.UpdateBook)
	V1.DELETE("/book/:id", bookHandler.DeleteBook)

	router.Run(":8888")
}
