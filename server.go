package main

import (
	"be_goperpus/config"
	"be_goperpus/handler"
	books "be_goperpus/repositories/Books"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db := config.ConnDb()

	datarequest := books.NewRepository(db)

	dataJson := books.Books{
		Title:    "buku baru saja di tambah",
		Author:   "sarada uzumaki",
		Desc:     "buku perjalanan hidup petapa genit",
		Image:    "saradaimuet.jpg",
		Price:    430000,
		Discound: 10,
	}

	err := datarequest.Create(dataJson)

	fmt.Println(err)

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
