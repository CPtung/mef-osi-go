package server

import (
	"log"
	"net"
	"os"
	"path"
	"syscall"

	serial "github.com/MOXA-IPC/mef-osi-go/interface/serial"
	rpc_serial "github.com/MOXA-IPC/mef-osi-go/rpc/serial"
	"google.golang.org/grpc"
)

var sockfd = "/run/mem/osi.sock"

type Server struct {
	listener net.Listener
	server   *grpc.Server
}

func New() *Server {
	if err := os.MkdirAll(path.Dir(sockfd), os.ModePerm); err != nil {
		log.Printf("mkdirall error %s\n", err.Error())
		return nil
	}

	// start grpc reverse proxy socket
	syscall.Unlink(sockfd)
	sockProxy, err := net.Listen("unix", sockfd)
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
