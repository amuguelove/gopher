/**
 * @author liangbo
 * @email  liangbo@codoon.com
 * @date   2018/2/6 下午11:03
 */
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/rpc"
	"runtime"
	"time"

	codec "github.com/mars9/codec"
)

var (
	host      = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	delay     = flag.Duration("delay", 0, "delay to mock business processing")
	debugAddr = flag.String("d", "127.0.0.1:9981", "server ip and port")
)

func main() {
	flag.Parse()

	go func() {
		log.Println(http.ListenAndServe(*debugAddr, nil))
	}()

	server := rpc.NewServer()
	server.Register(new(Hello))

	ln, err := net.Listen("tcp", *host)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print("rpc.Serve: accept:", err.Error())
			return
		}
		go serveConn(server, conn)
	}
}

func serveConn(server *rpc.Server, conn io.ReadWriteCloser) {
	srv := codec.NewServerCodec(conn)
	server.ServeCodec(srv)
}

type Hello int

func (t *Hello) Say(args *BenchmarkMessage, reply *BenchmarkMessage) error {
	args.Field1 = "OK"
	args.Field2 = 100
	*reply = *args
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return nil
}
