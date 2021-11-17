package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"lmf.mortal.com/ThriftDemo/example"
	"log"
	"net"
	"testing"
)

func TestFormatDataImpl_DoFormat(t *testing.T) {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := transportFactory.GetTransport(tSocket)
	fmt.Println("err", err)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	data := example.Data{Text: "hello,world!"}
	d, err := client.DoFormat(context.Background(), &data)
	fmt.Println(d.Text)
}

//const (
//	HOST = "localhost"
//	PORT = "8080"
//)

//func main()  {

//}
