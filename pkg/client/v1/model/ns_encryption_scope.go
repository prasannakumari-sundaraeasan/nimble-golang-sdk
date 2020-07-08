// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsEncryptionScope Enum.

type NsEncryptionScope string

const (
	cNsEncryptionScopeVolume NsEncryptionScope = "volume"
	cNsEncryptionScopePool   NsEncryptionScope = "pool"
	cNsEncryptionScopeNone   NsEncryptionScope = "none"
	cNsEncryptionScopeGroup  NsEncryptionScope = "group"
)

var pNsEncryptionScopeVolume NsEncryptionScope
var pNsEncryptionScopePool NsEncryptionScope
var pNsEncryptionScopeNone NsEncryptionScope
var pNsEncryptionScopeGroup NsEncryptionScope

// Export
var NsEncryptionScopeVolume *NsEncryptionScope
var NsEncryptionScopePool *NsEncryptionScope
var NsEncryptionScopeNone *NsEncryptionScope
var NsEncryptionScopeGroup *NsEncryptionScope

func init() {
	pNsEncryptionScopeVolume = cNsEncryptionScopeVolume
	NsEncryptionScopeVolume = &pNsEncryptionScopeVolume

	pNsEncryptionScopePool = cNsEncryptionScopePool
	NsEncryptionScopePool = &pNsEncryptionScopePool

	pNsEncryptionScopeNone = cNsEncryptionScopeNone
	NsEncryptionScopeNone = &pNsEncryptionScopeNone

	pNsEncryptionScopeGroup = cNsEncryptionScopeGroup
	NsEncryptionScopeGroup = &pNsEncryptionScopeGroup

}
