package repositories

import (
	"errors"
	"final-project-backend/models"

	"gorm.io/gorm"
)

type RecycleCenterRepository interface {
	CreateRecycleCenter(recycleCenter models.RecyclingCenter) error
	GetAllRecycleCenter() ([]models.RecyclingCenter, error)
	GetRecycleCenterById(id string) (models.RecyclingCenter, error)
	UpdateRecycleCenter(id string, recycleCenter models.RecyclingCenter) error
	DeleteRecycleCenterById(id string) error
}

type recycleCenterRepo struct {
	db *gorm.DB
}

func NewRecycleCenterRepository(db *gorm.DB) RecycleCenterRepository {
	return &recycleCenterRepo{
		db: db,
	}
}

func (rc *recycleCenterRepo) CreateRecycleCenter(recycleCenter models.RecyclingCenter) error {
	err := rc.db.Create(&recycleCenter).Error
	if err != nil {
		return err
	}
	return nil
}

func (rc *recycleCenterRepo) GetAllRecycleCenter() ([]models.RecyclingCenter, error) {
	var recycleCenter []models.RecyclingCenter
	result := rc.db.Find(&recycleCenter)
	return recycleCenter, result.Error
}

func (rc *recycleCenterRepo) GetRecycleCenterById(id string) (models.RecyclingCenter, error) {
	var recycleCenter models.RecyclingCenter
	result := rc.db.First(&recycleCenter, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return recycleCenter, gorm.ErrRecordNotFound
		}
		return recycleCenter, result.Error
	}
	return recycleCenter, nil

}
func (rc *recycleCenterRepo) UpdateRecycleCenter(id string, recycleCenter models.RecyclingCenter) error {
	var exsistingRecycleCenter models.RecyclingCenter

	err := rc.db.First(&exsistingRecycleCenter, id).Error
	if err != nil {
		return err
	}

	exsistingRecycleCenter.Name = recycleCenter.Name
	exsistingRecycleCenter.Addres = recycleCenter.Addres
	exsistingRecycleCenter.ContactNumber = recycleCenter.ContactNumber

	err = rc.db.Save(&exsistingRecycleCenter).Error
	if err != nil {
		return err
	}

	return nil
}

func (rc *recycleCenterRepo) DeleteRecycleCenterById(id string) error {
	recycleCenter, err := rc.GetRecycleCenterById(id)

	if err != nil {
		return err
	}

	err = rc.db.Delete(recycleCenter).Error
	if err != nil {
		return err
	}

	return nil
}
