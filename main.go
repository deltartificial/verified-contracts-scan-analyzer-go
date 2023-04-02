package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
)

type Config struct {
	URL              string `json:"verified_contracts_scan_url"`
	DiscordWebhookURL       string `json:"discord_webhook_url"`
	RefreshTime      int    `json:"refresh_time_seconds"`
}

type WebhookPayload struct {
	Content string `json:"content"`
}

func main() {

	configBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	var config Config
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatalln(err)
	}

	lastAddress := ""

	ticker := time.NewTicker(time.Duration(config.RefreshTime) * time.Second)

	for {
		client := &http.Client{}

		req, err := http.NewRequest("GET", config.URL, nil)
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:105.0) Gecko/20100101 Firefox/105.0")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		blue := color.FgBlue.Render

		s := doc.Find("table.table > tbody > tr").First()

		hash := s.Find("td > a.hash-tag").Text()
		if hash != "" && hash != lastAddress {

			color.Println("<fg=255;bg=blue>SUCCESS</> <fg=255>Last Verified Contract</>", blue("fetched successfully"))
			name := s.Find("td:nth-child(2)").Text()
			compiler := s.Find("td:nth-child(3)").Text()
			version := s.Find("td:nth-child(4)").Text()
			balance := s.Find("td:nth-child(5)").Text()
			transactions := s.Find("td:nth-child(6)").Text()

			fmt.Printf("Address: %s\n", hash)
			fmt.Printf("Contract Name: %s\n", name)
			fmt.Printf("Compiler: %s\n", compiler)
			fmt.Printf("Version: %s\n", version)
			fmt.Printf("Balance: %s\n", balance)
			fmt.Printf("Transactions: %s\n\n", transactions)

			payload := WebhookPayload{
				Content: fmt.Sprintf("New verified contract:\nAddress: %s\nContract Name: %s\nCompiler: %s\nVersion: %s\nBalance: %s\nTransactions: %s", hash, name, compiler, version, balance, transactions),
			}
			payloadBytes, err := json.Marshal(payload)
			if err != nil {
				log.Fatalln(err)
			}
			resp, err := http.Post(config.DiscordWebhookURL, "application/json", bytes.NewBuffer(payloadBytes))
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			
			lastAddress = hash
		}
		<-ticker.C
	}
}
