// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsDiscoveredGroupInfo - Discovered group details.
// Export NsDiscoveredGroupInfoFields for advance operations like search filter etc.
var NsDiscoveredGroupInfoFields *NsDiscoveredGroupInfo

func init() {

	NsDiscoveredGroupInfoFields = &NsDiscoveredGroupInfo{
		ID:             "id",
		Name:           "name",
		VersionCurrent: "version_current",
		ManagementIp:   "management_ip",
		DiscoveryIp:    "discovery_ip",
	}
}

type NsDiscoveredGroupInfo struct {
	// ID - Identifier of the group.
	ID string `json:"id,omitempty"`
	// Name - Name of the group.
	Name string `json:"name,omitempty"`
	// VersionCurrent - Version of software running on the group.
	VersionCurrent string `json:"version_current,omitempty"`
	// CountOfMembers - Number of member arrays in the group.
	CountOfMembers int64 `json:"count_of_members,omitempty"`
	// ManagementIp - The management IP address of the group.
	ManagementIp string `json:"management_ip,omitempty"`
	// DiscoveryIp - The discovery IP address of the group.
	DiscoveryIp string `json:"discovery_ip,omitempty"`
	// IscsiEnabled - Whether iscsi is enabled on the group.
	IscsiEnabled *bool `json:"iscsi_enabled,omitempty"`
	// FcEnabled - Whether fibre channel is enabled on the group.
	FcEnabled *bool `json:"fc_enabled,omitempty"`
	// ArrayMemberList - Members of this group.
	ArrayMemberList []*NsArraySummaryInfo `json:"array_member_list,omitempty"`
}

// sdk internal struct
type nsDiscoveredGroupInfo struct {
	// ID - Identifier of the group.
	ID *string `json:"id,omitempty"`
	// Name - Name of the group.
	Name *string `json:"name,omitempty"`
	// VersionCurrent - Version of software running on the group.
	VersionCurrent *string `json:"version_current,omitempty"`
	// CountOfMembers - Number of member arrays in the group.
	CountOfMembers *int64 `json:"count_of_members,omitempty"`
	// ManagementIp - The management IP address of the group.
	ManagementIp *string `json:"management_ip,omitempty"`
	// DiscoveryIp - The discovery IP address of the group.
	DiscoveryIp *string `json:"discovery_ip,omitempty"`
	// IscsiEnabled - Whether iscsi is enabled on the group.
	IscsiEnabled *bool `json:"iscsi_enabled,omitempty"`
	// FcEnabled - Whether fibre channel is enabled on the group.
	FcEnabled *bool `json:"fc_enabled,omitempty"`
	// ArrayMemberList - Members of this group.
	ArrayMemberList []*NsArraySummaryInfo `json:"array_member_list,omitempty"`
}

// EncodeNsDiscoveredGroupInfo - Transform NsDiscoveredGroupInfo to nsDiscoveredGroupInfo type
func EncodeNsDiscoveredGroupInfo(request interface{}) (*nsDiscoveredGroupInfo, error) {
	reqNsDiscoveredGroupInfo := request.(*NsDiscoveredGroupInfo)
	byte, err := json.Marshal(reqNsDiscoveredGroupInfo)
	objPtr := &nsDiscoveredGroupInfo{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsDiscoveredGroupInfo Transform nsDiscoveredGroupInfo to NsDiscoveredGroupInfo type
func DecodeNsDiscoveredGroupInfo(request interface{}) (*NsDiscoveredGroupInfo, error) {
	reqNsDiscoveredGroupInfo := request.(*nsDiscoveredGroupInfo)
	byte, err := json.Marshal(reqNsDiscoveredGroupInfo)
	obj := &NsDiscoveredGroupInfo{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
