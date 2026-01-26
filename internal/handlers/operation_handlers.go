package handlers

import (
	"rgb-storage/api"
	"rgb-storage/internal/protocol"
	"rgb-storage/internal/storage"
	"strings"
)

type Handler interface {
	HandleGet(data []byte)
	HandleSet(data []byte)
	HandleDelete(data []byte)
}

type CommonHandler struct{}

func (h CommonHandler) HandleGet(data []byte) api.Response {
	propertyName := protocol.DeserializeBody(data)
	m := storage.GetMap()
	val, ok := m[propertyName]

	if !ok {
		return api.Response{Err: "Property does no exist"}
	}

	return api.Response{Val: val}
}

func (h CommonHandler) HandleSet(data []byte) api.Response {
	body := protocol.DeserializeBody(data)
	res := strings.Split(body, " ")

	if len(res) != 2 {
		return api.Response{Err: "Invalid pair name/value"}
	}

	propertyName, propertyValue := res[0], res[1]
	m := storage.GetMap()
	m[propertyName] = propertyValue

	return api.Response{}
}

func (h CommonHandler) HandleDelete(data []byte) api.Response {
	return api.Response{}
}
