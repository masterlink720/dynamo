// +build darwin

package netinfo

import (
	"net"
)

// PrimaryIP returns IP used for external lookups
func PrimaryIP() (net.IP, error) {
	// TODO: ensure interface/IP is external and primary
	// See IPNet.Contains()
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	ip := conn.LocalAddr().(*net.UDPAddr)
	return ip.IP, nil
}
