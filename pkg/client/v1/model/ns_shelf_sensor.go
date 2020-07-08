// Copyright 2020 Hewlett Packard Enterprise Development LP

package model

import (
	"encoding/json"
)

// NsShelfSensor - A shelf sensor data.
// Export NsShelfSensorFields for advance operations like search filter etc.
var NsShelfSensorFields *NsShelfSensor

func init() {

	NsShelfSensorFields = &NsShelfSensor{
		Name:        "name",
		DisplayName: "display_name",
		Location:    "location",
	}
}

type NsShelfSensor struct {
	// Type - Type of the sensor.
	Type *NsShelfSensorType `json:"type,omitempty"`
	// Name - Internal name of the sensor.
	Name string `json:"name,omitempty"`
	// DisplayName - Name for display purpose.
	DisplayName string `json:"display_name,omitempty"`
	// Location - Location of the sensor.
	Location string `json:"location,omitempty"`
	// Cid - Which controller this sensor applies to.
	Cid *NsShelfCtrlrSide `json:"cid,omitempty"`
	// Value - Value of the sensor reading.
	Value int64 `json:"value,omitempty"`
	// Status - Sensor status.
	Status *NsShelfSensorState `json:"status,omitempty"`
}

// sdk internal struct
type nsShelfSensor struct {
	// Type - Type of the sensor.
	Type *NsShelfSensorType `json:"type,omitempty"`
	// Name - Internal name of the sensor.
	Name *string `json:"name,omitempty"`
	// DisplayName - Name for display purpose.
	DisplayName *string `json:"display_name,omitempty"`
	// Location - Location of the sensor.
	Location *string `json:"location,omitempty"`
	// Cid - Which controller this sensor applies to.
	Cid *NsShelfCtrlrSide `json:"cid,omitempty"`
	// Value - Value of the sensor reading.
	Value *int64 `json:"value,omitempty"`
	// Status - Sensor status.
	Status *NsShelfSensorState `json:"status,omitempty"`
}

// EncodeNsShelfSensor - Transform NsShelfSensor to nsShelfSensor type
func EncodeNsShelfSensor(request interface{}) (*nsShelfSensor, error) {
	reqNsShelfSensor := request.(*NsShelfSensor)
	byte, err := json.Marshal(reqNsShelfSensor)
	objPtr := &nsShelfSensor{}
	err = json.Unmarshal(byte, objPtr)
	return objPtr, err
}

// DecodeNsShelfSensor Transform nsShelfSensor to NsShelfSensor type
func DecodeNsShelfSensor(request interface{}) (*NsShelfSensor, error) {
	reqNsShelfSensor := request.(*nsShelfSensor)
	byte, err := json.Marshal(reqNsShelfSensor)
	obj := &NsShelfSensor{}
	err = json.Unmarshal(byte, obj)
	return obj, err
}
