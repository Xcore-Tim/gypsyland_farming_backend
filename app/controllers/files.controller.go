package controllers

import (
	"gypsyland_farming/app/services"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	FileService services.FileService
}

type Response struct {
	Success        bool   `json:"success"`
	FileMessage    string `json:"message"`
	AccountMessage string `json:"accountRequestMessage"`
	Name           string `json:"name"`
}

func NewFileController(fileService services.FileService) FileController {
	return FileController{
		FileService: fileService,
	}
}

func (ctrl FileController) UploadFile(ctx *gin.Context) {

	var f *os.File

	response := &Response{
		Success:     false,
		FileMessage: "Keep uploading chunks",
	}

	file, uploadFile, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contentRangeHeader := ctx.Request.Header.Get("C-Range")
	rangeAndSize := strings.Split(contentRangeHeader, "/")
	rangeParts := strings.Split(rangeAndSize[0], "-")

	rangeMax, err := strconv.Atoi(rangeParts[1])

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing range in Content-Range header"})
		return
	}

	filesize, err := strconv.Atoi(rangeAndSize[1])

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing file size in Content-Range header"})
		return
	}

	if filesize > 1024*1024*1024 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file size must be less than 100 MB"})
		return
	}

	basePath := ctrl.FileService.ReturnBasePath()

	if _, err := os.Stat(basePath); os.IsNotExist(err) {
		if err := os.Mkdir(basePath, 0777); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error creating temporary directory"})
			return
		}
	}

	if f == nil {
		f, err = os.OpenFile(basePath+"/"+uploadFile.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	if _, err := io.Copy(f, file); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error writing to a file"})
		return
	}

	f.Close()

	response.Success = true

	if rangeMax >= filesize-1 {

		fileExt := filepath.Ext(ctx.Request.FormValue("fileName"))

		fileName, filePath, err := ctrl.FileService.CreateNewFile(uploadFile, fileExt)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		oid := ctx.Query("oid")

		oldFile, isFound, err := ctrl.FileService.CheckPreviousFile(oid)

		if err != nil {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		}

		if isFound {
			if err := ctrl.FileService.DeletePreviousFile(oldFile); err != nil {
				ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			}
		}

		if err := ctrl.FileService.UpdateDownloadLink(fileName, oid); err != nil {
			response.AccountMessage = "Error updating account request"
		}

		response.Success = true
		response.Name = filePath
		response.FileMessage = "File uploaded successfuly"
		response.AccountMessage = "Account request updated successfuly"

		ctx.JSON(http.StatusOK, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl FileController) DownloadFile(ctx *gin.Context) {

	fileName := ctx.Query("filename")

	downloadPath := ctrl.FileService.ReturnDownloadPath(fileName)

	ctx.String(http.StatusOK, downloadPath)

}

func (ctrl FileController) CheckPreviousFile(ctx *gin.Context) {
	oid := ctx.Query("oid")

	oldFile, isFound, err := ctrl.FileService.CheckPreviousFile(oid)
	host, _ := os.Hostname()
	ctx.JSON(http.StatusAccepted, gin.H{"host": host, "filepath": oldFile, "found": isFound, "err": err})
	ctx.JSON(http.StatusAccepted, host+"/"+oldFile)
}

func (ctrl FileController) DeleteFilesInDir(ctx *gin.Context) {

	homeDir, _ := os.UserHomeDir()
	fullpath := homeDir + "/static/"
	dir, _ := os.ReadDir(fullpath)

	for i := range dir {
		file := dir[i]
		fileName := file.Name()
		filePath := fullpath + fileName

		os.Remove(filePath)

	}

	ctx.JSON(http.StatusOK, "success")
}

func (ctrl FileController) RegisterUserRoutes(rg *gin.RouterGroup) {
	fileGroup := rg.Group("/file")
	fileGroup.POST("/upload", ctrl.UploadFile)
	fileGroup.POST("/download", ctrl.DownloadFile)
	fileGroup.POST("/check", ctrl.CheckPreviousFile)
	fileGroup.POST("/dir", ctrl.DeleteFilesInDir)
}
