package db

import (
	"bytes"
	"sync"
)

// MemoryDB holds the key and value for
type MemoryDB struct {
	Name          string
	KeyV          map[string]string
	ReadWriteLock sync.RWMutex
}

// Get function returns the value
func (mdb *MemoryDB) Get(key string) string {
	mdb.ReadWriteLock.RLock()
	defer mdb.ReadWriteLock.RUnlock()
	if value, ok := mdb.KeyV[key]; ok {
		return value
	}
	return "(nil)"

}

// Set function saves a key value pair in the memory
func (mdb *MemoryDB) Set(key string, value string) string {
	mdb.ReadWriteLock.Lock()
	defer mdb.ReadWriteLock.Unlock()
	mdb.KeyV[key] = value
	return "OK"
}

// WriteToFile writes the encoded MemoryDB struct to the file
func WriteToFile(encDB bytes.Buffer) {

}
