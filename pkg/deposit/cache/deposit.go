package cache

import (
	"sync"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/deposit"
)

type depositCache struct {
	cache map[uint64]*models.Deposit
	sync.RWMutex
}

var backup backupDeposit

func NewDepositCacheMap() deposit.Cache {
	depositCache := &depositCache{
		cache: make(map[uint64]*models.Deposit),
	}

	backup = backupDeposit{
		ids:     make(map[uint64]struct{}),
		RWMutex: sync.RWMutex{},
	}

	return depositCache
}

func (dc *depositCache) AddDeposit(deposit *models.Deposit) {
	dc.Lock()
	defer dc.Unlock()

	dc.cache[deposit.ID] = deposit
	backup.addIDToBackup(deposit.ID)
}

func (dc *depositCache) PutDepositsToCache(deposits []*models.Deposit) {
	dc.Lock()
	defer dc.Unlock()

	for _, d := range deposits {
		dc.cache[d.ID] = d
	}
}

func (dc *depositCache) IsDepositExist(depositID uint64) bool {
	dc.RLock()
	defer dc.RUnlock()

	if _, exist := dc.cache[depositID]; exist {
		return true
	}

	return false
}

func (dc *depositCache) GetDepositCountAndSum(userID uint64) []float64 {
	dc.RLock()
	defer dc.RUnlock()

	var depositValues []float64

	for _, v := range dc.cache {
		if v.UserID == userID {
			depositValues = append(depositValues, v.Amount)
		}
	}

	return depositValues
}

func (dc *depositCache) GetBackupDeposits() []*models.Deposit {
	dc.RLock()
	defer dc.RUnlock()

	ids := backup.getChangedUserIDs()
	result := make([]*models.Deposit, 0, len(ids))

	for _, id := range ids {
		result = append(result, dc.cache[id])
	}

	return result
}

func (dc *depositCache) CleanBackupDeposits() {
	backup.cleanBackupList()
}
