package main

import (
	"log"
	"os"
	"path"
)

type Service interface {
	UploadFile(path string)
}

type ServiceImpl struct{}

func NewServiceImpl() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) UploadFile(filename string) {
	log.Println("uploading file...")
	file, err := os.Open(filename)
	if err != nil {
		log.Println("failedto locate file : %s", filename)
	}

	defer file.Close()

	fileInfo, err := file.Stat()

	log.Printf("Mode : %v", fileInfo.Mode())
	if err != nil {
		log.Printf("error occurred file fetch file stats. error : %v", err)
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		log.Print(err)
	}

	log.Printf("bytes read : %v", bytesread)
	log.Printf("bytes to string : %s", string(buffer))
	log.Printf("File type : %s", path.Ext(filename))
}
