package rpcx

import (
	"context"
	"flag"
	"log"

	"github.com/gethinyan/sso-server/rpcx/services"
	"github.com/smallnest/rpcx/client"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

// Authoration 登录验证
func Authoration() (reply *services.Reply, err error) {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := &services.Args{
		A: 10,
		B: 20,
	}

	err = xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Printf("%d * %d = %d", args.A, args.B, reply.C)

	return
}