package main

import (
	"dataparse"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.uber.org/multierr"
)

func decodeInput(in, out string) error {
	if err := os.MkdirAll(out, 0o777); err != nil {
		return err
	}

	stat, err := os.Stat(in)
	if err != nil {
		return err
	}

	if stat.IsDir() {
		files, err := dataparse.GetAllFiles(in)
		if err != nil || len(files) < 1 {
			return fmt.Errorf("bad input folder %w", err)
		}

		var errr error
		for _, f := range files {
			multierr.AppendInto(&errr, decodeFile(f, out))
		}

		return errr
	}

	return decodeFile(in, out)
}

func decodeFile(name, out string) error {
	fileData, err := dataparse.DecodeFile(name)
	if err != nil {
		log.Println("error with "+name+" : ", err.Error())
	}

	outpath := filepath.Join(out, filepath.Base(name))
	if err := os.WriteFile(outpath, fileData, os.ModePerm); err != nil {
		log.Println("can't write file " + name + " : " + err.Error())
	}

	return nil
}
