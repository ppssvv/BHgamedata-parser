package dataparse

import (
	"dataparse/dump"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	bin "github.com/streamingfast/binary"
	"go.uber.org/multierr"
)

func Parse(name, out string) error {
	stat, err := os.Stat(name)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return ParseFolder(name, out)
	}

	return ParseFile(name, out)
}

func ParseFolder(name, out string) error {
	files, err := GetAllFiles(name)
	if err != nil || len(files) < 1 {
		return fmt.Errorf("bad input folder %w", err)
	}

	var errr error
	for _, f := range files {
		multierr.AppendInto(&errr, ParseFile(f, out))
	}

	return errr
}

func ParseFile(f string, out string) error {
	if err := os.MkdirAll(out, 0o777); err != nil {
		return err
	}

	id := GetIdFromName(f)

	asset, ok := dump.AssetName[id]
	if !ok {
		return fmt.Errorf("%d is not supported yet", id)
	}

	if fileExist(out, asset.Name) {
		// already exist
		return nil
	}

	data, err := prepareFileData(f)
	if err != nil {
		return err
	}

	dec := bin.NewDecoder(data)

	if err := dec.Decode(asset.Parser); err != nil {
		return err
	}

	// normal files end early
	if asset.Name != "CombinedBytes" {
		return writeJsonToFile(out, asset.Name, asset.Parser)
	}

	var combinedErr error

	// CombinedBytes is a container of files, need to parse them all individually
	combinedContainer, ok := asset.Parser.(*dump.CombinedBytesReader)
	if !ok {
		return fmt.Errorf("CombinedBytes has wrong reader type")
	}

	for _, info := range combinedContainer.GetFiles() {
		if fileExist(out, info.AssetInfo.Name) {
			// already exist
			continue
		}

		decSub := bin.NewDecoder(info.AssetData)
		if err := decSub.Decode(info.AssetInfo.Parser); err != nil {
			multierr.AppendInto(&combinedErr, fmt.Errorf("%s - %w", info.AssetInfo.Name, err))
			continue
		}

		outdata, err := json.MarshalIndent(info.AssetInfo.Parser, "", "  ")
		if err != nil {
			return err
		}

		outpath := filepath.Join(out, fmt.Sprintf("%s.json", info.AssetInfo.Name))
		if err := os.WriteFile(outpath, outdata, 0o777); err != nil {
			multierr.AppendInto(&combinedErr, fmt.Errorf("%s - %w", info.AssetInfo.Name, err))
		}
	}

	return combinedErr
}

func prepareFileData(f string) ([]byte, error) {
	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}

	needDecrypt, err := checkEncryption(fd)
	if err != nil {
		// smth wrong with file
		return nil, err
	}

	if needDecrypt {
		return Decode(fd)
	}

	return io.ReadAll(fd)
}
