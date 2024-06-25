package storage

import (
	"fmt"
	"github.com/colinmarc/hdfs/v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type FileStorage struct {
	Client *hdfs.Client
	logger *log.Logger
}

func New(logger *log.Logger) (*FileStorage, error) {

	hdfsUri := os.Getenv("HDFS_URI")
	fmt.Println("HDFS_URI:", hdfsUri)

	client, err := hdfs.New(hdfsUri)
	if err != nil {
		logger.Panic(err)
		return nil, err
	}

	// Return storage handler with logger and HDFS client
	return &FileStorage{
		Client: client,
		logger: logger,
	}, nil
}

func (fs *FileStorage) Close() {
	// Close all underlying connections to the HDFS server
	fs.Client.Close()
}

func (fs *FileStorage) CreateDirectories() error {

	err := fs.Client.MkdirAll(hdfsRoot, 0644)
	if err != nil {
		fs.logger.Println(err)
		return err
	}

	return nil
}

func (fs *FileStorage) CreateDirectory(folderName string) error {
	folderPath := path.Join(hdfsRoot, folderName)
	err := fs.Client.MkdirAll(folderPath, 0644)
	if err != nil {
		fs.logger.Printf("Error creating directory %s: %v", folderPath, err)
		return err
	}
	return nil
}

func (fs *FileStorage) SaveImage(folderName, imageName string, imageContent []byte) error {

	folderPath := path.Join(hdfsRoot, folderName)
	imagePath := path.Join(folderPath, imageName)

	if err := fs.CreateDirectory(folderName); err != nil {
		fs.logger.Printf("Error creating directory: %v", err)
		return err
	}

	file, err := fs.Client.Create(imagePath)
	if err != nil {
		fs.logger.Printf("Error creating file %s: %v", imagePath, err)
		return err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fs.logger.Printf("Error closing file: %v", closeErr)
		}
	}()

	if _, err := file.Write(imageContent); err != nil {
		fs.logger.Printf("Error writing image content: %v", err)
		return err
	}

	return nil
}

func (fs *FileStorage) WalkDirectories() []string {
	// Walk all files in HDFS root directory and all subdirectories
	var paths []string
	callbackFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fs.logger.Printf("Directory: %s\n", path)
			path = fmt.Sprintf("Directory: %s\n", path)
			paths = append(paths, path)
		} else {
			fs.logger.Printf("File: %s\n", path)
			path = fmt.Sprintf("File: %s\n", path)
			paths = append(paths, path)
		}
		return nil
	}
	fs.Client.Walk(hdfsRoot, callbackFunc)
	return paths
}

func (fs *FileStorage) GetImageURLS(folderName string) ([]string, error) {
	folderPath := path.Join(hdfsRoot, folderName)
	var imageNames []string

	callbackFunc := func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fs.logger.Println(err)
			return err
		}
		if !info.IsDir() {
			imageName := path.Base(filePath)
			imageNames = append(imageNames, imageName)
		}
		return nil
	}

	err := fs.Client.Walk(folderPath, callbackFunc)
	if err != nil {
		fs.logger.Println(err)
		return nil, err
	}

	return imageNames, nil
}

func (fs *FileStorage) GetImageContent(imagePath string) ([]byte, error) {

	fullPath := path.Join(hdfsRoot, "/", imagePath)

	file, err := fs.Client.Open(fullPath)
	if err != nil {
		fs.logger.Println(err)
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		fs.logger.Println(err)
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return imageData, nil
}
