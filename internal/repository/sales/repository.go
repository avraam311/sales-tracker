package sales

import (
	"errors"

	"github.com/wb-go/wbf/dbpg"
)

var (
	ErrSaleNotFound = errors.New("sale not found")
)

type Repository struct {
	db *dbpg.DB
}

func NewRepository(db *dbpg.DB) *Repository {
	return &Repository{
		db: db,
	}
}
