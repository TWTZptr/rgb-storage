package handlers

import (
	"fmt"
	"rgb-storage/api"
	"rgb-storage/internal/protocol"
)

func HandleClient(buf []byte, readBytesCount int) api.Response {
	operationType := protocol.Operation(buf[0])
	handler := CommonHandler{}
	payload := buf[1:readBytesCount]

	switch operationType {
	case protocol.OpGet:
		return handler.HandleGet(payload)

	case protocol.OpSet:
		return handler.HandleSet(payload)

	case protocol.OpDelete:
		return handler.HandleDelete(payload)

	default:
		fmt.Printf("Invalid operation: %d\n", operationType)
		return api.Response{Err: "Invalid OpType"}
	}
}
