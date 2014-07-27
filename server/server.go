package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/**
 *
 */
func listener(oConnection *net.TCPConn) {
	// fmt.Println("message received from the client : " + sClientMessage)
	oReader := bufio.NewReader(oConnection)
	defer oConnection.Close()
	aByts, oError := oReader.ReadBytes('\n')
	if oError == io.EOF {
		sMessage := string(aByts)
		fmt.Println("received : " + sMessage)
	} else {
		fmt.Println("not ununderstable message")
	}

}

/**
 * program entrance
 */
func main() {
	oAddr, _ := net.ResolveTCPAddr("tcp", ":31234")
	oSocket, oErr := net.ListenTCP("tcp", oAddr)

	if oErr != nil {
		panic("unable to start daemon " + oErr.Error())
	}

	log.Print("starting daemon")
	defer oSocket.Close()

	for {

		// blocking method
		oConnection, err := oSocket.AcceptTCP()
		if err != nil {
			// fmt.Println(err)
			panic("unable to start daemon " + err.Error())
		}

		go listener(oConnection)
	}

}
