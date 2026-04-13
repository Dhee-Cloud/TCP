package main

import (
	"fmt"
	"net"
	"log"
)
func handleConnection(connecting net.Conn) {
	defer connecting.Close()

message := "Hello World! \n"
connecting.Write([]byte(message))

}
func main () {
	listening, error := net.Listen("tcp", "0.0.0.0:8080")
	if error != nil {
		fmt.Println("error in connecting")
		log.Fatal(error)
	}
	defer listening.Close()
	fmt.Println("sever connected on 8080")

for {
	connecting, error := listening.Accept()
	if error != nil {
		fmt.Println("error in listening")
		log.Fatal(error)
	}
go handleConnection(connecting)
}
}
