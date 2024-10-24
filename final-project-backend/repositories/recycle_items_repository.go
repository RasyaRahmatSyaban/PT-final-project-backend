package repositories

import (
	"errors"
	"final-project-backend/models"

	"gorm.io/gorm"
)

type RecycleItemRepository interface {
	CreateRecycleItem(recycleItem models.RecyclableItems) error
	GetAllRecycleItem() ([]models.RecyclableItems, error)
	GetRecycleItemById(id string) (models.RecyclableItems, error)
	GetRecycleItemByName(name string) (models.RecyclableItems, error)
	Update(id string, recycleItem models.RecyclableItems) error
	DeleteRecycleItemById(id string) error
	DeleteRecycleItemByName(name string) error
}

type recycleItemRepo struct {
	db *gorm.DB
}

func NewRecycleItemRepository(db *gorm.DB) RecycleItemRepository {
	return &recycleItemRepo{
		db: db,
	}
}

func (ri *recycleItemRepo) CreateRecycleItem(recycleItem models.RecyclableItems) error {
	err := ri.db.Create(&recycleItem).Error
	if err != nil {
		return err
	}

	return nil

}

// GetAll implements UserRepository.
func (ri *recycleItemRepo) GetAllRecycleItem() ([]models.RecyclableItems, error) {
	var recycleItems []models.RecyclableItems
	result := ri.db.Find(&recycleItems)
	return recycleItems, result.Error

}

// GetRecycleItemByName implements UserRepository.
func (ri *recycleItemRepo) GetRecycleItemByName(name string) (models.RecyclableItems, error) {
	var recycleItems models.RecyclableItems

	result := ri.db.Debug().Where("name = ?", name).First(&recycleItems)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return recycleItems, gorm.ErrRecordNotFound
		}
		return recycleItems, result.Error
	}
	return recycleItems, nil
}

// GetRecycleItemById implements UserRepository.
func (ri *recycleItemRepo) GetRecycleItemById(id string) (models.RecyclableItems, error) {
	var recycleItems models.RecyclableItems
	result := ri.db.First(&recycleItems, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return recycleItems, gorm.ErrRecordNotFound
		}
		return recycleItems, result.Error
	}
	return recycleItems, nil
}

func (ri *recycleItemRepo) Update(id string, recycleItem models.RecyclableItems) error {
	var exsistingRecycleItem models.RecyclableItems

	err := ri.db.First(&exsistingRecycleItem, id).Error
	if err != nil {
		return err
	}

	exsistingRecycleItem.Name = recycleItem.Name
	exsistingRecycleItem.PricePerKg = recycleItem.PricePerKg
	exsistingRecycleItem.Description = recycleItem.Description

	err = ri.db.Save(&exsistingRecycleItem).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecycleItemById implements UserRepository.
func (ri *recycleItemRepo) DeleteRecycleItemById(id string) error {
	recycleItem, err := ri.GetRecycleItemById(id)

	if err != nil {
		return err
	}

	err = ri.db.Delete(&recycleItem).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecycleItemByName implements UserRepository.
func (ri *recycleItemRepo) DeleteRecycleItemByName(name string) error {
	recycleItem, err := ri.GetRecycleItemByName(name)

	if err != nil {
		return err
	}

	err = ri.db.Delete(&recycleItem).Error
	if err != nil {
		return err
	}

	return nil
}
