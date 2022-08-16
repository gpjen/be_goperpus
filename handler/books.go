package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// struct bookRequest
type bookRequest struct {
	Title    string      `json:"title" binding:"required"`
	Author   string      `json:"author" binding:"required"`
	Desc     string      `json:"desc" binding:"required"`
	Image    string      `json:"image"`
	Price    json.Number `json:"price" binding:"required"`
	Discound json.Number `json:"discound"`
}

// create
func NewBook(c *gin.Context) {

	var bookInput bookRequest

	err := c.ShouldBindJSON(&bookInput)
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

	data := bookRequest{
		Title:    bookInput.Title,
		Author:   bookInput.Author,
		Desc:     bookInput.Desc,
		Image:    bookInput.Image,
		Price:    bookInput.Price,
		Discound: bookInput.Discound,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "create new book",
		"data":    data,
	})
}

// read
func GetBooks(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get books",
		"data":    "not handle",
	})
}

// read by id
func GetBookById(c *gin.Context) {
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
func UpdateBook(c *gin.Context) {
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
func DeleteBook(c *gin.Context) {
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
