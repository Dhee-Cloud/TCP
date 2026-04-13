package main

import ( 
	"fmt"
    "log"
    "net"
)


func main () {
	ln, err := net.Listen("tcp", "localhost:8080")
// net package from Go creates a TCP listener on the specified address and port,
// in this case "localhost:8080".
 	if err !=nil {
//if there is an error while creating the TCP listener, it will log the error and exit the program.
//err is the standard idiomatic name for a variable of the built-in error type
	log.fatal(err)
	//log.fatal(err) statement is a panic-and-exit mechanism that terminates your entire program 
	//immediately when an error occurs
	panic(err)
	//statement forces the Go runtime to stop execution immediately and display an error message in the form of a stack trace.
 }
 	defer ln.close()
//The defer ln.Close() statement schedules the Close() method to execute automatically when the main  function exits.
//it closes all underlying sockets,frees up the ports and rejects new connections.
//defer close is basically saying when the function finishes running close the listener and free up the resources automatically.

//next enter an infinite loop that waits for incoming client connections and accepts the connections
for {
conn, err := ln.Accept()
//when incoming request arrives use the accept function to accept the connection.
//This returns a connect object representing the clients connection
if ! := nil {
	log.fatal(err)
//if there is an error while accepting the connection, it will log the error and exit the program.
}
//once a connection is established we pass it to handle connections in a new goroutine using go handleConnection(conn)
go handleconnection(conn)
}

}

  