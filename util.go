package dataparse

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetIdFromName(name string) int {
	idText := strings.Split(filepath.Base(name), "_")[0]
	id, err := strconv.Atoi(idText)
	if err != nil {
		return 0
	}

	return id
}

func writeJsonToFile(dir, name string, data interface{}) error {
	outdata, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	outpath := filepath.Join(dir, fmt.Sprintf("%s.json", name))
	return os.WriteFile(outpath, outdata, 0o777)
}

// checkEncryption return true if need to decrypt data, or false if data is ready
// error if something wrong with file
func checkEncryption(fd io.ReadSeeker) (bool, error) {
	var size int32
	if err := binary.Read(fd, binary.LittleEndian, &size); err != nil {
		return false, err
	}

	realSize, err := fd.Seek(0, io.SeekEnd)
	if err != nil {
		return false, err
	}

	// move to the start of file
	if _, err := fd.Seek(0, io.SeekStart); err != nil {
		return false, err
	}

	return realSize != int64(size), nil
}

func fileExist(dir, name string) bool {
	res, err := filepath.Glob(filepath.Join(dir, fmt.Sprintf("%s.json", name)))
	if err == nil && len(res) > 0 {
		return true
	}

	return false
}

func GetAllFiles(dir string) ([]string, error) {
	var result []string

	res, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range res {
		if f.IsDir() {
			continue
		}

		// ignore hidden files like .gitignore and etc
		if f.Name()[0] == '.' {
			continue
		}

		result = append(result, filepath.Join(dir, f.Name()))
	}

	return result, nil
}
