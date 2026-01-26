package protocol

import (
	"rgb-storage/api"
)

func DeserializeBody(data []byte) string {
	return string(data)
}

func SerializeResponse(response api.Response) []byte {
	val := []byte(response.Val)
	err := []byte(response.Err)
	result := make([]byte, len(val)+len(err)+1)
	result = append(result, val...)
	result = append(result, 0)
	result = append(result, err...)
	return result
}
