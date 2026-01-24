package api

type Operation int

const (
	OpGet Operation = iota
	OpSet
	OpDelete
)
