package services

import (
	"mime/multipart"
	"os"
)

type FileService interface {
	UploadFile(*multipart.File, *multipart.FileHeader) (*os.File, error)
}
