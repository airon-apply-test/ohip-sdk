/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
	"fmt"
)

// HotelInterfaceTypeType XML Posting Interface
type HotelInterfaceTypeType string

// List of hotelInterfaceTypeType
const (
	HOTELINTERFACETYPETYPE_BMS HotelInterfaceTypeType = "Bms"
	HOTELINTERFACETYPETYPE_CAS HotelInterfaceTypeType = "Cas"
	HOTELINTERFACETYPETYPE_CCW HotelInterfaceTypeType = "Ccw"
	HOTELINTERFACETYPETYPE_DLS HotelInterfaceTypeType = "Dls"
	HOTELINTERFACETYPETYPE_EFT HotelInterfaceTypeType = "Eft"
	HOTELINTERFACETYPETYPE_EXP HotelInterfaceTypeType = "Exp"
	HOTELINTERFACETYPETYPE_MAK HotelInterfaceTypeType = "Mak"
	HOTELINTERFACETYPETYPE_MBS HotelInterfaceTypeType = "Mbs"
	HOTELINTERFACETYPETYPE_MSC HotelInterfaceTypeType = "Msc"
	HOTELINTERFACETYPETYPE_PBX HotelInterfaceTypeType = "Pbx"
	HOTELINTERFACETYPETYPE_POS HotelInterfaceTypeType = "Pos"
	HOTELINTERFACETYPETYPE_SVS HotelInterfaceTypeType = "Svs"
	HOTELINTERFACETYPETYPE_TIK HotelInterfaceTypeType = "Tik"
	HOTELINTERFACETYPETYPE_VID HotelInterfaceTypeType = "Vid"
	HOTELINTERFACETYPETYPE_VMS HotelInterfaceTypeType = "Vms"
	HOTELINTERFACETYPETYPE_WWW HotelInterfaceTypeType = "Www"
	HOTELINTERFACETYPETYPE_XML HotelInterfaceTypeType = "Xml"
)

// All allowed values of HotelInterfaceTypeType enum
var AllowedHotelInterfaceTypeTypeEnumValues = []HotelInterfaceTypeType{
	"Bms",
	"Cas",
	"Ccw",
	"Dls",
	"Eft",
	"Exp",
	"Mak",
	"Mbs",
	"Msc",
	"Pbx",
	"Pos",
	"Svs",
	"Tik",
	"Vid",
	"Vms",
	"Www",
	"Xml",
}

func (v *HotelInterfaceTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HotelInterfaceTypeType(value)
	for _, existing := range AllowedHotelInterfaceTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HotelInterfaceTypeType", value)
}

// NewHotelInterfaceTypeTypeFromValue returns a pointer to a valid HotelInterfaceTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHotelInterfaceTypeTypeFromValue(v string) (*HotelInterfaceTypeType, error) {
	ev := HotelInterfaceTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HotelInterfaceTypeType: valid values are %v", v, AllowedHotelInterfaceTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HotelInterfaceTypeType) IsValid() bool {
	for _, existing := range AllowedHotelInterfaceTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to hotelInterfaceTypeType value
func (v HotelInterfaceTypeType) Ptr() *HotelInterfaceTypeType {
	return &v
}

type NullableHotelInterfaceTypeType struct {
	value *HotelInterfaceTypeType
	isSet bool
}

func (v NullableHotelInterfaceTypeType) Get() *HotelInterfaceTypeType {
	return v.value
}

func (v *NullableHotelInterfaceTypeType) Set(val *HotelInterfaceTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableHotelInterfaceTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableHotelInterfaceTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHotelInterfaceTypeType(val *HotelInterfaceTypeType) *NullableHotelInterfaceTypeType {
	return &NullableHotelInterfaceTypeType{value: val, isSet: true}
}

func (v NullableHotelInterfaceTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHotelInterfaceTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

