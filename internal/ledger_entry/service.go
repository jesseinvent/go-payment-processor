package ledgerentry

import "fmt"

type LedgerEntryService interface {
		Create(
			userId uint, 
			transactionId uint, 
			currencyId uint, 
			entryType LedgerEntryType, 
			amount int,
		) (*LedgerEntry, error)
}

type ledgerEntryService struct {
	ledgerEntryStore LedgerEntryStore
}

func NewLedgerEntryService(ledgerEntryStore LedgerEntryStore) LedgerEntryService {
	return &ledgerEntryService{ledgerEntryStore: ledgerEntryStore}
}

func (s *ledgerEntryService) Create(
			userId uint, 
			transactionId uint, 
			currencyId uint, 
			entryType LedgerEntryType, 
			amount int,
	) (*LedgerEntry, error) {

	ledgerEntry := &LedgerEntry{
		UserId: userId,
		TransactionId: transactionId,
		CurrencyId: currencyId,
		EntryType: entryType,
		Amount: amount,
	}

	err := s.ledgerEntryStore.Create(ledgerEntry)

	if err != nil {
		return nil, fmt.Errorf("failed to create ledger: %v", err)
	}

	return ledgerEntry, nil
}