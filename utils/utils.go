package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
	"uploadCSV/models"
)

func EncodeJSON200Body(resp http.ResponseWriter, data interface{}) {
	err := json.NewEncoder(resp).Encode(data)
	if err != nil {
		logrus.Errorf("Error encoding response %v", err)
	}
}

func EncodeJSONBody(resp http.ResponseWriter, statusCode int, data interface{}) {
	logrus.Info("EncodeJSONBody: in this function with this data: ", data)
	resp.WriteHeader(statusCode)
	if err := json.NewEncoder(resp).Encode(data); err != nil {
		logrus.Error(err)
	}
	err := json.NewEncoder(resp).Encode(data)
	if err != nil {
		logrus.Errorf("Error encoding response %v", err)
	}
}

func DownLoadFileFromURL(uploadImage models.UploadImage) (string, string, error) {
	response, err := http.Get(uploadImage.URL)
	if err != nil {
		return "", "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if response.StatusCode != 200 {
		return "", "", errors.New("received non 200 response code")
	}

	fileName := fmt.Sprintf("%v-%v.jpg", uploadImage.Name, time.Now().Unix())
	file, err := os.Create(fileName)
	if err != nil {
		return "", "", err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	contentType := response.Header.Get("Content-Type")
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", contentType, err
	}

	return fileName, contentType, err
}
