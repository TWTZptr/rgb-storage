package storage

var m map[string]string = nil

func GetMap() map[string]string {
	if m == nil {
		m = make(map[string]string)
	}

	return m
}
