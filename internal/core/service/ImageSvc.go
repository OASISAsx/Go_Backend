package service

import (
	"collection/internal/core/domain"
	"collection/utils"

	"github.com/go-playground/validator/v10"
	"fmt"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file domain.File) (string, error)
	RemoteUpload(url domain.Url) (string, error)
	MultiFileUpload(files []domain.File) ([]string, error)
}

type media struct{}


func NewMediaUpload() mediaUpload {
	
	return &media{}
	
}


func (*media) FileUpload(file domain.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := utils.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url domain.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := utils.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
func (*media) MultiFileUpload(files []domain.File) ([]string, error) {
	var uploadUrls []string

	for _, file := range files {
		// อัพโหลดแต่ละไฟล์ด้วยฟังก์ชัน FileUpload
		uploadUrl, err := utils.ImageUploadHelper(file.File)
		if err != nil {
			// จัดการข้อผิดพลาดเมื่อเกิดข้อผิดพลาดในการอัพโหลด
			return nil, fmt.Errorf("upload failed: %w", err)
		}
		uploadUrls = append(uploadUrls, uploadUrl)
	}
	return uploadUrls, nil
}
