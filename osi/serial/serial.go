package serial

import (
	"context"

	rpc "github.com/MOXA-IPC/mef-osi-go/rpc/serial"
)

type SerialImpl struct {
}

func NewService() rpc.SerialServer {
	return &SerialImpl{}
}

func (s *SerialImpl) GetSerial(ctx context.Context, empty *rpc.SerialEmptyRequest) (*rpc.SerialReply, error) {
	return &rpc.SerialReply{
		Profiles: []*rpc.Profile{{
			Name: "COM1",
			Path: "/dev/ttyM0",
			Mode: "RS232",
		}},
	}, nil
}

func (s *SerialImpl) SetSerial(ctx context.Context, request *rpc.SerialRequest) (*rpc.SerialReply, error) {
	return nil, nil
}
