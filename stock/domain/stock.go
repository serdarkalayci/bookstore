// Package domain describes the domain model of the document management system.
package domain

// BookStock is the domain model of a book.
type BookStock struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string
	// Stock is the number of books in stock.
	Stock 	  int
}



