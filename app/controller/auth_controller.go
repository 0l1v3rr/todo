package controller

import (
	"net/http"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// binding the user from the body
	var user model.User

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid user.",
		})
		return
	}

	// creating the user with the model
	created, err := model.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// return with json
	c.JSON(http.StatusCreated, created)
}

func Login(c *gin.Context) {
	// binding the user from the body
	var user model.User

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide a valid user.",
		})
		return
	}

	// getting the user with the given email from the db
	foundUser, err := model.GetUserByEmail(user.Email)

	// if there is no user with this email
	if foundUser.Id == 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User with this email does not exist.",
		})
		return
	}

	// if the user is not enabled
	if !foundUser.IsEnabled {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "This user is not activated.",
		})
		return
	}

	// if the password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Incorrect password.",
		})
		return
	}

	// successful login
	c.JSON(http.StatusOK, foundUser)
}
