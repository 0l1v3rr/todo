package model

import (
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// defining a model struct
type User struct {
	Id        int    `json:"id" gorm:"primaryKey" example:"1"`
	Name      string `json:"name" gorm:"not null" example:"John Doe"`
	Email     string `json:"email" gorm:"not null;unique" example:"johndoe@gmail.com"`
	Password  string `json:"password,omitempty" gorm:"not null;column:password" example:"secret"`
	IsEnabled bool   `json:"isEnabled" gorm:"not null;column:is_enabled" example:"true"`
}

// defining a LoginUser for the documentation
type LoginUser struct {
	Email    string `json:"email" example:"johndoe@gmail.com"`
	Password string `json:"password" example:"SuperSecret69"`
}

func (user User) Validate() (bool, string) {
	// if the length of the username is less than 6 characters
	if len(user.Name) < 6 {
		return false, "Your name has to be at least 6 characters long."
	}

	// if the length of the username is more than 64 characters
	if len(user.Name) > 64 {
		return false, "The length of your name should be maximum of 64 characters long."
	}

	// creating the regexp for the email validation
	r, _ := regexp.Compile("^[a-zA-Z0-9+_.-]+@[a-zA-Z0-9.-]+$")

	// validating the email
	if !r.MatchString(user.Email) {
		return false, "Please provide a valid email address!"
	}

	// if the user is valid
	return true, ""
}

func ExistsByEmail(email string) bool {
	// getting the user from the db
	user, err := GetUserByEmail(email)

	// if there's an error, the user does not exist
	if err != nil {
		return false
	}

	// if the id is 0, the user does not exist
	if user.Id == 0 {
		return false
	}

	// the user does exist
	return true
}

func Register(user User) (User, error) {
	// encrypting the password with bcrypt
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	// overriding the values
	user.IsEnabled = true
	user.Password = string(encrypted)

	// creating the user
	tx := DB.Create(&user)
	return user, tx.Error
}

// if the specified id is an int
func GetUserById(id int) (User, error) {
	// getting the user form the db by id
	var user User
	tx := DB.Where("id = ?", id).First(&user)
	return user, tx.Error
}

// if the specified id is a string
func GetUserByStringId(id string) (User, error) {
	// getting the user form the db by id
	var user User
	tx := DB.Where("id = ?", id).First(&user)
	return user, tx.Error
}

func GetUserByEmail(email string) (User, error) {
	// getting the user form the db by the specified email
	var user User
	tx := DB.Where("email = ?", email).First(&user)
	return user, tx.Error
}

func GetLoggedInUser(c gin.Context) (User, error) {
	// getting the cookie from the request
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		return User{}, err
	}

	// parsing the token from the cookie
	token, err := jwt.ParseWithClaims(
		cookie.Value,
		&jwt.StandardClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return User{}, err
	}

	// getting the claims from the token
	claims := token.Claims.(*jwt.StandardClaims)

	// getting the user from the db
	user, err := GetUserByStringId(claims.Issuer)
	if err != nil {
		return User{}, err
	}

	// if the user is logged in
	return user, nil
}

func IsLoggedIn(c gin.Context) bool {
	// getting the logged in user
	_, err := GetLoggedInUser(c)

	// if the error is nil, the user is logged in
	return err == nil
}
