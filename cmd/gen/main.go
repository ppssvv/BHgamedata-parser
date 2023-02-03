package main

import (
	"bufio"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	f := os.Args[1]

	f, err := prepareFile(f, "dump")
	if err != nil {
		log.Fatal("can't prepare file: ", err)
	}

	dir := filepath.Dir(f)
	packagename, list, err := parseFile(f)
	if err != nil || len(list) < 1 {
		log.Fatal("nothing found")
	}

	funcstxt, err := buildFuncs(packagename, list)
	if err != nil {
		log.Fatal("can't build funcs file: ", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "util.go"), []byte(funcstxt), fs.ModePerm); err != nil {
		log.Fatal("can't write functions: ", err)
	}

	namestxt, err := buildNames(packagename, list)
	if err != nil {
		log.Fatal("can't build names file: ", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "names.go"), []byte(namestxt), fs.ModePerm); err != nil {
		log.Fatal("can't write names: ", err)
	}

	if err := copyCommon(dir, packagename); err != nil {
		log.Fatal("can't copy file with common code: ", err)
	}
}

func prepareFile(f string, packagename string) (string, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(data), "\n")

	entryRE := regexp.MustCompile(`\tEntryCount\s{1,}uint32`)

	for i, line := range lines {
		// could also check for first line
		if strings.HasPrefix(line, "package ") {
			lines[i] = "package " + packagename
			continue
		}

		line = strings.Replace(line, `json:" - "`, `json:"-"`, 1)

		// all field names should be Uppercase
		if strings.HasPrefix(line, "\t") {
			line = "\t" + strings.ToTitle(line[1:2]) + line[2:]
		}

		// add bin tag to EntryCount fields
		if entryRE.MatchString(line) {
			line = line + " `bin:\"sizeof=Keys,ItemOffsets,Data\"`"
		}

		lines[i] = line
	}

	filedata, err := format.Source([]byte(strings.Join(lines, "\n")))
	if err != nil {
		return "", err
	}
	dir, filename := filepath.Split(f)
	newfilename := filepath.Join(dir, strings.TrimSuffix(filename, filepath.Ext(filename))+".go")
	os.WriteFile(newfilename, []byte(filedata), fs.ModePerm)

	return newfilename, nil
}

func parseFile(f string) (string, []string, error) {
	fd, err := os.Open(f)
	if err != nil {
		return "", nil, err
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	packagename := "s"
	result := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "package ") {
			packagename = strings.TrimPrefix(line, "package ")
		}

		if !strings.HasPrefix(line, "type ") { // type declaration
			continue
		}

		declaration := strings.Fields(line)
		// ["type", $name, "struct", "{"]
		if len(declaration) < 4 {
			fmt.Printf("\tweird declaration: %s\n", line)
		}

		name := declaration[1]

		if !strings.HasSuffix(name, "Reader") {
			continue
		}

		result = append(result, name)
	}

	return packagename, result, nil
}
