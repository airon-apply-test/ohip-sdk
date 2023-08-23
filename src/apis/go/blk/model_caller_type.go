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

// checks if the CallerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallerType{}

// CallerType Holds caller information.
type CallerType struct {
	// Caller First Name.
	GivenName *string `json:"givenName,omitempty"`
	// Caller Last Name.
	Surname *string `json:"surname,omitempty"`
	// Caller Phone Number.
	Phone *string `json:"phone,omitempty"`
	// Caller Email.
	Email *string `json:"email,omitempty"`
	// Caller Fax Number.
	Fax *string `json:"fax,omitempty"`
	// URL that identifies the location associated with the record identified by the UniqueID.
	Url *string `json:"url,omitempty"`
	// A reference to the type of object defined by the UniqueID element.
	Type *string `json:"type,omitempty"`
	// The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
	Instance *string `json:"instance,omitempty"`
	// Used to identify the source of the identifier (e.g., IATA, ABTA).
	IdContext *string `json:"idContext,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// Additional identifying value assigned by the creating system.
	IdExtension *int32 `json:"idExtension,omitempty"`
}

// NewCallerType instantiates a new CallerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallerType() *CallerType {
	this := CallerType{}
	return &this
}

// NewCallerTypeWithDefaults instantiates a new CallerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallerTypeWithDefaults() *CallerType {
	this := CallerType{}
	return &this
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *CallerType) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *CallerType) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *CallerType) SetGivenName(v string) {
	o.GivenName = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *CallerType) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *CallerType) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *CallerType) SetSurname(v string) {
	o.Surname = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *CallerType) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *CallerType) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *CallerType) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CallerType) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CallerType) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CallerType) SetEmail(v string) {
	o.Email = &v
}

// GetFax returns the Fax field value if set, zero value otherwise.
func (o *CallerType) GetFax() string {
	if o == nil || IsNil(o.Fax) {
		var ret string
		return ret
	}
	return *o.Fax
}

// GetFaxOk returns a tuple with the Fax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetFaxOk() (*string, bool) {
	if o == nil || IsNil(o.Fax) {
		return nil, false
	}
	return o.Fax, true
}

// HasFax returns a boolean if a field has been set.
func (o *CallerType) HasFax() bool {
	if o != nil && !IsNil(o.Fax) {
		return true
	}

	return false
}

// SetFax gets a reference to the given string and assigns it to the Fax field.
func (o *CallerType) SetFax(v string) {
	o.Fax = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CallerType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CallerType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CallerType) SetUrl(v string) {
	o.Url = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CallerType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CallerType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CallerType) SetType(v string) {
	o.Type = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *CallerType) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *CallerType) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *CallerType) SetInstance(v string) {
	o.Instance = &v
}

// GetIdContext returns the IdContext field value if set, zero value otherwise.
func (o *CallerType) GetIdContext() string {
	if o == nil || IsNil(o.IdContext) {
		var ret string
		return ret
	}
	return *o.IdContext
}

// GetIdContextOk returns a tuple with the IdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetIdContextOk() (*string, bool) {
	if o == nil || IsNil(o.IdContext) {
		return nil, false
	}
	return o.IdContext, true
}

// HasIdContext returns a boolean if a field has been set.
func (o *CallerType) HasIdContext() bool {
	if o != nil && !IsNil(o.IdContext) {
		return true
	}

	return false
}

// SetIdContext gets a reference to the given string and assigns it to the IdContext field.
func (o *CallerType) SetIdContext(v string) {
	o.IdContext = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CallerType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CallerType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CallerType) SetId(v string) {
	o.Id = &v
}

// GetIdExtension returns the IdExtension field value if set, zero value otherwise.
func (o *CallerType) GetIdExtension() int32 {
	if o == nil || IsNil(o.IdExtension) {
		var ret int32
		return ret
	}
	return *o.IdExtension
}

// GetIdExtensionOk returns a tuple with the IdExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallerType) GetIdExtensionOk() (*int32, bool) {
	if o == nil || IsNil(o.IdExtension) {
		return nil, false
	}
	return o.IdExtension, true
}

// HasIdExtension returns a boolean if a field has been set.
func (o *CallerType) HasIdExtension() bool {
	if o != nil && !IsNil(o.IdExtension) {
		return true
	}

	return false
}

// SetIdExtension gets a reference to the given int32 and assigns it to the IdExtension field.
func (o *CallerType) SetIdExtension(v int32) {
	o.IdExtension = &v
}

func (o CallerType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Fax) {
		toSerialize["fax"] = o.Fax
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.IdContext) {
		toSerialize["idContext"] = o.IdContext
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdExtension) {
		toSerialize["idExtension"] = o.IdExtension
	}
	return toSerialize, nil
}

type NullableCallerType struct {
	value *CallerType
	isSet bool
}

func (v NullableCallerType) Get() *CallerType {
	return v.value
}

func (v *NullableCallerType) Set(val *CallerType) {
	v.value = val
	v.isSet = true
}

func (v NullableCallerType) IsSet() bool {
	return v.isSet
}

func (v *NullableCallerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallerType(val *CallerType) *NullableCallerType {
	return &NullableCallerType{value: val, isSet: true}
}

func (v NullableCallerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


