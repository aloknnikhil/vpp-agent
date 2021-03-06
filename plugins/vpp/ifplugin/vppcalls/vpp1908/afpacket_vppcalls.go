//  Copyright (c) 2019 Cisco and/or its affiliates.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at:
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package vpp1908

import (
	"net"

	"go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp1908/af_packet"
	interfaces "go.ligato.io/vpp-agent/v3/proto/ligato/vpp/interfaces"
)

// AddAfPacketInterface implements AfPacket handler.
func (h *InterfaceVppHandler) AddAfPacketInterface(ifName string, hwAddr string, afPacketIntf *interfaces.AfpacketLink) (swIndex uint32, err error) {
	req := &af_packet.AfPacketCreate{
		HostIfName: []byte(afPacketIntf.HostIfName),
	}
	if hwAddr == "" {
		req.UseRandomHwAddr = 1
	} else {
		mac, err := net.ParseMAC(hwAddr)
		if err != nil {
			return 0, err
		}
		req.HwAddr = mac
	}
	reply := &af_packet.AfPacketCreateReply{}

	if err = h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, err
	}

	return reply.SwIfIndex, h.SetInterfaceTag(ifName, reply.SwIfIndex)
}

// DeleteAfPacketInterface implements AfPacket handler.
func (h *InterfaceVppHandler) DeleteAfPacketInterface(ifName string, idx uint32, afPacketIntf *interfaces.AfpacketLink) error {
	req := &af_packet.AfPacketDelete{
		HostIfName: []byte(afPacketIntf.HostIfName),
	}
	reply := &af_packet.AfPacketDeleteReply{}

	if err := h.callsChannel.SendRequest(req).ReceiveReply(reply); err != nil {
		return err
	}

	return h.RemoveInterfaceTag(ifName, idx)
}
