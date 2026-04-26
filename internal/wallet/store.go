package wallet

import "gorm.io/gorm"

type WalletStore struct {
	db *gorm.DB
}