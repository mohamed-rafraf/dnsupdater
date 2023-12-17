package dns

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/mohamed-rafraf/dnsupdater/pkg/config"
)

// UpdateDNS updates the DNS record to the new IP address
func UpdateDNS(cfg *config.AppConfig, ipAddress string) error {
	api, err := cloudflare.New(cfg.APIKey, cfg.Email)
	if err != nil {
		return fmt.Errorf("error creating Cloudflare API client: %v", err)
	}

	zoneID, err := api.ZoneIDByName(cfg.Domain)
	if err != nil {
		return fmt.Errorf("error getting Zone ID: %v", err)
	}

	records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{Name: cfg.Subdomain + "." + cfg.Domain})
	if err != nil {
		return fmt.Errorf("error listing DNS records: %v", err)
	}
	if len(records) != 1 {
		return fmt.Errorf("DNS record not found or multiple records returned")
	}

	rr := cloudflare.UpdateDNSRecordParams{
		ID:      records[0].ID,
		Type:    records[0].Type,
		Name:    records[0].Name,
		Content: ipAddress,
		TTL:     records[0].TTL,
		Proxied: records[0].Proxied,
	}

	_, err = api.UpdateDNSRecord(context.Background(), cloudflare.ZoneIdentifier(zoneID), rr)
	if err != nil {
		return fmt.Errorf("error updating DNS record: %v", err)
	}

	return nil
}
