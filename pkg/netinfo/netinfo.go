package netinfo

import (
	"log"
	"net"
	"strings"
)

// TODO
type IntDetails struct {
}

// TODO
func InterfaceInfo(intName string) *IntDetails {
	// iface, err := net.InterfaceByName(intName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// addrs, err := iface.Addrs()
	// if err != nil {
	// 	log.Fatalf("No addresses found on interface: %s", err)
	// }

	// ips, err := netinfo.IPv4Addr(addrs)
	// if err != nil {
	// 	log.Fatal(err)
	// } else if len(ips) == 0 {
	// 	log.Fatalf("No IPv4 IPs found")
	// }

	// for _, ip := range ips {
	// 	log.Println(ip)
	// }
	return nil
}

func ValidateIP(addrs []net.Addr) ([]net.IP, error) {
	// TODO: Ensure IP is valid IPv4 address
	var (
		err     error
		ip      net.IP
		subnet  *net.IPNet
		ipv4IPs []net.IP
	)

	for _, addr := range addrs {
		ip, subnet, err = net.ParseCIDR(addr.String())

		if err != nil {
			// Log error and continue
			log.Printf("%s: %s", ip.String(), err)
			continue
		}

		// Validate IPv4 address
		if ip.To4() == nil || strings.Contains(subnet.String(), "::") {
			// Bad IP, skip it
			log.Printf("%s: Non-IPv4 IP address, skipping.", ip.String())
			continue
		}
		ipv4IPs = append(ipv4IPs, ip)
	}
	return ipv4IPs, err
}
