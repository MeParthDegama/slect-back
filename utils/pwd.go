package utils

// #cgo LDFLAGS: -lcrypt
/*
extern int checkSystemPassword(char *user, char *pass);
*/
import "C"
import (
	"fmt"
	"os"
)

func CheckSystemPassword(name string, pass string) {

	_, err := os.Stat("/home/" + name)
	if err != nil {
		fmt.Println("user not found")
		return
	}

	username := C.CString(name)
	password := C.CString(pass)
	result := C.checkSystemPassword(username, password)

	if result == 1 {
		fmt.Println("loginsuccess")
	} else {
		fmt.Println("loginerror")
	}

}
