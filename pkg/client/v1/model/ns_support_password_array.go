// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsSupportPasswordArray - Support password blobs for an array.
// Export NsSupportPasswordArrayFields for advance operations like search filter etc.
var NsSupportPasswordArrayFields *NsSupportPasswordArray

func init() {

	NsSupportPasswordArrayFields = &NsSupportPasswordArray{
		ArrayName: "array_name",
		ArraySn:   "array_sn",
		Model:     "model",
	}
}

type NsSupportPasswordArray struct {
	// ArrayName - The name of the array.
	ArrayName string `json:"array_name,omitempty"`
	// ArraySn - The serial number of the array.
	ArraySn string `json:"array_sn,omitempty"`
	// Model - The model of the array.
	Model string `json:"model,omitempty"`
	// BlobCount - The number of blobs stored for the array.
	BlobCount int64 `json:"blob_count,omitempty"`
	// BlobList - The blobs stored for the array.
	BlobList []*NsSupportPasswordObject `json:"blob_list,omitempty"`
}

// sdk internal struct
type nsSupportPasswordArray struct {
	// ArrayName - The name of the array.
	ArrayName *string `json:"array_name,omitempty"`
	// ArraySn - The serial number of the array.
	ArraySn *string `json:"array_sn,omitempty"`
	// Model - The model of the array.
	Model *string `json:"model,omitempty"`
	// BlobCount - The number of blobs stored for the array.
	BlobCount *int64 `json:"blob_count,omitempty"`
	// BlobList - The blobs stored for the array.
	BlobList []*NsSupportPasswordObject `json:"blob_list,omitempty"`
}

// EncodeNsSupportPasswordArray - Transform NsSupportPasswordArray to nsSupportPasswordArray type
func EncodeNsSupportPasswordArray(request interface{}) (*nsSupportPasswordArray, error) {
	reqNsSupportPasswordArray := request.(*NsSupportPasswordArray)
	byte, err := json.Marshal(reqNsSupportPasswordArray)
	objPtr := &nsSupportPasswordArray{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsSupportPasswordArray Transform nsSupportPasswordArray to NsSupportPasswordArray type
func DecodeNsSupportPasswordArray(request interface{}) (*NsSupportPasswordArray, error) {
	reqNsSupportPasswordArray := request.(*nsSupportPasswordArray)
	byte, err := json.Marshal(reqNsSupportPasswordArray)
	obj := &NsSupportPasswordArray{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
