package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ctrl AccountRequestController) DeleteAccountRequest(ctx *gin.Context) {
	oid, err := primitive.ObjectIDFromHex(ctx.Query("oid"))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.WriteAccountRequestService.DeleteAccountRequest(oid); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": true})
}

func (ctrl AccountRequestController) DeleteAllAccountRequests(ctx *gin.Context) {

	requestCount, err := ctrl.WriteAccountRequestService.DeleteAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	fullpath := "/var/www/html/react/downloads"

	dir, _ := os.ReadDir(fullpath)

	for i := range dir {
		file := dir[i]
		fileName := file.Name()
		filePath := fullpath + "/" + fileName

		if err := os.Remove(filePath); err != nil {
			ctx.JSON(http.StatusOK, err.Error())
			return
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"request tasks": requestCount, "documents": true})
}
