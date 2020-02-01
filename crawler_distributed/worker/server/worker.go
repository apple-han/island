package main

import (
	"fmt"
	"log"
	"net/http"
	"island/crawler_distributed/rpcsupport"

	"flag"

	"island/crawler/fetcher"
)

var (
	port = flag.Int("port", 0,
	"the port for me to listen on")
	httpPort = flag.Int("httpPort", 0,
		"the port for me to listen on")
	)

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	go func() {
		log.Fatal(rpcsupport.ServeRpc(
			fmt.Sprintf(":%d", *port),
			&rpcsupport.RPCService{}))
	}()

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request){
		_, err := res.Write([]byte("pong"));
		if err != nil{
			log.Fatal("write err--->",err)
		}
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil{
		log.Fatal("open http err--->",err)
	}
}
