package repositories

import (
	"errors"
	"final-project-backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) error
	GetAll() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	Update(id string, user models.User) error
	DeleteUserById(id string) error
	DeleteUserByEmail(email string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// Create New User
func (u *userRepo) CreateUser(user models.User) error {
	err := u.db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil

}

// GetAll implements UserRepository.
func (u *userRepo) GetAll() ([]models.User, error) {
	var users []models.User
	result := u.db.Find(&users)
	return users, result.Error

}

// GetUserByEmail implements UserRepository.
func (u *userRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	result := u.db.Debug().Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, gorm.ErrRecordNotFound
		}
		return user, result.Error
	}
	return user, nil
}

// GetUserById implements UserRepository.
func (u *userRepo) GetUserById(id string) (models.User, error) {
	var user models.User
	result := u.db.First(&user, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, gorm.ErrRecordNotFound
		}
		return user, result.Error
	}
	return user, nil
}

// Update user data
func (u *userRepo) Update(id string, user models.User) error {
	var existingUser models.User

	err := u.db.First(&existingUser, id).Error
	if err != nil {
		return err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Balance = user.Balance
	existingUser.UpdatedAt = user.UpdatedAt

	err = u.db.Save(&existingUser).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteUserById implements UserRepository.
func (u *userRepo) DeleteUserById(id string) error {
	user, err := u.GetUserById(id)

	if err != nil {
		return err
	}

	err = u.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil

}

// DeleteUserByEmail implements UserRepository.
func (u *userRepo) DeleteUserByEmail(email string) error {
	user, err := u.GetUserByEmail(email)

	if err != nil {
		return err
	}

	err = u.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil

}
