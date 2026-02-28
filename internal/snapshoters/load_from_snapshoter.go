package snapshoters

import (
	"fmt"
	storage "rgb-storage/internal/storage"
)

func LoadDataFromSnapshoter() {
	s := storage.GetMap()

	if snapshoter == nil {
		fmt.Println("snapshoter cannot load data: snapshoter == nil")
		return
	}

	snapshoter.LoadSnapshot(s)
}
