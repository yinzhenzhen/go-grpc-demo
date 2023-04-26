package common

import (
	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"
	"github.com/yinzhenzhen/go-grpc-demo/pkg/constant"
	"io/ioutil"
	"log"
)

func LoadConfig(isGm bool, isClient bool) (config *gmtls.Config) {
	var err error
	var cert gmtls.Certificate
	var caCert []byte
	config = new(gmtls.Config)
	if isGm {
		cert, err = gmtls.LoadX509KeyPair(constant.GM_CERT, constant.GM_KEY)
		if err != nil {
			log.Fatal(err)
		}
		caCert, err = ioutil.ReadFile(constant.GM_CA)
		if err != nil {
			log.Fatal(err)
		}
		config.GMSupport = &gmtls.GMSupport{}
		config.ServerName = "localhost"
	} else {
		cert, err = gmtls.LoadX509KeyPair(constant.ECC_CERT, constant.ECC_KEY)
		if err != nil {
			log.Fatal(err)
		}
		caCert, err = ioutil.ReadFile(constant.ECC_CA)
		if err != nil {
			log.Fatal(err)
		}
		config.ServerName = "orderer0.example.com"
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	if isClient {
		config.RootCAs = certPool
	} else {
		config.ClientCAs = certPool
	}

	config.Certificates = []gmtls.Certificate{cert, cert}
	config.ClientAuth = gmtls.RequireAndVerifyClientCert

	return
}
