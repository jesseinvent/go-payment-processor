package ledgerentry

import "gorm.io/gorm"

type LedgerEntryStore interface {
	Create(ledgerEntry *LedgerEntry) error
}

type ledgerEntryStore struct {
	db *gorm.DB
}

func NewLedgerEntryStore(db *gorm.DB) LedgerEntryStore {
	return &ledgerEntryStore{db: db}
}

func (s *ledgerEntryStore) Create(ledgerEntry *LedgerEntry) error {
	return s.db.Create(ledgerEntry).Error
}