package config

import "github.com/orvice/utils/env"

var (
	NetInterfaceName string
	WebAddr          string
	GrpcAddr         string
	PostUrl          string
	PostKey          string
)

func InitEnv() {
	NetInterfaceName = env.Get("NET_INTERFACE", "eth0")
	WebAddr = env.Get("WEB_ADDR", ":8080")
	GrpcAddr = env.Get("GRPC_ADDR", ":8090")
	PostUrl = env.Get("POST_URL")
	PostKey = env.Get("POST_KEY")
}
