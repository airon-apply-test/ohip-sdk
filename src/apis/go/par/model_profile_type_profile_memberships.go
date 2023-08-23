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

// checks if the ProfileTypeProfileMemberships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTypeProfileMemberships{}

// ProfileTypeProfileMemberships List of loyalty program(s) the profile is subscribed to.
type ProfileTypeProfileMemberships struct {
	// Collection of Detailed information on memberships for the customer.
	ProfileMembership []ProfileMembershipType `json:"profileMembership,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewProfileTypeProfileMemberships instantiates a new ProfileTypeProfileMemberships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTypeProfileMemberships() *ProfileTypeProfileMemberships {
	this := ProfileTypeProfileMemberships{}
	return &this
}

// NewProfileTypeProfileMembershipsWithDefaults instantiates a new ProfileTypeProfileMemberships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTypeProfileMembershipsWithDefaults() *ProfileTypeProfileMemberships {
	this := ProfileTypeProfileMemberships{}
	return &this
}

// GetProfileMembership returns the ProfileMembership field value if set, zero value otherwise.
func (o *ProfileTypeProfileMemberships) GetProfileMembership() []ProfileMembershipType {
	if o == nil || IsNil(o.ProfileMembership) {
		var ret []ProfileMembershipType
		return ret
	}
	return o.ProfileMembership
}

// GetProfileMembershipOk returns a tuple with the ProfileMembership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeProfileMemberships) GetProfileMembershipOk() ([]ProfileMembershipType, bool) {
	if o == nil || IsNil(o.ProfileMembership) {
		return nil, false
	}
	return o.ProfileMembership, true
}

// HasProfileMembership returns a boolean if a field has been set.
func (o *ProfileTypeProfileMemberships) HasProfileMembership() bool {
	if o != nil && !IsNil(o.ProfileMembership) {
		return true
	}

	return false
}

// SetProfileMembership gets a reference to the given []ProfileMembershipType and assigns it to the ProfileMembership field.
func (o *ProfileTypeProfileMemberships) SetProfileMembership(v []ProfileMembershipType) {
	o.ProfileMembership = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ProfileTypeProfileMemberships) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeProfileMemberships) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ProfileTypeProfileMemberships) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ProfileTypeProfileMemberships) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ProfileTypeProfileMemberships) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeProfileMemberships) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ProfileTypeProfileMemberships) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ProfileTypeProfileMemberships) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ProfileTypeProfileMemberships) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTypeProfileMemberships) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ProfileTypeProfileMemberships) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ProfileTypeProfileMemberships) SetCount(v int32) {
	o.Count = &v
}

func (o ProfileTypeProfileMemberships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTypeProfileMemberships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileMembership) {
		toSerialize["profileMembership"] = o.ProfileMembership
	}
	if !IsNil(o.HasMore) {
		toSerialize["hasMore"] = o.HasMore
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableProfileTypeProfileMemberships struct {
	value *ProfileTypeProfileMemberships
	isSet bool
}

func (v NullableProfileTypeProfileMemberships) Get() *ProfileTypeProfileMemberships {
	return v.value
}

func (v *NullableProfileTypeProfileMemberships) Set(val *ProfileTypeProfileMemberships) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTypeProfileMemberships) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTypeProfileMemberships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTypeProfileMemberships(val *ProfileTypeProfileMemberships) *NullableProfileTypeProfileMemberships {
	return &NullableProfileTypeProfileMemberships{value: val, isSet: true}
}

func (v NullableProfileTypeProfileMemberships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTypeProfileMemberships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


