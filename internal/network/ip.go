package network

import (
	"errors"
	"net"
)

// GetLocalIP returns the non-loopback local IP of the host.
// It iterates through all network interfaces to find an active IPv4 address.
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// Check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("cannot find local IP address, are you connected to a network?")
}
