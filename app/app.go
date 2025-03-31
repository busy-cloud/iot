package app

import (
	"archive/zip"
	_ "embed"
	"encoding/json"
	"io"
	"path/filepath"
)

type Manifest struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
	Author      string `json:"author,omitempty"`
	Copyright   string `json:"copyright,omitempty"`
	Url         string `json:"url,omitempty"`
}

const APP_PATH = "apps"
const APP_EXT = ".zip"
const APP_MANIFEST = "manifest.json"
const APP_ICON = "icon.png"

//go:embed icon.png
var icon []byte

func ReadManifest(name string) (*Manifest, error) {
	reader, err := zip.OpenReader(filepath.Join(APP_PATH, name))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	file, err := reader.Open(APP_MANIFEST)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var app Manifest
	err = json.Unmarshal(buf, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func ReadIcon(name string) ([]byte, error) {
	reader, err := zip.OpenReader(filepath.Join(APP_PATH, name))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	file, err := reader.Open(APP_ICON)
	if err != nil {
		//return nil, err
		return icon, nil //使用默认图片
	}
	defer file.Close()

	return io.ReadAll(file)
}
