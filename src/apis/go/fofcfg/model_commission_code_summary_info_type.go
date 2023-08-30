/*
OPERA Cloud Front Desk Configuration API

APIs to cater for Front Desk Configuration in OPERA Cloud. Here you can find operations to get, post, put and delete front desk codes such as commission codes, transaction groups, codes & subgroups, articles, payment methods and credit card types.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fofcfg

import (
	"encoding/json"
)

// checks if the CommissionCodeSummaryInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommissionCodeSummaryInfoType{}

// CommissionCodeSummaryInfoType Commission codes summary details.
type CommissionCodeSummaryInfoType struct {
	// Property of the commission codes.
	HotelId *string `json:"hotelId,omitempty"`
	// Commission code details.
	CommissionCode *string `json:"commissionCode,omitempty"`
	// Commission code description for which details are fetched.
	Description *string `json:"description,omitempty"`
	// Commission code sequence number.
	Sequence *int32 `json:"sequence,omitempty"`
	BasedOn *CommissionBasedOnType `json:"basedOn,omitempty"`
	CommissionAmount *CurrencyAmountType `json:"commissionAmount,omitempty"`
	// Commission paid Tax details.
	TaxPercentage *float32 `json:"taxPercentage,omitempty"`
	Default *bool `json:"default,omitempty"`
}

// NewCommissionCodeSummaryInfoType instantiates a new CommissionCodeSummaryInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommissionCodeSummaryInfoType() *CommissionCodeSummaryInfoType {
	this := CommissionCodeSummaryInfoType{}
	return &this
}

// NewCommissionCodeSummaryInfoTypeWithDefaults instantiates a new CommissionCodeSummaryInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommissionCodeSummaryInfoTypeWithDefaults() *CommissionCodeSummaryInfoType {
	this := CommissionCodeSummaryInfoType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *CommissionCodeSummaryInfoType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetCommissionCode returns the CommissionCode field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetCommissionCode() string {
	if o == nil || IsNil(o.CommissionCode) {
		var ret string
		return ret
	}
	return *o.CommissionCode
}

// GetCommissionCodeOk returns a tuple with the CommissionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetCommissionCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CommissionCode) {
		return nil, false
	}
	return o.CommissionCode, true
}

// HasCommissionCode returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasCommissionCode() bool {
	if o != nil && !IsNil(o.CommissionCode) {
		return true
	}

	return false
}

// SetCommissionCode gets a reference to the given string and assigns it to the CommissionCode field.
func (o *CommissionCodeSummaryInfoType) SetCommissionCode(v string) {
	o.CommissionCode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommissionCodeSummaryInfoType) SetDescription(v string) {
	o.Description = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence) {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *CommissionCodeSummaryInfoType) SetSequence(v int32) {
	o.Sequence = &v
}

// GetBasedOn returns the BasedOn field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetBasedOn() CommissionBasedOnType {
	if o == nil || IsNil(o.BasedOn) {
		var ret CommissionBasedOnType
		return ret
	}
	return *o.BasedOn
}

// GetBasedOnOk returns a tuple with the BasedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetBasedOnOk() (*CommissionBasedOnType, bool) {
	if o == nil || IsNil(o.BasedOn) {
		return nil, false
	}
	return o.BasedOn, true
}

// HasBasedOn returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasBasedOn() bool {
	if o != nil && !IsNil(o.BasedOn) {
		return true
	}

	return false
}

// SetBasedOn gets a reference to the given CommissionBasedOnType and assigns it to the BasedOn field.
func (o *CommissionCodeSummaryInfoType) SetBasedOn(v CommissionBasedOnType) {
	o.BasedOn = &v
}

// GetCommissionAmount returns the CommissionAmount field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetCommissionAmount() CurrencyAmountType {
	if o == nil || IsNil(o.CommissionAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.CommissionAmount
}

// GetCommissionAmountOk returns a tuple with the CommissionAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetCommissionAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.CommissionAmount) {
		return nil, false
	}
	return o.CommissionAmount, true
}

// HasCommissionAmount returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasCommissionAmount() bool {
	if o != nil && !IsNil(o.CommissionAmount) {
		return true
	}

	return false
}

// SetCommissionAmount gets a reference to the given CurrencyAmountType and assigns it to the CommissionAmount field.
func (o *CommissionCodeSummaryInfoType) SetCommissionAmount(v CurrencyAmountType) {
	o.CommissionAmount = &v
}

// GetTaxPercentage returns the TaxPercentage field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetTaxPercentage() float32 {
	if o == nil || IsNil(o.TaxPercentage) {
		var ret float32
		return ret
	}
	return *o.TaxPercentage
}

// GetTaxPercentageOk returns a tuple with the TaxPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetTaxPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxPercentage) {
		return nil, false
	}
	return o.TaxPercentage, true
}

// HasTaxPercentage returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasTaxPercentage() bool {
	if o != nil && !IsNil(o.TaxPercentage) {
		return true
	}

	return false
}

// SetTaxPercentage gets a reference to the given float32 and assigns it to the TaxPercentage field.
func (o *CommissionCodeSummaryInfoType) SetTaxPercentage(v float32) {
	o.TaxPercentage = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *CommissionCodeSummaryInfoType) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionCodeSummaryInfoType) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *CommissionCodeSummaryInfoType) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *CommissionCodeSummaryInfoType) SetDefault(v bool) {
	o.Default = &v
}

func (o CommissionCodeSummaryInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommissionCodeSummaryInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.CommissionCode) {
		toSerialize["commissionCode"] = o.CommissionCode
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.BasedOn) {
		toSerialize["basedOn"] = o.BasedOn
	}
	if !IsNil(o.CommissionAmount) {
		toSerialize["commissionAmount"] = o.CommissionAmount
	}
	if !IsNil(o.TaxPercentage) {
		toSerialize["taxPercentage"] = o.TaxPercentage
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableCommissionCodeSummaryInfoType struct {
	value *CommissionCodeSummaryInfoType
	isSet bool
}

func (v NullableCommissionCodeSummaryInfoType) Get() *CommissionCodeSummaryInfoType {
	return v.value
}

func (v *NullableCommissionCodeSummaryInfoType) Set(val *CommissionCodeSummaryInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableCommissionCodeSummaryInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableCommissionCodeSummaryInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommissionCodeSummaryInfoType(val *CommissionCodeSummaryInfoType) *NullableCommissionCodeSummaryInfoType {
	return &NullableCommissionCodeSummaryInfoType{value: val, isSet: true}
}

func (v NullableCommissionCodeSummaryInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommissionCodeSummaryInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


