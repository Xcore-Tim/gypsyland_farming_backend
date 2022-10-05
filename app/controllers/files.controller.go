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
	"github.com/google/uuid"
)

type FileController struct {
	FileService           services.FileService
	AccountRequestService services.WriteAccountRequestService
}

type Input struct {
	Model string `form:"model,omitempty" binding:"required"`
}

type Response struct {
	Success        bool   `json:"success"`
	FileMessage    string `json:"message"`
	AccountMessage string `json:"accountRequestMessage"`
	Name           string `json:"name"`
}

func NewFileController(fileService services.FileService, accountRequestService services.WriteAccountRequestService) FileController {
	return FileController{
		FileService:           fileService,
		AccountRequestService: accountRequestService,
	}
}

func (ctrl FileController) UploadFile(ctx *gin.Context) {
	var f *os.File

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

	var input Input

	input.Model = "static"
	tempDir, err := os.UserHomeDir()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error resolving home directory"})
	}

	if _, err := os.Stat(tempDir + "/" + input.Model); os.IsNotExist(err) {
		if err := os.Mkdir(tempDir+"/"+input.Model, 0777); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error creating temporary directory"})
			return
		}
	}

	if f == nil {
		f, err = os.OpenFile(tempDir+"/"+input.Model+"/"+uploadFile.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating file"})
			return
		}
	}

	if _, err := io.Copy(f, file); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error writing to a file"})
		return
	}

	f.Close()

	response := &Response{
		Success:     true,
		FileMessage: "Keep uploading chunks",
	}

	if rangeMax >= filesize-1 {

		builtFile := tempDir + "/" + input.Model + "/" + uploadFile.Filename

		finalFile, err := os.Open(builtFile)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "failed to upload file"})
			return
		}

		finalFile.Close()

		uploadFileName := ctx.Request.FormValue("fileName")
		fileExt := filepath.Ext(uploadFileName)

		fileName := uuid.NewString()
		newName := tempDir + "/" + input.Model + "/" + fileName + fileExt

		if err := os.Rename(builtFile, newName); err != nil {
			newName = builtFile
		}

		oid := ctx.Query("oid")
		if err := ctrl.AccountRequestService.UpdateDownloadLink(fileName, oid); err != nil {
			response.AccountMessage = "Error updating account request"
		}

		response.Success = true
		response.Name = newName
		response.AccountMessage = "Account request updated successfuly"

		ctx.JSON(http.StatusOK, response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ctrl FileController) DownloadFile(ctx *gin.Context) {

	tempDir, err := os.UserHomeDir()

	if err != nil {
		return
	}

	basePath := tempDir + "/static/"
	fileName := ctx.Query("filename")
	targetPath := filepath.Join(basePath, fileName)

	ctx.JSON(http.StatusOK, targetPath)

	if !strings.HasPrefix(filepath.Clean(targetPath), basePath) {
		ctx.JSON(http.StatusForbidden, "forbidden")
		return
	}

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(targetPath)
}

func (ctrl FileController) RegisterUserRoutes(rg *gin.RouterGroup) {
	fileGroup := rg.Group("/file")
	fileGroup.POST("/upload", ctrl.UploadFile)
	fileGroup.POST("/download", ctrl.DownloadFile)
}
