package books

import (
	"encoding/json"
	"time"
)

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

// struct bookResponse
type BookResponse struct {
	ID        uint        `json:"id"`
	Title     string      `json:"title"`
	Author    string      `json:"author"`
	Desc      string      `json:"desc"`
	Image     string      `json:"image"`
	Price     json.Number `json:"price"`
	Discound  json.Number `json:"discound"`
	Rating    json.Number `json:"rating"`
	CreatedAt time.Time   `json:"created_at"`
}
