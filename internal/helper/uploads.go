package helper

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func UploadFile(file *multipart.FileHeader, uploadsFolder string) (string, error) {
	// Pastikan folder uploads ada
	if _, err := os.Stat(uploadsFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadsFolder, os.ModePerm); err != nil {
			return "", errors.New("failed to create uploads folder")
		}
	}

	// Buat path untuk file yang akan disimpan
	savePath := filepath.Join(uploadsFolder, file.Filename)

	// Buka file yang diunggah
	src, err := file.Open()
	if err != nil {
		return "", errors.New("failed to open uploaded file")
	}
	defer src.Close()

	// Buat file baru di server untuk menyimpan data
	dst, err := os.Create(savePath)
	if err != nil {
		return "", errors.New("failed to create file on server")
	}
	defer dst.Close()

	// Salin konten dari file yang diunggah ke file baru di server
	if _, err := io.Copy(dst, src); err != nil {
		return "", errors.New("failed to copy file content")
	}

	return savePath, nil
}

func UploadFile2(file *multipart.FileHeader, uploadsFolder string) (string, error) {
	// Pastikan folder uploads ada
	if _, err := os.Stat(uploadsFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadsFolder, os.ModePerm); err != nil {
			return "", errors.New("failed to create uploads folder")
		}
	}

	// Buat path untuk file yang akan disimpan
	savePath := filepath.Join(uploadsFolder, file.Filename)

	// Buka file yang diunggah
	src, err := file.Open()
	if err != nil {
		return "", errors.New("failed to open uploaded file")
	}
	defer src.Close()

	// Buat file baru di server untuk menyimpan data
	dst, err := os.Create(savePath)
	if err != nil {
		return "", errors.New("failed to create file on server")
	}
	defer dst.Close()

	// Salin konten dari file yang diunggah ke file baru di server
	if _, err := io.Copy(dst, src); err != nil {
		return "", errors.New("failed to copy file content")
	}

	return savePath, nil
}