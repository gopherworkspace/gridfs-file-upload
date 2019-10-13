package models

import (
	"mime/multipart"
)

type MultiPartData struct {
	File multipart.File `json:"file" bson:"file"`
}

type FileData struct {
	FileName      string         `json:"file_name" bson:"file_name"`
	FileExtension string         `json:"file extension" bson:"file extension"`
	File          multipart.File `json:"file" bson:"file"`
}

type Document struct {
	FileId string `json:"file_id"`
}

type Response struct {
	FileId   interface{}
	FileName string
	Code     int
	Success  bool
	Message  string
}
