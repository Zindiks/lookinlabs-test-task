package service

import (
	"github.com/Zindiks/lookinlabs-test-task/model"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUsers() ([]model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUser(user *model.User) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.db.Create(user).Error
}

func (s *userService) GetUsers() ([]model.User, error) {
	var users []model.User
	err := s.db.Find(&users).Error
	return users, err
}

func (s *userService) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.db.Save(user).Error
}
