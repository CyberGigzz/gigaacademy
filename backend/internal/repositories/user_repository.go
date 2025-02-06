package repositories

import (
	// "context"
	"database/sql"
	// "errors"
	// "github.com/CyberGigzz/gigaacademy/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
// 	query := "INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)"
// 	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
// 	query := "SELECT id, name, email, password FROM users WHERE id = ?"
// 	row := r.db.QueryRowContext(ctx, query, id)

// 	var user models.User
// 	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
// 	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
// 	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
// 	query := "DELETE FROM users WHERE id = ?"
// 	_, err := r.db.ExecContext(ctx, query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
