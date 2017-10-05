FreePort
========

Get a free open TCP (or UDP) port that is ready to use

```bash
# Ask the kernel to give us an open port.
export port=$(freeport)

# Start standalone httpd server for testing
httpd -X -c "Listen $port" &

# Curl local server on the selected port
curl localhost:$port
```

#### Usage (Go)
```go
package main

import(
    "fmt"

    "github.com/gabstv/freeport"
)

func main(){
    tcpp, err := freeport.TCP()
    if err != nil {
        fmt.Println("Error getting a free TCP port:", err.Error())
        return
    }
    fmt.Println("TCP:", tcpp)

    udpp, err := freeport.UDP()
    if err != nil {
        fmt.Println("Error getting a free UDP port:", err.Error())
        return
    }
    fmt.Println("UDP:", udpp)
}
```

#### Building From Source
```bash
sudo apt-get install golang                    # Download go. Alternativly build from source: https://golang.org/doc/install/source
mkdir ~/.gopath && export GOPATH=~/.gopath     # Replace with desired GOPATH
export PATH=$PATH:$GOPATH/bin                  # For convenience, add go's bin dir to your PATH
go get github.com/gabstv/freeport/cmd/freeport
```
