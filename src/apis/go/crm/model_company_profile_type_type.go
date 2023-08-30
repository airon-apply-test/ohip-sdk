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

// CompanyProfileTypeType The types of Profile handled by the web service.
type CompanyProfileTypeType string

// List of companyProfileTypeType
const (
	COMPANYPROFILETYPETYPE_AGENT CompanyProfileTypeType = "Agent"
	COMPANYPROFILETYPETYPE_COMPANY CompanyProfileTypeType = "Company"
	COMPANYPROFILETYPETYPE_GROUP CompanyProfileTypeType = "Group"
	COMPANYPROFILETYPETYPE_SOURCE CompanyProfileTypeType = "Source"
)

// All allowed values of CompanyProfileTypeType enum
var AllowedCompanyProfileTypeTypeEnumValues = []CompanyProfileTypeType{
	"Agent",
	"Company",
	"Group",
	"Source",
}

func (v *CompanyProfileTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CompanyProfileTypeType(value)
	for _, existing := range AllowedCompanyProfileTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CompanyProfileTypeType", value)
}

// NewCompanyProfileTypeTypeFromValue returns a pointer to a valid CompanyProfileTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCompanyProfileTypeTypeFromValue(v string) (*CompanyProfileTypeType, error) {
	ev := CompanyProfileTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CompanyProfileTypeType: valid values are %v", v, AllowedCompanyProfileTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CompanyProfileTypeType) IsValid() bool {
	for _, existing := range AllowedCompanyProfileTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to companyProfileTypeType value
func (v CompanyProfileTypeType) Ptr() *CompanyProfileTypeType {
	return &v
}

type NullableCompanyProfileTypeType struct {
	value *CompanyProfileTypeType
	isSet bool
}

func (v NullableCompanyProfileTypeType) Get() *CompanyProfileTypeType {
	return v.value
}

func (v *NullableCompanyProfileTypeType) Set(val *CompanyProfileTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyProfileTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyProfileTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyProfileTypeType(val *CompanyProfileTypeType) *NullableCompanyProfileTypeType {
	return &NullableCompanyProfileTypeType{value: val, isSet: true}
}

func (v NullableCompanyProfileTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyProfileTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

