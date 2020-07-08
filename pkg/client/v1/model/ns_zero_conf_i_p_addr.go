// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsZeroConfIPAddr - Zero Conf of array.
// Export NsZeroConfIPAddrFields for advance operations like search filter etc.
var NsZeroConfIPAddrFields *NsZeroConfIPAddr

func init() {

	NsZeroConfIPAddrFields = &NsZeroConfIPAddr{
		Nic:          "nic",
		LocalIpaddr:  "local_ipaddr",
		RemoteIpaddr: "remote_ipaddr",
	}
}

type NsZeroConfIPAddr struct {
	// Nic - Nic of array.
	Nic string `json:"nic,omitempty"`
	// LocalIpaddr - Local IP address of array.
	LocalIpaddr string `json:"local_ipaddr,omitempty"`
	// RemoteIpaddr - Remote IP address of array.
	RemoteIpaddr string `json:"remote_ipaddr,omitempty"`
}

// sdk internal struct
type nsZeroConfIPAddr struct {
	// Nic - Nic of array.
	Nic *string `json:"nic,omitempty"`
	// LocalIpaddr - Local IP address of array.
	LocalIpaddr *string `json:"local_ipaddr,omitempty"`
	// RemoteIpaddr - Remote IP address of array.
	RemoteIpaddr *string `json:"remote_ipaddr,omitempty"`
}

// EncodeNsZeroConfIPAddr - Transform NsZeroConfIPAddr to nsZeroConfIPAddr type
func EncodeNsZeroConfIPAddr(request interface{}) (*nsZeroConfIPAddr, error) {
	reqNsZeroConfIPAddr := request.(*NsZeroConfIPAddr)
	byte, err := json.Marshal(reqNsZeroConfIPAddr)
	objPtr := &nsZeroConfIPAddr{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsZeroConfIPAddr Transform nsZeroConfIPAddr to NsZeroConfIPAddr type
func DecodeNsZeroConfIPAddr(request interface{}) (*NsZeroConfIPAddr, error) {
	reqNsZeroConfIPAddr := request.(*nsZeroConfIPAddr)
	byte, err := json.Marshal(reqNsZeroConfIPAddr)
	obj := &NsZeroConfIPAddr{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
