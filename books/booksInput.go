package books

import "encoding/json"

// struct bookRequest
type BookRequest struct {
	Title    string      `json:"title" binding:"required"`
	Author   string      `json:"author" binding:"required"`
	Desc     string      `json:"desc" binding:"required"`
	Image    string      `json:"image"`
	Price    json.Number `json:"price" binding:"required"`
	Discound json.Number `json:"discound"`
	Rating   json.Number `json:"rating"`
}
