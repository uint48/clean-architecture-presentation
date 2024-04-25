package mysql

import (
	"database/sql"
	"errors"
	"myapp/internal/entity/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) FindByID(id string) (*user.User, error) {
	var u user.User
	query := "SELECT id, username, password, email, is_active, role, balance FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.IsActive, &u.Role, &u.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Save(u *user.User) error {
	query := "INSERT INTO users (id, username, password, email, is_active, role, balance) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := r.db.Exec(query, u.ID, u.Username, u.Password, u.Email, u.IsActive, u.Role)
	return err
}

func (r *UserRepository) Delete(id string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UserRepository) Update(u *user.User) error {
	query := "UPDATE users SET username = ?, password = ?, email = ?, is_active = ?, role = ?, balance = ? WHERE id = ?"
	_, err := r.db.Exec(query, u.Username, u.Password, u.Email, u.IsActive, u.Role, u.Balance, u.ID)
	return err
}

func (r *UserRepository) Get(username string) (*user.User, error) {
	var u user.User
	query := "SELECT id, username, password, email, is_active, role, balance FROM users WHERE username = ?"
	err := r.db.QueryRow(query, username).Scan(&u.ID, &u.Username, &u.Password, &u.Email, &u.IsActive, &u.Role, &u.Balance)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}
