package osi_test

import (
	"context"
	"log"
	"time"

	"github.com/MOXA-IPC/mef-osi-go/pkg/client"
	"github.com/stretchr/testify/assert"

	pb "github.com/MOXA-IPC/mef-osi-go/rpc/serial"
)

func (s *OsiV1TestSuite) TestSerial() {
	cli := client.New("localhost:8880")
	assert.NotNil(s.T(), cli)
	defer cli.Close()

	serial := pb.NewSerialClient(cli)
	assert.NotNil(s.T(), serial)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(3)*time.Second)
	defer cancel()
	r, err := serial.GetSerial(ctx, &pb.SerialEmptyRequest{})
	if err != nil {
		log.Fatalf("could not get serial: %v", err)
	}
	log.Printf("Greeting: %v", r.GetProfiles())
}
