package snapshoters

type Snapshoter interface {
	TakeSnapshot(storage map[string]string)
	LoadSnapshot(storage map[string]string)
}
