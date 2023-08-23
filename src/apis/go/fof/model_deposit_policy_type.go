/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DepositPolicyType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepositPolicyType{}

// DepositPolicyType Used to define the deposit policy, guarantees policy, and/or accepted forms of payment.
type DepositPolicyType struct {
	AmountPercent *PolicyAmountPercentType `json:"amountPercent,omitempty"`
	Deadline *PolicyDeadlineType `json:"deadline,omitempty"`
	// Text description of the Payment in a given language.
	Description *string `json:"description,omitempty"`
	// Receipt number associated with the deposit policy
	DepositReceiptNo *int32 `json:"depositReceiptNo,omitempty"`
	// Transaction Date associated with the deposit policy
	TransactionDate *string `json:"transactionDate,omitempty"`
	// Flag to indicate if deposit policy is reversed
	DepositReqReversed *bool `json:"depositReqReversed,omitempty"`
	// Formatted Text Rule of the deposit policy.
	FormattedRule *string `json:"formattedRule,omitempty"`
	TypeOfCharges *DepositCancelRevenueType `json:"typeOfCharges,omitempty"`
	// Deposit Policy Code
	PolicyCode *string `json:"policyCode,omitempty"`
	// Flag to indicate if the cancellation policy is manual.
	Manual *bool `json:"manual,omitempty"`
}

// NewDepositPolicyType instantiates a new DepositPolicyType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositPolicyType() *DepositPolicyType {
	this := DepositPolicyType{}
	return &this
}

// NewDepositPolicyTypeWithDefaults instantiates a new DepositPolicyType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositPolicyTypeWithDefaults() *DepositPolicyType {
	this := DepositPolicyType{}
	return &this
}

// GetAmountPercent returns the AmountPercent field value if set, zero value otherwise.
func (o *DepositPolicyType) GetAmountPercent() PolicyAmountPercentType {
	if o == nil || IsNil(o.AmountPercent) {
		var ret PolicyAmountPercentType
		return ret
	}
	return *o.AmountPercent
}

// GetAmountPercentOk returns a tuple with the AmountPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetAmountPercentOk() (*PolicyAmountPercentType, bool) {
	if o == nil || IsNil(o.AmountPercent) {
		return nil, false
	}
	return o.AmountPercent, true
}

// HasAmountPercent returns a boolean if a field has been set.
func (o *DepositPolicyType) HasAmountPercent() bool {
	if o != nil && !IsNil(o.AmountPercent) {
		return true
	}

	return false
}

// SetAmountPercent gets a reference to the given PolicyAmountPercentType and assigns it to the AmountPercent field.
func (o *DepositPolicyType) SetAmountPercent(v PolicyAmountPercentType) {
	o.AmountPercent = &v
}

// GetDeadline returns the Deadline field value if set, zero value otherwise.
func (o *DepositPolicyType) GetDeadline() PolicyDeadlineType {
	if o == nil || IsNil(o.Deadline) {
		var ret PolicyDeadlineType
		return ret
	}
	return *o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetDeadlineOk() (*PolicyDeadlineType, bool) {
	if o == nil || IsNil(o.Deadline) {
		return nil, false
	}
	return o.Deadline, true
}

// HasDeadline returns a boolean if a field has been set.
func (o *DepositPolicyType) HasDeadline() bool {
	if o != nil && !IsNil(o.Deadline) {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given PolicyDeadlineType and assigns it to the Deadline field.
func (o *DepositPolicyType) SetDeadline(v PolicyDeadlineType) {
	o.Deadline = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DepositPolicyType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DepositPolicyType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DepositPolicyType) SetDescription(v string) {
	o.Description = &v
}

// GetDepositReceiptNo returns the DepositReceiptNo field value if set, zero value otherwise.
func (o *DepositPolicyType) GetDepositReceiptNo() int32 {
	if o == nil || IsNil(o.DepositReceiptNo) {
		var ret int32
		return ret
	}
	return *o.DepositReceiptNo
}

// GetDepositReceiptNoOk returns a tuple with the DepositReceiptNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetDepositReceiptNoOk() (*int32, bool) {
	if o == nil || IsNil(o.DepositReceiptNo) {
		return nil, false
	}
	return o.DepositReceiptNo, true
}

// HasDepositReceiptNo returns a boolean if a field has been set.
func (o *DepositPolicyType) HasDepositReceiptNo() bool {
	if o != nil && !IsNil(o.DepositReceiptNo) {
		return true
	}

	return false
}

// SetDepositReceiptNo gets a reference to the given int32 and assigns it to the DepositReceiptNo field.
func (o *DepositPolicyType) SetDepositReceiptNo(v int32) {
	o.DepositReceiptNo = &v
}

// GetTransactionDate returns the TransactionDate field value if set, zero value otherwise.
func (o *DepositPolicyType) GetTransactionDate() string {
	if o == nil || IsNil(o.TransactionDate) {
		var ret string
		return ret
	}
	return *o.TransactionDate
}

// GetTransactionDateOk returns a tuple with the TransactionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetTransactionDateOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionDate) {
		return nil, false
	}
	return o.TransactionDate, true
}

// HasTransactionDate returns a boolean if a field has been set.
func (o *DepositPolicyType) HasTransactionDate() bool {
	if o != nil && !IsNil(o.TransactionDate) {
		return true
	}

	return false
}

// SetTransactionDate gets a reference to the given string and assigns it to the TransactionDate field.
func (o *DepositPolicyType) SetTransactionDate(v string) {
	o.TransactionDate = &v
}

