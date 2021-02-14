package db

// MemoryDB holds the key and value for
type MemoryDB struct {
	Name string
	KeyV map[string]string
}

// Get function returns the value
func (mdb *MemoryDB) Get(key string) string {
	if value, ok := mdb.KeyV[key]; ok {
		return value
	}
	return "(nil)"

}

// Set function saves a key value pair in the memory
func (mdb *MemoryDB) Set(key string, value string) string {
	mdb.KeyV[key] = value
	return "OK"
}
