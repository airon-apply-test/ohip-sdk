/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fof

import (
	"encoding/json"
)

// checks if the FixedChargeDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FixedChargeDetailType{}

// FixedChargeDetailType Fixed charge amount could be specified by flat fee or be a percentage of the rate amount.
type FixedChargeDetailType struct {
	Transaction *CodeDescriptionType `json:"transaction,omitempty"`
	// Quantity of the product.
	Quantity *int32 `json:"quantity,omitempty"`
	ChargeAmount *CurrencyAmountType `json:"chargeAmount,omitempty"`
	// Percentage of the rate amount.
	Percent *float32 `json:"percent,omitempty"`
	// Additional information regarding the fixed charge.
	Supplement *string `json:"supplement,omitempty"`
	Article *CodeDescriptionType `json:"article,omitempty"`
	// Holds number of comp or cash room night to allocate.
	RoomNights *int32 `json:"roomNights,omitempty"`
}

// NewFixedChargeDetailType instantiates a new FixedChargeDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFixedChargeDetailType() *FixedChargeDetailType {
	this := FixedChargeDetailType{}
	return &this
}

// NewFixedChargeDetailTypeWithDefaults instantiates a new FixedChargeDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFixedChargeDetailTypeWithDefaults() *FixedChargeDetailType {
	this := FixedChargeDetailType{}
	return &this
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetTransaction() CodeDescriptionType {
	if o == nil || IsNil(o.Transaction) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetTransactionOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given CodeDescriptionType and assigns it to the Transaction field.
func (o *FixedChargeDetailType) SetTransaction(v CodeDescriptionType) {
	o.Transaction = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *FixedChargeDetailType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetChargeAmount() CurrencyAmountType {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetChargeAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given CurrencyAmountType and assigns it to the ChargeAmount field.
func (o *FixedChargeDetailType) SetChargeAmount(v CurrencyAmountType) {
	o.ChargeAmount = &v
}

// GetPercent returns the Percent field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetPercent() float32 {
	if o == nil || IsNil(o.Percent) {
		var ret float32
		return ret
	}
	return *o.Percent
}

// GetPercentOk returns a tuple with the Percent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.Percent) {
		return nil, false
	}
	return o.Percent, true
}

// HasPercent returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasPercent() bool {
	if o != nil && !IsNil(o.Percent) {
		return true
	}

	return false
}

// SetPercent gets a reference to the given float32 and assigns it to the Percent field.
func (o *FixedChargeDetailType) SetPercent(v float32) {
	o.Percent = &v
}

// GetSupplement returns the Supplement field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetSupplement() string {
	if o == nil || IsNil(o.Supplement) {
		var ret string
		return ret
	}
	return *o.Supplement
}

// GetSupplementOk returns a tuple with the Supplement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetSupplementOk() (*string, bool) {
	if o == nil || IsNil(o.Supplement) {
		return nil, false
	}
	return o.Supplement, true
}

// HasSupplement returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasSupplement() bool {
	if o != nil && !IsNil(o.Supplement) {
		return true
	}

	return false
}

// SetSupplement gets a reference to the given string and assigns it to the Supplement field.
func (o *FixedChargeDetailType) SetSupplement(v string) {
	o.Supplement = &v
}

// GetArticle returns the Article field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetArticle() CodeDescriptionType {
	if o == nil || IsNil(o.Article) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Article
}

// GetArticleOk returns a tuple with the Article field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetArticleOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Article) {
		return nil, false
	}
	return o.Article, true
}

// HasArticle returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasArticle() bool {
	if o != nil && !IsNil(o.Article) {
		return true
	}

	return false
}

// SetArticle gets a reference to the given CodeDescriptionType and assigns it to the Article field.
func (o *FixedChargeDetailType) SetArticle(v CodeDescriptionType) {
	o.Article = &v
}

// GetRoomNights returns the RoomNights field value if set, zero value otherwise.
func (o *FixedChargeDetailType) GetRoomNights() int32 {
	if o == nil || IsNil(o.RoomNights) {
		var ret int32
		return ret
	}
	return *o.RoomNights
}

// GetRoomNightsOk returns a tuple with the RoomNights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FixedChargeDetailType) GetRoomNightsOk() (*int32, bool) {
	if o == nil || IsNil(o.RoomNights) {
		return nil, false
	}
	return o.RoomNights, true
}

// HasRoomNights returns a boolean if a field has been set.
func (o *FixedChargeDetailType) HasRoomNights() bool {
	if o != nil && !IsNil(o.RoomNights) {
		return true
	}

	return false
}

// SetRoomNights gets a reference to the given int32 and assigns it to the RoomNights field.
func (o *FixedChargeDetailType) SetRoomNights(v int32) {
	o.RoomNights = &v
}

func (o FixedChargeDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FixedChargeDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.ChargeAmount) {
		toSerialize["chargeAmount"] = o.ChargeAmount
	}
	if !IsNil(o.Percent) {
		toSerialize["percent"] = o.Percent
	}
	if !IsNil(o.Supplement) {
		toSerialize["supplement"] = o.Supplement
	}
	if !IsNil(o.Article) {
		toSerialize["article"] = o.Article
	}
	if !IsNil(o.RoomNights) {
		toSerialize["roomNights"] = o.RoomNights
	}
	return toSerialize, nil
}

type NullableFixedChargeDetailType struct {
	value *FixedChargeDetailType
	isSet bool
}

func (v NullableFixedChargeDetailType) Get() *FixedChargeDetailType {
	return v.value
}

func (v *NullableFixedChargeDetailType) Set(val *FixedChargeDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableFixedChargeDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableFixedChargeDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFixedChargeDetailType(val *FixedChargeDetailType) *NullableFixedChargeDetailType {
	return &NullableFixedChargeDetailType{value: val, isSet: true}
}

func (v NullableFixedChargeDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFixedChargeDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


