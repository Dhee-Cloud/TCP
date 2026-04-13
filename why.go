package main

import (
	"fmt"
	"log"
	"net"
	"io"
)

func main () {
	connecting, error := net.Dial("tcp", "localhost:8080")
	if error != nil {
		fmt.Println("error in listening")
		log.Fatal(error)

		defer connecting.Close()
fmt.Println("connection established")
}


connection, error := io.ReadAll(connecting)
if error != nil {
	fmt.Println("error in connecting")
	log.Fatal(error)
}
fmt.Println(string(connection))
	
	}
