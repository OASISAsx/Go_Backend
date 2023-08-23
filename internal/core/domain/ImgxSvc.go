package domain

import "mime/multipart"

type Filex struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}
type Urlx struct {
	Url string `json:"url,omitempty" validate:"required"`
}


type MediaResponsex struct {
	StatusCode int                    `json:"statusCode"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}
