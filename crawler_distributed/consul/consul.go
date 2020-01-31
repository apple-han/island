package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	c "reptiles/crawler_distributed/config"
	"strconv"
)

func Find(name string) []string{
	var lastIndex uint64
	config := api.DefaultConfig()
	config.Address = c.ConsulHost

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println("api new client is failed, err:", err)
		return nil
	}
	services, metainfo, err := client.Health().Service(name, "", true, &api.QueryOptions{
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
	return addrs
}
