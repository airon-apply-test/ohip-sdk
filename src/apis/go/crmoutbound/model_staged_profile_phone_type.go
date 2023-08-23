/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the StagedProfilePhoneType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StagedProfilePhoneType{}

// StagedProfilePhoneType Information on a telephone number for the customer.
type StagedProfilePhoneType struct {
	Telephone *TelephoneType `json:"telephone,omitempty"`
	// The error in telephone information in a staged profile with an invalid status.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// Inactivation date of the record.
	InactiveDate *string `json:"inactiveDate,omitempty"`
	// Indicates whether the phone is the default confirmation method.
	DefaultConfirmation *bool `json:"defaultConfirmation,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
}

// NewStagedProfilePhoneType instantiates a new StagedProfilePhoneType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagedProfilePhoneType() *StagedProfilePhoneType {
	this := StagedProfilePhoneType{}
	return &this
}

// NewStagedProfilePhoneTypeWithDefaults instantiates a new StagedProfilePhoneType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagedProfilePhoneTypeWithDefaults() *StagedProfilePhoneType {
	this := StagedProfilePhoneType{}
	return &this
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *StagedProfilePhoneType) GetTelephone() TelephoneType {
	if o == nil || IsNil(o.Telephone) {
		var ret TelephoneType
		return ret
	}
	return *o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfilePhoneType) GetTelephoneOk() (*TelephoneType, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *StagedProfilePhoneType) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given TelephoneType and assigns it to the Telephone field.
func (o *StagedProfilePhoneType) SetTelephone(v TelephoneType) {
	o.Telephone = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *StagedProfilePhoneType) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfilePhoneType) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *StagedProfilePhoneType) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *StagedProfilePhoneType) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetInactiveDate returns the InactiveDate field value if set, zero value otherwise.
func (o *StagedProfilePhoneType) GetInactiveDate() string {
	if o == nil || IsNil(o.InactiveDate) {
		var ret string
		return ret
	}
	return *o.InactiveDate
}

// GetInactiveDateOk returns a tuple with the InactiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfilePhoneType) GetInactiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.InactiveDate) {
		return nil, false
	}
	return o.InactiveDate, true
}

// HasInactiveDate returns a boolean if a field has been set.
func (o *StagedProfilePhoneType) HasInactiveDate() bool {
	if o != nil && !IsNil(o.InactiveDate) {
		return true
	}

	return false
}

// SetInactiveDate gets a reference to the given string and assigns it to the InactiveDate field.
func (o *StagedProfilePhoneType) SetInactiveDate(v string) {
	o.InactiveDate = &v
}

// GetDefaultConfirmation returns the DefaultConfirmation field value if set, zero value otherwise.
func (o *StagedProfilePhoneType) GetDefaultConfirmation() bool {
	if o == nil || IsNil(o.DefaultConfirmation) {
		var ret bool
		return ret
	}
	return *o.DefaultConfirmation
}

// GetDefaultConfirmationOk returns a tuple with the DefaultConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfilePhoneType) GetDefaultConfirmationOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultConfirmation) {
		return nil, false
	}
	return o.DefaultConfirmation, true
}

// HasDefaultConfirmation returns a boolean if a field has been set.
func (o *StagedProfilePhoneType) HasDefaultConfirmation() bool {
	if o != nil && !IsNil(o.DefaultConfirmation) {
		return true
	}

	return false
}

// SetDefaultConfirmation gets a reference to the given bool and assigns it to the DefaultConfirmation field.
func (o *StagedProfilePhoneType) SetDefaultConfirmation(v bool) {
	o.DefaultConfirmation = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StagedProfilePhoneType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfilePhoneType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StagedProfilePhoneType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StagedProfilePhoneType) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StagedProfilePhoneType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfilePhoneType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StagedProfilePhoneType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StagedProfilePhoneType) SetType(v string) {
	o.Type = &v
}

func (o StagedProfilePhoneType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StagedProfilePhoneType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Telephone) {
		toSerialize["telephone"] = o.Telephone
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if !IsNil(o.InactiveDate) {
		toSerialize["inactiveDate"] = o.InactiveDate
	}
	if !IsNil(o.DefaultConfirmation) {
		toSerialize["defaultConfirmation"] = o.DefaultConfirmation
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableStagedProfilePhoneType struct {
	value *StagedProfilePhoneType
	isSet bool
}

func (v NullableStagedProfilePhoneType) Get() *StagedProfilePhoneType {
	return v.value
}

func (v *NullableStagedProfilePhoneType) Set(val *StagedProfilePhoneType) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedProfilePhoneType) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedProfilePhoneType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedProfilePhoneType(val *StagedProfilePhoneType) *NullableStagedProfilePhoneType {
	return &NullableStagedProfilePhoneType{value: val, isSet: true}
}

func (v NullableStagedProfilePhoneType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedProfilePhoneType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


