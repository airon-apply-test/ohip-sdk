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

// checks if the TrxInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrxInfoType{}

// TrxInfoType Transaction codes info.
type TrxInfoType struct {
	// Transaction codes info.
	Description *string `json:"description,omitempty"`
	// Category of the transaction code.
	TransactionGroup *string `json:"transactionGroup,omitempty"`
	// Sub category of the transaction code.
	TransactionSubGroup *string `json:"transactionSubGroup,omitempty"`
	// Unique Universal product code of the transaction code.
	UniversalProductCode *string `json:"universalProductCode,omitempty"`
	// This is the Routing Instruction Id attached with Reservation. It is only used for internal purpose. It should not be used by external vendor or consumer.
	RoutingInstructionsId *float32 `json:"routingInstructionsId,omitempty"`
	// The List of Articles defined for this transaction code, when using the Articles functionality.
	Articles []ArticleInfoType `json:"articles,omitempty"`
	// Contains service type for transaction code.
	TrxServiceType *string `json:"trxServiceType,omitempty"`
	// Unique identifier for the Transaction code.
	TransactionCode *string `json:"transactionCode,omitempty"`
	// Hotel context of the Transaction code.
	HotelId *string `json:"hotelId,omitempty"`
	// Print receipt flag that tells whether the transaction receipt is to be printed or not. This is based on the transaction code.
	PrintTrxReceipt *bool `json:"printTrxReceipt,omitempty"`
}

// NewTrxInfoType instantiates a new TrxInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrxInfoType() *TrxInfoType {
	this := TrxInfoType{}
	return &this
}

// NewTrxInfoTypeWithDefaults instantiates a new TrxInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrxInfoTypeWithDefaults() *TrxInfoType {
	this := TrxInfoType{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TrxInfoType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TrxInfoType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TrxInfoType) SetDescription(v string) {
	o.Description = &v
}

// GetTransactionGroup returns the TransactionGroup field value if set, zero value otherwise.
func (o *TrxInfoType) GetTransactionGroup() string {
	if o == nil || IsNil(o.TransactionGroup) {
		var ret string
		return ret
	}
	return *o.TransactionGroup
}

// GetTransactionGroupOk returns a tuple with the TransactionGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetTransactionGroupOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionGroup) {
		return nil, false
	}
	return o.TransactionGroup, true
}

// HasTransactionGroup returns a boolean if a field has been set.
func (o *TrxInfoType) HasTransactionGroup() bool {
	if o != nil && !IsNil(o.TransactionGroup) {
		return true
	}

	return false
}

// SetTransactionGroup gets a reference to the given string and assigns it to the TransactionGroup field.
func (o *TrxInfoType) SetTransactionGroup(v string) {
	o.TransactionGroup = &v
}

// GetTransactionSubGroup returns the TransactionSubGroup field value if set, zero value otherwise.
func (o *TrxInfoType) GetTransactionSubGroup() string {
	if o == nil || IsNil(o.TransactionSubGroup) {
		var ret string
		return ret
	}
	return *o.TransactionSubGroup
}

// GetTransactionSubGroupOk returns a tuple with the TransactionSubGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetTransactionSubGroupOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionSubGroup) {
		return nil, false
	}
	return o.TransactionSubGroup, true
}

// HasTransactionSubGroup returns a boolean if a field has been set.
func (o *TrxInfoType) HasTransactionSubGroup() bool {
	if o != nil && !IsNil(o.TransactionSubGroup) {
		return true
	}

	return false
}

// SetTransactionSubGroup gets a reference to the given string and assigns it to the TransactionSubGroup field.
func (o *TrxInfoType) SetTransactionSubGroup(v string) {
	o.TransactionSubGroup = &v
}

// GetUniversalProductCode returns the UniversalProductCode field value if set, zero value otherwise.
func (o *TrxInfoType) GetUniversalProductCode() string {
	if o == nil || IsNil(o.UniversalProductCode) {
		var ret string
		return ret
	}
	return *o.UniversalProductCode
}

// GetUniversalProductCodeOk returns a tuple with the UniversalProductCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetUniversalProductCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UniversalProductCode) {
		return nil, false
	}
	return o.UniversalProductCode, true
}

// HasUniversalProductCode returns a boolean if a field has been set.
func (o *TrxInfoType) HasUniversalProductCode() bool {
	if o != nil && !IsNil(o.UniversalProductCode) {
		return true
	}

	return false
}

// SetUniversalProductCode gets a reference to the given string and assigns it to the UniversalProductCode field.
func (o *TrxInfoType) SetUniversalProductCode(v string) {
	o.UniversalProductCode = &v
}

// GetRoutingInstructionsId returns the RoutingInstructionsId field value if set, zero value otherwise.
func (o *TrxInfoType) GetRoutingInstructionsId() float32 {
	if o == nil || IsNil(o.RoutingInstructionsId) {
		var ret float32
		return ret
	}
	return *o.RoutingInstructionsId
}

// GetRoutingInstructionsIdOk returns a tuple with the RoutingInstructionsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetRoutingInstructionsIdOk() (*float32, bool) {
	if o == nil || IsNil(o.RoutingInstructionsId) {
		return nil, false
	}
	return o.RoutingInstructionsId, true
}

