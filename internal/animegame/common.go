package animegame

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func getStream(filename string) (*kaitai.Stream, error) {
	if !filepath.IsAbs(filename) {
		filename = filepath.Join("testdata", filename)
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return kaitai.NewStream(f), nil
}

func GetAsset(f string) Asset {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := AssetName[id]
	if ok {
		return a
	}

	return Asset{
		Name:   f,
		Parser: "",
	}
}

func resultDir(name string) (string, error) {
	s := filepath.Join("result", name)

	return s, os.MkdirAll(s, os.ModePerm)
}
