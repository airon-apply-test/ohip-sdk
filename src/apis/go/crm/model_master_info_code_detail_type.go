/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm

import (
	"encoding/json"
	"fmt"
)

// MasterInfoCodeDetailType the model 'MasterInfoCodeDetailType'
type MasterInfoCodeDetailType string

// List of masterInfoCodeDetailType
const (
	MASTERINFOCODEDETAILTYPE_LONG_DESCRIPTION MasterInfoCodeDetailType = "LongDescription"
	MASTERINFOCODEDETAILTYPE_SHORT_DESCRIPTION MasterInfoCodeDetailType = "ShortDescription"
)

// All allowed values of MasterInfoCodeDetailType enum
var AllowedMasterInfoCodeDetailTypeEnumValues = []MasterInfoCodeDetailType{
	"LongDescription",
	"ShortDescription",
}

func (v *MasterInfoCodeDetailType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MasterInfoCodeDetailType(value)
	for _, existing := range AllowedMasterInfoCodeDetailTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MasterInfoCodeDetailType", value)
}

// NewMasterInfoCodeDetailTypeFromValue returns a pointer to a valid MasterInfoCodeDetailType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMasterInfoCodeDetailTypeFromValue(v string) (*MasterInfoCodeDetailType, error) {
	ev := MasterInfoCodeDetailType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MasterInfoCodeDetailType: valid values are %v", v, AllowedMasterInfoCodeDetailTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MasterInfoCodeDetailType) IsValid() bool {
	for _, existing := range AllowedMasterInfoCodeDetailTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to masterInfoCodeDetailType value
func (v MasterInfoCodeDetailType) Ptr() *MasterInfoCodeDetailType {
	return &v
}

type NullableMasterInfoCodeDetailType struct {
	value *MasterInfoCodeDetailType
	isSet bool
}

func (v NullableMasterInfoCodeDetailType) Get() *MasterInfoCodeDetailType {
	return v.value
}

func (v *NullableMasterInfoCodeDetailType) Set(val *MasterInfoCodeDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterInfoCodeDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterInfoCodeDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterInfoCodeDetailType(val *MasterInfoCodeDetailType) *NullableMasterInfoCodeDetailType {
	return &NullableMasterInfoCodeDetailType{value: val, isSet: true}
}

func (v NullableMasterInfoCodeDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterInfoCodeDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

