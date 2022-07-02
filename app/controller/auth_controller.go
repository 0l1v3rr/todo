package controller

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/0l1v3rr/todo/app/model"
	"github.com/0l1v3rr/todo/app/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Registration
// @Description  Registers a new user into the database.
// @Tags         User endpoints
// @Accept       json
// @Produce      json
// @Param 		 user body model.User false "User to register"
// @Success      201  {object}  model.User "If the user has been created successfully."
// @Failure      400  {object}  util.Error "If the provided user is not valid."
// @Failure      409  {object}  util.Error "If the specified email already exists."
// @Failure      500  {object}  util.Error "If there was a server error while creating the user."
// @Router       /register [post]
func Register(c *gin.Context) {
	// binding the user from the body
	var user model.User

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please provide a valid user."})
		return
	}

	// validating the user
	ok, msg := user.Validate()
	if !ok {
		c.JSON(http.StatusBadRequest, util.Error{Message: msg})
		return
	}

	// checking if the email is already in the db
	exists := model.ExistsByEmail(user.Email)
	if exists {
		c.JSON(http.StatusConflict, util.Error{Message: "This email is already registered."})
	}

	// creating the user with the model
	created, err := model.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: err.Error()})
		return
	}

	// return with json
	c.JSON(http.StatusCreated, created)
}

// @Summary      Login
// @Description  Logs in a user and saves the JWT in cookies.
// @Tags         User endpoints
// @Accept       json
// @Produce      json
// @Param 		 user body model.LoginUser false "User to log in"
// @Success      200  {object}  util.Success "If the login was successful."
// @Failure      400  {object}  util.Error "If the provided user is not valid."
// @Failure      404  {object}  util.Error "If the user with the specified email does not exist."
// @Failure      403  {object}  util.Error "If the password is incorrect or the user is not activated."
// @Router       /login [post]
func Login(c *gin.Context) {
	// binding the user from the body
	var user model.LoginUser

	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, util.Error{Message: "Please provide a valid user."})
		return
	}

	// getting the user with the given email from the db
	foundUser, err := model.GetUserByEmail(user.Email)

	// if there is no user with this email
	if foundUser.Id == 0 || err != nil {
		c.JSON(http.StatusNotFound, util.Error{Message: "User with this email does not exist."})
		return
	}

	// if the user is not enabled
	if !foundUser.IsEnabled {
		c.JSON(http.StatusForbidden, util.Error{Message: "This user is not activated."})
		return
	}

	// if the password is incorrect
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusForbidden, util.Error{Message: "Incorrect password."})
		return
	}

	// creating the jwt claims
	// it expires in 30 days
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(foundUser.Id),
		ExpiresAt: time.Now().Add((time.Hour * 24) * 30).Unix(), // 30 days
	})

	// creating the token from the claims
	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.Error{Message: "Failed to log in."})
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
	c.JSON(http.StatusOK, util.Success{Message: "Successful login!"})
}

// @Summary      Logged In User
// @Description  Returns the currently logged-in user.
// @Tags         User endpoints
// @Produce      json
// @Success      200  {object}  model.User "If the user is logged in."
// @Failure      401  {object}  util.Error "If the user is not logged in."
// @Router       /user [get]
func GetLoggedInUser(c *gin.Context) {
	// getting the logged in user
	user, err := model.GetLoggedInUser(*c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, util.Error{Message: "You are not logged in!"})
		return
	}

	// if the user is logged in
	c.JSON(http.StatusOK, user)
}

// @Summary      Logout
// @Description  Logs out the currently logged-in user.
// @Tags         User endpoints
// @Produce      json
// @Success      200  {object}  util.Success "If the logout was success."
// @Router       /logout [post]
func Logout(c *gin.Context) {
	// removing the cookie
	c.SetCookie(
		"jwt",
		"",
		-3600,
		"/",
		"localhost",
		false,
		true,
	)

	// json response
	c.JSON(http.StatusOK, util.Success{Message: "Successful logout!"})
}
