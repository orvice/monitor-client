package vnstat

import (
	"encoding/json"
	"errors"
	"github.com/orvice/monitor-client/mod"
	"net"
	"os/exec"
)

// fork form https://github.com/stnight/go-vnstat/blob/master/functions.go

// This function will execute vnstat command
func VN(netinterface string) mod.VNResult {
	cmd := exec.Command("vnstat", "-m", "-i", netinterface, "--json")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
	err = cmd.Start()
	if err != nil {
		err = errors.New("COMMAND_ERROR")
	}
	defer cmd.Wait()
	b := []byte(stdout)
	var vnRes mod.VNResult
	err = json.Unmarshal(b, &vnRes)
	if err != nil {
		panic(err)
	}
	return vnRes
}

// This function will execute a command that lists all available network interfaces
func GetAllNetInterfaces() []mod.NetInterface {
	var allNetInterfaces []mod.NetInterface
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err.Error())
	}
	for key := range interfaces {
		allNetInterfaces = append(allNetInterfaces, mod.NetInterface{
			Index: interfaces[key].Index,
			MTU:   interfaces[key].MTU,
			Name:  interfaces[key].Name,
		})
	}
	return allNetInterfaces
}
