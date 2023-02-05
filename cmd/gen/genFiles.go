package main

import (
	"fmt"
	"go/format"
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

	for _, v := range list {
		name := strings.TrimSuffix(v, "Reader")
		result.WriteString(fmt.Sprintf("\t\"%s\": {\"%s\", &%s{}},\n", name, name, v))
	}

	result.WriteString("}\n")

	return format.Source([]byte(result.String()))
}

var funcsHeader = `package %s

type ReaderWrapper interface {
	GetData() any
}
`

var funcTmpl = `func (s *%s) GetData() any {
	return s.Data
}`

var namesHeader = `package %s

type Asset struct {
	Name string
	Parser ReaderWrapper
}

`
