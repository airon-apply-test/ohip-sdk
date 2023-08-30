/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
	"fmt"
)

// ResourceDiscountTypeType Indicates the option to select resources that will be discounted with the Resource Discount Percentage.
type ResourceDiscountTypeType string

// List of resourceDiscountTypeType
const (
	RESOURCEDISCOUNTTYPETYPE_ALL_DISCOUNTABLE_RESOURCES ResourceDiscountTypeType = "AllDiscountableResources"
	RESOURCEDISCOUNTTYPETYPE_RESOURCES_WITH_SAME_DISCOUNT_PERCENTAGE ResourceDiscountTypeType = "ResourcesWithSameDiscountPercentage"
	RESOURCEDISCOUNTTYPETYPE_ALL_DISCOUNTABLE_ITEMS ResourceDiscountTypeType = "AllDiscountableItems"
	RESOURCEDISCOUNTTYPETYPE_ITEMS_WITH_SAME_DISCOUNT_PERCENTAGE ResourceDiscountTypeType = "ItemsWithSameDiscountPercentage"
	RESOURCEDISCOUNTTYPETYPE_ALL_DISCOUNTABLE_MENUS_AND_MENU_ITEMS ResourceDiscountTypeType = "AllDiscountableMenusAndMenuItems"
	RESOURCEDISCOUNTTYPETYPE_MENUS_AND_MENU_ITEMS_WITH_SAME_DISCOUNT_PERCENTAGE ResourceDiscountTypeType = "MenusAndMenuItemsWithSameDiscountPercentage"
	RESOURCEDISCOUNTTYPETYPE_ALL_DISCOUNTABLE_SPACES ResourceDiscountTypeType = "AllDiscountableSpaces"
	RESOURCEDISCOUNTTYPETYPE_SPACES_WITH_SAME_DISCOUNT_PERCENTAGE ResourceDiscountTypeType = "SpacesWithSameDiscountPercentage"
	RESOURCEDISCOUNTTYPETYPE_NONE ResourceDiscountTypeType = "None"
)

// All allowed values of ResourceDiscountTypeType enum
var AllowedResourceDiscountTypeTypeEnumValues = []ResourceDiscountTypeType{
	"AllDiscountableResources",
	"ResourcesWithSameDiscountPercentage",
	"AllDiscountableItems",
	"ItemsWithSameDiscountPercentage",
	"AllDiscountableMenusAndMenuItems",
	"MenusAndMenuItemsWithSameDiscountPercentage",
	"AllDiscountableSpaces",
	"SpacesWithSameDiscountPercentage",
	"None",
}

func (v *ResourceDiscountTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ResourceDiscountTypeType(value)
	for _, existing := range AllowedResourceDiscountTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ResourceDiscountTypeType", value)
}

// NewResourceDiscountTypeTypeFromValue returns a pointer to a valid ResourceDiscountTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewResourceDiscountTypeTypeFromValue(v string) (*ResourceDiscountTypeType, error) {
	ev := ResourceDiscountTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ResourceDiscountTypeType: valid values are %v", v, AllowedResourceDiscountTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ResourceDiscountTypeType) IsValid() bool {
	for _, existing := range AllowedResourceDiscountTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceDiscountTypeType value
func (v ResourceDiscountTypeType) Ptr() *ResourceDiscountTypeType {
	return &v
}

type NullableResourceDiscountTypeType struct {
	value *ResourceDiscountTypeType
	isSet bool
}

func (v NullableResourceDiscountTypeType) Get() *ResourceDiscountTypeType {
	return v.value
}

func (v *NullableResourceDiscountTypeType) Set(val *ResourceDiscountTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceDiscountTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceDiscountTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceDiscountTypeType(val *ResourceDiscountTypeType) *NullableResourceDiscountTypeType {
	return &NullableResourceDiscountTypeType{value: val, isSet: true}
}

func (v NullableResourceDiscountTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceDiscountTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

