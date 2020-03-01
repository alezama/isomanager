package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {

	tcpAddress := flag.String("a", "", "IP address sending the message e.g. 127.0.0.1:9010")
	isoMsgString := flag.String("m", "", "ISO message to send")

	flag.Parse()

	if *tcpAddress == "" || *isoMsgString == "" {
		flag.Usage()
	}

	fmt.Println("address ", *tcpAddress )

	tcpAddr, _ := net.ResolveTCPAddr("tcp4", *tcpAddress)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()

	if err != nil {
		fmt.Println("Error connecting to the port check is listening")
		return
	}
	msgLength := len(*isoMsgString)
	msgRet := make([]byte, 2)
	binary.BigEndian.PutUint16(msgRet, uint16(msgLength))
	msgRet = append(msgRet, []byte(*isoMsgString)...)

	conn.Write(msgRet)

	conn.SetReadDeadline(time.Now().Add(7*time.Second) )

	msgResponse := make([]byte, 1024)
	n, err := conn.Read(msgResponse)

	if n > 0 {
		fmt.Println(string(msgResponse))
	}
	if e, ok := err.(interface{Timeout() bool}); ok && e.Timeout() {
		fmt.Println("Response timed out")

	} else if err != nil {
		fmt.Println("Error found ")
	}

}


