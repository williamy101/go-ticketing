package repository

import (
	"go-ticketing/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Register(user *entity.Users) error
	GetUserByEmail(email string) (*entity.Users, error)
	GetUserByID(userID int) (*entity.Users, error)
	GetAllUsers() ([]entity.Users, error)
	UpdateUserRole(userID int, role string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Register(user *entity.Users) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*entity.Users, error) {
	var user entity.Users
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) GetUserByID(userID int) (*entity.Users, error) {
	var user entity.Users
	err := r.db.First(&user, userID).Error
	return &user, err
}

func (r *userRepository) GetAllUsers() ([]entity.Users, error) {
	var users []entity.Users
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserRole(userID int, role string) error {
	return r.db.Model(&entity.Users{}).Where("user_id = ?", userID).Update("role", role).Error
}
