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

// checks if the StagedProfileIdentificationsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StagedProfileIdentificationsType{}

// StagedProfileIdentificationsType Information on the identification of the customer.
type StagedProfileIdentificationsType struct {
	Identification *IdentificationType `json:"identification,omitempty"`
	// The error in negotiated rate information in a staged profile with an invalid status
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// Hotel code to which the document belongs.
	DocumentResort *string `json:"documentResort,omitempty"`
	// The date when the record was inactivated.
	InactiveDate *string `json:"inactiveDate,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
}

// NewStagedProfileIdentificationsType instantiates a new StagedProfileIdentificationsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStagedProfileIdentificationsType() *StagedProfileIdentificationsType {
	this := StagedProfileIdentificationsType{}
	return &this
}

// NewStagedProfileIdentificationsTypeWithDefaults instantiates a new StagedProfileIdentificationsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStagedProfileIdentificationsTypeWithDefaults() *StagedProfileIdentificationsType {
	this := StagedProfileIdentificationsType{}
	return &this
}

// GetIdentification returns the Identification field value if set, zero value otherwise.
func (o *StagedProfileIdentificationsType) GetIdentification() IdentificationType {
	if o == nil || IsNil(o.Identification) {
		var ret IdentificationType
		return ret
	}
	return *o.Identification
}

// GetIdentificationOk returns a tuple with the Identification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileIdentificationsType) GetIdentificationOk() (*IdentificationType, bool) {
	if o == nil || IsNil(o.Identification) {
		return nil, false
	}
	return o.Identification, true
}

// HasIdentification returns a boolean if a field has been set.
func (o *StagedProfileIdentificationsType) HasIdentification() bool {
	if o != nil && !IsNil(o.Identification) {
		return true
	}

	return false
}

// SetIdentification gets a reference to the given IdentificationType and assigns it to the Identification field.
func (o *StagedProfileIdentificationsType) SetIdentification(v IdentificationType) {
	o.Identification = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *StagedProfileIdentificationsType) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileIdentificationsType) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *StagedProfileIdentificationsType) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *StagedProfileIdentificationsType) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetDocumentResort returns the DocumentResort field value if set, zero value otherwise.
func (o *StagedProfileIdentificationsType) GetDocumentResort() string {
	if o == nil || IsNil(o.DocumentResort) {
		var ret string
		return ret
	}
	return *o.DocumentResort
}

// GetDocumentResortOk returns a tuple with the DocumentResort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileIdentificationsType) GetDocumentResortOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentResort) {
		return nil, false
	}
	return o.DocumentResort, true
}

// HasDocumentResort returns a boolean if a field has been set.
func (o *StagedProfileIdentificationsType) HasDocumentResort() bool {
	if o != nil && !IsNil(o.DocumentResort) {
		return true
	}

	return false
}

// SetDocumentResort gets a reference to the given string and assigns it to the DocumentResort field.
func (o *StagedProfileIdentificationsType) SetDocumentResort(v string) {
	o.DocumentResort = &v
}

// GetInactiveDate returns the InactiveDate field value if set, zero value otherwise.
func (o *StagedProfileIdentificationsType) GetInactiveDate() string {
	if o == nil || IsNil(o.InactiveDate) {
		var ret string
		return ret
	}
	return *o.InactiveDate
}

// GetInactiveDateOk returns a tuple with the InactiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileIdentificationsType) GetInactiveDateOk() (*string, bool) {
	if o == nil || IsNil(o.InactiveDate) {
		return nil, false
	}
	return o.InactiveDate, true
}

// HasInactiveDate returns a boolean if a field has been set.
func (o *StagedProfileIdentificationsType) HasInactiveDate() bool {
	if o != nil && !IsNil(o.InactiveDate) {
		return true
	}

	return false
}

// SetInactiveDate gets a reference to the given string and assigns it to the InactiveDate field.
func (o *StagedProfileIdentificationsType) SetInactiveDate(v string) {
	o.InactiveDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StagedProfileIdentificationsType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileIdentificationsType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StagedProfileIdentificationsType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StagedProfileIdentificationsType) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StagedProfileIdentificationsType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StagedProfileIdentificationsType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StagedProfileIdentificationsType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StagedProfileIdentificationsType) SetType(v string) {
	o.Type = &v
}

func (o StagedProfileIdentificationsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StagedProfileIdentificationsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identification) {
		toSerialize["identification"] = o.Identification
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if !IsNil(o.DocumentResort) {
		toSerialize["documentResort"] = o.DocumentResort
	}
	if !IsNil(o.InactiveDate) {
		toSerialize["inactiveDate"] = o.InactiveDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableStagedProfileIdentificationsType struct {
	value *StagedProfileIdentificationsType
	isSet bool
}

func (v NullableStagedProfileIdentificationsType) Get() *StagedProfileIdentificationsType {
	return v.value
}

func (v *NullableStagedProfileIdentificationsType) Set(val *StagedProfileIdentificationsType) {
	v.value = val
	v.isSet = true
}

func (v NullableStagedProfileIdentificationsType) IsSet() bool {
	return v.isSet
}

func (v *NullableStagedProfileIdentificationsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStagedProfileIdentificationsType(val *StagedProfileIdentificationsType) *NullableStagedProfileIdentificationsType {
	return &NullableStagedProfileIdentificationsType{value: val, isSet: true}
}

func (v NullableStagedProfileIdentificationsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStagedProfileIdentificationsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


