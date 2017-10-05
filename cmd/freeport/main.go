package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gabstv/freeport"
)

func main() {
	istcp := true
	if len(os.Args) > 1 {
		if strings.ToLower(os.Args[1]) == "udp" {
			istcp = false
		}
	}
	var port int
	var err error
	if istcp {
		port, err = freeport.GetPortTCP()
	} else {
		port, err = freeport.GetPortUDP()
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Print(port)
}
