package services

import (
	"context"
	"errors"
	accounts "gypsylandFarming/app/models/accounts"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileServiceImpl struct {
	accountRequestTaskCollection *mongo.Collection
	ctx                          context.Context
}

func NewFileService(accountRequestTaskCollection *mongo.Collection, ctx context.Context) FileService {
	return &FileServiceImpl{
		accountRequestTaskCollection: accountRequestTaskCollection,
		ctx:                          ctx,
	}
}

func (srvc FileServiceImpl) CreateNewFile(uploadFile *multipart.FileHeader, fileExt string) (string, string, error) {

	basePath := srvc.ReturnBasePath()

	constructedFile := basePath + "/" + uploadFile.Filename

	finalFile, err := os.Open(constructedFile)

	if err != nil {
		return "", "", errors.New("failed to upload file")
	}

	finalFile.Close()

	fileName := uuid.NewString() + fileExt

	filePath := basePath + "/" + fileName

	if err := os.Rename(constructedFile, filePath); err != nil {
		filePath = constructedFile
	}

	return fileName, filePath, err
}

func (srvc FileServiceImpl) CheckPreviousFile(oid string) (string, bool, error) {

	requestID, err := primitive.ObjectIDFromHex(oid)

	if err != nil {
		return "", false, err
	}

	var accountRequest accounts.AccountRequestTask

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}

	if err := srvc.accountRequestTaskCollection.FindOne(srvc.ctx, filter).Decode(&accountRequest); err != nil {
		return "", false, err
	}

	if accountRequest.DownloadLink != "" {
		basePath := srvc.ReturnBasePath()

		result := basePath + "/" + accountRequest.FileName
		return result, true, nil
	}

	return "", false, err

}

func (srvc FileServiceImpl) UpdateDownloadLink(fileName, oid string) error {

	requestID, err := primitive.ObjectIDFromHex(oid)

	if err != nil {
		return err
	}

	downloadLink := srvc.ReturnDownloadPath(fileName)

	filter := bson.D{bson.E{Key: "_id", Value: requestID}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "downloadLink", Value: downloadLink},
		bson.E{Key: "fileName", Value: fileName},
	}}}

	result := srvc.accountRequestTaskCollection.FindOneAndUpdate(srvc.ctx, filter, update)

	if result.Err() != nil {
		return errors.New("no matched documents found for update")
	}

	return nil
}

func (srvc FileServiceImpl) DeletePreviousFile(filePath string) error {
	err := os.Remove(filePath)
	return err
}

func (srvc FileServiceImpl) DeleteFilesInDir() error {

	return nil
}

func (srvc FileServiceImpl) ReturnBasePath() string {

	basepath := "/var/www/html/react/downloads"

	return basepath
}

func (srvc FileServiceImpl) ReturnDownloadPath(fileName string) string {

	basepath := "http://84.201.138.62/downloads/"
	downloadPath := basepath + fileName

	return downloadPath
}
