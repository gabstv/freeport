package freeport

import (
	"log"
	"testing"
)

func TestGetPort(t *testing.T) {
	p, err := TCP()
	if err != nil {
		t.Error(err)
	}
	log.Println("New TCP port:", p)
	p, err = UDP()
	if err != nil {
		t.Error(err)
	}
	log.Println("New UDP port:", p)
}
