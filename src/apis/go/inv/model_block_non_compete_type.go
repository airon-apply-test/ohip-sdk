/*
OPERA Cloud Inventory API

APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package inv

import (
	"encoding/json"
)

// checks if the BlockNonCompeteType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockNonCompeteType{}

// BlockNonCompeteType Block Non Compete code information.
type BlockNonCompeteType struct {
	// Specifies the Non-Compete Industry.>
	Industry *string `json:"industry,omitempty"`
	// Specifies the Non-Compete Industry description.>
	IndustryDescription *string `json:"industryDescription,omitempty"`
	Criteria *RateProtectionType `json:"criteria,omitempty"`
	// Specifies a single date.
	ProtectedDates []string `json:"protectedDates,omitempty"`
}

// NewBlockNonCompeteType instantiates a new BlockNonCompeteType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockNonCompeteType() *BlockNonCompeteType {
	this := BlockNonCompeteType{}
	return &this
}

// NewBlockNonCompeteTypeWithDefaults instantiates a new BlockNonCompeteType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockNonCompeteTypeWithDefaults() *BlockNonCompeteType {
	this := BlockNonCompeteType{}
	return &this
}

// GetIndustry returns the Industry field value if set, zero value otherwise.
func (o *BlockNonCompeteType) GetIndustry() string {
	if o == nil || IsNil(o.Industry) {
		var ret string
		return ret
	}
	return *o.Industry
}

// GetIndustryOk returns a tuple with the Industry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockNonCompeteType) GetIndustryOk() (*string, bool) {
	if o == nil || IsNil(o.Industry) {
		return nil, false
	}
	return o.Industry, true
}

// HasIndustry returns a boolean if a field has been set.
func (o *BlockNonCompeteType) HasIndustry() bool {
	if o != nil && !IsNil(o.Industry) {
		return true
	}

	return false
}

// SetIndustry gets a reference to the given string and assigns it to the Industry field.
func (o *BlockNonCompeteType) SetIndustry(v string) {
	o.Industry = &v
}

// GetIndustryDescription returns the IndustryDescription field value if set, zero value otherwise.
func (o *BlockNonCompeteType) GetIndustryDescription() string {
	if o == nil || IsNil(o.IndustryDescription) {
		var ret string
		return ret
	}
	return *o.IndustryDescription
}

// GetIndustryDescriptionOk returns a tuple with the IndustryDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockNonCompeteType) GetIndustryDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.IndustryDescription) {
		return nil, false
	}
	return o.IndustryDescription, true
}

// HasIndustryDescription returns a boolean if a field has been set.
func (o *BlockNonCompeteType) HasIndustryDescription() bool {
	if o != nil && !IsNil(o.IndustryDescription) {
		return true
	}

	return false
}

// SetIndustryDescription gets a reference to the given string and assigns it to the IndustryDescription field.
func (o *BlockNonCompeteType) SetIndustryDescription(v string) {
	o.IndustryDescription = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *BlockNonCompeteType) GetCriteria() RateProtectionType {
	if o == nil || IsNil(o.Criteria) {
		var ret RateProtectionType
		return ret
	}
	return *o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockNonCompeteType) GetCriteriaOk() (*RateProtectionType, bool) {
	if o == nil || IsNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *BlockNonCompeteType) HasCriteria() bool {
	if o != nil && !IsNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given RateProtectionType and assigns it to the Criteria field.
func (o *BlockNonCompeteType) SetCriteria(v RateProtectionType) {
	o.Criteria = &v
}

// GetProtectedDates returns the ProtectedDates field value if set, zero value otherwise.
func (o *BlockNonCompeteType) GetProtectedDates() []string {
	if o == nil || IsNil(o.ProtectedDates) {
		var ret []string
		return ret
	}
	return o.ProtectedDates
}

// GetProtectedDatesOk returns a tuple with the ProtectedDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockNonCompeteType) GetProtectedDatesOk() ([]string, bool) {
	if o == nil || IsNil(o.ProtectedDates) {
		return nil, false
	}
	return o.ProtectedDates, true
}

// HasProtectedDates returns a boolean if a field has been set.
func (o *BlockNonCompeteType) HasProtectedDates() bool {
	if o != nil && !IsNil(o.ProtectedDates) {
		return true
	}

	return false
}

// SetProtectedDates gets a reference to the given []string and assigns it to the ProtectedDates field.
func (o *BlockNonCompeteType) SetProtectedDates(v []string) {
	o.ProtectedDates = v
}

func (o BlockNonCompeteType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockNonCompeteType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Industry) {
		toSerialize["industry"] = o.Industry
	}
	if !IsNil(o.IndustryDescription) {
		toSerialize["industryDescription"] = o.IndustryDescription
	}
	if !IsNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}
	if !IsNil(o.ProtectedDates) {
		toSerialize["protectedDates"] = o.ProtectedDates
	}
	return toSerialize, nil
}

type NullableBlockNonCompeteType struct {
	value *BlockNonCompeteType
	isSet bool
}

func (v NullableBlockNonCompeteType) Get() *BlockNonCompeteType {
	return v.value
}

func (v *NullableBlockNonCompeteType) Set(val *BlockNonCompeteType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockNonCompeteType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockNonCompeteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockNonCompeteType(val *BlockNonCompeteType) *NullableBlockNonCompeteType {
	return &NullableBlockNonCompeteType{value: val, isSet: true}
}

func (v NullableBlockNonCompeteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockNonCompeteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


