package dao

import (
	"time"
)

// BookInfoDAO represents the struct of document type to be stored in mongoDB
type BookInfoDAO struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string `bson:"isbn"`
	// Title is the title of the book.
	Title        string `bson:"title"`
	// Author is the author of the book.
	Author      string `bson:"author"`
	// Price is the price of the book.
	Price       float64 `bson:"price"`
	// PublishDate is the date when the book was published.
	PublishDate time.Time `bson:"publishDate"`
}


