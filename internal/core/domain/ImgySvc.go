package domain

import "mime/multipart"

type Filey struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}
type Urly struct {
	Url string `json:"url,omitempty" validate:"required"`
}


type MediaResponsey struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}
