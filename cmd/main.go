package main

import (
	"github.com/MOXA-IPC/mef-osi-go/pkg/server"
	"github.com/MOXA-IPC/mef-osi-go/pkg/types"
)

func main() {
	service := server.New(types.SockTCP, "localhost:8880")
	if service != nil {
		service.Serve()
	}
}
