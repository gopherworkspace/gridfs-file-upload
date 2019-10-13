package controllers

import (
	"net/http"
	"fmt"
	"github.com/gridfs-upload/backend/golang/services"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"strconv"
	"io"
)

const (
	mongoUrl        = "localhost:27017"
	Database        = "GridFsDB"
	signupCollecion = "FileDump"
)

//Controller ...
type Controller struct {
	Repository services.Service
}

type Response struct {
	FileId   interface{}
	FileName string
	Code     int
	Success  bool
	Message  string
}

// Login Page GET /

func (c *Controller) Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadGateway)
		w.Header().Set("Content-Type", "application/json")
		respErr, _ := json.Marshal(Response{Message: "File is empty", Code: 500})
		w.Write(respErr)
		return
	}

	if header.Filename == "" {
		w.WriteHeader(http.StatusBadGateway)
		w.Header().Set("Content-Type", "application/json")
		respErr, _ := json.Marshal(Response{Message: "File is empty", Code: 500})
		w.Write(respErr)
		return

	}

	fileData := services.Service.UploadService(services.Service{}, header, file)

	resp := Response{
		FileId:   fileData.Id(),
		FileName: fileData.Name(),
		Code:     200,
		Success:  true,
		Message:  "File uploaded Successfully",
	}
	respJson, err := json.Marshal(resp)

	if err != nil {
		fmt.Println(respJson)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJson)
	return
}

func (c *Controller) Download(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	fileId := r.Form["id"]
	fileName := r.Form["fileName"]

	fmt.Println("============================================")
	fmt.Println("file name :", fileName)
	fmt.Println("file Id :", fileId)
	fmt.Println("============================================")

	session, err := mgo.Dial("127.0.0.1")

	if (err != nil) {
		fmt.Println("err :", err)
	}

	defer session.Close()

	fileDb := session.DB(Database)

	file, err := fileDb.GridFS(signupCollecion).Open(r.Form["fileName"][0])


	fileHeader := make([]byte, 512)
	file.Read(fileHeader)
	fileContentType := http.DetectContentType(fileHeader)
	fileSize := file.Size()

	w.Header().Set("Content-Disposition", "attachment; filename="+r.Form["fileName"][0])
	w.Header().Set("Content-Type", fileContentType)
	w.Header().Set("Content-Length", strconv.FormatInt(fileSize, 10))

	file.Seek(0, 0)
	io.Copy(w, file)

	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
