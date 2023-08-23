/*
OPERA Cloud Event Configuration API

This API caters for Event Configuration in OPERA Cloud. <br /><There are operations to post, update, fetch and delete codes such as item inventory, function spaces, menu items and many more.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CateringMenuRevType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CateringMenuRevType{}

// CateringMenuRevType struct for CateringMenuRevType
type CateringMenuRevType struct {
	// This supports all Revenue Types
	RevenueType *string `json:"revenueType,omitempty"`
	Type *MenuTypeType `json:"type,omitempty"`
	Cost *CurrencyAmountType `json:"cost,omitempty"`
	InternalQuote *CurrencyAmountType `json:"internalQuote,omitempty"`
	// This type holds cost of the Menu.
	CostMargin *float32 `json:"costMargin,omitempty"`
	Price *CurrencyAmountType `json:"price,omitempty"`
	// This type holds price margin for the given Menu.
	PriceMargin *float32 `json:"priceMargin,omitempty"`
}

// NewCateringMenuRevType instantiates a new CateringMenuRevType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCateringMenuRevType() *CateringMenuRevType {
	this := CateringMenuRevType{}
	return &this
}

// NewCateringMenuRevTypeWithDefaults instantiates a new CateringMenuRevType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCateringMenuRevTypeWithDefaults() *CateringMenuRevType {
	this := CateringMenuRevType{}
	return &this
}

// GetRevenueType returns the RevenueType field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetRevenueType() string {
	if o == nil || IsNil(o.RevenueType) {
		var ret string
		return ret
	}
	return *o.RevenueType
}

// GetRevenueTypeOk returns a tuple with the RevenueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetRevenueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RevenueType) {
		return nil, false
	}
	return o.RevenueType, true
}

// HasRevenueType returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasRevenueType() bool {
	if o != nil && !IsNil(o.RevenueType) {
		return true
	}

	return false
}

// SetRevenueType gets a reference to the given string and assigns it to the RevenueType field.
func (o *CateringMenuRevType) SetRevenueType(v string) {
	o.RevenueType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetType() MenuTypeType {
	if o == nil || IsNil(o.Type) {
		var ret MenuTypeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetTypeOk() (*MenuTypeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given MenuTypeType and assigns it to the Type field.
func (o *CateringMenuRevType) SetType(v MenuTypeType) {
	o.Type = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetCost() CurrencyAmountType {
	if o == nil || IsNil(o.Cost) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetCostOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given CurrencyAmountType and assigns it to the Cost field.
func (o *CateringMenuRevType) SetCost(v CurrencyAmountType) {
	o.Cost = &v
}

// GetInternalQuote returns the InternalQuote field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetInternalQuote() CurrencyAmountType {
	if o == nil || IsNil(o.InternalQuote) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.InternalQuote
}

// GetInternalQuoteOk returns a tuple with the InternalQuote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetInternalQuoteOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.InternalQuote) {
		return nil, false
	}
	return o.InternalQuote, true
}

// HasInternalQuote returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasInternalQuote() bool {
	if o != nil && !IsNil(o.InternalQuote) {
		return true
	}

	return false
}

// SetInternalQuote gets a reference to the given CurrencyAmountType and assigns it to the InternalQuote field.
func (o *CateringMenuRevType) SetInternalQuote(v CurrencyAmountType) {
	o.InternalQuote = &v
}

// GetCostMargin returns the CostMargin field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetCostMargin() float32 {
	if o == nil || IsNil(o.CostMargin) {
		var ret float32
		return ret
	}
	return *o.CostMargin
}

// GetCostMarginOk returns a tuple with the CostMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetCostMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.CostMargin) {
		return nil, false
	}
	return o.CostMargin, true
}

// HasCostMargin returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasCostMargin() bool {
	if o != nil && !IsNil(o.CostMargin) {
		return true
	}

	return false
}

// SetCostMargin gets a reference to the given float32 and assigns it to the CostMargin field.
func (o *CateringMenuRevType) SetCostMargin(v float32) {
	o.CostMargin = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetPrice() CurrencyAmountType {
	if o == nil || IsNil(o.Price) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetPriceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given CurrencyAmountType and assigns it to the Price field.
func (o *CateringMenuRevType) SetPrice(v CurrencyAmountType) {
	o.Price = &v
}

// GetPriceMargin returns the PriceMargin field value if set, zero value otherwise.
func (o *CateringMenuRevType) GetPriceMargin() float32 {
	if o == nil || IsNil(o.PriceMargin) {
		var ret float32
		return ret
	}
	return *o.PriceMargin
}

// GetPriceMarginOk returns a tuple with the PriceMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringMenuRevType) GetPriceMarginOk() (*float32, bool) {
	if o == nil || IsNil(o.PriceMargin) {
		return nil, false
	}
	return o.PriceMargin, true
}

// HasPriceMargin returns a boolean if a field has been set.
func (o *CateringMenuRevType) HasPriceMargin() bool {
	if o != nil && !IsNil(o.PriceMargin) {
		return true
	}

	return false
}

// SetPriceMargin gets a reference to the given float32 and assigns it to the PriceMargin field.
func (o *CateringMenuRevType) SetPriceMargin(v float32) {
	o.PriceMargin = &v
}

func (o CateringMenuRevType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CateringMenuRevType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RevenueType) {
		toSerialize["revenueType"] = o.RevenueType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.InternalQuote) {
		toSerialize["internalQuote"] = o.InternalQuote
	}
	if !IsNil(o.CostMargin) {
		toSerialize["costMargin"] = o.CostMargin
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.PriceMargin) {
		toSerialize["priceMargin"] = o.PriceMargin
	}
	return toSerialize, nil
}

type NullableCateringMenuRevType struct {
	value *CateringMenuRevType
	isSet bool
}

func (v NullableCateringMenuRevType) Get() *CateringMenuRevType {
	return v.value
}

func (v *NullableCateringMenuRevType) Set(val *CateringMenuRevType) {
	v.value = val
	v.isSet = true
}

func (v NullableCateringMenuRevType) IsSet() bool {
	return v.isSet
}

func (v *NullableCateringMenuRevType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCateringMenuRevType(val *CateringMenuRevType) *NullableCateringMenuRevType {
	return &NullableCateringMenuRevType{value: val, isSet: true}
}

func (v NullableCateringMenuRevType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCateringMenuRevType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


