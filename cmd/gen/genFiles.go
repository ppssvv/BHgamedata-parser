package main

import (
	"fmt"
	"go/format"
	"strconv"
	"strings"
)

func buildFuncs(packagename string, list []string) ([]byte, error) {
	result := strings.Builder{}

	result.WriteString(fmt.Sprintf(funcsHeader, packagename))

	for _, v := range list {
		result.WriteString(fmt.Sprintf(funcTmpl, v) + "\n")
	}

	return format.Source([]byte(result.String()))
}

func buildNames(packagename string, list []string) ([]byte, error) {
	result := strings.Builder{}

	result.WriteString(fmt.Sprintf(namesHeader, packagename))

	result.WriteString("var AssetName = map[string]Asset{\n")

	for k, v := range list {
		result.WriteString("\t\"" + strconv.Itoa(k) + "\": {\"" +
			strings.TrimSuffix(v, "Reader") +
			"\", &" + v + "{}},\n")
	}

	result.WriteString("}\n")

	return format.Source([]byte(result.String()))
}

var funcsHeader = `package %s

import (
	"encoding/json"
)

`

var funcTmpl = `func (s *%s) JSON() ([]byte, error) {
	return json.MarshalIndent(s.Data, "", "  ")
}`

var namesHeader = `package %s

type GameStruct interface {
	JSON() ([]byte, error)
}

type Asset struct {
	Name string
	Parser GameStruct
}

`
