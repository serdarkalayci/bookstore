// Package mappers contains the funtions that maps DAO objects to domain objects and visa versa.
package mappers

import (
	"github.com/serdarkalayci/bookstore/info/adapters/data/mongodb/dao"
	"github.com/serdarkalayci/bookstore/info/domain"
)

// MapBookInfoDAO2BookInfo maps dao bookInfo to domain bookInfo
func MapBookInfoDAO2BookInfo(pd dao.BookInfoDAO) domain.BookInfo {
	return domain.BookInfo{
		ISBN: 	  pd.ISBN,
		Title:        pd.Title,
		Author: 	pd.Author,
		Price:       pd.Price,
		PublishDate: pd.PublishDate,
	}
}

// MapBookInfoDAOs2BookInfos maps dao bookInfo slice to domain bookInfo slice
func MapBookInfoDAOs2BookInfos(pds []dao.BookInfoDAO) []domain.BookInfo {
	var ps []domain.BookInfo
	for _, pd := range pds {
		ps = append(ps, MapBookInfoDAO2BookInfo(pd))
	}
	return ps
}

// MapBookInfo2BookInfoDAO maps domain bookInfo to dao bookInfo
func MapBookInfo2BookInfoDAO(p domain.BookInfo) dao.BookInfoDAO {
	return dao.BookInfoDAO{
		ISBN: 	  p.ISBN,
		Title:        p.Title,
		Author: 	p.Author,
		Price:       p.Price,
		PublishDate: p.PublishDate,
	}
}

