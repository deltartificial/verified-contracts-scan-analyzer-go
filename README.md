# Verified Contracts Scan Analyzer Go

🛰️ - Verified Contracts Scan Analyzer in Golang with Discord webhook alerts. (etherscan, arbiscan, bscscan, ftmscan, snowtrace, polygonscan etc).

This is a Golang script that enables you to analyze verified contracts scans across multiple blockchains like Ethereum, Binance Smart Chain, Arbitrum, Fantom, Snowtrace, and Polygon.

### Features

- Multi-blockchain support (Ethereum, Binance Smart Chain, Arbitrum, Fantom, Snowtrace, and Polygon)
- Integration with etherscan, arbiscan, bscscan, ftmscan, snowtrace, polygonscan, etc.
- Discord webhook alerts for updated verified contract scans

### Installation

Clone the repository using git clone `https://github.com/deltartificial/verified-contracts-scan-analyzer-go.git`
Navigate to the cloned repository using `cd verified-contracts-scan-analyzer-go`
Install the required dependencies using `go get`

### Configuration

The config.json file contains the following configuration options:

- verified_contracts_scan_url : The URL of the verified contracts scan.
- discord_webhook_url : The Discord webhook URL to send alerts to.
- refresh_time_seconds : The time interval in seconds to refresh the verified contracts scan.

### Usage

Open the config.json file and add your Discord webhook URL, verified contracts scan URL, and refresh time interval.
Run the script using `go run main.go`
The script will start analyzing the verified contract scans across the selected blockchains and send alerts to your Discord webhook URL.

### License

This project is licensed under the MIT License. See the LICENSE file for details.
