package main

import (
	"fmt"
	"net"
	"rgb-storage/api"
	"rgb-storage/internal/handlers"
)

func handleClient(clientConn net.Conn) api.Response {
	defer func() {
		if err := clientConn.Close(); err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}()

	buf := make([]byte, 65536)
	readBytesCount, err := clientConn.Read(buf)

	if err != nil {
		fmt.Printf("Error on client read: %v", err)
		return api.Response{}
	}

	operationType := api.Operation(buf[0])
	handler := handlers.CommonHandler{}
	payload := buf[1:readBytesCount]

	switch operationType {
	case api.OpGet:
		return handler.HandleGet(payload)

	case api.OpSet:
		return handler.HandleSet(payload)

	case api.OpDelete:
		return handler.HandleDelete(payload)

	default:
		fmt.Printf("Invalid operation: %d", operationType)
		return api.Response{Err: "Invalid OpType"}
	}
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")

	a := api.Response{}
	fmt.Printf("%v", a)

	defer (func() {
		if err := ln.Close(); err != nil {
			panic(err)
		}
	})()

	for {
		clientConn, _ := ln.Accept()
		go handleClient(clientConn)
	}
}
