package gmcredentials

import (
	"github.com/yinzhenzhen/go-grpc-demo/pkg/http"
	"testing"
)

func TestHttpGMServer(t *testing.T) {
	end = make(chan bool, 1)
	go http.Run(true, false, end)
	<-end
}

func TestHttpGMClient(t *testing.T) {
	end = make(chan bool, 1)
	go http.Run(true, true, end)
	<-end
}

func TestHttpECCServer(t *testing.T) {
	end = make(chan bool, 1)
	go http.Run(false, false, end)
	<-end
}

func TestHttpECCClient(t *testing.T) {
	end = make(chan bool, 1)
	go http.Run(false, true, end)
	<-end
}
