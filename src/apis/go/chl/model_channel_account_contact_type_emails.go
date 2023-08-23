/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ChannelAccountContactTypeEmails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelAccountContactTypeEmails{}

// ChannelAccountContactTypeEmails List of email address for the customer.
type ChannelAccountContactTypeEmails struct {
	// Collection of Detailed information on an eMail address for the customer.
	EmailInfo []EmailInfoType `json:"emailInfo,omitempty"`
	// Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
	HasMore *bool `json:"hasMore,omitempty"`
	// Total number of rows queried
	TotalResults *int32 `json:"totalResults,omitempty"`
	// Total number of rows returned
	Count *int32 `json:"count,omitempty"`
}

// NewChannelAccountContactTypeEmails instantiates a new ChannelAccountContactTypeEmails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelAccountContactTypeEmails() *ChannelAccountContactTypeEmails {
	this := ChannelAccountContactTypeEmails{}
	return &this
}

// NewChannelAccountContactTypeEmailsWithDefaults instantiates a new ChannelAccountContactTypeEmails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelAccountContactTypeEmailsWithDefaults() *ChannelAccountContactTypeEmails {
	this := ChannelAccountContactTypeEmails{}
	return &this
}

// GetEmailInfo returns the EmailInfo field value if set, zero value otherwise.
func (o *ChannelAccountContactTypeEmails) GetEmailInfo() []EmailInfoType {
	if o == nil || IsNil(o.EmailInfo) {
		var ret []EmailInfoType
		return ret
	}
	return o.EmailInfo
}

// GetEmailInfoOk returns a tuple with the EmailInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContactTypeEmails) GetEmailInfoOk() ([]EmailInfoType, bool) {
	if o == nil || IsNil(o.EmailInfo) {
		return nil, false
	}
	return o.EmailInfo, true
}

// HasEmailInfo returns a boolean if a field has been set.
func (o *ChannelAccountContactTypeEmails) HasEmailInfo() bool {
	if o != nil && !IsNil(o.EmailInfo) {
		return true
	}

	return false
}

// SetEmailInfo gets a reference to the given []EmailInfoType and assigns it to the EmailInfo field.
func (o *ChannelAccountContactTypeEmails) SetEmailInfo(v []EmailInfoType) {
	o.EmailInfo = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ChannelAccountContactTypeEmails) GetHasMore() bool {
	if o == nil || IsNil(o.HasMore) {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContactTypeEmails) GetHasMoreOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMore) {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ChannelAccountContactTypeEmails) HasHasMore() bool {
	if o != nil && !IsNil(o.HasMore) {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ChannelAccountContactTypeEmails) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetTotalResults returns the TotalResults field value if set, zero value otherwise.
func (o *ChannelAccountContactTypeEmails) GetTotalResults() int32 {
	if o == nil || IsNil(o.TotalResults) {
		var ret int32
		return ret
	}
	return *o.TotalResults
}

// GetTotalResultsOk returns a tuple with the TotalResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContactTypeEmails) GetTotalResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResults) {
		return nil, false
	}
	return o.TotalResults, true
}

// HasTotalResults returns a boolean if a field has been set.
func (o *ChannelAccountContactTypeEmails) HasTotalResults() bool {
	if o != nil && !IsNil(o.TotalResults) {
		return true
	}

	return false
}

// SetTotalResults gets a reference to the given int32 and assigns it to the TotalResults field.
func (o *ChannelAccountContactTypeEmails) SetTotalResults(v int32) {
	o.TotalResults = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ChannelAccountContactTypeEmails) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelAccountContactTypeEmails) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ChannelAccountContactTypeEmails) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ChannelAccountContactTypeEmails) SetCount(v int32) {
	o.Count = &v
}

func (o ChannelAccountContactTypeEmails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelAccountContactTypeEmails) ToMap() (map[string]interface{}, error) {
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

type NullableChannelAccountContactTypeEmails struct {
	value *ChannelAccountContactTypeEmails
	isSet bool
}

func (v NullableChannelAccountContactTypeEmails) Get() *ChannelAccountContactTypeEmails {
	return v.value
}

func (v *NullableChannelAccountContactTypeEmails) Set(val *ChannelAccountContactTypeEmails) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelAccountContactTypeEmails) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelAccountContactTypeEmails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelAccountContactTypeEmails(val *ChannelAccountContactTypeEmails) *NullableChannelAccountContactTypeEmails {
	return &NullableChannelAccountContactTypeEmails{value: val, isSet: true}
}

func (v NullableChannelAccountContactTypeEmails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelAccountContactTypeEmails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