// GetDepositReqReversed returns the DepositReqReversed field value if set, zero value otherwise.
func (o *DepositPolicyType) GetDepositReqReversed() bool {
	if o == nil || IsNil(o.DepositReqReversed) {
		var ret bool
		return ret
	}
	return *o.DepositReqReversed
}

// GetDepositReqReversedOk returns a tuple with the DepositReqReversed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetDepositReqReversedOk() (*bool, bool) {
	if o == nil || IsNil(o.DepositReqReversed) {
		return nil, false
	}
	return o.DepositReqReversed, true
}

// HasDepositReqReversed returns a boolean if a field has been set.
func (o *DepositPolicyType) HasDepositReqReversed() bool {
	if o != nil && !IsNil(o.DepositReqReversed) {
		return true
	}

	return false
}

// SetDepositReqReversed gets a reference to the given bool and assigns it to the DepositReqReversed field.
func (o *DepositPolicyType) SetDepositReqReversed(v bool) {
	o.DepositReqReversed = &v
}

// GetFormattedRule returns the FormattedRule field value if set, zero value otherwise.
func (o *DepositPolicyType) GetFormattedRule() string {
	if o == nil || IsNil(o.FormattedRule) {
		var ret string
		return ret
	}
	return *o.FormattedRule
}

// GetFormattedRuleOk returns a tuple with the FormattedRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetFormattedRuleOk() (*string, bool) {
	if o == nil || IsNil(o.FormattedRule) {
		return nil, false
	}
	return o.FormattedRule, true
}

// HasFormattedRule returns a boolean if a field has been set.
func (o *DepositPolicyType) HasFormattedRule() bool {
	if o != nil && !IsNil(o.FormattedRule) {
		return true
	}

	return false
}

// SetFormattedRule gets a reference to the given string and assigns it to the FormattedRule field.
func (o *DepositPolicyType) SetFormattedRule(v string) {
	o.FormattedRule = &v
}

// GetTypeOfCharges returns the TypeOfCharges field value if set, zero value otherwise.
func (o *DepositPolicyType) GetTypeOfCharges() DepositCancelRevenueType {
	if o == nil || IsNil(o.TypeOfCharges) {
		var ret DepositCancelRevenueType
		return ret
	}
	return *o.TypeOfCharges
}

// GetTypeOfChargesOk returns a tuple with the TypeOfCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetTypeOfChargesOk() (*DepositCancelRevenueType, bool) {
	if o == nil || IsNil(o.TypeOfCharges) {
		return nil, false
	}
	return o.TypeOfCharges, true
}

// HasTypeOfCharges returns a boolean if a field has been set.
func (o *DepositPolicyType) HasTypeOfCharges() bool {
	if o != nil && !IsNil(o.TypeOfCharges) {
		return true
	}

	return false
}

// SetTypeOfCharges gets a reference to the given DepositCancelRevenueType and assigns it to the TypeOfCharges field.
func (o *DepositPolicyType) SetTypeOfCharges(v DepositCancelRevenueType) {
	o.TypeOfCharges = &v
}

// GetPolicyCode returns the PolicyCode field value if set, zero value otherwise.
func (o *DepositPolicyType) GetPolicyCode() string {
	if o == nil || IsNil(o.PolicyCode) {
		var ret string
		return ret
	}
	return *o.PolicyCode
}

// GetPolicyCodeOk returns a tuple with the PolicyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetPolicyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyCode) {
		return nil, false
	}
	return o.PolicyCode, true
}

// HasPolicyCode returns a boolean if a field has been set.
func (o *DepositPolicyType) HasPolicyCode() bool {
	if o != nil && !IsNil(o.PolicyCode) {
		return true
	}

	return false
}

// SetPolicyCode gets a reference to the given string and assigns it to the PolicyCode field.
func (o *DepositPolicyType) SetPolicyCode(v string) {
	o.PolicyCode = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *DepositPolicyType) GetManual() bool {
	if o == nil || IsNil(o.Manual) {
		var ret bool
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositPolicyType) GetManualOk() (*bool, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *DepositPolicyType) HasManual() bool {
	if o != nil && !IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given bool and assigns it to the Manual field.
func (o *DepositPolicyType) SetManual(v bool) {
	o.Manual = &v
}

func (o DepositPolicyType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepositPolicyType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AmountPercent) {
		toSerialize["amountPercent"] = o.AmountPercent
	}
	if !IsNil(o.Deadline) {
		toSerialize["deadline"] = o.Deadline
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DepositReceiptNo) {
		toSerialize["depositReceiptNo"] = o.DepositReceiptNo
	}
	if !IsNil(o.TransactionDate) {
		toSerialize["transactionDate"] = o.TransactionDate
	}
	if !IsNil(o.DepositReqReversed) {
		toSerialize["depositReqReversed"] = o.DepositReqReversed
	}
	if !IsNil(o.FormattedRule) {
		toSerialize["formattedRule"] = o.FormattedRule
	}
	if !IsNil(o.TypeOfCharges) {
		toSerialize["typeOfCharges"] = o.TypeOfCharges
	}
	if !IsNil(o.PolicyCode) {
		toSerialize["policyCode"] = o.PolicyCode
	}
	if !IsNil(o.Manual) {
		toSerialize["manual"] = o.Manual
	}
	return toSerialize, nil
}

type NullableDepositPolicyType struct {
	value *DepositPolicyType
	isSet bool
}

func (v NullableDepositPolicyType) Get() *DepositPolicyType {
	return v.value
}

func (v *NullableDepositPolicyType) Set(val *DepositPolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositPolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositPolicyType(val *DepositPolicyType) *NullableDepositPolicyType {
	return &NullableDepositPolicyType{value: val, isSet: true}
}

func (v NullableDepositPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


