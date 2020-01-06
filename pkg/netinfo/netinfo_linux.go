// +build linux

package netinfo

import (
	"github.com/google/gopacket/routing"
	log "github.com/sirupsen/logrus"
	"net"
)

func PrimaryIP(ips []net.IP) (net.IP, error) {
	// TODO: ensure interface/IP is external and primary
	// See IPNet.Contains()
	router, err := routing.New()
	if err != nil {
		log.Fatal(err, "error while creating routing object")
	}

	_, _, extIP, err := router.Route(ips[0])
	if err != nil {
		return nil, err
	}

	return extIP, nil
}
