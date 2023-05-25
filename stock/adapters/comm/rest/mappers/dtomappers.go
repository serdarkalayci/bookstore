package mappers

import (
	"github.com/serdarkalayci/bookstore/stock/adapters/comm/rest/dto"
	"github.com/serdarkalayci/bookstore/stock/domain"
)

// MapBookStock2BookStockResponseDTO maps BookStock to BookStockDTO
func MapBookStock2BookStockResponseDTO(b domain.BookStock) dto.BookStockResponseDTO {
	return dto.BookStockResponseDTO{
		ISBN:        b.ISBN,
		Stock: 	 	b.Stock,
	}
}

// MapBookStocks2BookStockDTOs maps BookStockDAOs to BookStocks
func MapBookStocks2BookStockDTOs(pds []domain.BookStock) []dto.BookStockResponseDTO {
	var ps []dto.BookStockResponseDTO
	for _, pd := range pds {
		ps = append(ps, MapBookStock2BookStockResponseDTO(pd))
	}
	return ps
}