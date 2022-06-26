package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/0l1v3rr/todo/app/data"
	"github.com/0l1v3rr/todo/app/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
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

	// creating the jwt claims
	// it expires in 30 days
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(foundUser.Id),
		ExpiresAt: time.Now().Add((time.Hour * 24) * 30).Unix(), // 30 days
	})

	// creating the token from the claims
	token, err := claims.SignedString([]byte(data.Env["JWT_SECRET"]))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Log In. Try again later.",
		})
		return
	}

	// creating the cookie for the token
	c.SetCookie(
		"jwt",
		token,
		3600*24*30, // 30 days
		"/",
		"localhost",
		false,
		true,
	)

	// successful login
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful login!",
	})
}

func GetLoggedInUser(c *gin.Context) {
	// getting the cookie from the request
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not logged in.",
		})
		return
	}

	// parsing the token from the cookie
	token, err := jwt.ParseWithClaims(
		cookie.Value,
		&jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(data.Env["JWT_SECRET"]), nil
		},
	)

	// error handling
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not logged in.",
		})
		return
	}

	// getting the claims from the token
	claims := token.Claims.(*jwt.StandardClaims)

	// getting the user from the db
	user, err := model.GetUserByStringId(claims.Issuer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// if the user is logged in
	c.JSON(http.StatusOK, user)
}

func Logout(c *gin.Context) {
	// removing the cookie
	c.SetCookie(
		"jwt",
		"",
		-3600, // 30 days
		"/",
		"localhost",
		false,
		true,
	)

	// json response
	c.JSON(http.StatusOK, gin.H{
		"message": "Successful logout!",
	})
}
