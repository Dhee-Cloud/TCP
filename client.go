package main 

import (
    "fmt"
    "io"
    "log"
    "net"
)

func main () {
    connect, error := net.Dial("tcp", "localhost:8080")
    if error != nil {
        fmt.Println("error in connecting")
        log.Fatal(error)
    }

    connecting, error := io.ReadAll(connect)
    if error != nil {
        fmt.Println("error in reading")
        log.Fatal(error)
    }
    fmt.Println(string(connecting))
}