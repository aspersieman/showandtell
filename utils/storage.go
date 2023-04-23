package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

var baseFolder = "./storage"
var cacheFolder = baseFolder + "/cache"

type Storage struct {
	BaseFolder  string
	CacheFolder string
}

var Store Storage

func Setup() {
	Store = Storage{
		BaseFolder:  baseFolder,
		CacheFolder: cacheFolder,
	}
	if _, err := os.Stat(Store.BaseFolder); os.IsNotExist(err) {
		os.MkdirAll(Store.BaseFolder, 0777)
	}
	if _, err := os.Stat(Store.CacheFolder); os.IsNotExist(err) {
		os.MkdirAll(Store.CacheFolder, 0777)
	}
}

func CacheFolder(folder string) string {
	Setup()
	dir := filepath.Join(Store.CacheFolder, folder)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("Storage#CacheFolder: Creating folder %s\n", dir)
		os.MkdirAll(dir, 0777)
	}
	return dir
}
