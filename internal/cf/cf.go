package cf

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/akston/dynamo/pkg/netinfo"
	"github.com/cloudflare/cloudflare-go"
)

func record(apiToken, zoneName, recordName string) (*cloudflare.API, *cloudflare.DNSRecord, string, error) {
	// TODO: return struct rather than a bunch of separate values
	var rec []cloudflare.DNSRecord

	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return nil, nil, "", err
	}

	id, err := api.ZoneIDByName(zoneName)
	if err != nil {
		return nil, nil, "", err
	}

	rec, err = api.DNSRecords(id, cloudflare.DNSRecord{Name: recordName})
	if err != nil {
		return nil, nil, "", err
	}

	if len(rec) != 1 {
		return nil, nil, "", fmt.Errorf("expected one record, got %d. exiting", len(rec))
	}

	return api, &rec[0], id, nil
}

func UpdateRecord() error {
	// TODO: pull these values from config or commandline flags
	apiToken := os.Getenv("CF_API_TOKEN")
	zoneName := os.Getenv("CF_ZONE_NAME")
	recordName := os.Getenv("CF_RECORD_NAME") + "." + zoneName

	api, rec, id, err := record(apiToken, zoneName, recordName)
	if err != nil {
		return err
	}

	recIP := net.ParseIP(rec.Content)

	// TODO: Validate extIP
	extIP, err := netinfo.PrimaryIP()
	if err != nil {
		return err
	}

	if recIP.Equal(extIP) {
		log.Printf("IP %s is already current, exiting without changes", recIP)
		return nil
	}

	log.Printf("Updating %s to %s. Used to be %s", rec.Name, extIP, recIP)

	rec.Content = extIP.String()
	if err := api.UpdateDNSRecord(id, rec.ID, *rec); err != nil {
		log.Fatal(err)
	}

	log.Println("Update complete.")
	return nil
}
