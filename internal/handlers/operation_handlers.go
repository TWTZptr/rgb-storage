package handlers

import (
	"rgb-storage/api"
	"rgb-storage/internal/protocol"
	"rgb-storage/internal/storage"
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
	return api.Response{}
}

func (h CommonHandler) HandleDelete(data []byte) api.Response {
	return api.Response{}
}
