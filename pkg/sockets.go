package pkg

import (
	"fmt"
	"net"
)

type connectionCallbackFunction func(net.Conn)

func StartServer(listenAddress string, connectionCallbackFunction connectionCallbackFunction) {
	listen, err := net.Listen("tcp", listenAddress)
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		return
	}

	fmt.Println("Waiting for client connection to", listenAddress)

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error connecting to client:", err.Error())
		}
		fmt.Println("Client " + connection.RemoteAddr().String() + " connected.")
		connectionCallbackFunction(connection)
	}

}

func ConnectToServer(serverAddress string, connectionCallbackFunction connectionCallbackFunction) {
	connection, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
	}
	connectionCallbackFunction(connection)
}
