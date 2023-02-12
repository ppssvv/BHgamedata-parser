package dataparse

import (
	"dataparse/dump"
	"path/filepath"
	"strings"
)

func GetAsset(f string) ([]dump.Asset, bool) {
	f = filepath.Base(f)
	id, _, _ := strings.Cut(f, "_")
	a, ok := NewAssets[id]
	return a, ok
}

func EventDialogMultiLang(lang string) []dump.Asset {
	ass := dump.AssetName["EventDialogDataMetaData"]
	ass.Name += "_" + lang

	return []dump.Asset{ass}
}

func TextMapMultiLang(lang string) []dump.Asset {
	ass := dump.AssetName["TextMapMultiLangMetaData"]
	ass.Name += "_" + lang

	return []dump.Asset{ass}
}
