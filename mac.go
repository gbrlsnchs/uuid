package uuid

import (
	"errors"
	"net"
)

func macAddr() (net.HardwareAddr, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range interfaces {
		mac := i.HardwareAddr
		if i.Flags&net.FlagUp > 0 && string(mac) != "" {
			return mac, nil
		}
	}
	return nil, errors.New("uuid: unable to retrieve a MAC Address")
}
