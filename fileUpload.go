package bullet

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type fileUpload struct {
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	FileUrl   string `json:"file_url"`
	UploadUrl string `json:"upload_url"`
	DeviceID  string `json:"device_iden"`
}

func newFileUpload(filename, fileType string) fileUpload {
	f := fileUpload{Type: "file", FileName: filename, FileType: fileType}
	return f
}

func (f fileUpload) getReader() (*bytes.Buffer, error) {
	return getReader(f)
}

func getFileRequest(filePath string, upload fileUpload) (*http.Request, error) {
	buffer := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(buffer)

	file, errOpen := os.Open(filePath)
	if errOpen != nil {
		return nil, errOpen
	}
	defer file.Close()

	writer, errWriter := multipartWriter.CreateFormFile("file", upload.FileName)
	if errWriter != nil {
		return nil, errWriter
	}
	_, errCopy := io.Copy(writer, file)
	if errCopy != nil {
		return nil, errCopy
	}

	errMultipartClose := multipartWriter.Close()
	if errMultipartClose != nil {
		return nil, errMultipartClose
	}

	request, errRequest := http.NewRequest(http.MethodPost, upload.UploadUrl, buffer)
	if errRequest != nil {
		return nil, errRequest
	}
	request.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	return request, nil
}
