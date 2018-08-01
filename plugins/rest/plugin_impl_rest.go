// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"fmt"
	"sync"

	"git.fd.io/govpp.git/api"
	"github.com/ligato/cn-infra/infra"
	"github.com/ligato/cn-infra/rpc/rest"
	"github.com/ligato/cn-infra/utils/safeclose"
	"github.com/ligato/vpp-agent/plugins/govppmux"
	"github.com/ligato/vpp-agent/plugins/rest/resturl"
	"github.com/ligato/vpp-agent/plugins/vpp"
	aclvppcalls "github.com/ligato/vpp-agent/plugins/vpp/aclplugin/vppcalls"
	ifvppcalls "github.com/ligato/vpp-agent/plugins/vpp/ifplugin/vppcalls"
	ipsecvppcalls "github.com/ligato/vpp-agent/plugins/vpp/ipsecplugin/vppcalls"
	l2vppcalls "github.com/ligato/vpp-agent/plugins/vpp/l2plugin/vppcalls"
	l3vppcalls "github.com/ligato/vpp-agent/plugins/vpp/l3plugin/vppcalls"
	l4vppcalls "github.com/ligato/vpp-agent/plugins/vpp/l4plugin/vppcalls"
)

// REST api methods
const (
	GET  = "GET"
	POST = "POST"
)

// Plugin registers Rest Plugin
type Plugin struct {
	Deps

	// Index page
	indexes *indexes

	// Channels
	vppChan  api.Channel
	dumpChan api.Channel

	// Handlers
	aclHandler   aclvppcalls.AclVppRead
	ifHandler    ifvppcalls.IfVppRead
	bfdHandler   ifvppcalls.BfdVppRead
	natHandler   ifvppcalls.NatVppRead
	stnHandler   ifvppcalls.StnVppRead
	ipSecHandler ipsecvppcalls.IPSecVPPRead
	bdHandler    l2vppcalls.BridgeDomainVppRead
	fibHandler   l2vppcalls.FibVppRead
	xcHandler    l2vppcalls.XConnectVppRead
	arpHandler   l3vppcalls.ArpVppRead
	pArpHandler  l3vppcalls.ProxyArpVppRead
	rtHandler    l3vppcalls.RouteVppRead
	l4Handler    l4vppcalls.L4VppRead

	govppmux sync.Mutex
}

// Deps represents dependencies of Rest Plugin
type Deps struct {
	infra.PluginDeps
	HTTPHandlers rest.HTTPHandlers
	GoVppmux     govppmux.API
	VPP          vpp.API
}

type indexes struct {
	IdxMap map[string][]indexItem
}

type indexItem struct {
	Name string
	Path string
}

