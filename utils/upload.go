package utils

import (
	"encoding/json"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	imagePath = "http://localhost:10000/images"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	unixFileName := handler.Filename + cast.ToString(time.Now().Unix())
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	f, err := os.OpenFile("./uploads/"+ unixFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	io.Copy(f, file)
	json.NewEncoder(w).Encode(map[string]string{
		"image_url": imagePath + cast.ToString(time.Now().Unix()) + unixFileName,
	})
}
