// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsKmipProtocol Enum.

type NsKmipProtocol string

const (
	cNsKmipProtocolKmip11 NsKmipProtocol = "KMIP1_1"
	cNsKmipProtocolKmip12 NsKmipProtocol = "KMIP1_2"
	cNsKmipProtocolKmip10 NsKmipProtocol = "KMIP1_0"
	cNsKmipProtocolKmip13 NsKmipProtocol = "KMIP1_3"
)

var pNsKmipProtocolKmip11 NsKmipProtocol
var pNsKmipProtocolKmip12 NsKmipProtocol
var pNsKmipProtocolKmip10 NsKmipProtocol
var pNsKmipProtocolKmip13 NsKmipProtocol

// Export
var NsKmipProtocolKmip11 *NsKmipProtocol
var NsKmipProtocolKmip12 *NsKmipProtocol
var NsKmipProtocolKmip10 *NsKmipProtocol
var NsKmipProtocolKmip13 *NsKmipProtocol

func init() {
	pNsKmipProtocolKmip11 = cNsKmipProtocolKmip11
	NsKmipProtocolKmip11 = &pNsKmipProtocolKmip11

	pNsKmipProtocolKmip12 = cNsKmipProtocolKmip12
	NsKmipProtocolKmip12 = &pNsKmipProtocolKmip12

	pNsKmipProtocolKmip10 = cNsKmipProtocolKmip10
	NsKmipProtocolKmip10 = &pNsKmipProtocolKmip10

	pNsKmipProtocolKmip13 = cNsKmipProtocolKmip13
	NsKmipProtocolKmip13 = &pNsKmipProtocolKmip13

}
