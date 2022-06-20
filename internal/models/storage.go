// Package models exposes a set of models
package models

import (
	"errors"

	"github.com/jinzhu/gorm"

	// Mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

var (
	// errNotFound is returned when a resource cannot be found
	// in the database.
	errNotFound = errors.New("wallet not found")
)

// Wallet represents a wallet object in our database.
type Wallet struct {
	gorm.Model
	WID     string          `gorm:"not null;unique_index"`
	Balance decimal.Decimal `sql:"type:decimal(20,8);"`
}

// walletGorm represents our database interaction layer
type walletGorm struct {
	db *gorm.DB
}

// Get queries the db for an object with the given wid
func (wg *walletGorm) Get(wid string) (*Wallet, error) {
	var wallet Wallet
	err := wg.db.Where("w_id = ?", wid).First(&wallet).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errNotFound
		}
		return nil, err
	}
	return &wallet, nil
}

// Create creates a wallet
func (wg *walletGorm) Create(wallet *Wallet) error {
	return wg.db.Create(wallet).Error
}

// Update updates a wallet
func (wg *walletGorm) Update(wallet *Wallet) error {
	return wg.db.Save(wallet).Error
}

// Delete deletes a wallet
func (wg *walletGorm) Delete(wallet *Wallet) error {
	return wg.db.Delete(wallet).Error
}
