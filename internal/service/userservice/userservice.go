package userservice

import (
	"errors"
	"myapp/internal/entity/user"
	"myapp/internal/repository"
)

type UserService interface {
	Register(user *user.User) error
	Activate(userID string) error
	Login(username, password string) (*user.User, error)
	CheckBalance(userID string) (float64, error)
}

type Service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}

func (s *Service) Register(user *user.User) error {
	// Validate the user data
	if err := validateUser(user); err != nil {
		return err
	}

	// saving the user
	return s.userRepository.Save(user)
}

func (s *Service) Activate(userID string) error {
	u, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}
	u.IsActive = true

	return s.userRepository.Update(u)
}
func (s *Service) Login(username, password string) (*user.User, error) {
	return s.userRepository.Get(username)
}

func (s *Service) CheckBalance(userID string) (float64, error) {
	u, err := s.userRepository.Get(userID)
	if err != nil {
		return 0, err
	}
	return u.Balance, nil
}

func validateUser(u *user.User) error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}
	// Add more validation rules as needed

	return nil
}
