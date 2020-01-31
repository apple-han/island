package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	"strconv"
)

func main() {
	var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = "http://localhost:8500" //consul server

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return
	}
	services, metainfo, err := client.Health().Service("car", "", true, &api.QueryOptions{
		WaitIndex: lastIndex,
	})
	if err != nil {
		log.Println("error retrieving instances from Consul: ", err)
	}
	lastIndex = metainfo.LastIndex

	addrs := make([]string, 0)
	for _, service := range services {
		addrs = append(addrs, net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port)))
	}
	fmt.Println("addr---->",addrs)
}
