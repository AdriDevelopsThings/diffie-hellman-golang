package pkg

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"net"

	"github.com/adridevelopsthings/diffie-hellman-golang/pkg/commands"
)

var pub_p *big.Int
var pub_g *big.Int
var pub_s *big.Int
var pub_my_key *big.Int
var pub_key *big.Int
var pub_shared_priv *big.Int

func SendPublicNumbers(conn net.Conn) {
	p, err := rand.Prime(rand.Reader, 16)
	if err != nil {
		fmt.Println("Error while generating public p: ", err.Error())
	}

	var g *big.Int
	g, err = rand.Int(rand.Reader, p)
	if err != nil {
		fmt.Println("Error while generating public g: ", err.Error())
	}

	var s *big.Int
	s, err = rand.Int(rand.Reader, p)
	if err != nil {
		fmt.Println("Error while generating public s: ", err.Error())
	}

	args := [][]byte{p.Bytes(), g.Bytes()}
	pub_p = p
	pub_g = g
	pub_s = s
	pub_my_key = new(big.Int)
	pub_my_key.Exp(g, s, p)

	commands.SendCommand(conn, commands.SEND_PUBLIC_NUMBERS, args)
	fmt.Println("Sent public numbers.")
	receivePublicKey(conn, true)
}

func ReceivePublicNumbers(conn net.Conn) {
	reader := bufio.NewReader(conn)
	command, err := commands.ReceiveCommand(reader)
	if err != nil {
		fmt.Println("Error while receiving command:", err.Error())
	}
	if command.CommandType == commands.SEND_PUBLIC_NUMBERS {
		fmt.Println("Received public numbers.")
		pub_p = new(big.Int)
		pub_g = new(big.Int)
		pub_p.SetBytes(command.Args[0].Value)
		pub_g.SetBytes(command.Args[1].Value)
		fmt.Println(pub_p, pub_g)
		var s *big.Int
		s, err = rand.Int(rand.Reader, pub_p)
		if err != nil {
			fmt.Println("Error while generating public s: ", err.Error())
		}
		pub_s = s
		pub_my_key = new(big.Int)
		pub_my_key.Exp(pub_g, pub_s, pub_p)
		sendPublicKey(conn, true)
	}

}

func sendPublicKey(conn net.Conn, rereceive bool) {
	args := [][]byte{pub_my_key.Bytes()}
	commands.SendCommand(conn, commands.SEND_PUBLIC_KEY, args)
	fmt.Println("Send Public key")
	if rereceive {
		receivePublicKey(conn, false)
	} else {
		printInformationAndClose(conn)
	}
}

func receivePublicKey(conn net.Conn, resend bool) {
	reader := bufio.NewReader(conn)
	command, err := commands.ReceiveCommand(reader)
	if err != nil {
		fmt.Println("Error while receiving command:", err.Error())
	}
	if command.CommandType == commands.SEND_PUBLIC_KEY {
		fmt.Println("Received Public key")

		pub_key = new(big.Int)
		pub_key.SetBytes(command.Args[0].Value)
		pub_shared_priv = new(big.Int)
		pub_shared_priv.Exp(pub_key, pub_s, pub_p)
		if resend {
			sendPublicKey(conn, false)
		} else {
			printInformationAndClose(conn)
		}
	}

}

func printInformationAndClose(conn net.Conn) {
	fmt.Println("Information:")
	fmt.Printf("Public shared p=%d\n", pub_p)
	fmt.Printf("Public shared g=%d\n", pub_g)
	fmt.Printf("My secret s=%d\n", pub_s)
	fmt.Printf("My public key K=%d\n", pub_my_key)
	fmt.Printf("Others public key K=%d\n", pub_key)
	fmt.Printf("Shared secret K=%d\n", pub_shared_priv)
	conn.Close()
}
