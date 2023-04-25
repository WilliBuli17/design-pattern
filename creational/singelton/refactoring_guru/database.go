package refactoring_guru

import (
	"fmt"
	"sync"
)

var dbMutex sync.Mutex
var onceDb sync.Once
var instanceDatabase *database

type database struct {
}

// saya lebih suka menggunakan cara ini
func NewDatabaseUsingOnce() *database {
	onceDb.Do(func() {
		instanceDatabase = &database{}
		fmt.Println("Creating single instance using once now.")
	})

	fmt.Println("Single instance using once already created.")

	return instanceDatabase
}

func NewDatabaseUsingMutex() *database {
	if instanceDatabase == nil {
		dbMutex.Lock()
		defer dbMutex.Unlock()
		fmt.Println("Creating single instance using mutex now.")
		instanceDatabase = &database{}
	} else {
		fmt.Println("Single instance using mutex already created.")
	}

	return instanceDatabase
}
