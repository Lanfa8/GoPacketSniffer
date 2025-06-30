# GoPacketSniffer

This repository contains a comprehensive packet sniffer written in Go (Golang), using the `gopacket` library. It's designed to capture and analyze various network protocols including IP, TCP, UDP, and ICMP packets on a configurable network interface.

## Introduction

The Go Packet Sniffer is a network analysis tool, similar in concept to Wireshark, but focused on educational purposes and simplicity. It captures and analyzes packets from multiple protocols, displaying detailed information about each packet type. The tool now supports configurable network interfaces through environment variables.

## Supported Protocols

- **IP (IPv4 and IPv6)**: Source and destination IP addresses, TTL/Hop Limit, and IP version
- **TCP**: Source and destination ports, sequence numbers, and payload content
- **UDP**: Source and destination ports, packet size, and payload content  
- **ICMP (ICMPv4 and ICMPv6)**: Message type, code, and checksum

## Prerequisites

Before you can run the packet sniffer, you need to have the following installed:

- Go (Golang) - [Installation guide](https://golang.org/doc/install)
- libpcap - On Unix-based systems, this can usually be installed via your package manager (e.g., `apt` for Ubuntu, `yum` for CentOS).
  - For Debian/Ubuntu: `sudo apt-get install libpcap-dev`
  - For RedHat/CentOS: `sudo yum install libpcap-devel`
  - For macOS: `brew install libpcap` (using Homebrew)
- WinPcap or Npcap for Windows users - [Npcap](https://nmap.org/npcap/)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/araujo88/GoPacketSniffer.git
   ```
2. Navigate to the cloned repository:
   ```bash
   cd GoPacketSniffer
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

## Configuration

The packet sniffer uses a configurable network interface through a `.env` file. By default, it uses `enp1s0`, but you can easily change this:


### Manual configuration
Edit the `.env` file directly:
```bash
# Edit .env file
echo "NETWORK_INTERFACE=wlan0" > .env
```


### Useful commands:
- **List available interfaces**: `make list-interfaces`

## Usage

1. **Build the packet sniffer**:
   ```bash
   make
   ```

2. **Configure the network interface** (optional, defaults to enp1s0):
   ```bash
   # Edit .env file
   echo "NETWORK_INTERFACE=wlan0" > .env
   ```

3. **Run the packet sniffer**:
   ```bash
   sudo ./sniffer
   ```

The sniffer will start capturing packets from all supported protocols on the configured network interface.

## Sample Output

```
--- New Packet ---
[IPv4] Source: 192.168.1.10 -> Destination: 8.8.8.8 | TTL: 64 | Version: 4
[UDP] Source Port: 53461 -> Destination Port: 53 | Length: 36 bytes
[UDP] Payload: DNS query data...

--- New Packet ---
[IPv4] Source: 10.0.0.1 -> Destination: 10.0.0.2 | TTL: 64 | Version: 4
[TCP] Source Port: 80 -> Destination Port: 54321 | Seq: 1234567890
[TCP] Payload: HTTP/1.1 200 OK...

--- New Packet ---
[IPv4] Source: 192.168.1.1 -> Destination: 192.168.1.255 | TTL: 64 | Version: 4
[ICMPv4] Type: 8 | Code: 0 | Checksum: 0x1234
```

## Features

- **Multi-protocol support**: Captures and analyzes IP, TCP, UDP, and ICMP packets
- **IPv4 and IPv6 support**: Full support for both IP versions
- **Configurable network interface**: Easy configuration through .env file or environment variables
- **Detailed packet information**: 
  - IP layer: Source/destination addresses, TTL, version
  - TCP layer: Source/destination ports, sequence numbers, payload
  - UDP layer: Source/destination ports, packet size, payload
  - ICMP layer: Message type, code, checksum
- **Real-time packet analysis**: Live capture and display of network traffic
- **Educational-focused**: Clean, readable output format for learning network protocols

## Disclaimer

This tool is for educational purposes only. Unauthorized packet sniffing or network analysis can be illegal and unethical. Always ensure you have permission to capture packets on the network you're monitoring.

## Contributing

Contributions to the Go Packet Sniffer are welcome!

## License

This project is licensed under the [GPL License](LICENSE) - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Special thanks to the contributors of the `gopacket` library.
- Inspired by the functionalities of Wireshark.
