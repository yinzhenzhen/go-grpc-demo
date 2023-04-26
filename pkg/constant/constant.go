package constant

import "time"

const (
	Port    = ":50051"
	Address = "localhost:50051"
)

const (
	DialTimeout        = 10 * time.Second
	MaxRecvMessageSize = 100 * 1024 * 1024 // 100 MiB
	MaxSendMessageSize = 100 * 1024 * 1024 // 100 MiB
)

const (
	GM_CA   = "/Users/yz/Desktop/gopath/src/github.com/yinzhenzhen/go-grpc-demo/certs/gm/root.crt"
	GM_KEY  = "/Users/yz/Desktop/gopath/src/github.com/yinzhenzhen/go-grpc-demo/certs/gm/sign.key"
	GM_CERT = "/Users/yz/Desktop/gopath/src/github.com/yinzhenzhen/go-grpc-demo/certs/gm/sign.crt"

	ECC_CA   = "/Users/yz/Desktop/gopath/src/github.com/yinzhenzhen/go-grpc-demo/certs/ecc/root.crt"
	ECC_KEY  = "/Users/yz/Desktop/gopath/src/github.com/yinzhenzhen/go-grpc-demo/certs/ecc/sign.key"
	ECC_CERT = "/Users/yz/Desktop/gopath/src/github.com/yinzhenzhen/go-grpc-demo/certs/ecc/sign.crt"
)
