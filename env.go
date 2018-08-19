package main

import "github.com/orvice/utils/env"

var (
	netInterfaceName string
	webAddr          string
)

func InitEnv() {
	netInterfaceName = env.Get("NET_INTERFACE", "eth0")
	webAddr = env.Get("WEB_ADDR", ":8080")
}
