package cache

import "sync"

type backupDeposit struct {
	ids map[uint64]struct{}
	sync.RWMutex
}

func (bd *backupDeposit) addIDToBackup(depositID uint64) {
	bd.Lock()
	defer bd.Unlock()

	bd.ids[depositID] = struct{}{}
}

func (bd *backupDeposit) getChangedUserIDs() []uint64 {
	bd.RLock()
	defer bd.RUnlock()

	var ids []uint64

	for id := range bd.ids {
		ids = append(ids, id)
	}

	return ids
}

func (bd *backupDeposit) cleanBackupList() {
	bd.Lock()
	defer bd.Unlock()

	bd.ids = make(map[uint64]struct{}, len(bd.ids))
}
