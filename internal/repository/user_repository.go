package repository

import (
	"database/sql"

	"github.com/greeflas/go_di_example/internal/model"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepository {
	return &UserRepository{conn: conn}
}

func (r *UserRepository) GetById(id int) (*model.User, error) {
	row := r.conn.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	user := new(model.User)

	if err := row.Scan(&user.Id, &user.Name); err != nil {
		return nil, err
	}

	return user, nil
}
