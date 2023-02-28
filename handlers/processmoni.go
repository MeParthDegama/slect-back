package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"golang.org/x/net/websocket"
)

type ProcMoniRes struct {
	CPU         float64 `json:"cpu"`
	MemoryTotal uint64  `json:"memory_total"`
	MemoryUsed  uint64  `json:"memory_used"`
}

func ProcMoni(ws *websocket.Conn) {
	for {

		before, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		time.Sleep(time.Duration(2) * time.Second)
		after, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		total := float64(after.Total - before.Total)

		memory, err := memory.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		user := &ProcMoniRes{
			CPU:         float64(after.User-before.User) / total * 100,
			MemoryTotal: memory.Total,
			MemoryUsed:  memory.Used,
		}
		b, err := json.Marshal(user)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}

		websocket.Message.Send(ws, string(b))
	}
}
