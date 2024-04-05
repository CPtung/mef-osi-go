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
	return nil, nil
}

func (s *SerialImpl) SetSerial(ctx context.Context, request *rpc.SerialRequest) (*rpc.SerialReply, error) {
	return nil, nil
}
