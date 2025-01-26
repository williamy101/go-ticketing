package service

import (
	"errors"
	"go-ticketing/entity"
	"go-ticketing/repository"
	"go-ticketing/util"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user *entity.Users) error
	Login(email, password string) (string, error)
	GetUserByID(userID int) (*entity.Users, error)
	GetAllUsers() ([]entity.Users, error)
	UpdateUserRole(userID int, role string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(user *entity.Users) error {
	user.Role = "Admin"

	if user.UserName == "" || user.Email == "" || user.Password == "" {
		return errors.New("username, email, and password are required")
	}

	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	existingUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Register(user)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}
	token, err := util.GenerateToken(user.UserID, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) GetUserByID(userID int) (*entity.Users, error) {
	return s.userRepo.GetUserByID(userID)
}

func (s *userService) GetAllUsers() ([]entity.Users, error) {
	return s.userRepo.GetAllUsers()
}

func (s *userService) UpdateUserRole(userID int, role string) error {
	if role != "Admin" && role != "User" {
		return errors.New("invalid role")
	}
	_, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}
	return s.userRepo.UpdateUserRole(userID, role)
}
