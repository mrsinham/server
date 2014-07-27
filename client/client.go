package main

import (
	"fmt"
	"net"
	// "strconv"
	//"time"
)

func main() {
	// iLimit := 1000000
	oAddr, _ := net.ResolveTCPAddr("tcp", "localhost:31234")
	oConn, _ := net.DialTCP("tcp", nil, oAddr)

	// for i := 0; i < iLimit; i++ {

	sStr := "askpicture"

	fmt.Println("sending : " + sStr)
	oByteStr := []byte(sStr)
	oConn.Write(oByteStr)

	oConn.Close()

}
