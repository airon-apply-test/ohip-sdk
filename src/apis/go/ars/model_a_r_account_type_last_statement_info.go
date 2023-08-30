/*
OPERA Cloud Accounts Receivables API

APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ars

import (
	"encoding/json"
)

// checks if the ARAccountTypeLastStatementInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ARAccountTypeLastStatementInfo{}

// ARAccountTypeLastStatementInfo Report History Type used as based type for Remiders and Statements History types.
type ARAccountTypeLastStatementInfo struct {
	// Report Name.
	ReportName *string `json:"reportName,omitempty"`
	// Report file name when exists to allow report re-printing.
	ReportFileName *string `json:"reportFileName,omitempty"`
	// The Reminder Letter name which is to be used for this Reminder based on the setup on the Account Type.
	DateSent *string `json:"dateSent,omitempty"`
	// When using Statement Numbering, a unique number is associated to the Statement.
	StatementNo *int32 `json:"statementNo,omitempty"`
	// Indicates that statement history exists.
	HistoryExists *bool `json:"historyExists,omitempty"`
}

// NewARAccountTypeLastStatementInfo instantiates a new ARAccountTypeLastStatementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewARAccountTypeLastStatementInfo() *ARAccountTypeLastStatementInfo {
	this := ARAccountTypeLastStatementInfo{}
	return &this
}

// NewARAccountTypeLastStatementInfoWithDefaults instantiates a new ARAccountTypeLastStatementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewARAccountTypeLastStatementInfoWithDefaults() *ARAccountTypeLastStatementInfo {
	this := ARAccountTypeLastStatementInfo{}
	return &this
}

// GetReportName returns the ReportName field value if set, zero value otherwise.
func (o *ARAccountTypeLastStatementInfo) GetReportName() string {
	if o == nil || IsNil(o.ReportName) {
		var ret string
		return ret
	}
	return *o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTypeLastStatementInfo) GetReportNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReportName) {
		return nil, false
	}
	return o.ReportName, true
}

// HasReportName returns a boolean if a field has been set.
func (o *ARAccountTypeLastStatementInfo) HasReportName() bool {
	if o != nil && !IsNil(o.ReportName) {
		return true
	}

	return false
}

// SetReportName gets a reference to the given string and assigns it to the ReportName field.
func (o *ARAccountTypeLastStatementInfo) SetReportName(v string) {
	o.ReportName = &v
}

// GetReportFileName returns the ReportFileName field value if set, zero value otherwise.
func (o *ARAccountTypeLastStatementInfo) GetReportFileName() string {
	if o == nil || IsNil(o.ReportFileName) {
		var ret string
		return ret
	}
	return *o.ReportFileName
}

// GetReportFileNameOk returns a tuple with the ReportFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTypeLastStatementInfo) GetReportFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReportFileName) {
		return nil, false
	}
	return o.ReportFileName, true
}

// HasReportFileName returns a boolean if a field has been set.
func (o *ARAccountTypeLastStatementInfo) HasReportFileName() bool {
	if o != nil && !IsNil(o.ReportFileName) {
		return true
	}

	return false
}

// SetReportFileName gets a reference to the given string and assigns it to the ReportFileName field.
func (o *ARAccountTypeLastStatementInfo) SetReportFileName(v string) {
	o.ReportFileName = &v
}

// GetDateSent returns the DateSent field value if set, zero value otherwise.
func (o *ARAccountTypeLastStatementInfo) GetDateSent() string {
	if o == nil || IsNil(o.DateSent) {
		var ret string
		return ret
	}
	return *o.DateSent
}

// GetDateSentOk returns a tuple with the DateSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTypeLastStatementInfo) GetDateSentOk() (*string, bool) {
	if o == nil || IsNil(o.DateSent) {
		return nil, false
	}
	return o.DateSent, true
}

// HasDateSent returns a boolean if a field has been set.
func (o *ARAccountTypeLastStatementInfo) HasDateSent() bool {
	if o != nil && !IsNil(o.DateSent) {
		return true
	}

	return false
}

// SetDateSent gets a reference to the given string and assigns it to the DateSent field.
func (o *ARAccountTypeLastStatementInfo) SetDateSent(v string) {
	o.DateSent = &v
}

// GetStatementNo returns the StatementNo field value if set, zero value otherwise.
func (o *ARAccountTypeLastStatementInfo) GetStatementNo() int32 {
	if o == nil || IsNil(o.StatementNo) {
		var ret int32
		return ret
	}
	return *o.StatementNo
}

// GetStatementNoOk returns a tuple with the StatementNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTypeLastStatementInfo) GetStatementNoOk() (*int32, bool) {
	if o == nil || IsNil(o.StatementNo) {
		return nil, false
	}
	return o.StatementNo, true
}

// HasStatementNo returns a boolean if a field has been set.
func (o *ARAccountTypeLastStatementInfo) HasStatementNo() bool {
	if o != nil && !IsNil(o.StatementNo) {
		return true
	}

	return false
}

// SetStatementNo gets a reference to the given int32 and assigns it to the StatementNo field.
func (o *ARAccountTypeLastStatementInfo) SetStatementNo(v int32) {
	o.StatementNo = &v
}

// GetHistoryExists returns the HistoryExists field value if set, zero value otherwise.
func (o *ARAccountTypeLastStatementInfo) GetHistoryExists() bool {
	if o == nil || IsNil(o.HistoryExists) {
		var ret bool
		return ret
	}
	return *o.HistoryExists
}

// GetHistoryExistsOk returns a tuple with the HistoryExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ARAccountTypeLastStatementInfo) GetHistoryExistsOk() (*bool, bool) {
	if o == nil || IsNil(o.HistoryExists) {
		return nil, false
	}
	return o.HistoryExists, true
}

// HasHistoryExists returns a boolean if a field has been set.
func (o *ARAccountTypeLastStatementInfo) HasHistoryExists() bool {
	if o != nil && !IsNil(o.HistoryExists) {
		return true
	}

	return false
}

// SetHistoryExists gets a reference to the given bool and assigns it to the HistoryExists field.
func (o *ARAccountTypeLastStatementInfo) SetHistoryExists(v bool) {
	o.HistoryExists = &v
}

func (o ARAccountTypeLastStatementInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ARAccountTypeLastStatementInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportName) {
		toSerialize["reportName"] = o.ReportName
	}
	if !IsNil(o.ReportFileName) {
		toSerialize["reportFileName"] = o.ReportFileName
	}
	if !IsNil(o.DateSent) {
		toSerialize["dateSent"] = o.DateSent
	}
	if !IsNil(o.StatementNo) {
		toSerialize["statementNo"] = o.StatementNo
	}
	if !IsNil(o.HistoryExists) {
		toSerialize["historyExists"] = o.HistoryExists
	}
	return toSerialize, nil
}

type NullableARAccountTypeLastStatementInfo struct {
	value *ARAccountTypeLastStatementInfo
	isSet bool
}

func (v NullableARAccountTypeLastStatementInfo) Get() *ARAccountTypeLastStatementInfo {
	return v.value
}

func (v *NullableARAccountTypeLastStatementInfo) Set(val *ARAccountTypeLastStatementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableARAccountTypeLastStatementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableARAccountTypeLastStatementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableARAccountTypeLastStatementInfo(val *ARAccountTypeLastStatementInfo) *NullableARAccountTypeLastStatementInfo {
	return &NullableARAccountTypeLastStatementInfo{value: val, isSet: true}
}

func (v NullableARAccountTypeLastStatementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableARAccountTypeLastStatementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


