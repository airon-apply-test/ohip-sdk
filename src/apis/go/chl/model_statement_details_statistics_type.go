/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the StatementDetailsStatisticsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatementDetailsStatisticsType{}

// StatementDetailsStatisticsType Channel statement details statistics .
type StatementDetailsStatisticsType struct {
	// Holds the statistic details for the statement details.
	DetailsByResort []StatementDetailsStatisticType `json:"detailsByResort,omitempty"`
	// Holds the statistic details for the statement details.
	DetailsByChannel []StatementDetailsStatisticType `json:"detailsByChannel,omitempty"`
	// Holds the statistic details for the statement details.
	DetailsByChannelType []StatementDetailsStatisticType `json:"detailsByChannelType,omitempty"`
	// Holds the statistic details for the statement details.
	DetailsByFeeType []StatementDetailsStatisticType `json:"detailsByFeeType,omitempty"`
}

// NewStatementDetailsStatisticsType instantiates a new StatementDetailsStatisticsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatementDetailsStatisticsType() *StatementDetailsStatisticsType {
	this := StatementDetailsStatisticsType{}
	return &this
}

// NewStatementDetailsStatisticsTypeWithDefaults instantiates a new StatementDetailsStatisticsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatementDetailsStatisticsTypeWithDefaults() *StatementDetailsStatisticsType {
	this := StatementDetailsStatisticsType{}
	return &this
}

// GetDetailsByResort returns the DetailsByResort field value if set, zero value otherwise.
func (o *StatementDetailsStatisticsType) GetDetailsByResort() []StatementDetailsStatisticType {
	if o == nil || IsNil(o.DetailsByResort) {
		var ret []StatementDetailsStatisticType
		return ret
	}
	return o.DetailsByResort
}

// GetDetailsByResortOk returns a tuple with the DetailsByResort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementDetailsStatisticsType) GetDetailsByResortOk() ([]StatementDetailsStatisticType, bool) {
	if o == nil || IsNil(o.DetailsByResort) {
		return nil, false
	}
	return o.DetailsByResort, true
}

// HasDetailsByResort returns a boolean if a field has been set.
func (o *StatementDetailsStatisticsType) HasDetailsByResort() bool {
	if o != nil && !IsNil(o.DetailsByResort) {
		return true
	}

	return false
}

// SetDetailsByResort gets a reference to the given []StatementDetailsStatisticType and assigns it to the DetailsByResort field.
func (o *StatementDetailsStatisticsType) SetDetailsByResort(v []StatementDetailsStatisticType) {
	o.DetailsByResort = v
}

// GetDetailsByChannel returns the DetailsByChannel field value if set, zero value otherwise.
func (o *StatementDetailsStatisticsType) GetDetailsByChannel() []StatementDetailsStatisticType {
	if o == nil || IsNil(o.DetailsByChannel) {
		var ret []StatementDetailsStatisticType
		return ret
	}
	return o.DetailsByChannel
}

// GetDetailsByChannelOk returns a tuple with the DetailsByChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementDetailsStatisticsType) GetDetailsByChannelOk() ([]StatementDetailsStatisticType, bool) {
	if o == nil || IsNil(o.DetailsByChannel) {
		return nil, false
	}
	return o.DetailsByChannel, true
}

// HasDetailsByChannel returns a boolean if a field has been set.
func (o *StatementDetailsStatisticsType) HasDetailsByChannel() bool {
	if o != nil && !IsNil(o.DetailsByChannel) {
		return true
	}

	return false
}

// SetDetailsByChannel gets a reference to the given []StatementDetailsStatisticType and assigns it to the DetailsByChannel field.
func (o *StatementDetailsStatisticsType) SetDetailsByChannel(v []StatementDetailsStatisticType) {
	o.DetailsByChannel = v
}

