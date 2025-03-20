package repositories

import "database/sql"

type Models struct {
	Users UserRepository
	// ProductRepository ProductRepository
	// OrderRepository   OrderRepository
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users: UserRepository{DB: db},
		// ProductRepository: productRepo,
		// OrderRepository:   orderRepo,
	}
}
