package v1

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 // 10MB
	path            = "./images/"
)

// todo :refactor/ decopmpose
// else try - base64  https://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang

// helper
func listFile(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func (h *Handler) Download(c *gin.Context) {
	machine_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}
	// find all images - by MachineID
	files := listFile(path)

	// result := []*os.File{}
	// result := [][]byte{}
	// c.Writer.Header().Set("Content-Type", "application/octet-stream") // image/*
	c.Writer.Header().Set("Content-Type", "application/octet-stream") // image/*

	for idx, file := range files {
		if file.Name() == fmt.Sprintf("%d_%d%s", machine_id, idx, filepath.Ext(file.Name())) {
			// add in pull images, send to client
			fileBytes, err := ioutil.ReadFile(path + file.Name())
			if err != nil {
				panic(err)
			}
			// result = append(result, fileBytes)
			c.Writer.Write(fileBytes)
			// defer img.Close()
			// img, err := os.Open(path + file.Name())
			// result = append(result, img)
		}
	}

	// io.Copy(c.Writer, result[1])
	c.Status(200)
	// todo :DetectContentType, customFunc ?
}

func (h *Handler) Upload(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("req, parseMultipart", err)
		return
	}
	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := c.Request.MultipartForm.File["image"]
	// async; create machine -> proccessing photo  ? js
	// setPhoto - user_id, machine_id

	machine_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	for idx, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		// To prevent the aggregate size from exceeding
		// a specified value, use the http.MaxBytesReader() method
		// before calling ParseMultipartForm()
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			log.Println("max size, large file")
			// http.Error(w, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			log.Println("open file err", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			log.Println("file read ", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpeg" {
			log.Println("file type err  ", err)
			// http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			log.Println("file seek err  ", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("./images", os.ModePerm)
		if err != nil {
			log.Println("mkdir create  ", err)
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		f, err := os.Create(fmt.Sprintf("./images/%d_%d%s", machine_id, idx, filepath.Ext(fileHeader.Filename)))
		if err != nil {
			log.Println("os create  file ", err)
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			log.Println("copy ", err)
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	responseWithStatus(c, http.StatusOK, "success uploaded images", "OK", nil)
}
