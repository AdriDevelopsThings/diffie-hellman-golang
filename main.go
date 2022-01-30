package main

import (
	"fmt"
	"os"

	"github.com/adridevelopsthings/diffie-hellman-golang/pkg"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("diffie-hellman", "Diffie hellman key exchange over socket")
	mode := parser.String("m", "mode", &argparse.Options{Required: true, Help: "Run as 'server' or 'client'"})
	address := parser.String("a", "address", &argparse.Options{Required: true, Help: "Server (listen) address like 127.0.0.1:8008"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println("Error while parsing command args: ", err.Error())
		return
	}
	if *mode != "server" && *mode != "client" {
		fmt.Println("Mode argument must be server or client.")
		return
	}
	if *mode == "server" {
		pkg.StartServer(*address, pkg.ReceivePublicNumbers)
	} else {
		pkg.ConnectToServer(*address, pkg.SendPublicNumbers)
	}
}
