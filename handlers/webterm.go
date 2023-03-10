package handlers

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/creack/pty"
	"golang.org/x/net/websocket"
)

func WebTerm(ws *websocket.Conn) {
	c := exec.Command("zsh")
	f, err := pty.Start(c)
	if err != nil {
		ws.Write([]byte(fmt.Sprintf("Error creating pty: %s\r\n", err)))
		ws.Close()
		return
	}

	go io.Copy(ws, f)
	io.Copy(f, ws)
	ws.Close()
}
