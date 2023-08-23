/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelBillingStatementType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelBillingStatementType{}

// ChannelBillingStatementType To hold channel account statement detailed information.
type ChannelBillingStatementType struct {
	StatementId *UniqueIDType `json:"statementId,omitempty"`
	AccountCodeList []string `json:"accountCodeList,omitempty"`
	// Holds date on which statement is created.
	StatementDate *string `json:"statementDate,omitempty"`
	// Holds begin date of the statement.
	BeginDate *string `json:"beginDate,omitempty"`
	// Holds end date of the statement.
	EndDate *string `json:"endDate,omitempty"`
	// Holds date on which statement is locked.
	LockDate *string `json:"lockDate,omitempty"`
	// Holds Note for the statement.
	Note *string `json:"note,omitempty"`
	TotalAmount *CurrencyAmountType `json:"totalAmount,omitempty"`
	// Flag to identify Statements were generated or not.
	Generated *bool `json:"generated,omitempty"`
	// Flag to identify Statements generated were dirty or not.
	Dirty *bool `json:"dirty,omitempty"`
	// Provides detailed information regarding Channel Account contract.
	ChannelAccountStatements []ChannelStatementAccountType `json:"channelAccountStatements,omitempty"`
}

// NewChannelBillingStatementType instantiates a new ChannelBillingStatementType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelBillingStatementType() *ChannelBillingStatementType {
	this := ChannelBillingStatementType{}
	return &this
}

// NewChannelBillingStatementTypeWithDefaults instantiates a new ChannelBillingStatementType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelBillingStatementTypeWithDefaults() *ChannelBillingStatementType {
	this := ChannelBillingStatementType{}
	return &this
}

// GetStatementId returns the StatementId field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetStatementId() UniqueIDType {
	if o == nil || IsNil(o.StatementId) {
		var ret UniqueIDType
		return ret
	}
	return *o.StatementId
}

// GetStatementIdOk returns a tuple with the StatementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetStatementIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.StatementId) {
		return nil, false
	}
	return o.StatementId, true
}

// HasStatementId returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasStatementId() bool {
	if o != nil && !IsNil(o.StatementId) {
		return true
	}

	return false
}

// SetStatementId gets a reference to the given UniqueIDType and assigns it to the StatementId field.
func (o *ChannelBillingStatementType) SetStatementId(v UniqueIDType) {
	o.StatementId = &v
}

// GetAccountCodeList returns the AccountCodeList field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetAccountCodeList() []string {
	if o == nil || IsNil(o.AccountCodeList) {
		var ret []string
		return ret
	}
	return o.AccountCodeList
}

// GetAccountCodeListOk returns a tuple with the AccountCodeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetAccountCodeListOk() ([]string, bool) {
	if o == nil || IsNil(o.AccountCodeList) {
		return nil, false
	}
	return o.AccountCodeList, true
}

// HasAccountCodeList returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasAccountCodeList() bool {
	if o != nil && !IsNil(o.AccountCodeList) {
		return true
	}

	return false
}

// SetAccountCodeList gets a reference to the given []string and assigns it to the AccountCodeList field.
func (o *ChannelBillingStatementType) SetAccountCodeList(v []string) {
	o.AccountCodeList = v
}

// GetStatementDate returns the StatementDate field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetStatementDate() string {
	if o == nil || IsNil(o.StatementDate) {
		var ret string
		return ret
	}
	return *o.StatementDate
}

// GetStatementDateOk returns a tuple with the StatementDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetStatementDateOk() (*string, bool) {
	if o == nil || IsNil(o.StatementDate) {
		return nil, false
	}
	return o.StatementDate, true
}

// HasStatementDate returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasStatementDate() bool {
	if o != nil && !IsNil(o.StatementDate) {
		return true
	}

	return false
}

// SetStatementDate gets a reference to the given string and assigns it to the StatementDate field.
func (o *ChannelBillingStatementType) SetStatementDate(v string) {
	o.StatementDate = &v
}

// GetBeginDate returns the BeginDate field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetBeginDate() string {
	if o == nil || IsNil(o.BeginDate) {
		var ret string
		return ret
	}
	return *o.BeginDate
}

// GetBeginDateOk returns a tuple with the BeginDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetBeginDateOk() (*string, bool) {
	if o == nil || IsNil(o.BeginDate) {
		return nil, false
	}
	return o.BeginDate, true
}

// HasBeginDate returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasBeginDate() bool {
	if o != nil && !IsNil(o.BeginDate) {
		return true
	}

	return false
}

// SetBeginDate gets a reference to the given string and assigns it to the BeginDate field.
func (o *ChannelBillingStatementType) SetBeginDate(v string) {
	o.BeginDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ChannelBillingStatementType) SetEndDate(v string) {
	o.EndDate = &v
}

// GetLockDate returns the LockDate field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetLockDate() string {
	if o == nil || IsNil(o.LockDate) {
		var ret string
		return ret
	}
	return *o.LockDate
}

// GetLockDateOk returns a tuple with the LockDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetLockDateOk() (*string, bool) {
	if o == nil || IsNil(o.LockDate) {
		return nil, false
	}
	return o.LockDate, true
}

// HasLockDate returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasLockDate() bool {
	if o != nil && !IsNil(o.LockDate) {
		return true
	}

	return false
}