// HasRoutingInstructionsId returns a boolean if a field has been set.
func (o *TrxInfoType) HasRoutingInstructionsId() bool {
	if o != nil && !IsNil(o.RoutingInstructionsId) {
		return true
	}

	return false
}

// SetRoutingInstructionsId gets a reference to the given float32 and assigns it to the RoutingInstructionsId field.
func (o *TrxInfoType) SetRoutingInstructionsId(v float32) {
	o.RoutingInstructionsId = &v
}

// GetArticles returns the Articles field value if set, zero value otherwise.
func (o *TrxInfoType) GetArticles() []ArticleInfoType {
	if o == nil || IsNil(o.Articles) {
		var ret []ArticleInfoType
		return ret
	}
	return o.Articles
}

// GetArticlesOk returns a tuple with the Articles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetArticlesOk() ([]ArticleInfoType, bool) {
	if o == nil || IsNil(o.Articles) {
		return nil, false
	}
	return o.Articles, true
}

// HasArticles returns a boolean if a field has been set.
func (o *TrxInfoType) HasArticles() bool {
	if o != nil && !IsNil(o.Articles) {
		return true
	}

	return false
}

// SetArticles gets a reference to the given []ArticleInfoType and assigns it to the Articles field.
func (o *TrxInfoType) SetArticles(v []ArticleInfoType) {
	o.Articles = v
}

// GetTrxServiceType returns the TrxServiceType field value if set, zero value otherwise.
func (o *TrxInfoType) GetTrxServiceType() string {
	if o == nil || IsNil(o.TrxServiceType) {
		var ret string
		return ret
	}
	return *o.TrxServiceType
}

// GetTrxServiceTypeOk returns a tuple with the TrxServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetTrxServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TrxServiceType) {
		return nil, false
	}
	return o.TrxServiceType, true
}

// HasTrxServiceType returns a boolean if a field has been set.
func (o *TrxInfoType) HasTrxServiceType() bool {
	if o != nil && !IsNil(o.TrxServiceType) {
		return true
	}

	return false
}

// SetTrxServiceType gets a reference to the given string and assigns it to the TrxServiceType field.
func (o *TrxInfoType) SetTrxServiceType(v string) {
	o.TrxServiceType = &v
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise.
func (o *TrxInfoType) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode) {
		var ret string
		return ret
	}
	return *o.TransactionCode
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetTransactionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionCode) {
		return nil, false
	}
	return o.TransactionCode, true
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *TrxInfoType) HasTransactionCode() bool {
	if o != nil && !IsNil(o.TransactionCode) {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given string and assigns it to the TransactionCode field.
func (o *TrxInfoType) SetTransactionCode(v string) {
	o.TransactionCode = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *TrxInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *TrxInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *TrxInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetPrintTrxReceipt returns the PrintTrxReceipt field value if set, zero value otherwise.
func (o *TrxInfoType) GetPrintTrxReceipt() bool {
	if o == nil || IsNil(o.PrintTrxReceipt) {
		var ret bool
		return ret
	}
	return *o.PrintTrxReceipt
}

// GetPrintTrxReceiptOk returns a tuple with the PrintTrxReceipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrxInfoType) GetPrintTrxReceiptOk() (*bool, bool) {
	if o == nil || IsNil(o.PrintTrxReceipt) {
		return nil, false
	}
	return o.PrintTrxReceipt, true
}

// HasPrintTrxReceipt returns a boolean if a field has been set.
func (o *TrxInfoType) HasPrintTrxReceipt() bool {
	if o != nil && !IsNil(o.PrintTrxReceipt) {
		return true
	}

	return false
}

// SetPrintTrxReceipt gets a reference to the given bool and assigns it to the PrintTrxReceipt field.
func (o *TrxInfoType) SetPrintTrxReceipt(v bool) {
	o.PrintTrxReceipt = &v
}

func (o TrxInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrxInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TransactionGroup) {
		toSerialize["transactionGroup"] = o.TransactionGroup
	}
	if !IsNil(o.TransactionSubGroup) {
		toSerialize["transactionSubGroup"] = o.TransactionSubGroup
	}
	if !IsNil(o.UniversalProductCode) {
		toSerialize["universalProductCode"] = o.UniversalProductCode
	}
	if !IsNil(o.RoutingInstructionsId) {
		toSerialize["routingInstructionsId"] = o.RoutingInstructionsId
	}
	if !IsNil(o.Articles) {
		toSerialize["articles"] = o.Articles
	}
	if !IsNil(o.TrxServiceType) {
		toSerialize["trxServiceType"] = o.TrxServiceType
	}
	if !IsNil(o.TransactionCode) {
		toSerialize["transactionCode"] = o.TransactionCode
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.PrintTrxReceipt) {
		toSerialize["printTrxReceipt"] = o.PrintTrxReceipt
	}
	return toSerialize, nil
}

type NullableTrxInfoType struct {
	value *TrxInfoType
	isSet bool
}

func (v NullableTrxInfoType) Get() *TrxInfoType {
	return v.value
}

func (v *NullableTrxInfoType) Set(val *TrxInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableTrxInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableTrxInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrxInfoType(val *TrxInfoType) *NullableTrxInfoType {
	return &NullableTrxInfoType{value: val, isSet: true}
}

func (v NullableTrxInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrxInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


