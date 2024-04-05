package main

import "github.com/MOXA-IPC/mef-osi-go/pkg/server"

func main() {
	service := server.New()
	if service != nil {
		service.Serve()
	}
}
