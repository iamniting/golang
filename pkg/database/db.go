package database

import "fmt"

// Database interface
type Database interface {
	// Set the key to given value
	Set(key string, value int)

	// Get the value for the given key, set 'ok' to true if key exists
	Get(key string) (value int, ok bool)

	// Unset the key, making it just like that key was never set
	Unset(key string)

	// Begin opens a new transaction
	Begin()

	// Commit closes all open transaction blocks, permanently apply the
	// changes made in them.
	Commit() error

	// Rollback undoes all of the commands issued in the most recent
	// transaction block, and closes the block.
	Rollback() error
}

type db struct {
	data map[string]int
	logs []map[string]int
}

// NewDatabase returns the Database interface
func NewDatabase() Database {
	return &db{data: make(map[string]int)}
}

// Set the key to given value
func (d *db) Set(key string, value int) {
	d.data[key] = value
}

// Get the value for the given key, set 'ok' to true if key exists
func (d *db) Get(key string) (int, bool) {
	value, ok := d.data[key]
	return value, ok
}

// Unset the key, making it just like that key was never set
func (d *db) Unset(key string) {
	delete(d.data, key)
}

// Begin opens a new transaction
func (d *db) Begin() {
	// Save a snapshot of the current state of the data to the logs.
	snapshot := make(map[string]int)
	for k, v := range d.data {
		snapshot[k] = v
	}
	d.logs = append(d.logs, snapshot)
}

// Commit closes all open transaction blocks, permanently apply the
// changes made in them.
func (d *db) Commit() error {
	// Check if there is anything to commit.
	if len(d.logs) == 0 {
		return fmt.Errorf("nothing to commit")
	}

	d.logs = nil
	return nil
}

// Rollback undoes all of the commands issued in the most recent
// transaction block, and closes the block.
func (d *db) Rollback() error {
	// Check if there is anything to rollback.
	if len(d.logs) == 0 {
		return fmt.Errorf("nothing to rollback")
	}

	// Restore the data to the last snapshot in the logs.
	d.data = d.logs[len(d.logs)-1]
	d.logs = d.logs[:len(d.logs)-1]
	return nil
}
