package v1

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
)

const (
	MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB
	path            = "./images/"
	name            = "hgthl5fys"
	api_key         = "534124638472856"
	api_secret      = "sixUMqMtBqLSDrP4LjG7YZwTLto"
	folderName      = "images/"
)

func GetCloudinary() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(name, api_key, api_secret)
	if err != nil {
		return nil, err
	}
	return cld, nil
}

// todo :refactor/ decopmpose -> to service side || pkg

// Upload photo godoc
// @Description  Upload images service, recieve  key(image) : multipartFormFile
// @Tags         Machine
// @Produce      json
// @Accept       multipart/form-data
// @Param        input  body  string  true "key : image, value : file"
// @Param        input  path  string  true "lastCarID"
// @Security BearerAuth
// @Failure      400,500  {object}  models.Response
// @Success      200      {object}  models.Response
// @Router       /v1/machine/upload/:lastCarId [post]
func (h *Handler) Upload(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		log.Println("req, parseMultipart", err)
		return
	}
	// Get a reference to the fileHeaders.
	// They are accessible only after ParseMultipartForm is called
	files := c.Request.MultipartForm.File["file"]

	machine_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}
	listSrc := []string{}

	for _, fileHeader := range files {
		// Restrict the size of each uploaded file to 1MB.
		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(c.Writer, fmt.Sprintf("The uploaded image is too big: %s. Please use an image less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}

		filetype := http.DetectContentType(buff)
		if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" {
			http.Error(c.Writer, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		cld, err := GetCloudinary()
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx := context.Background()

		uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
		if err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Writer.Header().Set("Content-Type", "application/json")
		c.AbortWithStatus(200)

		// json.NewEncoder(c.Writer).Encode(uploadResult.SecureURL)

		// err = os.MkdirAll("./images", os.ModePerm)
		// if err != nil {
		//  http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		//  return
		// }
		// // f, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
		// src := fmt.Sprintf("./images/%s", fileHeader.Filename)

		// f, err := os.Create(src)
		// if err != nil {
		//  http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		//  return
		// }

		// defer f.Close()

		// _, err = io.Copy(f, file)
		// if err != nil {
		//  http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		//  return
		// }
		listSrc = append(listSrc, uploadResult.SecureURL)

	}
	// save in db src
	err = h.useCases.FileManagerUseCaseInterface.CreateSrc(listSrc, machine_id)
	if err != nil {
		log.Println("upload db ", err)
		responseWithStatus(c, http.StatusInternalServerError, "", "Failed", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success uploaded images", "OK", nil)
}