// SetLockDate gets a reference to the given string and assigns it to the LockDate field.
func (o *ChannelBillingStatementType) SetLockDate(v string) {
	o.LockDate = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ChannelBillingStatementType) SetNote(v string) {
	o.Note = &v
}

// GetTotalAmount returns the TotalAmount field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetTotalAmount() CurrencyAmountType {
	if o == nil || IsNil(o.TotalAmount) {
		var ret CurrencyAmountType
		return ret
	}
	return *o.TotalAmount
}

// GetTotalAmountOk returns a tuple with the TotalAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetTotalAmountOk() (*CurrencyAmountType, bool) {
	if o == nil || IsNil(o.TotalAmount) {
		return nil, false
	}
	return o.TotalAmount, true
}

// HasTotalAmount returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasTotalAmount() bool {
	if o != nil && !IsNil(o.TotalAmount) {
		return true
	}

	return false
}

// SetTotalAmount gets a reference to the given CurrencyAmountType and assigns it to the TotalAmount field.
func (o *ChannelBillingStatementType) SetTotalAmount(v CurrencyAmountType) {
	o.TotalAmount = &v
}

// GetGenerated returns the Generated field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetGenerated() bool {
	if o == nil || IsNil(o.Generated) {
		var ret bool
		return ret
	}
	return *o.Generated
}

// GetGeneratedOk returns a tuple with the Generated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetGeneratedOk() (*bool, bool) {
	if o == nil || IsNil(o.Generated) {
		return nil, false
	}
	return o.Generated, true
}

// HasGenerated returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasGenerated() bool {
	if o != nil && !IsNil(o.Generated) {
		return true
	}

	return false
}

// SetGenerated gets a reference to the given bool and assigns it to the Generated field.
func (o *ChannelBillingStatementType) SetGenerated(v bool) {
	o.Generated = &v
}

// GetDirty returns the Dirty field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetDirty() bool {
	if o == nil || IsNil(o.Dirty) {
		var ret bool
		return ret
	}
	return *o.Dirty
}

// GetDirtyOk returns a tuple with the Dirty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetDirtyOk() (*bool, bool) {
	if o == nil || IsNil(o.Dirty) {
		return nil, false
	}
	return o.Dirty, true
}

// HasDirty returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasDirty() bool {
	if o != nil && !IsNil(o.Dirty) {
		return true
	}

	return false
}

// SetDirty gets a reference to the given bool and assigns it to the Dirty field.
func (o *ChannelBillingStatementType) SetDirty(v bool) {
	o.Dirty = &v
}

// GetChannelAccountStatements returns the ChannelAccountStatements field value if set, zero value otherwise.
func (o *ChannelBillingStatementType) GetChannelAccountStatements() []ChannelStatementAccountType {
	if o == nil || IsNil(o.ChannelAccountStatements) {
		var ret []ChannelStatementAccountType
		return ret
	}
	return o.ChannelAccountStatements
}

// GetChannelAccountStatementsOk returns a tuple with the ChannelAccountStatements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelBillingStatementType) GetChannelAccountStatementsOk() ([]ChannelStatementAccountType, bool) {
	if o == nil || IsNil(o.ChannelAccountStatements) {
		return nil, false
	}
	return o.ChannelAccountStatements, true
}

// HasChannelAccountStatements returns a boolean if a field has been set.
func (o *ChannelBillingStatementType) HasChannelAccountStatements() bool {
	if o != nil && !IsNil(o.ChannelAccountStatements) {
		return true
	}

	return false
}

// SetChannelAccountStatements gets a reference to the given []ChannelStatementAccountType and assigns it to the ChannelAccountStatements field.
func (o *ChannelBillingStatementType) SetChannelAccountStatements(v []ChannelStatementAccountType) {
	o.ChannelAccountStatements = v
}

func (o ChannelBillingStatementType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelBillingStatementType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatementId) {
		toSerialize["statementId"] = o.StatementId
	}
	if !IsNil(o.AccountCodeList) {
		toSerialize["accountCodeList"] = o.AccountCodeList
	}
	if !IsNil(o.StatementDate) {
		toSerialize["statementDate"] = o.StatementDate
	}
	if !IsNil(o.BeginDate) {
		toSerialize["beginDate"] = o.BeginDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.LockDate) {
		toSerialize["lockDate"] = o.LockDate
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.TotalAmount) {
		toSerialize["totalAmount"] = o.TotalAmount
	}
	if !IsNil(o.Generated) {
		toSerialize["generated"] = o.Generated
	}
	if !IsNil(o.Dirty) {
		toSerialize["dirty"] = o.Dirty
	}
	if !IsNil(o.ChannelAccountStatements) {
		toSerialize["channelAccountStatements"] = o.ChannelAccountStatements
	}
	return toSerialize, nil
}

type NullableChannelBillingStatementType struct {
	value *ChannelBillingStatementType
	isSet bool
}

func (v NullableChannelBillingStatementType) Get() *ChannelBillingStatementType {
	return v.value
}

func (v *NullableChannelBillingStatementType) Set(val *ChannelBillingStatementType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelBillingStatementType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelBillingStatementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelBillingStatementType(val *ChannelBillingStatementType) *NullableChannelBillingStatementType {
	return &NullableChannelBillingStatementType{value: val, isSet: true}
}

func (v NullableChannelBillingStatementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelBillingStatementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


