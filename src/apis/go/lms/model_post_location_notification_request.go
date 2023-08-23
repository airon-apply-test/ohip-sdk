/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PostLocationNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLocationNotificationRequest{}

// PostLocationNotificationRequest struct for PostLocationNotificationRequest
type PostLocationNotificationRequest struct {
	ReservationId *UniqueIDType `json:"reservationId,omitempty"`
	ActivityTime *DateTimeSpanType `json:"activityTime,omitempty"`
	LocationText *FormattedTextTextType `json:"locationText,omitempty"`
	ProfileId *UniqueIDType `json:"profileId,omitempty"`
	Description *FormattedTextTextType `json:"description,omitempty"`
	LocationNotificationStatus *LocationNotificationStatus `json:"locationNotificationStatus,omitempty"`
	OtherLocationNotificationStatus *string `json:"otherLocationNotificationStatus,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewPostLocationNotificationRequest instantiates a new PostLocationNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLocationNotificationRequest() *PostLocationNotificationRequest {
	this := PostLocationNotificationRequest{}
	return &this
}

// NewPostLocationNotificationRequestWithDefaults instantiates a new PostLocationNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLocationNotificationRequestWithDefaults() *PostLocationNotificationRequest {
	this := PostLocationNotificationRequest{}
	return &this
}

// GetReservationId returns the ReservationId field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetReservationId() UniqueIDType {
	if o == nil || IsNil(o.ReservationId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ReservationId
}

// GetReservationIdOk returns a tuple with the ReservationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetReservationIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ReservationId) {
		return nil, false
	}
	return o.ReservationId, true
}

// HasReservationId returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasReservationId() bool {
	if o != nil && !IsNil(o.ReservationId) {
		return true
	}

	return false
}

// SetReservationId gets a reference to the given UniqueIDType and assigns it to the ReservationId field.
func (o *PostLocationNotificationRequest) SetReservationId(v UniqueIDType) {
	o.ReservationId = &v
}

// GetActivityTime returns the ActivityTime field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetActivityTime() DateTimeSpanType {
	if o == nil || IsNil(o.ActivityTime) {
		var ret DateTimeSpanType
		return ret
	}
	return *o.ActivityTime
}

// GetActivityTimeOk returns a tuple with the ActivityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetActivityTimeOk() (*DateTimeSpanType, bool) {
	if o == nil || IsNil(o.ActivityTime) {
		return nil, false
	}
	return o.ActivityTime, true
}

// HasActivityTime returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasActivityTime() bool {
	if o != nil && !IsNil(o.ActivityTime) {
		return true
	}

	return false
}

// SetActivityTime gets a reference to the given DateTimeSpanType and assigns it to the ActivityTime field.
func (o *PostLocationNotificationRequest) SetActivityTime(v DateTimeSpanType) {
	o.ActivityTime = &v
}

// GetLocationText returns the LocationText field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetLocationText() FormattedTextTextType {
	if o == nil || IsNil(o.LocationText) {
		var ret FormattedTextTextType
		return ret
	}
	return *o.LocationText
}

// GetLocationTextOk returns a tuple with the LocationText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetLocationTextOk() (*FormattedTextTextType, bool) {
	if o == nil || IsNil(o.LocationText) {
		return nil, false
	}
	return o.LocationText, true
}

// HasLocationText returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasLocationText() bool {
	if o != nil && !IsNil(o.LocationText) {
		return true
	}

	return false
}

// SetLocationText gets a reference to the given FormattedTextTextType and assigns it to the LocationText field.
func (o *PostLocationNotificationRequest) SetLocationText(v FormattedTextTextType) {
	o.LocationText = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetProfileId() UniqueIDType {
	if o == nil || IsNil(o.ProfileId) {
		var ret UniqueIDType
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetProfileIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given UniqueIDType and assigns it to the ProfileId field.
func (o *PostLocationNotificationRequest) SetProfileId(v UniqueIDType) {
	o.ProfileId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetDescription() FormattedTextTextType {
	if o == nil || IsNil(o.Description) {
		var ret FormattedTextTextType
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetDescriptionOk() (*FormattedTextTextType, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given FormattedTextTextType and assigns it to the Description field.
func (o *PostLocationNotificationRequest) SetDescription(v FormattedTextTextType) {
	o.Description = &v
}

// GetLocationNotificationStatus returns the LocationNotificationStatus field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetLocationNotificationStatus() LocationNotificationStatus {
	if o == nil || IsNil(o.LocationNotificationStatus) {
		var ret LocationNotificationStatus
		return ret
	}
	return *o.LocationNotificationStatus
}

// GetLocationNotificationStatusOk returns a tuple with the LocationNotificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetLocationNotificationStatusOk() (*LocationNotificationStatus, bool) {
	if o == nil || IsNil(o.LocationNotificationStatus) {
		return nil, false
	}
	return o.LocationNotificationStatus, true
}

// HasLocationNotificationStatus returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasLocationNotificationStatus() bool {
	if o != nil && !IsNil(o.LocationNotificationStatus) {
		return true
	}

	return false
}

// SetLocationNotificationStatus gets a reference to the given LocationNotificationStatus and assigns it to the LocationNotificationStatus field.
func (o *PostLocationNotificationRequest) SetLocationNotificationStatus(v LocationNotificationStatus) {
	o.LocationNotificationStatus = &v
}

// GetOtherLocationNotificationStatus returns the OtherLocationNotificationStatus field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetOtherLocationNotificationStatus() string {
	if o == nil || IsNil(o.OtherLocationNotificationStatus) {
		var ret string
		return ret
	}
	return *o.OtherLocationNotificationStatus
}

// GetOtherLocationNotificationStatusOk returns a tuple with the OtherLocationNotificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetOtherLocationNotificationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OtherLocationNotificationStatus) {
		return nil, false
	}
	return o.OtherLocationNotificationStatus, true
}

// HasOtherLocationNotificationStatus returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasOtherLocationNotificationStatus() bool {
	if o != nil && !IsNil(o.OtherLocationNotificationStatus) {
		return true
	}

	return false
}

// SetOtherLocationNotificationStatus gets a reference to the given string and assigns it to the OtherLocationNotificationStatus field.
func (o *PostLocationNotificationRequest) SetOtherLocationNotificationStatus(v string) {
	o.OtherLocationNotificationStatus = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *PostLocationNotificationRequest) SetLinks(v []InstanceLink) {
	o.Links = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *PostLocationNotificationRequest) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLocationNotificationRequest) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *PostLocationNotificationRequest) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *PostLocationNotificationRequest) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o PostLocationNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLocationNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReservationId) {
		toSerialize["reservationId"] = o.ReservationId
	}
	if !IsNil(o.ActivityTime) {
		toSerialize["activityTime"] = o.ActivityTime
	}
	if !IsNil(o.LocationText) {
		toSerialize["locationText"] = o.LocationText
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.LocationNotificationStatus) {
		toSerialize["locationNotificationStatus"] = o.LocationNotificationStatus
	}
	if !IsNil(o.OtherLocationNotificationStatus) {
		toSerialize["otherLocationNotificationStatus"] = o.OtherLocationNotificationStatus
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullablePostLocationNotificationRequest struct {
	value *PostLocationNotificationRequest
	isSet bool
}

func (v NullablePostLocationNotificationRequest) Get() *PostLocationNotificationRequest {
	return v.value
}

func (v *NullablePostLocationNotificationRequest) Set(val *PostLocationNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLocationNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLocationNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLocationNotificationRequest(val *PostLocationNotificationRequest) *NullablePostLocationNotificationRequest {
	return &NullablePostLocationNotificationRequest{value: val, isSet: true}
}

func (v NullablePostLocationNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLocationNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


