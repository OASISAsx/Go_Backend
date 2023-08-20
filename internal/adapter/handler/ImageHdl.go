package handler

import (
	"collection/internal/core/domain"
	"collection/internal/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

func FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		//upload
		formfile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}
		uploadUrl, err := service.NewMediaUpload().FileUpload(domain.File{File: formfile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			domain.MediaResponse{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
func MultiFileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error parsing form data", err)
			return
		}

		files := form.File["files"]
		if len(files) == 0 {
			respondWithError(c, http.StatusBadRequest, "No files uploaded", nil)
			return
		}

		var domainFiles []domain.File
		for _, fileHeader := range files {
			file, err := fileHeader.Open()
			if err != nil {
				respondWithError(c, http.StatusInternalServerError, "Error opening uploaded file", err)
				return
			}
			defer file.Close()

			domainFiles = append(domainFiles, domain.File{File: file})
		}

		// เรียกใช้ MultiFileUpload ใน service เพื่ออัพโหลดไฟล์หลาย ๆ ไฟล์
		uploadUrls, err := service.NewMediaUpload().MultiFileUpload(domainFiles)
		if err != nil {
			respondWithError(c, http.StatusInternalServerError, "Error uploading files", err)
			return
		}

		// ตอบกลับด้วย URL ที่อัพโหลดแล้ว
		respond(c, http.StatusOK, "success", gin.H{"urls": uploadUrls})
	}
}

// ฟังก์ชันสำหรับการตอบกลับ
func respond(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{
		"message": message,
		"data":    data,
	})
}

// ฟังก์ชันสำหรับการตอบกลับแสดงข้อผิดพลาด
func respondWithError(c *gin.Context, statusCode int, errorMessage string, err error) {
	c.JSON(statusCode, gin.H{
		"error":   errorMessage,
		"details": err.Error(),
	})
}



func RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url domain.Url

		//validate the request body
		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				domain.MediaResponse{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := service.NewMediaUpload().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaResponse{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			domain.MediaResponse{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
