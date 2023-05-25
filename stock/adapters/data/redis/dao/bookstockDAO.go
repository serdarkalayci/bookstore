// Package dao contains the data access objects for the bookStock management system.
package dao

// BookStockDAO represents the struct of document type to be stored in mongoDB
type BookStockDAO struct {
	// ISBN is the unique identifier of the book.
	ISBN 	  string `bson:"isbn"`
	// Stock is the number of books in stock.
	Stock 	  int `bson:"stock"`
}


