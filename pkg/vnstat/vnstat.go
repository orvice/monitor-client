package vnstat

import (
	"encoding/json"
	"errors"
	"github.com/orvice/monitor-client/mod"
	"github.com/weeon/log"
	"net"
	"os/exec"
)

// fork form https://github.com/stnight/go-vnstat/blob/master/functions.go

// This function will execute vnstat command
func VN(netInterface string) mod.VNResult {
	cmd := exec.Command("vnstat", "-m", "-i", netInterface, "--json")
	stdout, err := cmd.Output()
	if err != nil {
		log.Errorf("get vnstat ret error %s", err.Error())
		return mod.VNResult{}
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
		log.Errorf("json.Unmarshal error %s", err.Error())
		return mod.VNResult{}
	}
	return vnRes
}

// This function will execute a command that lists all available network interfaces
func GetAllNetInterfaces() []mod.NetInterface {
	var allNetInterfaces []mod.NetInterface
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Errorf("get net interfaces error %s", err.Error())
		return nil
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
