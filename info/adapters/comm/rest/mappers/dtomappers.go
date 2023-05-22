package mappers

import (
	"github.com/serdarkalayci/bookstore/info/adapters/comm/rest/dto"
	"github.com/serdarkalayci/bookstore/info/domain"
)

// MapBookInfoRequestDTO2BookInfo maps BookInfoDTO to BookInfo
func MapBookInfoRequestDTO2BookInfo(b dto.BookInfoRequestDTO) domain.BookInfo {
	return domain.BookInfo{
		ISBN:        b.ISBN,
		Title:       b.Title,
		Author: 	b.Author,
		Price:       b.Price,
		PublishDate: b.PublishDate,
	}
}

// MapBookInfo2BookInfoResponseDTO maps BookInfo to BookInfoDTO
func MapBookInfo2BookInfoResponseDTO(b domain.BookInfo) dto.BookInfoResponseDTO {
	return dto.BookInfoResponseDTO{
		ISBN:        b.ISBN,
		Title:       b.Title,
		Author: 	b.Author,
		Price:       b.Price,
		PublishDate: b.PublishDate,
	}
}

// MapBookInfos2BookInfoDTOs maps BookInfoDAOs to BookInfos
func MapBookInfos2BookInfoDTOs(pds []domain.BookInfo) []dto.BookInfoResponseDTO {
	var ps []dto.BookInfoResponseDTO
	for _, pd := range pds {
		ps = append(ps, MapBookInfo2BookInfoResponseDTO(pd))
	}
	return ps
}