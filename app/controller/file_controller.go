package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/0l1v3rr/todo/app/util"
	"github.com/gin-gonic/gin"
)

// @Summary      Upload file
// @Description  Uploads a new file into the images/ folder
// @Tags         File endpoints
// @Produce      json
// @Success      201  {object}  util.Success
// @Failure      400  {object}  util.Error "If the file is not valid."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Failure      500  {object}  util.Error "If there was a file error."
// @Router       /files [post]
func UploadFile(c *gin.Context) {
	// checking whether the user is logged in
	_, err := model.GetLoggedInUser(*c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in."})
		return
	}

	// getting the file from the request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please provide a file!"})
		return
	}

	// filename
	filename := fmt.Sprintf("%s-%s", util.GenerateHash(16), header.Filename)

	// creating the file
	out, err := os.Create("images/" + filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: "Failed to upload the file!"})
		return
	}
	defer out.Close()

	// copying the file
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: "Failed to copy the file!"})
		return
	}

	// filepath
	filepath := fmt.Sprintf("/assets/images/%s", filename)

	// success file upload
	c.JSON(http.StatusCreated, gin.H{"filepath": filepath})
}
