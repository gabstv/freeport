package freeport

import (
	"log"
	"testing"
)

func TestGetPort(t *testing.T) {
	p, err := GetPortTCP()
	if err != nil {
		t.Error(err)
	}
	log.Println("New TCP port:", p)
	p, err = GetPortUDP()
	if err != nil {
		t.Error(err)
	}
	log.Println("New UDP port:", p)
}
