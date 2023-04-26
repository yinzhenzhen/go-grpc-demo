package http

import (
	"fmt"
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/yinzhenzhen/go-grpc-demo/pkg/common"
	"log"
	"net/http"
)

func Run(isGm bool, isClient bool, end chan bool) {

	config := common.LoadConfig(isGm, isClient)

	if isClient {
		conn, err := gmtls.Dial("tcp", "localhost:50051", config)
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		req := []byte("GET / HTTP/1.1\r\n" +
			"Host: localhost\r\n" +
			"Connection: close\r\n\r\n")
		_, err = conn.Write(req)
		if err != nil {
			panic(err)
		}
		buff := make([]byte, 1024)
		for {
			n, _ := conn.Read(buff)
			if n <= 0 {
				break
			} else {
				fmt.Printf("%s", buff[0:n])
			}
		}
		fmt.Println(">> [PASS]")
		end <- true

	} else {
		ln, err := gmtls.Listen("tcp", ":50051", config)
		if err != nil {
			log.Println(err)
			return
		}
		defer ln.Close()

		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintf(writer, "hello yin\n")
		})
		//http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		//	fmt.Fprintf(writer, "hello hello yin\n")
		//})
		fmt.Println(">> HTTP Over running...")
		err = http.Serve(ln, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(">> HTTP Over end")
	}

}
