package util

import (
	"io"
	"mime/multipart"
	"os"
)

var dest = "C:\\tmp\\upload\\"

func CreateFile(file multipart.FileHeader) (err error) {
	src, err := file.Open()
	defer src.Close()
	if err != nil {
		return
	}
	dest, err := os.Create(dest + file.Filename)
	defer dest.Close()
	if err != nil {
		return
	}
	io.Copy(dest, src)
	return nil
}

func CreateFiles(files []*multipart.FileHeader) (success []string, err error) {
	success = []string{}
	for _, f := range files {
		err = CreateFile(*f)
		success = append(success, f.Filename)
		if err != nil {
			return
		}
	}
	return success, nil
}
