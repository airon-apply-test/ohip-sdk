/*
OPERA Cloud Block API

APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NameValueOriginType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameValueOriginType{}

// NameValueOriginType Contains origin details.
type NameValueOriginType struct {
	OriginName *NameValueModuleType `json:"originName,omitempty"`
	// Contains destination column for Origin.
	Destination *string `json:"destination,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
	// Used to identify the source of the identifier (e.g., IATA, ABTA).
	IdContext *string `json:"idContext,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
}

// NewNameValueOriginType instantiates a new NameValueOriginType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameValueOriginType() *NameValueOriginType {
	this := NameValueOriginType{}
	return &this
}

// NewNameValueOriginTypeWithDefaults instantiates a new NameValueOriginType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameValueOriginTypeWithDefaults() *NameValueOriginType {
	this := NameValueOriginType{}
	return &this
}

// GetOriginName returns the OriginName field value if set, zero value otherwise.
func (o *NameValueOriginType) GetOriginName() NameValueModuleType {
	if o == nil || IsNil(o.OriginName) {
		var ret NameValueModuleType
		return ret
	}
	return *o.OriginName
}

// GetOriginNameOk returns a tuple with the OriginName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueOriginType) GetOriginNameOk() (*NameValueModuleType, bool) {
	if o == nil || IsNil(o.OriginName) {
		return nil, false
	}
	return o.OriginName, true
}

// HasOriginName returns a boolean if a field has been set.
func (o *NameValueOriginType) HasOriginName() bool {
	if o != nil && !IsNil(o.OriginName) {
		return true
	}

	return false
}

// SetOriginName gets a reference to the given NameValueModuleType and assigns it to the OriginName field.
func (o *NameValueOriginType) SetOriginName(v NameValueModuleType) {
	o.OriginName = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *NameValueOriginType) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueOriginType) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *NameValueOriginType) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *NameValueOriginType) SetDestination(v string) {
	o.Destination = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NameValueOriginType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueOriginType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NameValueOriginType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NameValueOriginType) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NameValueOriginType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueOriginType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NameValueOriginType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NameValueOriginType) SetType(v string) {
	o.Type = &v
}

// GetIdContext returns the IdContext field value if set, zero value otherwise.
func (o *NameValueOriginType) GetIdContext() string {
	if o == nil || IsNil(o.IdContext) {
		var ret string
		return ret
	}
	return *o.IdContext
}

// GetIdContextOk returns a tuple with the IdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueOriginType) GetIdContextOk() (*string, bool) {
	if o == nil || IsNil(o.IdContext) {
		return nil, false
	}
	return o.IdContext, true
}

// HasIdContext returns a boolean if a field has been set.
func (o *NameValueOriginType) HasIdContext() bool {
	if o != nil && !IsNil(o.IdContext) {
		return true
	}

	return false
}

// SetIdContext gets a reference to the given string and assigns it to the IdContext field.
func (o *NameValueOriginType) SetIdContext(v string) {
	o.IdContext = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *NameValueOriginType) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValueOriginType) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *NameValueOriginType) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *NameValueOriginType) SetWarnings(v []WarningType) {
	o.Warnings = v
}

func (o NameValueOriginType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameValueOriginType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginName) {
		toSerialize["originName"] = o.OriginName
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.IdContext) {
		toSerialize["idContext"] = o.IdContext
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableNameValueOriginType struct {
	value *NameValueOriginType
	isSet bool
}

func (v NullableNameValueOriginType) Get() *NameValueOriginType {
	return v.value
}

func (v *NullableNameValueOriginType) Set(val *NameValueOriginType) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValueOriginType) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValueOriginType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValueOriginType(val *NameValueOriginType) *NullableNameValueOriginType {
	return &NullableNameValueOriginType{value: val, isSet: true}
}

func (v NullableNameValueOriginType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValueOriginType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


