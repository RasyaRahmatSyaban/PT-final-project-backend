package repositories

import (
	"errors"
	"final-project-backend/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transactions) error
	GetAllTransactions() ([]models.Transactions, error)
	GetTransactionById(id string) (models.Transactions, error)
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepo{
		db: db,
	}
}

func (t *transactionRepo) CreateTransaction(transaction models.Transactions) error {
	err := t.db.Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepo) GetAllTransactions() ([]models.Transactions, error) {
	var transaction []models.Transactions
	result := t.db.Find(&transaction)
	return transaction, result.Error
}

func (t *transactionRepo) GetTransactionById(id string) (models.Transactions, error) {
	var transaction models.Transactions
	result := t.db.First(&transaction, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return transaction, gorm.ErrRecordNotFound
		}
		return transaction, result.Error
	}
	return transaction, nil
}
