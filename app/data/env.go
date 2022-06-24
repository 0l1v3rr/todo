package data

import "github.com/joho/godotenv"

var Env map[string]string

func GetVariables() error {
	// reading the environment variables from the .env file
	var err error
	Env, err = godotenv.Read(".env")

	return err
}
