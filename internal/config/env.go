package config

import "os"

var (
	NetInterfaceName string
	WebAddr          string
	GrpcAddr         string
	PostUrl          string
	PostKey          string
)

func InitEnv() {
	NetInterfaceName = getEvn("NET_INTERFACE", "eth0")
	WebAddr = getEvn("WEB_ADDR", ":8080")
	GrpcAddr = getEvn("GRPC_ADDR", ":8090")
	PostUrl = getEvn("POST_URL")
	PostKey = getEvn("POST_KEY")
}

func getEvn(key string, df ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 && len(df) != 0 {
		return df[0]
	}
	return v
}
