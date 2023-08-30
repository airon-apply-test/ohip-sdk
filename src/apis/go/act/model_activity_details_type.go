/*
OPERA Cloud Activity API

APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package act

import (
	"encoding/json"
)

// checks if the ActivityDetailsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityDetailsType{}

// ActivityDetailsType Complete Activity Related Information.
type ActivityDetailsType struct {
	ActivityId *ActivityId `json:"activityId,omitempty"`
	ActivityDetail *ActivityInfoType `json:"activityDetail,omitempty"`
	LinkedAccounts *ActivityDetailsTypeLinkedAccounts `json:"linkedAccounts,omitempty"`
	LinkedContacts *ActivityDetailsTypeLinkedContacts `json:"linkedContacts,omitempty"`
	// List of Blocks that are linked to the Activity.
	LinkedBlocks []ActivityBlockInfoType `json:"linkedBlocks,omitempty"`
	// Attachment List.
	LinkedAttachments []AttachmentType `json:"linkedAttachments,omitempty"`
	// The list of activities associated with an activity.
	LinkedActivities []LinkedActivityDetailsType `json:"linkedActivities,omitempty"`
	// Collection of lamp indicators.
	Indicators []IndicatorType `json:"indicators,omitempty"`
}

// NewActivityDetailsType instantiates a new ActivityDetailsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityDetailsType() *ActivityDetailsType {
	this := ActivityDetailsType{}
	return &this
}

// NewActivityDetailsTypeWithDefaults instantiates a new ActivityDetailsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityDetailsTypeWithDefaults() *ActivityDetailsType {
	this := ActivityDetailsType{}
	return &this
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetActivityId() ActivityId {
	if o == nil || IsNil(o.ActivityId) {
		var ret ActivityId
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetActivityIdOk() (*ActivityId, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given ActivityId and assigns it to the ActivityId field.
func (o *ActivityDetailsType) SetActivityId(v ActivityId) {
	o.ActivityId = &v
}

// GetActivityDetail returns the ActivityDetail field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetActivityDetail() ActivityInfoType {
	if o == nil || IsNil(o.ActivityDetail) {
		var ret ActivityInfoType
		return ret
	}
	return *o.ActivityDetail
}

// GetActivityDetailOk returns a tuple with the ActivityDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetActivityDetailOk() (*ActivityInfoType, bool) {
	if o == nil || IsNil(o.ActivityDetail) {
		return nil, false
	}
	return o.ActivityDetail, true
}

// HasActivityDetail returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasActivityDetail() bool {
	if o != nil && !IsNil(o.ActivityDetail) {
		return true
	}

	return false
}

// SetActivityDetail gets a reference to the given ActivityInfoType and assigns it to the ActivityDetail field.
func (o *ActivityDetailsType) SetActivityDetail(v ActivityInfoType) {
	o.ActivityDetail = &v
}

// GetLinkedAccounts returns the LinkedAccounts field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetLinkedAccounts() ActivityDetailsTypeLinkedAccounts {
	if o == nil || IsNil(o.LinkedAccounts) {
		var ret ActivityDetailsTypeLinkedAccounts
		return ret
	}
	return *o.LinkedAccounts
}

// GetLinkedAccountsOk returns a tuple with the LinkedAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetLinkedAccountsOk() (*ActivityDetailsTypeLinkedAccounts, bool) {
	if o == nil || IsNil(o.LinkedAccounts) {
		return nil, false
	}
	return o.LinkedAccounts, true
}

// HasLinkedAccounts returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasLinkedAccounts() bool {
	if o != nil && !IsNil(o.LinkedAccounts) {
		return true
	}

	return false
}

// SetLinkedAccounts gets a reference to the given ActivityDetailsTypeLinkedAccounts and assigns it to the LinkedAccounts field.
func (o *ActivityDetailsType) SetLinkedAccounts(v ActivityDetailsTypeLinkedAccounts) {
	o.LinkedAccounts = &v
}

// GetLinkedContacts returns the LinkedContacts field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetLinkedContacts() ActivityDetailsTypeLinkedContacts {
	if o == nil || IsNil(o.LinkedContacts) {
		var ret ActivityDetailsTypeLinkedContacts
		return ret
	}
	return *o.LinkedContacts
}

// GetLinkedContactsOk returns a tuple with the LinkedContacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetLinkedContactsOk() (*ActivityDetailsTypeLinkedContacts, bool) {
	if o == nil || IsNil(o.LinkedContacts) {
		return nil, false
	}
	return o.LinkedContacts, true
}

// HasLinkedContacts returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasLinkedContacts() bool {
	if o != nil && !IsNil(o.LinkedContacts) {
		return true
	}

	return false
}

// SetLinkedContacts gets a reference to the given ActivityDetailsTypeLinkedContacts and assigns it to the LinkedContacts field.
func (o *ActivityDetailsType) SetLinkedContacts(v ActivityDetailsTypeLinkedContacts) {
	o.LinkedContacts = &v
}

// GetLinkedBlocks returns the LinkedBlocks field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetLinkedBlocks() []ActivityBlockInfoType {
	if o == nil || IsNil(o.LinkedBlocks) {
		var ret []ActivityBlockInfoType
		return ret
	}
	return o.LinkedBlocks
}

// GetLinkedBlocksOk returns a tuple with the LinkedBlocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetLinkedBlocksOk() ([]ActivityBlockInfoType, bool) {
	if o == nil || IsNil(o.LinkedBlocks) {
		return nil, false
	}
	return o.LinkedBlocks, true
}

// HasLinkedBlocks returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasLinkedBlocks() bool {
	if o != nil && !IsNil(o.LinkedBlocks) {
		return true
	}

	return false
}

// SetLinkedBlocks gets a reference to the given []ActivityBlockInfoType and assigns it to the LinkedBlocks field.
func (o *ActivityDetailsType) SetLinkedBlocks(v []ActivityBlockInfoType) {
	o.LinkedBlocks = v
}

// GetLinkedAttachments returns the LinkedAttachments field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetLinkedAttachments() []AttachmentType {
	if o == nil || IsNil(o.LinkedAttachments) {
		var ret []AttachmentType
		return ret
	}
	return o.LinkedAttachments
}

// GetLinkedAttachmentsOk returns a tuple with the LinkedAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetLinkedAttachmentsOk() ([]AttachmentType, bool) {
	if o == nil || IsNil(o.LinkedAttachments) {
		return nil, false
	}
	return o.LinkedAttachments, true
}

// HasLinkedAttachments returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasLinkedAttachments() bool {
	if o != nil && !IsNil(o.LinkedAttachments) {
		return true
	}

	return false
}

// SetLinkedAttachments gets a reference to the given []AttachmentType and assigns it to the LinkedAttachments field.
func (o *ActivityDetailsType) SetLinkedAttachments(v []AttachmentType) {
	o.LinkedAttachments = v
}

// GetLinkedActivities returns the LinkedActivities field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetLinkedActivities() []LinkedActivityDetailsType {
	if o == nil || IsNil(o.LinkedActivities) {
		var ret []LinkedActivityDetailsType
		return ret
	}
	return o.LinkedActivities
}

// GetLinkedActivitiesOk returns a tuple with the LinkedActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetLinkedActivitiesOk() ([]LinkedActivityDetailsType, bool) {
	if o == nil || IsNil(o.LinkedActivities) {
		return nil, false
	}
	return o.LinkedActivities, true
}

// HasLinkedActivities returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasLinkedActivities() bool {
	if o != nil && !IsNil(o.LinkedActivities) {
		return true
	}

	return false
}

// SetLinkedActivities gets a reference to the given []LinkedActivityDetailsType and assigns it to the LinkedActivities field.
func (o *ActivityDetailsType) SetLinkedActivities(v []LinkedActivityDetailsType) {
	o.LinkedActivities = v
}

// GetIndicators returns the Indicators field value if set, zero value otherwise.
func (o *ActivityDetailsType) GetIndicators() []IndicatorType {
	if o == nil || IsNil(o.Indicators) {
		var ret []IndicatorType
		return ret
	}
	return o.Indicators
}

// GetIndicatorsOk returns a tuple with the Indicators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityDetailsType) GetIndicatorsOk() ([]IndicatorType, bool) {
	if o == nil || IsNil(o.Indicators) {
		return nil, false
	}
	return o.Indicators, true
}

// HasIndicators returns a boolean if a field has been set.
func (o *ActivityDetailsType) HasIndicators() bool {
	if o != nil && !IsNil(o.Indicators) {
		return true
	}

	return false
}

// SetIndicators gets a reference to the given []IndicatorType and assigns it to the Indicators field.
func (o *ActivityDetailsType) SetIndicators(v []IndicatorType) {
	o.Indicators = v
}

func (o ActivityDetailsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityDetailsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityId) {
		toSerialize["activityId"] = o.ActivityId
	}
	if !IsNil(o.ActivityDetail) {
		toSerialize["activityDetail"] = o.ActivityDetail
	}
	if !IsNil(o.LinkedAccounts) {
		toSerialize["linkedAccounts"] = o.LinkedAccounts
	}
	if !IsNil(o.LinkedContacts) {
		toSerialize["linkedContacts"] = o.LinkedContacts
	}
	if !IsNil(o.LinkedBlocks) {
		toSerialize["linkedBlocks"] = o.LinkedBlocks
	}
	if !IsNil(o.LinkedAttachments) {
		toSerialize["linkedAttachments"] = o.LinkedAttachments
	}
	if !IsNil(o.LinkedActivities) {
		toSerialize["linkedActivities"] = o.LinkedActivities
	}
	if !IsNil(o.Indicators) {
		toSerialize["indicators"] = o.Indicators
	}
	return toSerialize, nil
}

type NullableActivityDetailsType struct {
	value *ActivityDetailsType
	isSet bool
}

func (v NullableActivityDetailsType) Get() *ActivityDetailsType {
	return v.value
}

func (v *NullableActivityDetailsType) Set(val *ActivityDetailsType) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityDetailsType) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityDetailsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityDetailsType(val *ActivityDetailsType) *NullableActivityDetailsType {
	return &NullableActivityDetailsType{value: val, isSet: true}
}

func (v NullableActivityDetailsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityDetailsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


