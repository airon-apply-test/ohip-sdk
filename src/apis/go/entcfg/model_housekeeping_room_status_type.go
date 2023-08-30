/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
	"fmt"
)

// HousekeepingRoomStatusType Simple type for room status instructions to be used in requests for fetching housekeeping rooms. Valid status values are Clean, Dirty, Pickup, Inspected, OutOfOrder, OutOfService.
type HousekeepingRoomStatusType string

// List of housekeepingRoomStatusType
const (
	HOUSEKEEPINGROOMSTATUSTYPE_CLEAN HousekeepingRoomStatusType = "Clean"
	HOUSEKEEPINGROOMSTATUSTYPE_DIRTY HousekeepingRoomStatusType = "Dirty"
	HOUSEKEEPINGROOMSTATUSTYPE_PICKUP HousekeepingRoomStatusType = "Pickup"
	HOUSEKEEPINGROOMSTATUSTYPE_INSPECTED HousekeepingRoomStatusType = "Inspected"
	HOUSEKEEPINGROOMSTATUSTYPE_OUT_OF_ORDER HousekeepingRoomStatusType = "OutOfOrder"
	HOUSEKEEPINGROOMSTATUSTYPE_OUT_OF_SERVICE HousekeepingRoomStatusType = "OutOfService"
)

// All allowed values of HousekeepingRoomStatusType enum
var AllowedHousekeepingRoomStatusTypeEnumValues = []HousekeepingRoomStatusType{
	"Clean",
	"Dirty",
	"Pickup",
	"Inspected",
	"OutOfOrder",
	"OutOfService",
}

func (v *HousekeepingRoomStatusType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HousekeepingRoomStatusType(value)
	for _, existing := range AllowedHousekeepingRoomStatusTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HousekeepingRoomStatusType", value)
}

// NewHousekeepingRoomStatusTypeFromValue returns a pointer to a valid HousekeepingRoomStatusType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHousekeepingRoomStatusTypeFromValue(v string) (*HousekeepingRoomStatusType, error) {
	ev := HousekeepingRoomStatusType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HousekeepingRoomStatusType: valid values are %v", v, AllowedHousekeepingRoomStatusTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HousekeepingRoomStatusType) IsValid() bool {
	for _, existing := range AllowedHousekeepingRoomStatusTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to housekeepingRoomStatusType value
func (v HousekeepingRoomStatusType) Ptr() *HousekeepingRoomStatusType {
	return &v
}

type NullableHousekeepingRoomStatusType struct {
	value *HousekeepingRoomStatusType
	isSet bool
}

func (v NullableHousekeepingRoomStatusType) Get() *HousekeepingRoomStatusType {
	return v.value
}

func (v *NullableHousekeepingRoomStatusType) Set(val *HousekeepingRoomStatusType) {
	v.value = val
	v.isSet = true
}

func (v NullableHousekeepingRoomStatusType) IsSet() bool {
	return v.isSet
}

func (v *NullableHousekeepingRoomStatusType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHousekeepingRoomStatusType(val *HousekeepingRoomStatusType) *NullableHousekeepingRoomStatusType {
	return &NullableHousekeepingRoomStatusType{value: val, isSet: true}
}

func (v NullableHousekeepingRoomStatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHousekeepingRoomStatusType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

