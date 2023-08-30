/*
OPERA Cloud Enterprise Configuration API

APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package entcfg

import (
	"encoding/json"
)

// checks if the ProfileTaxInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileTaxInfoType{}

// ProfileTaxInfoType Profile information related to tax.
type ProfileTaxInfoType struct {
	// The tax id of this profile. Usually issued by a government agency. Used by 1099 printing.
	Tax1No *string `json:"tax1No,omitempty"`
	// Tax 2 id of this profile.
	Tax2No *string `json:"tax2No,omitempty"`
	// Tax Category to be changed.
	TaxCategory *string `json:"taxCategory,omitempty"`
	// Tax Office to be changed.
	TaxOffice *string `json:"taxOffice,omitempty"`
	// Tax type to be changed.
	TaxType *string `json:"taxType,omitempty"`
	// Business ID. The maximum length of this element should not exceed 120 characters.
	BusinessId *string `json:"businessId,omitempty"`
	// Business Registration Code. The maximum length of this element should not exceed 120 characters.
	BusinessRegistration *string `json:"businessRegistration,omitempty"`
}

// NewProfileTaxInfoType instantiates a new ProfileTaxInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileTaxInfoType() *ProfileTaxInfoType {
	this := ProfileTaxInfoType{}
	return &this
}

// NewProfileTaxInfoTypeWithDefaults instantiates a new ProfileTaxInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileTaxInfoTypeWithDefaults() *ProfileTaxInfoType {
	this := ProfileTaxInfoType{}
	return &this
}

// GetTax1No returns the Tax1No field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetTax1No() string {
	if o == nil || IsNil(o.Tax1No) {
		var ret string
		return ret
	}
	return *o.Tax1No
}

// GetTax1NoOk returns a tuple with the Tax1No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetTax1NoOk() (*string, bool) {
	if o == nil || IsNil(o.Tax1No) {
		return nil, false
	}
	return o.Tax1No, true
}

// HasTax1No returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasTax1No() bool {
	if o != nil && !IsNil(o.Tax1No) {
		return true
	}

	return false
}

// SetTax1No gets a reference to the given string and assigns it to the Tax1No field.
func (o *ProfileTaxInfoType) SetTax1No(v string) {
	o.Tax1No = &v
}

// GetTax2No returns the Tax2No field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetTax2No() string {
	if o == nil || IsNil(o.Tax2No) {
		var ret string
		return ret
	}
	return *o.Tax2No
}

// GetTax2NoOk returns a tuple with the Tax2No field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetTax2NoOk() (*string, bool) {
	if o == nil || IsNil(o.Tax2No) {
		return nil, false
	}
	return o.Tax2No, true
}

// HasTax2No returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasTax2No() bool {
	if o != nil && !IsNil(o.Tax2No) {
		return true
	}

	return false
}

// SetTax2No gets a reference to the given string and assigns it to the Tax2No field.
func (o *ProfileTaxInfoType) SetTax2No(v string) {
	o.Tax2No = &v
}

// GetTaxCategory returns the TaxCategory field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetTaxCategory() string {
	if o == nil || IsNil(o.TaxCategory) {
		var ret string
		return ret
	}
	return *o.TaxCategory
}

// GetTaxCategoryOk returns a tuple with the TaxCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetTaxCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.TaxCategory) {
		return nil, false
	}
	return o.TaxCategory, true
}

// HasTaxCategory returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasTaxCategory() bool {
	if o != nil && !IsNil(o.TaxCategory) {
		return true
	}

	return false
}

// SetTaxCategory gets a reference to the given string and assigns it to the TaxCategory field.
func (o *ProfileTaxInfoType) SetTaxCategory(v string) {
	o.TaxCategory = &v
}

// GetTaxOffice returns the TaxOffice field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetTaxOffice() string {
	if o == nil || IsNil(o.TaxOffice) {
		var ret string
		return ret
	}
	return *o.TaxOffice
}

// GetTaxOfficeOk returns a tuple with the TaxOffice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetTaxOfficeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxOffice) {
		return nil, false
	}
	return o.TaxOffice, true
}

// HasTaxOffice returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasTaxOffice() bool {
	if o != nil && !IsNil(o.TaxOffice) {
		return true
	}

	return false
}

// SetTaxOffice gets a reference to the given string and assigns it to the TaxOffice field.
func (o *ProfileTaxInfoType) SetTaxOffice(v string) {
	o.TaxOffice = &v
}

// GetTaxType returns the TaxType field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetTaxType() string {
	if o == nil || IsNil(o.TaxType) {
		var ret string
		return ret
	}
	return *o.TaxType
}

// GetTaxTypeOk returns a tuple with the TaxType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetTaxTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TaxType) {
		return nil, false
	}
	return o.TaxType, true
}

// HasTaxType returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasTaxType() bool {
	if o != nil && !IsNil(o.TaxType) {
		return true
	}

	return false
}

// SetTaxType gets a reference to the given string and assigns it to the TaxType field.
func (o *ProfileTaxInfoType) SetTaxType(v string) {
	o.TaxType = &v
}

// GetBusinessId returns the BusinessId field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetBusinessId() string {
	if o == nil || IsNil(o.BusinessId) {
		var ret string
		return ret
	}
	return *o.BusinessId
}

// GetBusinessIdOk returns a tuple with the BusinessId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetBusinessIdOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessId) {
		return nil, false
	}
	return o.BusinessId, true
}

// HasBusinessId returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasBusinessId() bool {
	if o != nil && !IsNil(o.BusinessId) {
		return true
	}

	return false
}

// SetBusinessId gets a reference to the given string and assigns it to the BusinessId field.
func (o *ProfileTaxInfoType) SetBusinessId(v string) {
	o.BusinessId = &v
}

// GetBusinessRegistration returns the BusinessRegistration field value if set, zero value otherwise.
func (o *ProfileTaxInfoType) GetBusinessRegistration() string {
	if o == nil || IsNil(o.BusinessRegistration) {
		var ret string
		return ret
	}
	return *o.BusinessRegistration
}

// GetBusinessRegistrationOk returns a tuple with the BusinessRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProfileTaxInfoType) GetBusinessRegistrationOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessRegistration) {
		return nil, false
	}
	return o.BusinessRegistration, true
}

// HasBusinessRegistration returns a boolean if a field has been set.
func (o *ProfileTaxInfoType) HasBusinessRegistration() bool {
	if o != nil && !IsNil(o.BusinessRegistration) {
		return true
	}

	return false
}

// SetBusinessRegistration gets a reference to the given string and assigns it to the BusinessRegistration field.
func (o *ProfileTaxInfoType) SetBusinessRegistration(v string) {
	o.BusinessRegistration = &v
}

func (o ProfileTaxInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileTaxInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tax1No) {
		toSerialize["tax1No"] = o.Tax1No
	}
	if !IsNil(o.Tax2No) {
		toSerialize["tax2No"] = o.Tax2No
	}
	if !IsNil(o.TaxCategory) {
		toSerialize["taxCategory"] = o.TaxCategory
	}
	if !IsNil(o.TaxOffice) {
		toSerialize["taxOffice"] = o.TaxOffice
	}
	if !IsNil(o.TaxType) {
		toSerialize["taxType"] = o.TaxType
	}
	if !IsNil(o.BusinessId) {
		toSerialize["businessId"] = o.BusinessId
	}
	if !IsNil(o.BusinessRegistration) {
		toSerialize["businessRegistration"] = o.BusinessRegistration
	}
	return toSerialize, nil
}

type NullableProfileTaxInfoType struct {
	value *ProfileTaxInfoType
	isSet bool
}

func (v NullableProfileTaxInfoType) Get() *ProfileTaxInfoType {
	return v.value
}

func (v *NullableProfileTaxInfoType) Set(val *ProfileTaxInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileTaxInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileTaxInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileTaxInfoType(val *ProfileTaxInfoType) *NullableProfileTaxInfoType {
	return &NullableProfileTaxInfoType{value: val, isSet: true}
}

func (v NullableProfileTaxInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileTaxInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


