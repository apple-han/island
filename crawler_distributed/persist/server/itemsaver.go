package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
	"island/crawler/config"
	c "island/crawler_distributed/config"
	"island/crawler_distributed/rpcsupport"
	"log"
	"net/http"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	go func() {
		log.Fatal(serveRpc(
			fmt.Sprintf(":%d", *port),
			config.ElasticIndex))
	}()

	http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request){
		_, err := res.Write([]byte("pong"));
		if err != nil{
			log.Fatal("write err--->",err)
		}
	})
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal("open http err--->",err)
	}
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(c.ElasticHost),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&rpcsupport.RPCService{
			Client: client,
			Index:  index,
		})
}
