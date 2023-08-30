/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blk

import (
	"encoding/json"
)

// checks if the ChangedCateringStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangedCateringStatus{}

// ChangedCateringStatus Response object for the request to change catering status. Response contains information on the block whose status was successfully changed.
type ChangedCateringStatus struct {
	Block *BlockType `json:"block,omitempty"`
	CancellationDetails *CancellationDetailsType `json:"cancellationDetails,omitempty"`
	// Next catering status of the business block.
	CateringNextStatusList []BookingStatusDetailType `json:"cateringNextStatusList,omitempty"`
	// Collection of catering status history.
	CateringStatusChangeHistory []BookingStatusHistoryType `json:"cateringStatusChangeHistory,omitempty"`
	// Status/Info of the processed events.
	CateringEventsProcessedInfo []CateringEventsProcessedInfoType `json:"cateringEventsProcessedInfo,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewChangedCateringStatus instantiates a new ChangedCateringStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangedCateringStatus() *ChangedCateringStatus {
	this := ChangedCateringStatus{}
	return &this
}

// NewChangedCateringStatusWithDefaults instantiates a new ChangedCateringStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangedCateringStatusWithDefaults() *ChangedCateringStatus {
	this := ChangedCateringStatus{}
	return &this
}

// GetBlock returns the Block field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetBlock() BlockType {
	if o == nil || IsNil(o.Block) {
		var ret BlockType
		return ret
	}
	return *o.Block
}

// GetBlockOk returns a tuple with the Block field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetBlockOk() (*BlockType, bool) {
	if o == nil || IsNil(o.Block) {
		return nil, false
	}
	return o.Block, true
}

// HasBlock returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasBlock() bool {
	if o != nil && !IsNil(o.Block) {
		return true
	}

	return false
}

// SetBlock gets a reference to the given BlockType and assigns it to the Block field.
func (o *ChangedCateringStatus) SetBlock(v BlockType) {
	o.Block = &v
}

// GetCancellationDetails returns the CancellationDetails field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetCancellationDetails() CancellationDetailsType {
	if o == nil || IsNil(o.CancellationDetails) {
		var ret CancellationDetailsType
		return ret
	}
	return *o.CancellationDetails
}

// GetCancellationDetailsOk returns a tuple with the CancellationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetCancellationDetailsOk() (*CancellationDetailsType, bool) {
	if o == nil || IsNil(o.CancellationDetails) {
		return nil, false
	}
	return o.CancellationDetails, true
}

// HasCancellationDetails returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasCancellationDetails() bool {
	if o != nil && !IsNil(o.CancellationDetails) {
		return true
	}

	return false
}

// SetCancellationDetails gets a reference to the given CancellationDetailsType and assigns it to the CancellationDetails field.
func (o *ChangedCateringStatus) SetCancellationDetails(v CancellationDetailsType) {
	o.CancellationDetails = &v
}

// GetCateringNextStatusList returns the CateringNextStatusList field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetCateringNextStatusList() []BookingStatusDetailType {
	if o == nil || IsNil(o.CateringNextStatusList) {
		var ret []BookingStatusDetailType
		return ret
	}
	return o.CateringNextStatusList
}

// GetCateringNextStatusListOk returns a tuple with the CateringNextStatusList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetCateringNextStatusListOk() ([]BookingStatusDetailType, bool) {
	if o == nil || IsNil(o.CateringNextStatusList) {
		return nil, false
	}
	return o.CateringNextStatusList, true
}

// HasCateringNextStatusList returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasCateringNextStatusList() bool {
	if o != nil && !IsNil(o.CateringNextStatusList) {
		return true
	}

	return false
}

// SetCateringNextStatusList gets a reference to the given []BookingStatusDetailType and assigns it to the CateringNextStatusList field.
func (o *ChangedCateringStatus) SetCateringNextStatusList(v []BookingStatusDetailType) {
	o.CateringNextStatusList = v
}

// GetCateringStatusChangeHistory returns the CateringStatusChangeHistory field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetCateringStatusChangeHistory() []BookingStatusHistoryType {
	if o == nil || IsNil(o.CateringStatusChangeHistory) {
		var ret []BookingStatusHistoryType
		return ret
	}
	return o.CateringStatusChangeHistory
}

// GetCateringStatusChangeHistoryOk returns a tuple with the CateringStatusChangeHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetCateringStatusChangeHistoryOk() ([]BookingStatusHistoryType, bool) {
	if o == nil || IsNil(o.CateringStatusChangeHistory) {
		return nil, false
	}
	return o.CateringStatusChangeHistory, true
}

// HasCateringStatusChangeHistory returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasCateringStatusChangeHistory() bool {
	if o != nil && !IsNil(o.CateringStatusChangeHistory) {
		return true
	}

	return false
}

// SetCateringStatusChangeHistory gets a reference to the given []BookingStatusHistoryType and assigns it to the CateringStatusChangeHistory field.
func (o *ChangedCateringStatus) SetCateringStatusChangeHistory(v []BookingStatusHistoryType) {
	o.CateringStatusChangeHistory = v
}

// GetCateringEventsProcessedInfo returns the CateringEventsProcessedInfo field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetCateringEventsProcessedInfo() []CateringEventsProcessedInfoType {
	if o == nil || IsNil(o.CateringEventsProcessedInfo) {
		var ret []CateringEventsProcessedInfoType
		return ret
	}
	return o.CateringEventsProcessedInfo
}

// GetCateringEventsProcessedInfoOk returns a tuple with the CateringEventsProcessedInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetCateringEventsProcessedInfoOk() ([]CateringEventsProcessedInfoType, bool) {
	if o == nil || IsNil(o.CateringEventsProcessedInfo) {
		return nil, false
	}
	return o.CateringEventsProcessedInfo, true
}

// HasCateringEventsProcessedInfo returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasCateringEventsProcessedInfo() bool {
	if o != nil && !IsNil(o.CateringEventsProcessedInfo) {
		return true
	}

	return false
}

// SetCateringEventsProcessedInfo gets a reference to the given []CateringEventsProcessedInfoType and assigns it to the CateringEventsProcessedInfo field.
func (o *ChangedCateringStatus) SetCateringEventsProcessedInfo(v []CateringEventsProcessedInfoType) {
	o.CateringEventsProcessedInfo = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangedCateringStatus) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangedCateringStatus) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangedCateringStatus) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangedCateringStatus) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangedCateringStatus) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o ChangedCateringStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangedCateringStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Block) {
		toSerialize["block"] = o.Block
	}
	if !IsNil(o.CancellationDetails) {
		toSerialize["cancellationDetails"] = o.CancellationDetails
	}
	if !IsNil(o.CateringNextStatusList) {
		toSerialize["cateringNextStatusList"] = o.CateringNextStatusList
	}
	if !IsNil(o.CateringStatusChangeHistory) {
		toSerialize["cateringStatusChangeHistory"] = o.CateringStatusChangeHistory
	}
	if !IsNil(o.CateringEventsProcessedInfo) {
		toSerialize["cateringEventsProcessedInfo"] = o.CateringEventsProcessedInfo
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableChangedCateringStatus struct {
	value *ChangedCateringStatus
	isSet bool
}

func (v NullableChangedCateringStatus) Get() *ChangedCateringStatus {
	return v.value
}

func (v *NullableChangedCateringStatus) Set(val *ChangedCateringStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableChangedCateringStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableChangedCateringStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangedCateringStatus(val *ChangedCateringStatus) *NullableChangedCateringStatus {
	return &NullableChangedCateringStatus{value: val, isSet: true}
}

func (v NullableChangedCateringStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangedCateringStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


