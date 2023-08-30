/*
OPERA Cloud Customer Relationship Management Outbound API

APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crmoutbound

import (
	"encoding/json"
)

// checks if the OwnerType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OwnerType{}

// OwnerType Generic type for information about an owner.
type OwnerType struct {
	Hotel *CodeDescriptionType `json:"hotel,omitempty"`
	UserId *UniqueIDType `json:"userId,omitempty"`
	// Unique application user name of the owner.
	UserName *string `json:"userName,omitempty"`
	// Unique Code to identify the owner.
	OwnerCode *string `json:"ownerCode,omitempty"`
	ProfileId *ProfileId `json:"profileId,omitempty"`
	Name *PersonNameType `json:"name,omitempty"`
	Department *CodeDescriptionType `json:"department,omitempty"`
	Email *EmailInfoType `json:"email,omitempty"`
	Phone *TelephoneInfoType `json:"phone,omitempty"`
	Relationship *CodeDescriptionType `json:"relationship,omitempty"`
	// When true, this is a primary owner.
	Primary *bool `json:"primary,omitempty"`
}

// NewOwnerType instantiates a new OwnerType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOwnerType() *OwnerType {
	this := OwnerType{}
	return &this
}

// NewOwnerTypeWithDefaults instantiates a new OwnerType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOwnerTypeWithDefaults() *OwnerType {
	this := OwnerType{}
	return &this
}

// GetHotel returns the Hotel field value if set, zero value otherwise.
func (o *OwnerType) GetHotel() CodeDescriptionType {
	if o == nil || IsNil(o.Hotel) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Hotel
}

// GetHotelOk returns a tuple with the Hotel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetHotelOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Hotel) {
		return nil, false
	}
	return o.Hotel, true
}

// HasHotel returns a boolean if a field has been set.
func (o *OwnerType) HasHotel() bool {
	if o != nil && !IsNil(o.Hotel) {
		return true
	}

	return false
}

// SetHotel gets a reference to the given CodeDescriptionType and assigns it to the Hotel field.
func (o *OwnerType) SetHotel(v CodeDescriptionType) {
	o.Hotel = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *OwnerType) GetUserId() UniqueIDType {
	if o == nil || IsNil(o.UserId) {
		var ret UniqueIDType
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetUserIdOk() (*UniqueIDType, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *OwnerType) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given UniqueIDType and assigns it to the UserId field.
func (o *OwnerType) SetUserId(v UniqueIDType) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *OwnerType) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *OwnerType) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *OwnerType) SetUserName(v string) {
	o.UserName = &v
}

// GetOwnerCode returns the OwnerCode field value if set, zero value otherwise.
func (o *OwnerType) GetOwnerCode() string {
	if o == nil || IsNil(o.OwnerCode) {
		var ret string
		return ret
	}
	return *o.OwnerCode
}

// GetOwnerCodeOk returns a tuple with the OwnerCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetOwnerCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerCode) {
		return nil, false
	}
	return o.OwnerCode, true
}

// HasOwnerCode returns a boolean if a field has been set.
func (o *OwnerType) HasOwnerCode() bool {
	if o != nil && !IsNil(o.OwnerCode) {
		return true
	}

	return false
}

// SetOwnerCode gets a reference to the given string and assigns it to the OwnerCode field.
func (o *OwnerType) SetOwnerCode(v string) {
	o.OwnerCode = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *OwnerType) GetProfileId() ProfileId {
	if o == nil || IsNil(o.ProfileId) {
		var ret ProfileId
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetProfileIdOk() (*ProfileId, bool) {
	if o == nil || IsNil(o.ProfileId) {
		return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *OwnerType) HasProfileId() bool {
	if o != nil && !IsNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given ProfileId and assigns it to the ProfileId field.
func (o *OwnerType) SetProfileId(v ProfileId) {
	o.ProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OwnerType) GetName() PersonNameType {
	if o == nil || IsNil(o.Name) {
		var ret PersonNameType
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetNameOk() (*PersonNameType, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OwnerType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given PersonNameType and assigns it to the Name field.
func (o *OwnerType) SetName(v PersonNameType) {
	o.Name = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *OwnerType) GetDepartment() CodeDescriptionType {
	if o == nil || IsNil(o.Department) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetDepartmentOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Department) {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *OwnerType) HasDepartment() bool {
	if o != nil && !IsNil(o.Department) {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given CodeDescriptionType and assigns it to the Department field.
func (o *OwnerType) SetDepartment(v CodeDescriptionType) {
	o.Department = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *OwnerType) GetEmail() EmailInfoType {
	if o == nil || IsNil(o.Email) {
		var ret EmailInfoType
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetEmailOk() (*EmailInfoType, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *OwnerType) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given EmailInfoType and assigns it to the Email field.
func (o *OwnerType) SetEmail(v EmailInfoType) {
	o.Email = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *OwnerType) GetPhone() TelephoneInfoType {
	if o == nil || IsNil(o.Phone) {
		var ret TelephoneInfoType
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetPhoneOk() (*TelephoneInfoType, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *OwnerType) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given TelephoneInfoType and assigns it to the Phone field.
func (o *OwnerType) SetPhone(v TelephoneInfoType) {
	o.Phone = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *OwnerType) GetRelationship() CodeDescriptionType {
	if o == nil || IsNil(o.Relationship) {
		var ret CodeDescriptionType
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetRelationshipOk() (*CodeDescriptionType, bool) {
	if o == nil || IsNil(o.Relationship) {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *OwnerType) HasRelationship() bool {
	if o != nil && !IsNil(o.Relationship) {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given CodeDescriptionType and assigns it to the Relationship field.
func (o *OwnerType) SetRelationship(v CodeDescriptionType) {
	o.Relationship = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *OwnerType) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OwnerType) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *OwnerType) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *OwnerType) SetPrimary(v bool) {
	o.Primary = &v
}

func (o OwnerType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OwnerType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hotel) {
		toSerialize["hotel"] = o.Hotel
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.OwnerCode) {
		toSerialize["ownerCode"] = o.OwnerCode
	}
	if !IsNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Department) {
		toSerialize["department"] = o.Department
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Relationship) {
		toSerialize["relationship"] = o.Relationship
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	return toSerialize, nil
}

type NullableOwnerType struct {
	value *OwnerType
	isSet bool
}

func (v NullableOwnerType) Get() *OwnerType {
	return v.value
}

func (v *NullableOwnerType) Set(val *OwnerType) {
	v.value = val
	v.isSet = true
}

func (v NullableOwnerType) IsSet() bool {
	return v.isSet
}

func (v *NullableOwnerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOwnerType(val *OwnerType) *NullableOwnerType {
	return &NullableOwnerType{value: val, isSet: true}
}

func (v NullableOwnerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOwnerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


