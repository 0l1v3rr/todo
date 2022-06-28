package util

type Error struct {
	Message string `json:"error" example:"An unknown error occurred."`
}

type Success struct {
	Message string `json:"message" example:"Success!"`
}
