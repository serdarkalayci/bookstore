package dto

import "time"

// BookInfoResponseDTO represents the struct of document type
type BookInfoResponseDTO struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string `json:"isbn"`
	// Title is the title of the book.
	Title        string `json:"title"`
	// Author is the author of the book.
	Author      string `json:"author"`
	// Price is the price of the book.
	Price       float64 `json:"price"`
	// PublishDate is the date when the book was published.
	PublishDate time.Time `json:"publishdate"`
	// Stock is the stock of the book.
	Stock       int `json:"stock"`
}

// BookInfoListDTO represents the struct of document type which is stripped down a few fields 
type BookInfoListDTO struct {
		// ISBN is the unique identifier of the book.
		ISBN 	  string `json:"isbn"`
		// Title is the title of the book.
		Title        string `json:"title"`
		// Author is the author of the book.
		Author      string `json:"author"`
}

// BookInfoRequestDTO represents the struct of document type to be stored in the data source
type BookInfoRequestDTO struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string `json:"isbn" validate:"required"`
	// Title is the title of the book.
	Title        string `json:"title" validate:"required"`
	// Author is the author of the book.
	Author      string `json:"author" validate:"required"`
	// Price is the price of the book.
	Price       float64 `json:"price" validate:"required"`
	// PublishDate is the date when the book was published.
	PublishDate time.Time `json:"publishdate" validate:"required"`
}