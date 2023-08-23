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
)

// checks if the PkgInventoryItemType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PkgInventoryItemType{}

// PkgInventoryItemType struct for PkgInventoryItemType
type PkgInventoryItemType struct {
	// Article Number of the inventory item.
	ArticleNumber *string `json:"articleNumber,omitempty"`
	// Description of the inventory item.
	Description *string `json:"description,omitempty"`
	// Quantity of the inventory item allocated to the package.
	Quantity *int32 `json:"quantity,omitempty"`
	// Identifier for the inventory item.
	ItemId *int32 `json:"itemId,omitempty"`
}

// NewPkgInventoryItemType instantiates a new PkgInventoryItemType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPkgInventoryItemType() *PkgInventoryItemType {
	this := PkgInventoryItemType{}
	return &this
}

// NewPkgInventoryItemTypeWithDefaults instantiates a new PkgInventoryItemType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPkgInventoryItemTypeWithDefaults() *PkgInventoryItemType {
	this := PkgInventoryItemType{}
	return &this
}

// GetArticleNumber returns the ArticleNumber field value if set, zero value otherwise.
func (o *PkgInventoryItemType) GetArticleNumber() string {
	if o == nil || IsNil(o.ArticleNumber) {
		var ret string
		return ret
	}
	return *o.ArticleNumber
}

// GetArticleNumberOk returns a tuple with the ArticleNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgInventoryItemType) GetArticleNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ArticleNumber) {
		return nil, false
	}
	return o.ArticleNumber, true
}

// HasArticleNumber returns a boolean if a field has been set.
func (o *PkgInventoryItemType) HasArticleNumber() bool {
	if o != nil && !IsNil(o.ArticleNumber) {
		return true
	}

	return false
}

// SetArticleNumber gets a reference to the given string and assigns it to the ArticleNumber field.
func (o *PkgInventoryItemType) SetArticleNumber(v string) {
	o.ArticleNumber = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PkgInventoryItemType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgInventoryItemType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PkgInventoryItemType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PkgInventoryItemType) SetDescription(v string) {
	o.Description = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PkgInventoryItemType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgInventoryItemType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PkgInventoryItemType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PkgInventoryItemType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *PkgInventoryItemType) GetItemId() int32 {
	if o == nil || IsNil(o.ItemId) {
		var ret int32
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PkgInventoryItemType) GetItemIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *PkgInventoryItemType) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given int32 and assigns it to the ItemId field.
func (o *PkgInventoryItemType) SetItemId(v int32) {
	o.ItemId = &v
}

func (o PkgInventoryItemType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PkgInventoryItemType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArticleNumber) {
		toSerialize["articleNumber"] = o.ArticleNumber
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ItemId) {
		toSerialize["itemId"] = o.ItemId
	}
	return toSerialize, nil
}

type NullablePkgInventoryItemType struct {
	value *PkgInventoryItemType
	isSet bool
}

func (v NullablePkgInventoryItemType) Get() *PkgInventoryItemType {
	return v.value
}

func (v *NullablePkgInventoryItemType) Set(val *PkgInventoryItemType) {
	v.value = val
	v.isSet = true
}

func (v NullablePkgInventoryItemType) IsSet() bool {
	return v.isSet
}

func (v *NullablePkgInventoryItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePkgInventoryItemType(val *PkgInventoryItemType) *NullablePkgInventoryItemType {
	return &NullablePkgInventoryItemType{value: val, isSet: true}
}

func (v NullablePkgInventoryItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePkgInventoryItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


