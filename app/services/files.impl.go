package services

import (
	"context"
	"io"
	"mime/multipart"
	"os"
)

type FileServiceImpl struct {
	ctx context.Context
}

func NewFileService(ctx context.Context) FileService {
	return &FileServiceImpl{
		ctx: ctx,
	}
}

func (srvc FileServiceImpl) UploadFile(file *multipart.File, header *multipart.FileHeader) (*os.File, error) {

	f, err := os.OpenFile(header.Filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic("Error creating file on the filesystem: " + err.Error())
	}

	if _, err := io.Copy(f, *file); err != nil {
		defer f.Close()
		panic("Error during chunk write:" + err.Error())
	}

	return f, nil

}
