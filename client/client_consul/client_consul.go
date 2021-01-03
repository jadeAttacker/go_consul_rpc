package client_consul

import (
	"fmt"
	"net"
	"strconv"

	"github.com/hashicorp/consul/api"
)

var lastIndex uint64

func Lookup() (addr interface{}) {
	// consul寻址
	config := api.DefaultConfig()
	config.Address = "192.168.1.106:8500"

	client, err := api.NewClient(config)

	if err != nil {
		fmt.Println("新建consul客户端失败", err)
	}

	services, matainfo, err := client.Health().Service("serverNode", "v1000", true, &api.QueryOptions{
		WaitIndex: lastIndex,
	})

	if err != nil {
		fmt.Println("获取可用服务失败", err)
	}

	lastIndex = matainfo.LastIndex

	var addr2 interface{}
	for _, service := range services {
		addr2 = net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))
		fmt.Println("获取服务信息", service, lastIndex, addr2)
	}
	return addr2
}
