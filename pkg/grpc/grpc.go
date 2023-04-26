package grpc

import (
	"context"
	"fmt"
	"github.com/tjfoc/gmsm/gmtls/gmcredentials"
	"github.com/yinzhenzhen/go-grpc-demo/pkg/common"
	"github.com/yinzhenzhen/go-grpc-demo/pkg/constant"
	"github.com/yinzhenzhen/go-grpc-demo/pkg/greet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"time"
)

func Run(isGm bool, isClient bool, end chan bool) {

	config := common.LoadConfig(isGm, isClient)

	creds := gmcredentials.NewTLS(config)

	if isClient {
		dialOpts := []grpc.DialOption{
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                1 * time.Minute,
				Timeout:             5 * time.Second,
				PermitWithoutStream: true,
			}),
			//grpc.WithBlock(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithDefaultCallOptions(
				grpc.MaxCallRecvMsgSize(constant.MaxRecvMessageSize),
				grpc.MaxCallSendMsgSize(constant.MaxSendMessageSize),
			),
			grpc.WithTransportCredentials(creds),
		}

		ctx, cancel := context.WithTimeout(context.Background(), constant.DialTimeout)
		defer cancel()
		conn, err := grpc.DialContext(ctx, constant.Address, dialOpts...)
		if err != nil || conn == nil {
			log.Fatalf("cannot to connect: %v", err)
		}
		c := greet.NewGreetServiceClient(conn)
		greetApi(c)
		end <- true

	} else {
		s := grpc.NewServer(grpc.Creds(creds))
		greet.RegisterGreetServiceServer(s, &GreetServer{})

		lis, err := net.Listen("tcp", constant.Port)
		if err != nil {
			log.Fatalf("fail to listen: %v", err)
		}
		err = s.Serve(lis)
		if err != nil {
			log.Fatalf("Serve: %v", err)
		}
		fmt.Println("server end")
	}
}
