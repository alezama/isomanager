package main

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

func TestSendingMessage (t *testing.T) {
	os.Args = append(os.Args, "-a=127.0.0.1:9505")
	os.Args = append(os.Args, "-m=Estaeslacadenadeentrada")
	msgResponse := "This is the message response"
	go func () {

		listener, _ := net.Listen("tcp", ":9505")
		conn, _ :=	listener.Accept()
		defer conn.Close()
		inputBuf := make([]byte,1024)
		n, err := conn.Read(inputBuf)
		if err != nil {
			fmt.Println("Error reading from socket")
			return
		}
		if n == 0 {
			fmt.Println("Empty input buffer")
		}
		fmt.Println("Message received")
		n, err = conn.Write([]byte(msgResponse))
		if n == 0 {
			fmt.Println("Empty input buffer")
		}
		fmt.Println("Response sent")
	} ()
	main()

}


func TestSendingMessageTimeout (t *testing.T) {
	os.Args = append(os.Args, "-a=127.0.0.1:9505")
	os.Args = append(os.Args, "-m=Estaeslacadenadeentrada")
	msgResponse := "This is the message response"
	go func () {

		listener, _ := net.Listen("tcp", ":9505")
		conn, _ :=	listener.Accept()
		defer conn.Close()
		inputBuf := make([]byte,1024)
		n, err := conn.Read(inputBuf)
		if err != nil {
			fmt.Println("Error reading from socket")
			return
		}
		if n == 0 {
			fmt.Println("Empty input buffer")
		}
		fmt.Println("Message received")
		time.Sleep(10*time.Second)
		n, err = conn.Write([]byte(msgResponse))
		if n == 0 {
			fmt.Println("Empty input buffer")
		}
		fmt.Println("Response sent")
	} ()
	main()

}


