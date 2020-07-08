// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsFanStatus Enum.

type NsFanStatus string

const (
	cNsFanStatusFanFailed  NsFanStatus = "fan_failed"
	cNsFanStatusFanOkay    NsFanStatus = "fan_okay"
	cNsFanStatusFanAlerted NsFanStatus = "fan_alerted"
	cNsFanStatusFanUnknown NsFanStatus = "fan_unknown"
)

var pNsFanStatusFanFailed NsFanStatus
var pNsFanStatusFanOkay NsFanStatus
var pNsFanStatusFanAlerted NsFanStatus
var pNsFanStatusFanUnknown NsFanStatus

// Export
var NsFanStatusFanFailed *NsFanStatus
var NsFanStatusFanOkay *NsFanStatus
var NsFanStatusFanAlerted *NsFanStatus
var NsFanStatusFanUnknown *NsFanStatus

func init() {
	pNsFanStatusFanFailed = cNsFanStatusFanFailed
	NsFanStatusFanFailed = &pNsFanStatusFanFailed

	pNsFanStatusFanOkay = cNsFanStatusFanOkay
	NsFanStatusFanOkay = &pNsFanStatusFanOkay

	pNsFanStatusFanAlerted = cNsFanStatusFanAlerted
	NsFanStatusFanAlerted = &pNsFanStatusFanAlerted

	pNsFanStatusFanUnknown = cNsFanStatusFanUnknown
	NsFanStatusFanUnknown = &pNsFanStatusFanUnknown

}
