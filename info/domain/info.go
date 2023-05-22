// Package domain describes the domain model of the document management system.
package domain

import "time"

// BookInfo is the domain model of a book.
type BookInfo struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string
	// Title is the title of the book.
	Title        string
	// Author is the author of the book.
	Author      string
	// Price is the price of the book.
	Price       float64
	// PublishDate is the date when the book was published.
	PublishDate time.Time
}



