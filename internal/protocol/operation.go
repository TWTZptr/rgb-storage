package protocol

type Operation int

const (
	OpGet Operation = iota
	OpSet
	OpDelete
)
