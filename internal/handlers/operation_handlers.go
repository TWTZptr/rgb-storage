package handlers

import "rgb-storage/api"

type Handler interface {
	HandleGet(data []byte)
	HandleSet(data []byte)
	HandleDelete(data []byte)
}

type CommonHandler struct{}

func (h CommonHandler) HandleGet(data []byte) api.Response {
	return api.Response{}
}

func (h CommonHandler) HandleSet(data []byte) api.Response {
	return api.Response{}
}

func (h CommonHandler) HandleDelete(data []byte) api.Response {
	return api.Response{}
}
