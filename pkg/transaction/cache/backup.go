package cache

import "sync"

type backupTransactions struct {
	ids map[uint64]struct{}
	sync.RWMutex
}

func (bt *backupTransactions) addIDToBackup(transactionID uint64) {
	bt.Lock()
	defer bt.Unlock()

	bt.ids[transactionID] = struct{}{}
}

func (bt *backupTransactions) getChangedUserIDs() []uint64 {
	bt.RLock()
	defer bt.RUnlock()

	var ids []uint64

	for id := range bt.ids {
		ids = append(ids, id)
	}

	return ids
}

func (bt *backupTransactions) cleanBackupList() {
	bt.Lock()
	defer bt.Unlock()

	bt.ids = make(map[uint64]struct{}, len(bt.ids))
}
