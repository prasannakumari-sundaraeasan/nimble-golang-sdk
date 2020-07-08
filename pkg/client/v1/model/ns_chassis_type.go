// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsChassisType Enum.

type NsChassisType string

const (
	cNsChassisTypeChassisSmc4u24  NsChassisType = "chassis_smc_4u24"
	cNsChassisTypeChassisUnknown  NsChassisType = "chassis_unknown"
	cNsChassisTypeChassisSmc3u16  NsChassisType = "chassis_smc_3u16"
	cNsChassisTypeChassisNmbl2u12 NsChassisType = "chassis_nmbl_2u12"
	cNsChassisTypeChassisNmbl4u24 NsChassisType = "chassis_nmbl_4u24"
)

var pNsChassisTypeChassisSmc4u24 NsChassisType
var pNsChassisTypeChassisUnknown NsChassisType
var pNsChassisTypeChassisSmc3u16 NsChassisType
var pNsChassisTypeChassisNmbl2u12 NsChassisType
var pNsChassisTypeChassisNmbl4u24 NsChassisType

// Export
var NsChassisTypeChassisSmc4u24 *NsChassisType
var NsChassisTypeChassisUnknown *NsChassisType
var NsChassisTypeChassisSmc3u16 *NsChassisType
var NsChassisTypeChassisNmbl2u12 *NsChassisType
var NsChassisTypeChassisNmbl4u24 *NsChassisType

func init() {
	pNsChassisTypeChassisSmc4u24 = cNsChassisTypeChassisSmc4u24
	NsChassisTypeChassisSmc4u24 = &pNsChassisTypeChassisSmc4u24

	pNsChassisTypeChassisUnknown = cNsChassisTypeChassisUnknown
	NsChassisTypeChassisUnknown = &pNsChassisTypeChassisUnknown

	pNsChassisTypeChassisSmc3u16 = cNsChassisTypeChassisSmc3u16
	NsChassisTypeChassisSmc3u16 = &pNsChassisTypeChassisSmc3u16

	pNsChassisTypeChassisNmbl2u12 = cNsChassisTypeChassisNmbl2u12
	NsChassisTypeChassisNmbl2u12 = &pNsChassisTypeChassisNmbl2u12

	pNsChassisTypeChassisNmbl4u24 = cNsChassisTypeChassisNmbl4u24
	NsChassisTypeChassisNmbl4u24 = &pNsChassisTypeChassisNmbl4u24

}
