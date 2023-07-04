package main

import (
	"fmt"
	"net"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "2012"
)

func err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
func server_addr(server_host string, server_port string) string {
	return server_host + ":" + server_port
}

func handle_client(connection net.Conn) {
	// Inspired by: https://www.developer.com/languages/intro-socket-programming-go/

	buffer := make([]byte, 1024)

	messageLength, err := connection.Read(buffer)
	err_panic(err)

	fmt.Println("Received: " + string(buffer[:messageLength]))

	_, err = connection.Write([]byte("Received: " + string(buffer[:messageLength])))
	err_panic(err)

	connection.Close()
}

func main() {

	server_address := server_addr(SERVER_HOST, SERVER_PORT)

	server, err := net.Listen(SERVER_TYPE, server_address)

	if err != nil {
		panic(err)
	}

	defer server.Close()

	fmt.Println("Listening on " + server_address)

	for {
		connection, err := server.Accept()

		if err != nil {
			panic(err)
		}

		fmt.Println("Listening to a client.")

		go handle_client(connection)

	}

}