// GetDetailsByChannelType returns the DetailsByChannelType field value if set, zero value otherwise.
func (o *StatementDetailsStatisticsType) GetDetailsByChannelType() []StatementDetailsStatisticType {
	if o == nil || IsNil(o.DetailsByChannelType) {
		var ret []StatementDetailsStatisticType
		return ret
	}
	return o.DetailsByChannelType
}

// GetDetailsByChannelTypeOk returns a tuple with the DetailsByChannelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementDetailsStatisticsType) GetDetailsByChannelTypeOk() ([]StatementDetailsStatisticType, bool) {
	if o == nil || IsNil(o.DetailsByChannelType) {
		return nil, false
	}
	return o.DetailsByChannelType, true
}

// HasDetailsByChannelType returns a boolean if a field has been set.
func (o *StatementDetailsStatisticsType) HasDetailsByChannelType() bool {
	if o != nil && !IsNil(o.DetailsByChannelType) {
		return true
	}

	return false
}

// SetDetailsByChannelType gets a reference to the given []StatementDetailsStatisticType and assigns it to the DetailsByChannelType field.
func (o *StatementDetailsStatisticsType) SetDetailsByChannelType(v []StatementDetailsStatisticType) {
	o.DetailsByChannelType = v
}

// GetDetailsByFeeType returns the DetailsByFeeType field value if set, zero value otherwise.
func (o *StatementDetailsStatisticsType) GetDetailsByFeeType() []StatementDetailsStatisticType {
	if o == nil || IsNil(o.DetailsByFeeType) {
		var ret []StatementDetailsStatisticType
		return ret
	}
	return o.DetailsByFeeType
}

// GetDetailsByFeeTypeOk returns a tuple with the DetailsByFeeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatementDetailsStatisticsType) GetDetailsByFeeTypeOk() ([]StatementDetailsStatisticType, bool) {
	if o == nil || IsNil(o.DetailsByFeeType) {
		return nil, false
	}
	return o.DetailsByFeeType, true
}

// HasDetailsByFeeType returns a boolean if a field has been set.
func (o *StatementDetailsStatisticsType) HasDetailsByFeeType() bool {
	if o != nil && !IsNil(o.DetailsByFeeType) {
		return true
	}

	return false
}

// SetDetailsByFeeType gets a reference to the given []StatementDetailsStatisticType and assigns it to the DetailsByFeeType field.
func (o *StatementDetailsStatisticsType) SetDetailsByFeeType(v []StatementDetailsStatisticType) {
	o.DetailsByFeeType = v
}

func (o StatementDetailsStatisticsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatementDetailsStatisticsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DetailsByResort) {
		toSerialize["detailsByResort"] = o.DetailsByResort
	}
	if !IsNil(o.DetailsByChannel) {
		toSerialize["detailsByChannel"] = o.DetailsByChannel
	}
	if !IsNil(o.DetailsByChannelType) {
		toSerialize["detailsByChannelType"] = o.DetailsByChannelType
	}
	if !IsNil(o.DetailsByFeeType) {
		toSerialize["detailsByFeeType"] = o.DetailsByFeeType
	}
	return toSerialize, nil
}

type NullableStatementDetailsStatisticsType struct {
	value *StatementDetailsStatisticsType
	isSet bool
}

func (v NullableStatementDetailsStatisticsType) Get() *StatementDetailsStatisticsType {
	return v.value
}

func (v *NullableStatementDetailsStatisticsType) Set(val *StatementDetailsStatisticsType) {
	v.value = val
	v.isSet = true
}

func (v NullableStatementDetailsStatisticsType) IsSet() bool {
	return v.isSet
}

func (v *NullableStatementDetailsStatisticsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatementDetailsStatisticsType(val *StatementDetailsStatisticsType) *NullableStatementDetailsStatisticsType {
	return &NullableStatementDetailsStatisticsType{value: val, isSet: true}
}

func (v NullableStatementDetailsStatisticsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatementDetailsStatisticsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


