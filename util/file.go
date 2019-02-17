package util

import (
	"io"
	"mime/multipart"
	"os"
)

const dest = "C:\\tmp\\upload\\"

func CreateFile(id string, file multipart.FileHeader) (err error) {
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return
	}
	autoCreateDir(dest + id)
	dest, err := os.Create(dest + id + "\\" + file.Filename)
	defer dest.Close()
	if err != nil {
		return
	}
	io.Copy(dest, src)
	return nil
}

func CreateFiles(id string, files []*multipart.FileHeader) (success []string, err error) {
	success = []string{}
	for _, f := range files {
		err = CreateFile(id, *f)
		success = append(success, f.Filename)
		if err != nil {
			return
		}
	}
	return success, nil
}

func autoCreateDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 777)
	}
}
