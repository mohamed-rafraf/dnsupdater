package main

import (
	"log"
	"time"

	"github.com/mohamed-rafraf/dnsupdater/pkg/config"
	"github.com/mohamed-rafraf/dnsupdater/pkg/dns"
	"github.com/mohamed-rafraf/dnsupdater/pkg/file"
	"github.com/mohamed-rafraf/dnsupdater/pkg/ip"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	for {
		// Read IP address from file
		ipFromFile, err := file.ReadFile(cfg.FilePath)
		if err != nil {
			log.Printf("Error reading file: %v", err)
			continue
		}

		// Fetch current IP address
		currentIP, err := ip.GetIPAddress()
		if err != nil {
			log.Printf("Error fetching IP address: %v", err)
			continue
		}

		// Update DNS if IP address has changed
		if ipFromFile != currentIP {
			log.Println("IP Address is changed!")
			if err := dns.UpdateDNS(cfg, currentIP); err != nil {
				log.Printf("Error updating DNS: %v", err)
				continue
			}

			log.Println("DNS Record is changed successfully!")

			// Update the file with the new IP address
			if err := file.WriteFile(cfg.FilePath, currentIP); err != nil {
				log.Printf("Error writing to file: %v", err)
				continue
			}

			log.Println("The current IP address is saved!")
		}

		// Sleep for the specified interval
		time.Sleep(cfg.CheckInterval)
	}
}
