package rpcx

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gethinyan/sso-server/rpcx/services"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "172.16.21.71:8972", "server address")
	// Reply 响应参数
	Reply *services.Reply
)

// Authoration 登录验证
func Authoration(token string) {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Auth", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &services.Args{
		JsonWebToken: token,
	}
	fmt.Println(args)

	reply := &services.Reply{}
	err := xclient.Call(context.Background(), "CheckToken", args, reply)
	if err != nil {
		log.Printf("failed to call: %v\n", err)
	}
	fmt.Println(reply)
	Reply = reply
}
