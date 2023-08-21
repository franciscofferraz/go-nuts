package repositories

import (
	"database/sql"

	"github.com/franciscofferraz/go-struct/internal/api/models"
)

type UserRepositoryInterface interface {
	Create(*models.User) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := ur.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.Id)

	if err != nil {
		return err
	}
	return nil
}
