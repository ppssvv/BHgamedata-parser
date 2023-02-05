package dataparse

import (
	"dataparse/dump"
	"dataparse/internal/decode"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	bin "github.com/streamingfast/binary"
)

func readFile(f string) *bin.Decoder {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	return bin.NewDecoder(data)
}

func checkEncrypted(f string) bool {
	fd, err := os.Open(f)
	if err != nil {
		panic(err)
	}

	var size int32
	if err := binary.Read(fd, binary.LittleEndian, &size); err != nil {
		panic(err)
	}
	defer fd.Close()

	stats, err := fd.Stat()
	if err != nil {
		panic(err)
	}

	return stats.Size() != int64(size)
}

func GetTestData(f string) *bin.Decoder {
	if !filepath.IsAbs(f) {
		f = filepath.Join("testdata", f)
	}

	// before reading all file, check first 4 bytes
	if checkEncrypted(f) {
		log.Println("file is encrypted, dec")
		decodedData, err := decode.Parse(f)
		if err != nil {
			panic(err)
		}
		return bin.NewDecoder(decodedData)
	}

	return readFile(f)
}

func ProcessStructNew(f string, obj dump.ReaderWrapper) ([]byte, error) {

	if err := GetTestData(f).Decode(obj); err != nil {
		return nil, err
	}

	if obj.GetData() == nil {
		return nil, fmt.Errorf("no data")
	}

	return json.MarshalIndent(obj.GetData(), "", "  ")
	// return jettison.MarshalOpts(obj.GetData(), jettison.NilSliceEmpty(), jettison.NoCompact())
}
