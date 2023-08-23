/*
OPERA Cloud Rate API

APIs to cater for Rate Availability functionality in OPERA Cloud. <br /><br /> Rate Management provides all the tools you need to effectively define and manage the rate structures for a property in OPERA Cloud. Some of the things you can do include creating and managing rate codes, rate classes, rate categories, display sets, rate strategies, as well as managing promotion groups and codes.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MasterType the model 'MasterType'
type MasterType string

// List of masterType
const (
	COUNTRY MasterType = "Country"
	STATE MasterType = "State"
	ADDRESS_TYPE MasterType = "AddressType"
	PHONE_TYPE MasterType = "PhoneType"
	RATE_CATEGORY MasterType = "RateCategory"
	CALCULATION_RULE MasterType = "CalculationRule"
	POSTING_RYTHYM MasterType = "PostingRythym"
	BILLING_INSTRUCTION MasterType = "BillingInstruction"
	TRANSACTION_CODE MasterType = "TransactionCode"
	DISPLAY_SET MasterType = "DisplaySet"
	MAILING_ACTIONS MasterType = "MailingActions"
	DISTANCE_TYPE MasterType = "DistanceType"
	DISTRICT MasterType = "District"
	TERRITORY MasterType = "Territory"
	FISCAL_REGION MasterType = "FiscalRegion"
	INVENTORY_ITEM MasterType = "InventoryItem"
	PACKAGE MasterType = "Package"
	ROOM_FEATURE_PREFERENCE MasterType = "RoomFeaturePreference"
	SPECIAL_PREFERENCE MasterType = "SpecialPreference"
	PROMOTION MasterType = "Promotion"
	DEPARTMENT MasterType = "Department"
	RESERVATION_PREFERENCE MasterType = "ReservationPreference"
	FACILITY_TASK MasterType = "FacilityTask"
	ROOM_TYPE MasterType = "RoomType"
	RATE_CODE MasterType = "RateCode"
	OUT_OF_ORDER_REASON MasterType = "OutOfOrderReason"
	BLOCK MasterType = "Block"
)

// All allowed values of MasterType enum
var AllowedMasterTypeEnumValues = []MasterType{
	"Country",
	"State",
	"AddressType",
	"PhoneType",
	"RateCategory",
	"CalculationRule",
	"PostingRythym",
	"BillingInstruction",
	"TransactionCode",
	"DisplaySet",
	"MailingActions",
	"DistanceType",
	"District",
	"Territory",
	"FiscalRegion",
	"InventoryItem",
	"Package",
	"RoomFeaturePreference",
	"SpecialPreference",
	"Promotion",
	"Department",
	"ReservationPreference",
	"FacilityTask",
	"RoomType",
	"RateCode",
	"OutOfOrderReason",
	"Block",
}

func (v *MasterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MasterType(value)
	for _, existing := range AllowedMasterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MasterType", value)
}

// NewMasterTypeFromValue returns a pointer to a valid MasterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMasterTypeFromValue(v string) (*MasterType, error) {
	ev := MasterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MasterType: valid values are %v", v, AllowedMasterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MasterType) IsValid() bool {
	for _, existing := range AllowedMasterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to masterType value
func (v MasterType) Ptr() *MasterType {
	return &v
}

type NullableMasterType struct {
	value *MasterType
	isSet bool
}

func (v NullableMasterType) Get() *MasterType {
	return v.value
}

func (v *NullableMasterType) Set(val *MasterType) {
	v.value = val
	v.isSet = true
}

func (v NullableMasterType) IsSet() bool {
	return v.isSet
}

func (v *NullableMasterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMasterType(val *MasterType) *NullableMasterType {
	return &NullableMasterType{value: val, isSet: true}
}

func (v NullableMasterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMasterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

