package main

import (
	"net"
	"time"
)

func main() {
	iLimit := 1000000
	oAddr, _ := net.ResolveUDPAddr("udp", "localhost:31234")
	oConn, _ := net.DialUDP("udp", nil, oAddr)

	for i := 0; i < iLimit; i++ {

		sStr := "test" + string(i)
		oByteStr := []byte(sStr)
		oConn.Write(oByteStr)

		time.Sleep(time.Duration(200) * time.Microsecond)
	}
	oConn.Close()
}
