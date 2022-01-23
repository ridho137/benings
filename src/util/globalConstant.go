package util

import (
	"benings/model"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GenerateOtp(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// func Key() []byte {
// 	return []byte("")
// }

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

var DefaultOutErrorFailed = 1

const MaxFileSize = 3145728

func UploadHandler(w http.ResponseWriter, r *http.Request, photoPath string) model.UpdateMemberPhotoResponse {
	var response model.UpdateMemberPhotoResponse
	switch r.Method {
	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		log.Println(r.ContentLength)
		if r.ContentLength > MaxFileSize {
			response.OutError = 1
			response.OutMessage = "File size to longger"
			return response
		}
		r.Body = http.MaxBytesReader(w, r.Body, MaxFileSize)
		err := r.ParseMultipartForm(1024)
		if err != nil {
			response.OutError = 1
			response.OutMessage = err.Error()
			return response
		}
		//get a ref to the parsed multipart form
		m := r.MultipartForm
		//get the *fileheaders
		files := m.File["photoFile"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				return response
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create(photoPath)
			defer dst.Close()
			if err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				return response
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				response.OutError = 1
				response.OutMessage = err.Error()
				return response
			}
			fmt.Println("File " + photoPath + " successfully Updated")
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return response
}

func DeleteHandler(photoPath string) {
	err := os.Remove(photoPath)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File " + photoPath + " successfully deleted")
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
