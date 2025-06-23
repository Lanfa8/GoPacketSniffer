package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

// processIPLayer processes IPv4 and IPv6 layers
func processIPLayer(packet gopacket.Packet) {
	ipv4Layer := packet.Layer(layers.LayerTypeIPv4)
	if ipv4Layer != nil {
		ip, _ := ipv4Layer.(*layers.IPv4)
		fmt.Printf("[IPv4] Source: %s -> Destination: %s | TTL: %d | Version: %d\n",
			ip.SrcIP, ip.DstIP, ip.TTL, ip.Version)
	}

	ipv6Layer := packet.Layer(layers.LayerTypeIPv6)
	if ipv6Layer != nil {
		ip, _ := ipv6Layer.(*layers.IPv6)
		fmt.Printf("[IPv6] Source: %s -> Destination: %s | Hop Limit: %d | Version: %d\n",
			ip.SrcIP, ip.DstIP, ip.HopLimit, ip.Version)
	}
}

// processTCPLayer processes TCP packets
func processTCPLayer(packet gopacket.Packet) {
	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("[TCP] Source Port: %d -> Destination Port: %d | Seq: %d\n",
			tcp.SrcPort, tcp.DstPort, tcp.Seq)

		if len(tcp.Payload) > 0 {
			fmt.Printf("[TCP] Payload: %s\n", string(tcp.Payload))
		}
	}
}

// processUDPLayer processes UDP packets
func processUDPLayer(packet gopacket.Packet) {
	udpLayer := packet.Layer(layers.LayerTypeUDP)
	if udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		fmt.Printf("[UDP] Source Port: %d -> Destination Port: %d | Length: %d bytes\n",
			udp.SrcPort, udp.DstPort, udp.Length)

		if len(udp.Payload) > 0 {
			fmt.Printf("[UDP] Payload: %s\n", string(udp.Payload))
		}
	}
}

// processICMPLayer processes ICMP packets
func processICMPLayer(packet gopacket.Packet) {
	icmpLayer := packet.Layer(layers.LayerTypeICMPv4)
	if icmpLayer != nil {
		icmp, _ := icmpLayer.(*layers.ICMPv4)
		fmt.Printf("[ICMPv4] Type: %d | Code: %d | Checksum: 0x%x\n",
			icmp.TypeCode.Type(), icmp.TypeCode.Code(), icmp.Checksum)
	}

	icmpv6Layer := packet.Layer(layers.LayerTypeICMPv6)
	if icmpv6Layer != nil {
		icmp, _ := icmpv6Layer.(*layers.ICMPv6)
		fmt.Printf("[ICMPv6] Type: %d | Code: %d | Checksum: 0x%x\n",
			icmp.TypeCode.Type(), icmp.TypeCode.Code(), icmp.Checksum)
	}
}

func main() {
	handle, err := pcap.OpenLive("eth0", 1600, true, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println("\n--- New Packet ---")

		processIPLayer(packet)
		processTCPLayer(packet)
		processUDPLayer(packet)
		processICMPLayer(packet)
	}
}
