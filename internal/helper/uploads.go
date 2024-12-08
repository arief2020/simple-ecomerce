package helper

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func UploadFile(file *multipart.FileHeader, uploadsFolder string) (string, error) {
	if _, err := os.Stat(uploadsFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadsFolder, os.ModePerm); err != nil {
			return "", errors.New("failed to create uploads folder")
		}
	}

	savePath := filepath.Join(uploadsFolder, fmt.Sprintf("%d-%s", time.Now().UnixNano(), file.Filename))

	src, err := file.Open()
	if err != nil {
		return "", errors.New("failed to open uploaded file")
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return "", errors.New("failed to create file on server")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", errors.New("failed to copy file content")
	}

	return savePath, nil
}
