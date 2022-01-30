package commands

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

const (
	SEND_PUBLIC_NUMBERS = int8(0)
	SEND_PUBLIC_KEY     = int8(1)
)

type CommandArg struct {
	Length int8
	Value  []byte
}

type Command struct {
	CommandType int8
	ArgsLength  int8
	Args        []CommandArg
}

func SendCommand(connection net.Conn, commandType int8, rawArgs [][]byte) {
	responseBuffer := new(bytes.Buffer)
	binary.Write(responseBuffer, binary.LittleEndian, commandType)
	binary.Write(responseBuffer, binary.LittleEndian, int8(len(rawArgs)))
	for _, element := range rawArgs {
		binary.Write(responseBuffer, binary.LittleEndian, int8(binary.Size(element)))
		responseBuffer.Write(element)
	}
	connection.Write(responseBuffer.Bytes())
}

func ReceiveCommand(reader *bufio.Reader) (*Command, error) {
	command := Command{}
	binary.Read(reader, binary.LittleEndian, &command.CommandType)
	binary.Read(reader, binary.LittleEndian, &command.ArgsLength)
	for i := int8(0); i < command.ArgsLength; i++ {
		arg := CommandArg{}
		binary.Read(reader, binary.LittleEndian, &arg.Length)
		argValue := make([]byte, arg.Length)
		io.ReadFull(reader, argValue)
		arg.Value = argValue
		command.Args = append(command.Args, arg)
	}
	return &command, nil
}
