package handler

import (
	"be_goperpus/books"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookServices books.Services
}

func NewBookHandler(bookServices books.Services) *bookHandler {
	return &bookHandler{bookServices}
}

// data response
func convertToBookResponse(getBook books.Books) books.BookResponse {

	return books.BookResponse{
		ID:        getBook.ID,
		Title:     getBook.Title,
		Author:    getBook.Author,
		Desc:      getBook.Desc,
		Image:     getBook.Image,
		Price:     getBook.Price,
		Discound:  getBook.Discound,
		Rating:    getBook.Rating,
		CreatedAt: getBook.CreatedAt,
	}
}

// create
func (h *bookHandler) NewBook(c *gin.Context) {

	var bookRequest books.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		// var errFields []string
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
		// 	errFields = append(errFields, errMessage)
		// }
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	getData, err := h.bookServices.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	data := convertToBookResponse(getData)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "create new book",
		"data":    data,
	})
}

// read
func (h *bookHandler) GetBooks(c *gin.Context) {

	getData, err := h.bookServices.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := []books.BookResponse{}

	for _, d := range getData {
		dt := convertToBookResponse(d)
		data = append(data, dt)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get books",
		"data":    data,
	})
}

// read by id
func (h *bookHandler) GetBookById(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("id %s must be number", idString),
		})
		return
	}

	getData, err := h.bookServices.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := convertToBookResponse(getData)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("get book by id %d", id),
		"data":    data,
	})

}

// update book
func (h *bookHandler) UpdateBook(c *gin.Context) {

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("id %s must be number", idString),
		})
		return
	}

	var bookRequest books.BookRequest

	err = c.ShouldBindJSON(&bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	getData, err := h.bookServices.Update(bookRequest, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	data := convertToBookResponse(getData)

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf(" update book id %d", id),
		"data":    data,
	})
}

// delete book
func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": fmt.Sprintf("id %s must be number", idString),
		})
		return
	}

	data, err := h.bookServices.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf(" Delete book id %d", id),
		"data":    data,
	})
}
