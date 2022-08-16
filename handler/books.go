package handler

import (
	"be_goperpus/books"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookServices books.Serfices
}

func NewBookHandler(bookServices books.Serfices) *bookHandler {
	return &bookHandler{bookServices}
}

type bookResponse struct {
	Title    string      `json:"title"`
	Author   string      `json:"author"`
	Desc     string      `json:"desc"`
	Image    string      `json:"image"`
	Price    json.Number `json:"price"`
	Discound json.Number `json:"discound"`
}

// create
func (h *bookHandler) NewBook(c *gin.Context) {

	var bookRequest books.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		var errFields []string
		for _, e := range err.(validator.ValidationErrors) {
			errMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errFields = append(errFields, errMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  errFields,
		})
		return
	}

	data, err := h.bookServices.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "create new book",
		"data":    data,
	})
}

// read
func (h *bookHandler) GetBooks(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get books",
		"data":    "not handle",
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
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("get book by id %d", id),
		"data":    "not handle",
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

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf(" update book id %d", id),
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

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf(" Delete book id %d", id),
	})
}
