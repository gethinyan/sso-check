package main

import (
	"fmt"

	"github.com/gethinyan/sso-check/pkg/setting"
	"github.com/gethinyan/sso-check/routes"
)

func main() {
	listenAddr := fmt.Sprintf("0.0.0.0:%d", setting.Server.HTTPPort)
	router := routes.SetupRouter()
	router.Run(listenAddr)
}
