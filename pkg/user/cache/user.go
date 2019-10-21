package cache

import (
	"sync"

	"github.com/IamStubborN/test/models"
	"github.com/IamStubborN/test/pkg/user"
)

type usersCache struct {
	cache       map[uint64]*models.User
	backupUsers backupUsers
	sync.RWMutex
}

func NewUsersCacheMap() user.Cache {
	usersCache := &usersCache{
		cache: make(map[uint64]*models.User),
		backupUsers: backupUsers{
			ids:     make(map[uint64]struct{}),
			RWMutex: sync.RWMutex{},
		},
		RWMutex: sync.RWMutex{},
	}

	return usersCache
}

func (uc *usersCache) AddUser(user *models.User) {
	uc.Lock()
	defer uc.Unlock()

	uc.cache[user.ID] = user
	uc.backupUsers.addIDToBackup(user.ID)
}

func (uc *usersCache) PutUsersToCache(users []*models.User) {
	uc.Lock()
	defer uc.Unlock()

	for _, u := range users {
		uc.cache[u.ID] = u
	}
}

func (uc *usersCache) GetUser(userID uint64) *models.User {
	uc.RLock()
	defer uc.RUnlock()

	return uc.cache[userID]
}

func (uc *usersCache) IsUserExist(userID uint64) bool {
	uc.RLock()
	defer uc.RUnlock()

	if _, exist := uc.cache[userID]; exist {
		return true
	}

	return false
}

func (uc *usersCache) ChangeUserBalance(userID uint64, balance float64) {
	uc.Lock()
	defer uc.Unlock()

	uc.cache[userID].Balance = balance
	uc.backupUsers.addIDToBackup(userID)
}

func (uc *usersCache) GetBackupUsers() []*models.User {
	uc.RLock()
	defer uc.RUnlock()

	ids := uc.backupUsers.getChangedUserIDs()
	result := make([]*models.User, 0, len(ids))

	for _, id := range ids {
		result = append(result, uc.cache[id])
	}

	return result
}

func (uc *usersCache) CleanBackupUsers() {
	uc.backupUsers.cleanBackupList()
}
