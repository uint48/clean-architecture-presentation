package repository

import "myapp/internal/entity/user"

type UserRepository interface {
	FindByID(id string) (*user.User, error)
	Save(u *user.User) error
	Delete(id string) error
	Update(u *user.User) error
	Get(username string) (*user.User, error)
}
