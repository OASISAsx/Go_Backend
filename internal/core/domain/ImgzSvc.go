package domain

import "mime/multipart"

type Filez struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}
type Urlz struct {
	Url string `json:"url,omitempty" validate:"required"`
}


type MediaResponsez struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}
