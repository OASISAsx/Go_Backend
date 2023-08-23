package service

import (
	"collection/internal/core/domain"
	"collection/utils"

	"github.com/go-playground/validator/v10"
	
)

var (
	validatec = validator.New()
)


type mediaUploadz interface {
	FileUploadx(file domain.File) (string, error)
	RemoteUpload(url domain.Url) (string, error)
	
}

type mediaz struct{}


func NewMediaUploadz() mediaUpload {
	
	return &media{}
	
}


func (*media) FileUploadz(file domain.File) (string, error) {
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

func (*media) RemoteUploadz(url domain.Url) (string, error) {
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
