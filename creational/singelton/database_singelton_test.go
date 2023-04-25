package singelton

import (
	"design-patterns/creational/singelton/refactoring_guru"
	"fmt"
	"testing"
)

func TestDatabaseSingeltonOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		go refactoring_guru.NewDatabaseUsingOnce()
	}

	// Fungsi Scanln mirip dengan fungsi Scan, tetapi berhenti melakukan pemindaian pada karakter newline, dan setelah item terakhir harus diikuti oleh karakter newline atau EOF.
	_, err := fmt.Scanln()
	if err != nil {
		panic(err)
	}
}

func TestDatabaseSingeltonMutex(t *testing.T) {
	for i := 0; i < 10; i++ {
		go refactoring_guru.NewDatabaseUsingMutex()
	}

	// Fungsi Scanln mirip dengan fungsi Scan, tetapi berhenti melakukan pemindaian pada karakter newline, dan setelah item terakhir harus diikuti oleh karakter newline atau EOF.
	_, err := fmt.Scanln()
	if err != nil {
		panic(err)
	}
}
