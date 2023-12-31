package middleware

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(next gin.HandlerFunc) gin.HandlerFunc {
	return func (c *gin.Context)  {
		file, err := c.FormFile("image")
		if err != nil{
			c.JSON(http.StatusBadRequest, err)
			return
		}
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		defer src.Close()

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		defer tempFile.Close()

		if _, err := io.Copy(tempFile, src); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		data := tempFile.Name()
		fmt.Println("sadasd", data)
		
		filename := data[8:] // split uploads/		
		fmt.Println(filename)

		c.Set("dataFile", filename)
		next(c)
	}
}