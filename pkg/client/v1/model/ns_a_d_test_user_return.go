// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsADTestUserReturn - Active Directory user details.
// Export NsADTestUserReturnFields for advance operations like search filter etc.
var NsADTestUserReturnFields *NsADTestUserReturn

func init() {

	NsADTestUserReturnFields = &NsADTestUserReturn{
		Username:         "username",
		PrimaryGroupName: "primary_group_name",
		PrimaryGroupId:   "primary_group_id",
	}
}

type NsADTestUserReturn struct {
	// Username - Name of the Active Directory user.
	Username string `json:"username,omitempty"`
	// PrimaryGroupName - Name of the Active Directory group the user belongs to. RBAC is based on this primary group.
	PrimaryGroupName string `json:"primary_group_name,omitempty"`
	// PrimaryGroupId - ID of the Active Directory group the user belongs to. RBAC is based on this primary group.
	PrimaryGroupId string `json:"primary_group_id,omitempty"`
	// GroupCount - Number of Active Directory groups the user belongs to.
	GroupCount int64 `json:"group_count,omitempty"`
	// Groups - List of Active Directory groups the user belongs to.
	Groups []*string `json:"groups,omitempty"`
	// Role - The role the user belongs to.
	Role *NsUserRoles `json:"role,omitempty"`
}

// sdk internal struct
type nsADTestUserReturn struct {
	// Username - Name of the Active Directory user.
	Username *string `json:"username,omitempty"`
	// PrimaryGroupName - Name of the Active Directory group the user belongs to. RBAC is based on this primary group.
	PrimaryGroupName *string `json:"primary_group_name,omitempty"`
	// PrimaryGroupId - ID of the Active Directory group the user belongs to. RBAC is based on this primary group.
	PrimaryGroupId *string `json:"primary_group_id,omitempty"`
	// GroupCount - Number of Active Directory groups the user belongs to.
	GroupCount *int64 `json:"group_count,omitempty"`
	// Groups - List of Active Directory groups the user belongs to.
	Groups []*string `json:"groups,omitempty"`
	// Role - The role the user belongs to.
	Role *NsUserRoles `json:"role,omitempty"`
}

// EncodeNsADTestUserReturn - Transform NsADTestUserReturn to nsADTestUserReturn type
func EncodeNsADTestUserReturn(request interface{}) (*nsADTestUserReturn, error) {
	reqNsADTestUserReturn := request.(*NsADTestUserReturn)
	byte, err := json.Marshal(reqNsADTestUserReturn)
	objPtr := &nsADTestUserReturn{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsADTestUserReturn Transform nsADTestUserReturn to NsADTestUserReturn type
func DecodeNsADTestUserReturn(request interface{}) (*NsADTestUserReturn, error) {
	reqNsADTestUserReturn := request.(*nsADTestUserReturn)
	byte, err := json.Marshal(reqNsADTestUserReturn)
	obj := &NsADTestUserReturn{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
