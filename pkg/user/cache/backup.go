package cache

import "sync"

type backupUsers struct {
	ids map[uint64]struct{}
	sync.RWMutex
}

func (bu *backupUsers) addIDToBackup(userID uint64) {
	bu.Lock()
	defer bu.Unlock()

	bu.ids[userID] = struct{}{}
}

func (bu *backupUsers) getChangedUserIDs() []uint64 {
	bu.RLock()
	defer bu.RUnlock()

	var ids []uint64

	for id := range bu.ids {
		ids = append(ids, id)
	}

	return ids
}

func (bu *backupUsers) cleanBackupList() {
	bu.Lock()
	defer bu.Unlock()

	bu.ids = make(map[uint64]struct{}, len(bu.ids))
}