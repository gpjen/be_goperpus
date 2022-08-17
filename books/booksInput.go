package books

import (
	"time"
)

// struct bookRequest
type BookRequest struct {
	Title    string `json:"title" binding:"required,min=2,max=100"`
	Author   string `json:"author" binding:"required,min=2,max=50"`
	Desc     string `json:"desc" binding:"required"`
	Image    string `json:"image"`
	Price    uint   `json:"price" binding:"required,min=1"`
	Discound uint   `json:"discound" binding:"min=0"`
}

// struct bookResponse
type BookResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Desc      string    `json:"desc"`
	Image     string    `json:"image"`
	Price     uint      `json:"price"`
	Discound  uint      `json:"discound"`
	Rating    uint      `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
