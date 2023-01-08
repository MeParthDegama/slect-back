package utils

// #cgo LDFLAGS: -lcrypt
/*
extern int checkSystemPassword(char *user, char *pass);
*/
import "C"
import (
	"os/user"
)

func CheckSystemPassword(name string, pass string) bool {

	user, err := user.Lookup(name)
	if err != nil {
		return false
	}

	username := C.CString(user.Username)
	password := C.CString(pass)
	result := C.checkSystemPassword(username, password)

	if result == 1 {
		return true
	} else {
		return false
	}

}
