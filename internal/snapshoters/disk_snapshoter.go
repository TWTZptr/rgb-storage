package snapshoters

import (
	"encoding/json"
	"fmt"
	"os"
)

type DiskSnapshoter struct {
	path string
}

func (d DiskSnapshoter) TakeSnapshot(storage map[string]string) {
	buf, err := json.Marshal(storage)

	if err != nil {
		fmt.Printf("Error on marshalling storage %v", err)
		return
	}

	err = os.WriteFile(d.path, buf, 0666)

	if err != nil {
		fmt.Printf("Error on writing file %v\n", d.path)
	}
}

func (d DiskSnapshoter) LoadSnapshot(storage map[string]string) {
	b, err := os.ReadFile(d.path)

	if err != nil {
		fmt.Printf("Error reading snapshot %v\n", d.path)
		return
	}

	err = json.Unmarshal(b, &storage)

	if err != nil {
		fmt.Printf("Error on unmarshaling data from snapshot: %v\n", err)
	}
}
