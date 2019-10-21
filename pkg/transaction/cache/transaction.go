package cache

import (
	"sync"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/transaction"
)

type transactionCache struct {
	cache map[uint64]*models.Transaction
	sync.RWMutex
}

var backup backupTransactions

func NewTransactionCacheMap() transaction.Cache {
	transactionCache := &transactionCache{
		cache:   make(map[uint64]*models.Transaction),
		RWMutex: sync.RWMutex{},
	}

	backup = backupTransactions{
		ids:     make(map[uint64]struct{}),
		RWMutex: sync.RWMutex{},
	}

	return transactionCache
}

func (tc *transactionCache) AddTransaction(transaction *models.Transaction) {
	tc.Lock()
	defer tc.Unlock()

	tc.cache[transaction.ID] = transaction
	backup.addIDToBackup(transaction.ID)
}

func (tc *transactionCache) PutTransactionsToCache(transactions []*models.Transaction) {
	tc.Lock()
	defer tc.Unlock()

	for _, t := range transactions {
		tc.cache[t.ID] = t
	}
}

func (tc *transactionCache) IsTransactionExist(transactionID uint64) bool {
	tc.RLock()
	defer tc.RUnlock()

	if _, exist := tc.cache[transactionID]; exist {
		return true
	}

	return false
}

func (tc *transactionCache) GetWinCountAndSum(userID uint64) []float64 {
	tc.RLock()
	defer tc.RUnlock()

	var winValues []float64

	for _, v := range tc.cache {
		if v.UserID == userID && v.Name == "Win" {
			winValues = append(winValues, v.Amount)
		}
	}

	return winValues
}

func (tc *transactionCache) GetBetCountAndSum(userID uint64) []float64 {
	tc.RLock()
	defer tc.RUnlock()

	var betValues []float64

	for _, v := range tc.cache {
		if v.UserID == userID && v.Name == "Bet" {
			betValues = append(betValues, v.Amount)
		}
	}

	return betValues
}

func (tc *transactionCache) GetBackupTransactions() []*models.Transaction {
	tc.RLock()
	defer tc.RUnlock()

	ids := backup.getChangedUserIDs()
	result := make([]*models.Transaction, 0, len(ids))

	for _, id := range ids {
		result = append(result, tc.cache[id])
	}

	return result
}

func (tc *transactionCache) CleanBackupTransactions() {
	backup.cleanBackupList()
}
