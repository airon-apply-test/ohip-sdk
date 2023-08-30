/*
OPERA Cloud Xchange Interface OXI API

APIs to create and manage OPERA Xchange Interface (OXI) configurations for OPERA Cloud Exchange module functionality.<br /><br /> Compatible with OPERA Cloud release 23.0.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 23.0.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package oxi

import (
	"encoding/json"
)

// checks if the ExchangeStatReportType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExchangeStatReportType{}

// ExchangeStatReportType Concrete exchange interface report type.
type ExchangeStatReportType struct {
	Stat []ExchangeStatType `json:"stat,omitempty"`
	ReportCode *ExchangeStatReportCodeType `json:"reportCode,omitempty"`
	Start *string `json:"start,omitempty"`
	End *string `json:"end,omitempty"`
}

// NewExchangeStatReportType instantiates a new ExchangeStatReportType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeStatReportType() *ExchangeStatReportType {
	this := ExchangeStatReportType{}
	return &this
}

// NewExchangeStatReportTypeWithDefaults instantiates a new ExchangeStatReportType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeStatReportTypeWithDefaults() *ExchangeStatReportType {
	this := ExchangeStatReportType{}
	return &this
}

// GetStat returns the Stat field value if set, zero value otherwise.
func (o *ExchangeStatReportType) GetStat() []ExchangeStatType {
	if o == nil || IsNil(o.Stat) {
		var ret []ExchangeStatType
		return ret
	}
	return o.Stat
}

// GetStatOk returns a tuple with the Stat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStatReportType) GetStatOk() ([]ExchangeStatType, bool) {
	if o == nil || IsNil(o.Stat) {
		return nil, false
	}
	return o.Stat, true
}

// HasStat returns a boolean if a field has been set.
func (o *ExchangeStatReportType) HasStat() bool {
	if o != nil && !IsNil(o.Stat) {
		return true
	}

	return false
}

// SetStat gets a reference to the given []ExchangeStatType and assigns it to the Stat field.
func (o *ExchangeStatReportType) SetStat(v []ExchangeStatType) {
	o.Stat = v
}

// GetReportCode returns the ReportCode field value if set, zero value otherwise.
func (o *ExchangeStatReportType) GetReportCode() ExchangeStatReportCodeType {
	if o == nil || IsNil(o.ReportCode) {
		var ret ExchangeStatReportCodeType
		return ret
	}
	return *o.ReportCode
}

// GetReportCodeOk returns a tuple with the ReportCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStatReportType) GetReportCodeOk() (*ExchangeStatReportCodeType, bool) {
	if o == nil || IsNil(o.ReportCode) {
		return nil, false
	}
	return o.ReportCode, true
}

// HasReportCode returns a boolean if a field has been set.
func (o *ExchangeStatReportType) HasReportCode() bool {
	if o != nil && !IsNil(o.ReportCode) {
		return true
	}

	return false
}

// SetReportCode gets a reference to the given ExchangeStatReportCodeType and assigns it to the ReportCode field.
func (o *ExchangeStatReportType) SetReportCode(v ExchangeStatReportCodeType) {
	o.ReportCode = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ExchangeStatReportType) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStatReportType) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ExchangeStatReportType) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *ExchangeStatReportType) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ExchangeStatReportType) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeStatReportType) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ExchangeStatReportType) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *ExchangeStatReportType) SetEnd(v string) {
	o.End = &v
}

func (o ExchangeStatReportType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExchangeStatReportType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stat) {
		toSerialize["stat"] = o.Stat
	}
	if !IsNil(o.ReportCode) {
		toSerialize["reportCode"] = o.ReportCode
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableExchangeStatReportType struct {
	value *ExchangeStatReportType
	isSet bool
}

func (v NullableExchangeStatReportType) Get() *ExchangeStatReportType {
	return v.value
}

func (v *NullableExchangeStatReportType) Set(val *ExchangeStatReportType) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeStatReportType) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeStatReportType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeStatReportType(val *ExchangeStatReportType) *NullableExchangeStatReportType {
	return &NullableExchangeStatReportType{value: val, isSet: true}
}

func (v NullableExchangeStatReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeStatReportType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


