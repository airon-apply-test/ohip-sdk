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
	"time"
)

// checks if the ProfileSubscriptionType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileSubscriptionType{}

// ProfileSubscriptionType Contains details of the profile subscription. The subscription represents the link between the OPERA profile and the external profile within a particular external system
type ProfileSubscriptionType struct {
	ProfileId *ProfileId `json:"profileId,omitempty"`
	ExternalProfileId *UniqueIDType `json:"externalProfileId,omitempty"`
	ProfileInfo *ProfileSubscriptionTypeProfileInfo `json:"profileInfo,omitempty"`
	// Indicates if the profile was distributed to the external system.
	Distributed *bool `json:"distributed,omitempty"`
	// Indicates if the profile information should be overwritten by the external system.
	Force *bool `json:"force,omitempty"`
	// Timestamp when the profile was subscribed to.
	SubscriptionDate *time.Time `json:"subscriptionDate,omitempty"`
	// Timestamp of the most recent distribution of this profile to the external system.
	LastDistributionDate *time.Time `json:"lastDistributionDate,omitempty"`
	// Timestamp of the most recent update of the subscription information by the external system.
	LastExternalUpdateDate *time.Time `json:"lastExternalUpdateDate,omitempty"`
	// Indicates whether this subscription is active or inactive.
	Inactive *bool `json:"inactive,omitempty"`
	// Code of system where profile is subscribed to.
	SystemCode *string `json:"systemCode,omitempty"`
	// Type of system where profile is subscribed to.
	SystemType *string `json:"systemType,omitempty"`
}

// NewProfileSubscriptionType instantiates a new ProfileSubscriptionType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileSubscriptionType() *ProfileSubscriptionType {
	this := ProfileSubscriptionType{}
	return &this
}

// NewProfileSubscriptionTypeWithDefaults instantiates a new ProfileSubscriptionType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileSubscriptionTypeWithDefaults() *ProfileSubscriptionType {
	this := ProfileSubscriptionType{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *ProfileSubscriptionType) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

// GetExternalProfileId returns the ExternalProfileId field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetExternalProfileId() UniqueIDType {
	if o == nil || IsNil(o.ExternalProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ExternalProfileId
}

// GetExternalProfileIdOk returns a tuple with the ExternalProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetExternalProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ExternalProfileId) {
		return nil, false
	}
	return o.ExternalProfileId, true
}

// HasExternalProfileId returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasExternalProfileId() bool {
	if o != nil && !IsNil(o.ExternalProfileId) {
		return true
	}

	return false
}

// SetExternalProfileId gets a reference to the given UniqueIDType and assigns it to the ExternalProfileId field.
func (o *ProfileSubscriptionType) SetExternalProfileId(v UniqueIDType) {
	o.ExternalProfileId = &v
}

// GetProfileInfo returns the ProfileInfo field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetProfileInfo() ProfileSubscriptionTypeProfileInfo {
	if o == nil || IsNil(o.ProfileInfo) {
		var ret ProfileSubscriptionTypeProfileInfo
		return ret
	}
	return *o.ProfileInfo
}

// GetProfileInfoOk returns a tuple with the ProfileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetProfileInfoOk() (*ProfileSubscriptionTypeProfileInfo, bool) {
	if o == nil || IsNil(o.ProfileInfo) {
		return nil, false
	}
	return o.ProfileInfo, true
}

// HasProfileInfo returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasProfileInfo() bool {
	if o != nil && !IsNil(o.ProfileInfo) {
		return true
	}

	return false
}

// SetProfileInfo gets a reference to the given ProfileSubscriptionTypeProfileInfo and assigns it to the ProfileInfo field.
func (o *ProfileSubscriptionType) SetProfileInfo(v ProfileSubscriptionTypeProfileInfo) {
	o.ProfileInfo = &v
}

// GetDistributed returns the Distributed field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetDistributed() bool {
	if o == nil || IsNil(o.Distributed) {
		var ret bool
		return ret
	}
	return *o.Distributed
}

// GetDistributedOk returns a tuple with the Distributed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetDistributedOk() (*bool, bool) {
	if o == nil || IsNil(o.Distributed) {
		return nil, false
	}
	return o.Distributed, true
}

// HasDistributed returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasDistributed() bool {
	if o != nil && !IsNil(o.Distributed) {
		return true
	}

	return false
}

// SetDistributed gets a reference to the given bool and assigns it to the Distributed field.
func (o *ProfileSubscriptionType) SetDistributed(v bool) {
	o.Distributed = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *ProfileSubscriptionType) SetForce(v bool) {
	o.Force = &v
}

// GetSubscriptionDate returns the SubscriptionDate field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetSubscriptionDate() time.Time {
	if o == nil || IsNil(o.SubscriptionDate) {
		var ret time.Time
		return ret
	}
	return *o.SubscriptionDate
}

// GetSubscriptionDateOk returns a tuple with the SubscriptionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetSubscriptionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.SubscriptionDate) {
		return nil, false
	}
	return o.SubscriptionDate, true
}

// HasSubscriptionDate returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasSubscriptionDate() bool {
	if o != nil && !IsNil(o.SubscriptionDate) {
		return true
	}

	return false
}

