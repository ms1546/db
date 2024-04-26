package repository

import (
	"database/sql"

	"pq/internal/app/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2)`
	_, err := r.db.Exec(query, user.Name, user.Email)
	return err
}

func (r *UserRepository) GetUser(id int) (model.User, error) {
	var user model.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	query := `SELECT id, name, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(user model.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
