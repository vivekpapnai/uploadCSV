package server

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"uploadCSV/models"
	"uploadCSV/utils"
)

func (srv *Server) greet(resp http.ResponseWriter, req *http.Request) {
	utils.EncodeJSON200Body(resp, map[string]interface{}{
		"message": "welcome to image downloading service",
	})
}

func (srv *Server) uploadCSV(resp http.ResponseWriter, req *http.Request) {
	// multipart form data
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		utils.EncodeJSONBody(resp, http.StatusBadRequest, map[string]interface{}{
			"message": "file size should not be greater than 10 mb",
		})
		return
	}

	file, header, err := req.FormFile("imagesCsv")
	if err != nil {
		utils.EncodeJSONBody(resp, http.StatusBadRequest, map[string]interface{}{
			"message": "error data retrieving",
		})
		return
	}

	var email string

	email = req.FormValue("email")

	fmt.Println(email)
	typeOfUpload := "csvFiles"
	filePath := fmt.Sprintf(`%v/%v-%s`, typeOfUpload, time.Now().Unix(), header.Filename)
	url, err := srv.StorageProvider.Upload(req.Context(), file, filePath, "application/octet-stream")
	if err != nil {
		utils.EncodeJSONBody(resp, http.StatusInternalServerError, err)
		logrus.Errorf("uploadCSV: error in uploading csv: %v", err)
		return
	}

	type publishCSVFile struct {
		URL   string `json:"data"`
		Email string `json:"email"`
	}

	publishCSVFileData := publishCSVFile{URL: url, Email: email}

	csvFileMetaData, err := json.Marshal(&publishCSVFileData)
	if err != nil {
		utils.EncodeJSONBody(resp, http.StatusInternalServerError, err)
		logrus.Errorf("uplaodCSV: error in marshalling metadata: %v", err)
		return
	}

	srv.KafkaProvider.Publish(models.TopicCSVFileUpload, csvFileMetaData)

	utils.EncodeJSON200Body(resp, map[string]interface{}{
		"message": "success",
	})
}

//func (srv *Server) upload(resp http.ResponseWriter, req *http.Request) {
//
//	defer func(req *http.Request) {
//		if req.MultipartForm != nil { // prevent panic from nil pointer
//			if err := req.MultipartForm.RemoveAll(); err != nil {
//				logrus.Errorf("Unable to remove all multipart form. %+v", err)
//			}
//		}
//	}(req)
//
//	req.Body = http.MaxBytesReader(resp, req.Body, 51<<20)
//
//	if err := req.ParseMultipartForm(51 << 20); err != nil {
//		if err == io.EOF || err.Error() == "multipart: NextPart: unexpected EOF" {
//			logrus.Warn("EOF")
//		} else {
//			logrus.Errorf("[ParseMultipartForm] %s", err.Error())
//		}
//		return
//	}
//
//	file, header, err := req.FormFile("file")
//
//	defer func() {
//		if err = file.Close(); err != nil {
//			logrus.Errorf("Unable to close file multipart form. %+v", err)
//		}
//	}()
//
//	if err != nil {
//		if err == io.EOF {
//			logrus.Warn("EOF")
//		} else {
//			logrus.Error(err)
//		}
//		return
//	}
//
//	typeOfUpload := "image"
//	filePath := fmt.Sprintf(`images/%v/%v-%s`, typeOfUpload, time.Now().Unix(), header.Filename)
//
//	url, err := srv.StorageProvider.Upload(req.Context(), file, filePath, "application/octet-stream")
//	if err != nil {
//		utils.EncodeJSONBody(resp, http.StatusInternalServerError, "error in uploading image")
//		return
//	}
//
//	utils.EncodeJSON200Body(resp, map[string]interface{}{
//		"url": url,
//	})
//}
