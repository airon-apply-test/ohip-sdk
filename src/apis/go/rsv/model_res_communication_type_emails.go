/*
OPERA Cloud Reservation API

APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ResCommunicationTypeEmails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResCommunicationTypeEmails{}

// ResCommunicationTypeEmails List of email address for the customer.
type ResCommunicationTypeEmails struct {
	// Collection of Detailed information on an eMail address for the customer.
	EmailInfo []EmailInfoType `json:"emailInfo,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewResCommunicationTypeEmails instantiates a new ResCommunicationTypeEmails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResCommunicationTypeEmails() *ResCommunicationTypeEmails {
	this := ResCommunicationTypeEmails{}
	return &this
}

// NewResCommunicationTypeEmailsWithDefaults instantiates a new ResCommunicationTypeEmails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResCommunicationTypeEmailsWithDefaults() *ResCommunicationTypeEmails {
	this := ResCommunicationTypeEmails{}
	return &this
}

// GetEmailInfo returns the EmailInfo field value if set, zero value otherwise.
func (o *ResCommunicationTypeEmails) GetEmailInfo() []EmailInfoType {
	if o == nil || IsNil(o.EmailInfo) {
		var ret []EmailInfoType
		return ret
	}
	return o.EmailInfo
}

// GetEmailInfoOk returns a tuple with the EmailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeEmails) GetEmailInfoOk() ([]EmailInfoType, bool) {
	if o == nil || IsNil(o.EmailInfo) {
		return nil, false
	}
	return o.EmailInfo, true
}

// HasEmailInfo returns a boolean if a field has been set.
func (o *ResCommunicationTypeEmails) HasEmailInfo() bool {
	if o != nil && !IsNil(o.EmailInfo) {
		return true
	}

	return false
}

// SetEmailInfo gets a reference to the given []EmailInfoType and assigns it to the EmailInfo field.
func (o *ResCommunicationTypeEmails) SetEmailInfo(v []EmailInfoType) {
	o.EmailInfo = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ResCommunicationTypeEmails) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeEmails) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ResCommunicationTypeEmails) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ResCommunicationTypeEmails) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ResCommunicationTypeEmails) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeEmails) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ResCommunicationTypeEmails) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ResCommunicationTypeEmails) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ResCommunicationTypeEmails) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResCommunicationTypeEmails) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ResCommunicationTypeEmails) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ResCommunicationTypeEmails) SetCount(v int32) {
	o.Count = &v
}

func (o ResCommunicationTypeEmails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResCommunicationTypeEmails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailInfo) {
		toSerialize["emailInfo"] = o.EmailInfo
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

type NullableResCommunicationTypeEmails struct {
	value *ResCommunicationTypeEmails
	isSet bool
}

func (v NullableResCommunicationTypeEmails) Get() *ResCommunicationTypeEmails {
	return v.value
}

func (v *NullableResCommunicationTypeEmails) Set(val *ResCommunicationTypeEmails) {
	v.value = val
	v.isSet = true
}

func (v NullableResCommunicationTypeEmails) IsSet() bool {
	return v.isSet
}

func (v *NullableResCommunicationTypeEmails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResCommunicationTypeEmails(val *ResCommunicationTypeEmails) *NullableResCommunicationTypeEmails {
	return &NullableResCommunicationTypeEmails{value: val, isSet: true}
}

func (v NullableResCommunicationTypeEmails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResCommunicationTypeEmails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


