package dto

// BookStockResponseDTO represents the struct of document type
type BookStockResponseDTO struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string `json:"isbn"`
	// Stock is the number of books in stock.
	Stock        int `json:"stock"`
}

// BookStockRequestDTO represents the struct of document type to be stored in the data source
type BookStockRequestDTO struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string `json:"isbn" validate:"required"`
	// Stock is the number of books in stock.
	Stock        int `json:"stock" validate:"required"`
}