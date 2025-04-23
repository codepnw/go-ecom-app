package repository

import (
	"go-ecommerce-app/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUser(email string) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	if err := r.db.Create(&u).Error; err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r userRepository) FindUser(email string) (domain.User, error) {
	var user domain.User

	if err := r.db.First(&user, "email=?", email).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) FindUserByID(id uint) (domain.User, error) {
	var user domain.User

	if err := r.db.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(&u).Error
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