// Init initializes the Rest Plugin
func (plugin *Plugin) Init() (err error) {
	// Check VPP dependency
	if plugin.VPP == nil {
		return fmt.Errorf("REST plugin requires VPP plugin API")
	}
	// VPP channels
	if plugin.vppChan, err = plugin.GoVppmux.NewAPIChannel(); err != nil {
		return err
	}
	if plugin.dumpChan, err = plugin.GoVppmux.NewAPIChannel(); err != nil {
		return err
	}
	// Indexes
	ifIndexes := plugin.VPP.GetSwIfIndexes()
	bdIndexes := plugin.VPP.GetBDIndexes()
	spdIndexes := plugin.VPP.GetIPSecSPDIndexes()

	// Initialize handlers
	if plugin.aclHandler, err = aclvppcalls.NewAclVppHandler(plugin.vppChan, plugin.dumpChan, nil); err != nil {
		return err
	}
	if plugin.ifHandler, err = ifvppcalls.NewIfVppHandler(plugin.vppChan, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.bfdHandler, err = ifvppcalls.NewBfdVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.natHandler, err = ifvppcalls.NewNatVppHandler(plugin.vppChan, plugin.dumpChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.stnHandler, err = ifvppcalls.NewStnVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.ipSecHandler, err = ipsecvppcalls.NewIPsecVppHandler(plugin.vppChan, ifIndexes, spdIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.bdHandler, err = l2vppcalls.NewBridgeDomainVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.fibHandler, err = l2vppcalls.NewFibVppHandler(plugin.vppChan, plugin.dumpChan, make(chan *l2vppcalls.FibLogicalReq),
		ifIndexes, bdIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.xcHandler, err = l2vppcalls.NewXConnectVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.arpHandler, err = l3vppcalls.NewArpVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.pArpHandler, err = l3vppcalls.NewProxyArpVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.rtHandler, err = l3vppcalls.NewRouteVppHandler(plugin.vppChan, ifIndexes, plugin.Log, nil); err != nil {
		return err
	}
	if plugin.l4Handler, err = l4vppcalls.NewL4VppHandler(plugin.vppChan, plugin.Log, nil); err != nil {
		return err
	}

	// Fill index item lists
	idxMap := make(map[string][]indexItem)
	idxMap["ACL plugin"] = []indexItem{
		{Name: "IP-type access lists", Path: resturl.AclIP},
		{Name: "MACIP-type access lists", Path: resturl.AclMACIP},
	}
	idxMap["Interface plugin"] = []indexItem{
		{Name: "All interfaces", Path: resturl.Interface},
		{Name: "Loopbacks", Path: resturl.Loopback},
		{Name: "Ethernets", Path: resturl.Ethernet},
		{Name: "Memifs", Path: resturl.Memif},
		{Name: "Taps", Path: resturl.Tap},
		{Name: "VxLANs", Path: resturl.VxLan},
		{Name: "Af-packets", Path: resturl.AfPacket},
	}
	idxMap["IPSec plugin"] = []indexItem{
		{Name: "Security policy databases", Path: resturl.IPSecSpd},
		{Name: "Security associations", Path: resturl.IPSecSa},
		{Name: "Tunnel interfaces", Path: resturl.IPSecTnIf},
	}
	idxMap["L2 plugin"] = []indexItem{
		{Name: "Bridge domains", Path: resturl.Bd},
		{Name: "Bridge domain IDs", Path: resturl.BdId},
		{Name: "L2Fibs", Path: resturl.Fib},
		{Name: "Cross connects", Path: resturl.Xc},
	}
	idxMap["L3 plugin"] = []indexItem{
		{Name: "ARPs", Path: resturl.Arps},
		{Name: "Proxy ARPs", Path: resturl.ProxyArps},
		{Name: "Proxy ARP interfaces", Path: resturl.PArpIfs},
		{Name: "Proxy ARP ranges", Path: resturl.PArpRngs},
		{Name: "Static routes", Path: resturl.Routes},
	}
	idxMap["L4 plugin"] = []indexItem{
		{Name: "L4 sessions", Path: resturl.Sessions},
	}
	idxMap["Telemetry"] = []indexItem{
		{Name: "All data", Path: resturl.Telemetry},
		{Name: "Memory", Path: resturl.TMemory},
		{Name: "Runtime", Path: resturl.TRuntime},
		{Name: "Node count", Path: resturl.TNodeCount},
	}
	idxMap["Other"] = []indexItem{
		{Name: "CLI command", Path: resturl.Command},
	}

	plugin.indexes = &indexes{
		IdxMap: idxMap,
	}

	return nil
}

// AfterInit is used to register HTTP handlers
func (plugin *Plugin) AfterInit() (err error) {
	plugin.Log.Debug("REST API Plugin is up and running")

	plugin.registerAccessListHandlers()
	plugin.registerInterfaceHandlers()
	plugin.registerBfdHandlers()
	plugin.registerNatHandlers()
	plugin.registerStnHandlers()
	plugin.registerIPSecHandlers()
	plugin.registerL2Handlers()
	plugin.registerL3Handlers()
	plugin.registerL4Handlers()
	plugin.registerTelemetryHandlers()
	plugin.registerCommandHandler()
	plugin.registerIndexHandlers()

	return nil
}

// Close is used to clean up resources used by Plugin
func (plugin *Plugin) Close() (err error) {
	return safeclose.Close(plugin.vppChan, plugin.dumpChan)
}
