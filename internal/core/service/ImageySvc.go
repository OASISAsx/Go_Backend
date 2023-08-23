package service

import (
	"collection/internal/core/domain"
	"collection/utils"
	"github.com/go-playground/validator/v10"
	
)

var (
	validateb = validator.New()
)

type mediaUploady interface {
	FileUploadx(file domain.File) (string, error)
	RemoteUpload(url domain.Url) (string, error)

}

type mediay struct{}


func NewMediaUploady() mediaUpload {
	
	return &media{}
	
}


func (*media) FileUploady(file domain.File) (string, error) {
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

func (*media) RemoteUploady(url domain.Url) (string, error) {
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
