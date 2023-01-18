package animegame

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GetAsset(f string) Asset {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := AssetName[id]
	if ok {
		return a
	}

	return Asset{
		Name: f,
		Parser: func(f string) ([]byte, error) {
			return nil, fmt.Errorf("not implemeted")
		},
	}
}
