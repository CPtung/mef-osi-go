package osi_test

import (
	"fmt"
	"testing"

	"github.com/MOXA-IPC/mef-osi-go/pkg/server"
	"github.com/MOXA-IPC/mef-osi-go/pkg/types"
	"github.com/stretchr/testify/suite"
)

type OsiV1TestSuite struct {
	suite.Suite
}

func (suite *OsiV1TestSuite) SetupSuite() {
}

func (suite *OsiV1TestSuite) TearDownSuite() {
}

func (suite *OsiV1TestSuite) SetupTest() {
	go func() {
		srv := server.New(types.SockTCP, "localhost:8880")
		srv.Serve()
	}()
	fmt.Printf("Setup Server... \n")
}

func (suite *OsiV1TestSuite) TearDownTest() {
}

func (suite *OsiV1TestSuite) BeforeTest(suiteName, testName string) {
}

func (suite *OsiV1TestSuite) AfterTest(suiteName, testName string) {
}

func TestStart(t *testing.T) {
	suite.Run(t, new(OsiV1TestSuite))
}
