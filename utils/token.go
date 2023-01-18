package utils

import (
	"github.com/google/uuid"
)

func GenToken() string {
	token := uuid.NewString()
	return token
}
