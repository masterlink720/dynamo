package main

import (
	"net"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

func main() {
	// TODO: pull these values from config or commandline flags
	apiToken := os.Getenv("CF_API_TOKEN")
	zoneName := os.Getenv("CF_ZONE_NAME")
	recordName := os.Getenv("CF_RECORD_NAME") + "." + zoneName
	intName := os.Getenv("INT_NAME")

	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		log.Fatal(err)
	}

	id, err := api.ZoneIDByName(zoneName)
	if err != nil {
		log.Fatal(err)
	}

	recs, err := api.DNSRecords(id, cloudflare.DNSRecord{Name: recordName})
	if err != nil {
		log.Fatal(err)
	}

	if len(recs) != 1 {
		log.Fatalf("Expected one record, got %d. Exiting.", len(recs))
	}

	for _, r := range recs {
		log.Printf("%s: %s\n", r.Name, r.Content)
	}

	iface, err := net.InterfaceByName(intName)
	if err != nil {
		log.Fatal(err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		log.Fatalf("No addresses found on interface: %s", err)
	}

	ips, err := getIPv4Addr(addrs)
	if err != nil {
		log.Fatal(err)
	} else if len(ips) == 0 {
		log.Fatalf("No IPv4 IPs found")
	}

	for _, ip := range ips {
		log.Println(ip)
	}

	// TODO: ensure interface/IP is external and primary

	// TODO: if the two IPs differ, update A record with CF

}

func getIPv4Addr(addrs []net.Addr) ([]net.IP, error) {
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
