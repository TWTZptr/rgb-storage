package main

import (
	"fmt"
	"net"
	"rgb-storage/internal/handlers"
	"rgb-storage/internal/protocol"
	"rgb-storage/internal/snapshoters"
)

func handleClient(clientConn net.Conn) {
	defer func() {
		if err := clientConn.Close(); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}()

	buf := make([]byte, 65536)
	readBytesCount, err := clientConn.Read(buf)

	if err != nil {
		fmt.Printf("Error on client read: %v\n", err)
	}

	response := handlers.HandleClient(buf, readBytesCount)
	serializedResponse := protocol.SerializeResponse(response)
	wroteBytes, err := clientConn.Write(serializedResponse)

	if err != nil {
		fmt.Printf("Error on client write: %v\n", err)
	}

	if wroteBytes != len(serializedResponse) {
		fmt.Printf("Wrote bytes is not equal to serialized length: %d != %d\n", wroteBytes, serializedResponse)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("Error on start listen: %v\n", err)
		return
	}

	defer (func() {
		if err := ln.Close(); err != nil {
			panic(err)
		}
	})()

	snapshoters.RunSnapshoter("disk")
	snapshoters.LoadDataFromSnapshoter()

	fmt.Println("Storage started on :8080")

	for {
		clientConn, _ := ln.Accept()
		go handleClient(clientConn)
	}
}
