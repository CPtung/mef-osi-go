package server

import (
	"log"
	"net"
	"os"
	"path"
	"syscall"

	serial "github.com/MOXA-IPC/mef-osi-go/osi/serial"
	rpc_serial "github.com/MOXA-IPC/mef-osi-go/rpc/serial"
	"google.golang.org/grpc"
)

const (
	SockFD = "/run/mem/osi.sock"
)

type Server struct {
	listener net.Listener
	server   *grpc.Server
}

func New(sockType, sockPath string) *Server {
	if sockType == "unix" {
		if err := os.MkdirAll(path.Dir(sockPath), os.ModePerm); err != nil {
			log.Printf("mkdirall error %s\n", err.Error())
			return nil
		}
		syscall.Unlink(sockPath)
	} else if sockType != "tcp" {
		return nil
	}

	// start grpc reverse proxy socket
	sockProxy, err := net.Listen(sockType, sockPath)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return nil
	}

	// register all grpc services
	grpcServer := grpc.NewServer()
	rpc_serial.RegisterSerialServer(grpcServer, serial.NewService())
	/////////////////////////////

	return &Server{sockProxy, grpcServer}
}

func (s *Server) Serve() {
	if err := s.server.Serve(s.listener); err != nil {
		log.Printf("grpc serve error %s\n", err.Error())
	}
}
