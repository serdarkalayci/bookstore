// Package mappers contains the funtions that maps DAO objects to domain objects and visa versa.
package mappers

import (
	"github.com/serdarkalayci/bookstore/stock/domain"
)

// MapBookStockDAO2BookStock maps dao bookStock to domain bookStock
func MapBookStockDAO2BookStock(ISBN string, stock int) domain.BookStock {
	return domain.BookStock{
		ISBN: 	  ISBN,
		Stock: 	  stock,
	}
}



