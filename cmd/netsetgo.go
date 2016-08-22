package main

import (
	"flag"

	"github.com/teddyking/netsetgo"
)

func main() {
	var bridgeName, bridgeAddress, vethNamePrefix string
	var pid int

	flag.StringVar(&bridgeName, "bridgeName", "brg0", "Name to assign to bridge device")
	flag.StringVar(&bridgeAddress, "bridgeAddress", "10.10.10.1/24", "Address to assign to bridge device (CIDR notation)")
	flag.StringVar(&vethNamePrefix, "vethNamePrefix", "veth", "Name prefix for veth devices")
	flag.IntVar(&pid, "pid", 0, "pid of a process in the container's network namespace")
	flag.Parse()

	netsetgo.CreateBridge(bridgeName)
	netsetgo.AddAddressToBridge(bridgeName, bridgeAddress)
	netsetgo.SetBridgeUp(bridgeName)
	netsetgo.CreateVethPair(vethNamePrefix)
	netsetgo.AttachVethToBridge(bridgeName, vethNamePrefix)
	netsetgo.PlaceVethInNetworkNamespace(pid, vethNamePrefix)
}
