package util

import (
	"log"
	"os"
	"path/filepath"
)

func GetDownloadDirIfFolderIsNil(folder string) string {
	if folder == "" {
		dir, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return ""
		}
		folder = filepath.Join(dir, "download")
	}
	return folder
}
