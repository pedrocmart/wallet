package initdb

import (
	"github.com/jinzhu/gorm"
	"github.com/pedrocmart/wallet/internal/models"
	"github.com/shopspring/decimal"
)

func InitDB() {
	service, err := models.NewDBService()
	if err != nil {
		panic(err)
	}
	defer service.Close()
	service.AutoMigrate() // Initialize with wallet table

	b1, _ := decimal.NewFromString("0.01")
	w1 := &models.Wallet{
		gorm.Model{},
		"123",
		b1,
	}

	b2, _ := decimal.NewFromString("12.30")
	w2 := &models.Wallet{
		gorm.Model{},
		"456",
		b2,
	}

	// Delete wallets
	service.Wallet.Delete(w1)
	service.Wallet.Delete(w2)
	// Create
	err = service.Wallet.Create(w1)
	if err != nil {
		panic(err)
	}

	// Create
	err = service.Wallet.Create(w2)
	if err != nil {
		panic(err)
	}
}
