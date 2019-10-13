package services

import (
	"io/ioutil"
	"gopkg.in/mgo.v2"
	"fmt"
	"mime/multipart"
)

type Service struct{}

const (
	mongoUrl        = "127.0.0.1"
	Database        = "GridFsDB"
	signupCollecion = "FileDump"
	CHUNK_SIZE      = 16000000
)

func (r Service) UploadService(File *multipart.FileHeader, file1 multipart.File) *mgo.GridFile {

	session, err := mgo.Dial(mongoUrl)
	check(err)
	defer session.Close()
	db := session.DB(Database)
	file, err := db.GridFS(signupCollecion).Create(File.Filename)
	check(err)
	//default chunk size is 256 KB that is so less so I configured chunk size
	// configure chunk size is 16 MB
	file.SetChunkSize(CHUNK_SIZE)
	messages, err := ioutil.ReadAll(file1)
	check(err)
	_, err = file.Write(messages)
	check(err)
	check(err)
	err = file.Close()
	check(err)
	return file
}

func check(err error) {
	if err != nil {
		fmt.Println("err--->", err)
	}
}
