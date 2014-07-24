package main

import "net"
import "fmt"

/**
 *
 */
func processRequest(sClientMessage string, oSocket *net.UDPConn, oClientAddress *net.UDPAddr) {
	fmt.Println("message received from the client : " + sClientMessage)

}

/**
 * program entrance
 */
func main() {
	oAddr, _ := net.ResolveUDPAddr("udp", ":31234")
	oSocket, _ := net.ListenUDP("udp", oAddr)
	fmt.Println("launching daemon")

	for {
		buf := make([]byte, 1024)
		// blocking method
		iMessageLength, oClientAddress, err := oSocket.ReadFromUDP(buf)
		if err != nil {
			fmt.Println(err)
		}
		sClientMessage := string(buf[0:iMessageLength])
		go processRequest(sClientMessage, oSocket, oClientAddress)
	}

}
