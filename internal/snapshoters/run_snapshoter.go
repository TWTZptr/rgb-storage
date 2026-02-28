package snapshoters

import (
	"fmt"
	"rgb-storage/internal/storage"
	"time"
)

var snapshoter Snapshoter = nil

func RunSnapshoter(snapshoterType string) {
	fmt.Println("Starting snapshoter...")

	if snapshoter != nil {
		fmt.Println("Snapshoter already exists, snapshoter will be recreated")
	}

	switch snapshoterType {
	case "disk":
		snapshoter = DiskSnapshoter{path: defaultDiskSnapshoterPath}

	default:
		fmt.Printf("Invalid snapshoter specified: %s. Snapshots would not be enabled\n", snapshoterType)
		return
	}

	runSnapshoterRoutine()
	fmt.Println("Snapshoter started")
}

func runSnapshoterRoutine() {
	ticker := time.NewTicker(30 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C:
				snapshoter.TakeSnapshot(storage.GetMap())
			}
		}
	}()
}
