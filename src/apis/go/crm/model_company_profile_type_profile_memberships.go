/*
OPERA Cloud Customer Relationship Management API

APIs to cater for Customer Relationship Management (profile) functionality in OPERA Cloud.  There are different types of profiles in OPERA Cloud, including Guest, Company, Travel Agent, Source, Group, and Contact profile types.  A profile can store and display a wide range of information about the guest, company, travel agent etc., depending upon the type of profile.  For example, a guest profile can store the guest name, address, contact information, details on billing, membership benefits, preferences and much more.  All profiles in OPERA when created are assigned a ProfileID.  This ID will be used throughout the CRM APIs.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CompanyProfileTypeProfileMemberships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyProfileTypeProfileMemberships{}

// CompanyProfileTypeProfileMemberships List of loyalty program(s) the profile is subscribed to.
type CompanyProfileTypeProfileMemberships struct {
	// Collection of Detailed information on memberships for the customer.
	ProfileMembership []ProfileMembershipType `json:"profileMembership,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewCompanyProfileTypeProfileMemberships instantiates a new CompanyProfileTypeProfileMemberships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyProfileTypeProfileMemberships() *CompanyProfileTypeProfileMemberships {
	this := CompanyProfileTypeProfileMemberships{}
	return &this
}

// NewCompanyProfileTypeProfileMembershipsWithDefaults instantiates a new CompanyProfileTypeProfileMemberships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyProfileTypeProfileMembershipsWithDefaults() *CompanyProfileTypeProfileMemberships {
	this := CompanyProfileTypeProfileMemberships{}
	return &this
}

// GetProfileMembership returns the ProfileMembership field value if set, zero value otherwise.
func (o *CompanyProfileTypeProfileMemberships) GetProfileMembership() []ProfileMembershipType {
	if o == nil || IsNil(o.ProfileMembership) {
		var ret []ProfileMembershipType
		return ret
	}
	return o.ProfileMembership
}

// GetProfileMembershipOk returns a tuple with the ProfileMembership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileTypeProfileMemberships) GetProfileMembershipOk() ([]ProfileMembershipType, bool) {
	if o == nil || IsNil(o.ProfileMembership) {
		return nil, false
	}
	return o.ProfileMembership, true
}

// HasProfileMembership returns a boolean if a field has been set.
func (o *CompanyProfileTypeProfileMemberships) HasProfileMembership() bool {
	if o != nil && !IsNil(o.ProfileMembership) {
		return true
	}

	return false
}

// SetProfileMembership gets a reference to the given []ProfileMembershipType and assigns it to the ProfileMembership field.
func (o *CompanyProfileTypeProfileMemberships) SetProfileMembership(v []ProfileMembershipType) {
	o.ProfileMembership = v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *CompanyProfileTypeProfileMemberships) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileTypeProfileMemberships) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *CompanyProfileTypeProfileMemberships) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *CompanyProfileTypeProfileMemberships) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CompanyProfileTypeProfileMemberships) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfileTypeProfileMemberships) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CompanyProfileTypeProfileMemberships) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CompanyProfileTypeProfileMemberships) SetCount(v int32) {
	o.Count = &v
}

func (o CompanyProfileTypeProfileMemberships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyProfileTypeProfileMemberships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProfileMembership) {
		toSerialize["profileMembership"] = o.ProfileMembership
	}
	if !IsNil(o.TotalResults) {
		toSerialize["totalResults"] = o.TotalResults
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableCompanyProfileTypeProfileMemberships struct {
	value *CompanyProfileTypeProfileMemberships
	isSet bool
}

func (v NullableCompanyProfileTypeProfileMemberships) Get() *CompanyProfileTypeProfileMemberships {
	return v.value
}

func (v *NullableCompanyProfileTypeProfileMemberships) Set(val *CompanyProfileTypeProfileMemberships) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyProfileTypeProfileMemberships) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyProfileTypeProfileMemberships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyProfileTypeProfileMemberships(val *CompanyProfileTypeProfileMemberships) *NullableCompanyProfileTypeProfileMemberships {
	return &NullableCompanyProfileTypeProfileMemberships{value: val, isSet: true}
}

func (v NullableCompanyProfileTypeProfileMemberships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyProfileTypeProfileMemberships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


