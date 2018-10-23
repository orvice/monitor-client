package main

import "github.com/orvice/utils/env"

var (
	netInterfaceName string
	webAddr          string
	grpcAddr         string
	postUrl          string

	enableWS bool
)

func InitEnv() {
	netInterfaceName = env.Get("NET_INTERFACE", "eth0")
	webAddr = env.Get("WEB_ADDR", ":8080")
	grpcAddr = env.Get("GRPC_ADDR", ":8090")
	postUrl = env.Get("POST_URL")

	enableWs := env.Get("ENABLE_WS", "true")
	if enableWs == "true" {
		enableWS = true
	}
}