// SetSubscriptionDate gets a reference to the given time.Time and assigns it to the SubscriptionDate field.
func (o *ProfileSubscriptionType) SetSubscriptionDate(v time.Time) {
	o.SubscriptionDate = &v
}

// GetLastDistributionDate returns the LastDistributionDate field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetLastDistributionDate() time.Time {
	if o == nil || IsNil(o.LastDistributionDate) {
		var ret time.Time
		return ret
	}
	return *o.LastDistributionDate
}

// GetLastDistributionDateOk returns a tuple with the LastDistributionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetLastDistributionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastDistributionDate) {
		return nil, false
	}
	return o.LastDistributionDate, true
}

// HasLastDistributionDate returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasLastDistributionDate() bool {
	if o != nil && !IsNil(o.LastDistributionDate) {
		return true
	}

	return false
}

// SetLastDistributionDate gets a reference to the given time.Time and assigns it to the LastDistributionDate field.
func (o *ProfileSubscriptionType) SetLastDistributionDate(v time.Time) {
	o.LastDistributionDate = &v
}

// GetLastExternalUpdateDate returns the LastExternalUpdateDate field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetLastExternalUpdateDate() time.Time {
	if o == nil || IsNil(o.LastExternalUpdateDate) {
		var ret time.Time
		return ret
	}
	return *o.LastExternalUpdateDate
}

// GetLastExternalUpdateDateOk returns a tuple with the LastExternalUpdateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetLastExternalUpdateDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastExternalUpdateDate) {
		return nil, false
	}
	return o.LastExternalUpdateDate, true
}

// HasLastExternalUpdateDate returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasLastExternalUpdateDate() bool {
	if o != nil && !IsNil(o.LastExternalUpdateDate) {
		return true
	}

	return false
}

// SetLastExternalUpdateDate gets a reference to the given time.Time and assigns it to the LastExternalUpdateDate field.
func (o *ProfileSubscriptionType) SetLastExternalUpdateDate(v time.Time) {
	o.LastExternalUpdateDate = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *ProfileSubscriptionType) SetInactive(v bool) {
	o.Inactive = &v
}

// GetSystemCode returns the SystemCode field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetSystemCode() string {
	if o == nil || IsNil(o.SystemCode) {
		var ret string
		return ret
	}
	return *o.SystemCode
}

// GetSystemCodeOk returns a tuple with the SystemCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetSystemCodeOk() (*string, bool) {
	if o == nil || IsNil(o.SystemCode) {
		return nil, false
	}
	return o.SystemCode, true
}

// HasSystemCode returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasSystemCode() bool {
	if o != nil && !IsNil(o.SystemCode) {
		return true
	}

	return false
}

// SetSystemCode gets a reference to the given string and assigns it to the SystemCode field.
func (o *ProfileSubscriptionType) SetSystemCode(v string) {
	o.SystemCode = &v
}

// GetSystemType returns the SystemType field value if set, zero value otherwise.
func (o *ProfileSubscriptionType) GetSystemType() string {
	if o == nil || IsNil(o.SystemType) {
		var ret string
		return ret
	}
	return *o.SystemType
}

// GetSystemTypeOk returns a tuple with the SystemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileSubscriptionType) GetSystemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SystemType) {
		return nil, false
	}
	return o.SystemType, true
}

// HasSystemType returns a boolean if a field has been set.
func (o *ProfileSubscriptionType) HasSystemType() bool {
	if o != nil && !IsNil(o.SystemType) {
		return true
	}

	return false
}

// SetSystemType gets a reference to the given string and assigns it to the SystemType field.
func (o *ProfileSubscriptionType) SetSystemType(v string) {
	o.SystemType = &v
}

func (o ProfileSubscriptionType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileSubscriptionType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.ExternalProfileId) {
		toSerialize["externalProfileId"] = o.ExternalProfileId
	}
	if !IsNil(o.ProfileInfo) {
		toSerialize["profileInfo"] = o.ProfileInfo
	}
	if !IsNil(o.Distributed) {
		toSerialize["distributed"] = o.Distributed
	}
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.SubscriptionDate) {
		toSerialize["subscriptionDate"] = o.SubscriptionDate
	}
	if !IsNil(o.LastDistributionDate) {
		toSerialize["lastDistributionDate"] = o.LastDistributionDate
	}
	if !IsNil(o.LastExternalUpdateDate) {
		toSerialize["lastExternalUpdateDate"] = o.LastExternalUpdateDate
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.SystemCode) {
		toSerialize["systemCode"] = o.SystemCode
	}
	if !IsNil(o.SystemType) {
		toSerialize["systemType"] = o.SystemType
	}
	return toSerialize, nil
}

type NullableProfileSubscriptionType struct {
	value *ProfileSubscriptionType
	isSet bool
}

func (v NullableProfileSubscriptionType) Get() *ProfileSubscriptionType {
	return v.value
}

func (v *NullableProfileSubscriptionType) Set(val *ProfileSubscriptionType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileSubscriptionType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileSubscriptionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileSubscriptionType(val *ProfileSubscriptionType) *NullableProfileSubscriptionType {
	return &NullableProfileSubscriptionType{value: val, isSet: true}
}

func (v NullableProfileSubscriptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileSubscriptionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


