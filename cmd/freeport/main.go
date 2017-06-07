package main

import (
	"fmt"
	"github.com/gabstv/freeport"
	"os"
)

func main() {
	port, err := freeport.GetPort()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Print(port)
}
