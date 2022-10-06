package services

import "mime/multipart"

type FileService interface {
	CreateNewFile(*multipart.FileHeader, string) (string, string, error)
	CheckPreviousFile(string) (string, bool, error)
	DeletePreviousFile(string) error
	UpdateDownloadLink(string, string) error
	ReturnBasePath() string
	ReturnDownloadPath(string) string
	DeleteFilesInDir() error
}
