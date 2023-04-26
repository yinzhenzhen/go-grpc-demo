package gmcredentials

import (
	"github.com/yinzhenzhen/go-grpc-demo/pkg/grpc"
	"testing"
)

var end chan bool

func TestGMServer(t *testing.T) {
	end = make(chan bool, 1)
	go grpc.Run(true, false, end)
	<-end
}

func TestGMClient(t *testing.T) {
	end = make(chan bool, 1)
	go grpc.Run(true, true, end)
	<-end
}

func TestECCServer(t *testing.T) {
	end = make(chan bool, 1)
	go grpc.Run(false, false, end)
	<-end
}

func TestECCClient(t *testing.T) {
	end = make(chan bool, 1)
	go grpc.Run(false, true, end)
	<-end
}
