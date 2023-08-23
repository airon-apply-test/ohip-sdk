/*
OPERA Cloud Price Availability Rate API

APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PackageElementType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageElementType{}

// PackageElementType struct for PackageElementType
type PackageElementType struct {
	Amount *CurrencyAmountType `json:"amount,omitempty"`
	Allowance *CurrencyAmountType `json:"allowance,omitempty"`
	Description []string `json:"description,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate *string `json:"endDate,omitempty"`
	PackageCode *string `json:"packageCode,omitempty"`
	CalculationRule *string `json:"calculationRule,omitempty"`
	PostingRhythm *string `json:"postingRhythm,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	IncludedInRate *bool `json:"includedInRate,omitempty"`
	AddRateSeprateLine *bool `json:"addRateSeprateLine,omitempty"`
	AddRateCombinedLine *bool `json:"addRateCombinedLine,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
	EndTime *string `json:"endTime,omitempty"`
	SellSeparate *bool `json:"sellSeparate,omitempty"`
}

// NewPackageElementType instantiates a new PackageElementType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageElementType() *PackageElementType {
	this := PackageElementType{}
	return &this
}

// NewPackageElementTypeWithDefaults instantiates a new PackageElementType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageElementTypeWithDefaults() *PackageElementType {
	this := PackageElementType{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PackageElementType) GetAmount() CurrencyAmountType {
	if o == nil || IsNil(o.Amount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PackageElementType) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given CurrencyAmountType and assigns it to the Amount field.
func (o *PackageElementType) SetAmount(v CurrencyAmountType) {
	o.Amount = &v
}

// GetAllowance returns the Allowance field value if set, zero value otherwise.
func (o *PackageElementType) GetAllowance() CurrencyAmountType {
	if o == nil || IsNil(o.Allowance) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.Allowance
}

// GetAllowanceOk returns a tuple with the Allowance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetAllowanceOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.Allowance) {
		return nil, false
	}
	return o.Allowance, true
}

// HasAllowance returns a boolean if a field has been set.
func (o *PackageElementType) HasAllowance() bool {
	if o != nil && !IsNil(o.Allowance) {
		return true
	}

	return false
}

// SetAllowance gets a reference to the given CurrencyAmountType and assigns it to the Allowance field.
func (o *PackageElementType) SetAllowance(v CurrencyAmountType) {
	o.Allowance = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PackageElementType) GetDescription() []string {
	if o == nil || IsNil(o.Description) {
		var ret []string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetDescriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PackageElementType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []string and assigns it to the Description field.
func (o *PackageElementType) SetDescription(v []string) {
	o.Description = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PackageElementType) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PackageElementType) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *PackageElementType) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *PackageElementType) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *PackageElementType) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *PackageElementType) SetEndDate(v string) {
	o.EndDate = &v
}

// GetPackageCode returns the PackageCode field value if set, zero value otherwise.
func (o *PackageElementType) GetPackageCode() string {
	if o == nil || IsNil(o.PackageCode) {
		var ret string
		return ret
	}
	return *o.PackageCode
}

// GetPackageCodeOk returns a tuple with the PackageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetPackageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageCode) {
		return nil, false
	}
	return o.PackageCode, true
}

// HasPackageCode returns a boolean if a field has been set.
func (o *PackageElementType) HasPackageCode() bool {
	if o != nil && !IsNil(o.PackageCode) {
		return true
	}

	return false
}

// SetPackageCode gets a reference to the given string and assigns it to the PackageCode field.
func (o *PackageElementType) SetPackageCode(v string) {
	o.PackageCode = &v
}

// GetCalculationRule returns the CalculationRule field value if set, zero value otherwise.
func (o *PackageElementType) GetCalculationRule() string {
	if o == nil || IsNil(o.CalculationRule) {
		var ret string
		return ret
	}
	return *o.CalculationRule
}

// GetCalculationRuleOk returns a tuple with the CalculationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetCalculationRuleOk() (*string, bool) {
	if o == nil || IsNil(o.CalculationRule) {
		return nil, false
	}
	return o.CalculationRule, true
}

// HasCalculationRule returns a boolean if a field has been set.
func (o *PackageElementType) HasCalculationRule() bool {
	if o != nil && !IsNil(o.CalculationRule) {
		return true
	}

	return false
}

// SetCalculationRule gets a reference to the given string and assigns it to the CalculationRule field.
func (o *PackageElementType) SetCalculationRule(v string) {
	o.CalculationRule = &v
}

// GetPostingRhythm returns the PostingRhythm field value if set, zero value otherwise.
func (o *PackageElementType) GetPostingRhythm() string {
	if o == nil || IsNil(o.PostingRhythm) {
		var ret string
		return ret
	}
	return *o.PostingRhythm
}

// GetPostingRhythmOk returns a tuple with the PostingRhythm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetPostingRhythmOk() (*string, bool) {
	if o == nil || IsNil(o.PostingRhythm) {
		return nil, false
	}
	return o.PostingRhythm, true
}

// HasPostingRhythm returns a boolean if a field has been set.
func (o *PackageElementType) HasPostingRhythm() bool {
	if o != nil && !IsNil(o.PostingRhythm) {
		return true
	}

	return false
}

// SetPostingRhythm gets a reference to the given string and assigns it to the PostingRhythm field.
func (o *PackageElementType) SetPostingRhythm(v string) {
	o.PostingRhythm = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *PackageElementType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *PackageElementType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *PackageElementType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetIncludedInRate returns the IncludedInRate field value if set, zero value otherwise.
func (o *PackageElementType) GetIncludedInRate() bool {
	if o == nil || IsNil(o.IncludedInRate) {
		var ret bool
		return ret
	}
	return *o.IncludedInRate
}

// GetIncludedInRateOk returns a tuple with the IncludedInRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetIncludedInRateOk() (*bool, bool) {
	if o == nil || IsNil(o.IncludedInRate) {
		return nil, false
	}
	return o.IncludedInRate, true
}

// HasIncludedInRate returns a boolean if a field has been set.
func (o *PackageElementType) HasIncludedInRate() bool {
	if o != nil && !IsNil(o.IncludedInRate) {
		return true
	}

	return false
}

// SetIncludedInRate gets a reference to the given bool and assigns it to the IncludedInRate field.
func (o *PackageElementType) SetIncludedInRate(v bool) {
	o.IncludedInRate = &v
}

// GetAddRateSeprateLine returns the AddRateSeprateLine field value if set, zero value otherwise.
func (o *PackageElementType) GetAddRateSeprateLine() bool {
	if o == nil || IsNil(o.AddRateSeprateLine) {
		var ret bool
		return ret
	}
	return *o.AddRateSeprateLine
}

// GetAddRateSeprateLineOk returns a tuple with the AddRateSeprateLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetAddRateSeprateLineOk() (*bool, bool) {
	if o == nil || IsNil(o.AddRateSeprateLine) {
		return nil, false
	}
	return o.AddRateSeprateLine, true
}

// HasAddRateSeprateLine returns a boolean if a field has been set.
func (o *PackageElementType) HasAddRateSeprateLine() bool {
	if o != nil && !IsNil(o.AddRateSeprateLine) {
		return true
	}

	return false
}

// SetAddRateSeprateLine gets a reference to the given bool and assigns it to the AddRateSeprateLine field.
func (o *PackageElementType) SetAddRateSeprateLine(v bool) {
	o.AddRateSeprateLine = &v
}

// GetAddRateCombinedLine returns the AddRateCombinedLine field value if set, zero value otherwise.
func (o *PackageElementType) GetAddRateCombinedLine() bool {
	if o == nil || IsNil(o.AddRateCombinedLine) {
		var ret bool
		return ret
	}
	return *o.AddRateCombinedLine
}

// GetAddRateCombinedLineOk returns a tuple with the AddRateCombinedLine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetAddRateCombinedLineOk() (*bool, bool) {
	if o == nil || IsNil(o.AddRateCombinedLine) {
		return nil, false
	}
	return o.AddRateCombinedLine, true
}

// HasAddRateCombinedLine returns a boolean if a field has been set.
func (o *PackageElementType) HasAddRateCombinedLine() bool {
	if o != nil && !IsNil(o.AddRateCombinedLine) {
		return true
	}

	return false
}

// SetAddRateCombinedLine gets a reference to the given bool and assigns it to the AddRateCombinedLine field.
func (o *PackageElementType) SetAddRateCombinedLine(v bool) {
	o.AddRateCombinedLine = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *PackageElementType) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *PackageElementType) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *PackageElementType) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *PackageElementType) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *PackageElementType) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *PackageElementType) SetEndTime(v string) {
	o.EndTime = &v
}

// GetSellSeparate returns the SellSeparate field value if set, zero value otherwise.
func (o *PackageElementType) GetSellSeparate() bool {
	if o == nil || IsNil(o.SellSeparate) {
		var ret bool
		return ret
	}
	return *o.SellSeparate
}

// GetSellSeparateOk returns a tuple with the SellSeparate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageElementType) GetSellSeparateOk() (*bool, bool) {
	if o == nil || IsNil(o.SellSeparate) {
		return nil, false
	}
	return o.SellSeparate, true
}

// HasSellSeparate returns a boolean if a field has been set.
func (o *PackageElementType) HasSellSeparate() bool {
	if o != nil && !IsNil(o.SellSeparate) {
		return true
	}

	return false
}

// SetSellSeparate gets a reference to the given bool and assigns it to the SellSeparate field.
func (o *PackageElementType) SetSellSeparate(v bool) {
	o.SellSeparate = &v
}

func (o PackageElementType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageElementType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Allowance) {
		toSerialize["allowance"] = o.Allowance
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.PackageCode) {
		toSerialize["packageCode"] = o.PackageCode
	}
	if !IsNil(o.CalculationRule) {
		toSerialize["calculationRule"] = o.CalculationRule
	}
	if !IsNil(o.PostingRhythm) {
		toSerialize["postingRhythm"] = o.PostingRhythm
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.IncludedInRate) {
		toSerialize["includedInRate"] = o.IncludedInRate
	}
	if !IsNil(o.AddRateSeprateLine) {
		toSerialize["addRateSeprateLine"] = o.AddRateSeprateLine
	}
	if !IsNil(o.AddRateCombinedLine) {
		toSerialize["addRateCombinedLine"] = o.AddRateCombinedLine
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.SellSeparate) {
		toSerialize["sellSeparate"] = o.SellSeparate
	}
	return toSerialize, nil
}

type NullablePackageElementType struct {
	value *PackageElementType
	isSet bool
}

func (v NullablePackageElementType) Get() *PackageElementType {
	return v.value
}

func (v *NullablePackageElementType) Set(val *PackageElementType) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageElementType) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageElementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageElementType(val *PackageElementType) *NullablePackageElementType {
	return &NullablePackageElementType{value: val, isSet: true}
}

func (v NullablePackageElementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageElementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


