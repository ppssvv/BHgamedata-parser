package dataparse

import (
	"dataparse/dump"
	"path/filepath"
	"strings"
)

func GetAsset(f string) dump.Asset {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := NewAssets[id]
	if ok {
		return a
	}

	return dump.Asset{
		Name:   f,
		Parser: nil,
	}
}

func EventDialogMultiLang(lang string) dump.Asset {
	ass := dump.AssetName["EventDialogDataMetaData"]
	ass.Name += "_" + lang

	return ass
}

func TextMapMultiLang(lang string) dump.Asset {
	ass := dump.AssetName["TextMapMultiLangMetaData"]
	ass.Name += "_" + lang

	return ass
}
