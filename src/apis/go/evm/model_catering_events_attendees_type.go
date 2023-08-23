/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CateringEventsAttendeesType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CateringEventsAttendeesType{}

// CateringEventsAttendeesType All different attendees that a catering event can have.
type CateringEventsAttendeesType struct {
	// Number of expected attendees for the catering event.
	Expected *int32 `json:"expected,omitempty"`
	// Number of guaranteed attendees for the catering event.
	Guaranteed *int32 `json:"guaranteed,omitempty"`
	// Number of actual attendees attending to the catering event.
	Actual *int32 `json:"actual,omitempty"`
	// Setup for number of event attendees for resources (i.e. water, pads, tables, etc.).
	Set *int32 `json:"set,omitempty"`
	// This defines the count of billed attendees with respect to the menu.
	Billed *int32 `json:"billed,omitempty"`
	// Number of package expected attendees for the catering event.
	PackageExpected *int32 `json:"packageExpected,omitempty"`
	// Number of package guaranteed attendees for the catering event.
	PackageGuaranteed *int32 `json:"packageGuaranteed,omitempty"`
	// Number of actual package attendees attending the catering event.
	PackageActual *int32 `json:"packageActual,omitempty"`
	// This defines the count of package billed attendees with respect to the menu.
	PackageBilled *int32 `json:"packageBilled,omitempty"`
}

// NewCateringEventsAttendeesType instantiates a new CateringEventsAttendeesType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCateringEventsAttendeesType() *CateringEventsAttendeesType {
	this := CateringEventsAttendeesType{}
	return &this
}

// NewCateringEventsAttendeesTypeWithDefaults instantiates a new CateringEventsAttendeesType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCateringEventsAttendeesTypeWithDefaults() *CateringEventsAttendeesType {
	this := CateringEventsAttendeesType{}
	return &this
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetExpected() int32 {
	if o == nil || IsNil(o.Expected) {
		var ret int32
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetExpectedOk() (*int32, bool) {
	if o == nil || IsNil(o.Expected) {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasExpected() bool {
	if o != nil && !IsNil(o.Expected) {
		return true
	}

	return false
}

// SetExpected gets a reference to the given int32 and assigns it to the Expected field.
func (o *CateringEventsAttendeesType) SetExpected(v int32) {
	o.Expected = &v
}

// GetGuaranteed returns the Guaranteed field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetGuaranteed() int32 {
	if o == nil || IsNil(o.Guaranteed) {
		var ret int32
		return ret
	}
	return *o.Guaranteed
}

// GetGuaranteedOk returns a tuple with the Guaranteed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetGuaranteedOk() (*int32, bool) {
	if o == nil || IsNil(o.Guaranteed) {
		return nil, false
	}
	return o.Guaranteed, true
}

// HasGuaranteed returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasGuaranteed() bool {
	if o != nil && !IsNil(o.Guaranteed) {
		return true
	}

	return false
}

// SetGuaranteed gets a reference to the given int32 and assigns it to the Guaranteed field.
func (o *CateringEventsAttendeesType) SetGuaranteed(v int32) {
	o.Guaranteed = &v
}

// GetActual returns the Actual field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetActual() int32 {
	if o == nil || IsNil(o.Actual) {
		var ret int32
		return ret
	}
	return *o.Actual
}

// GetActualOk returns a tuple with the Actual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetActualOk() (*int32, bool) {
	if o == nil || IsNil(o.Actual) {
		return nil, false
	}
	return o.Actual, true
}

// HasActual returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasActual() bool {
	if o != nil && !IsNil(o.Actual) {
		return true
	}

	return false
}

// SetActual gets a reference to the given int32 and assigns it to the Actual field.
func (o *CateringEventsAttendeesType) SetActual(v int32) {
	o.Actual = &v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetSet() int32 {
	if o == nil || IsNil(o.Set) {
		var ret int32
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetSetOk() (*int32, bool) {
	if o == nil || IsNil(o.Set) {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasSet() bool {
	if o != nil && !IsNil(o.Set) {
		return true
	}

	return false
}

// SetSet gets a reference to the given int32 and assigns it to the Set field.
func (o *CateringEventsAttendeesType) SetSet(v int32) {
	o.Set = &v
}

// GetBilled returns the Billed field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetBilled() int32 {
	if o == nil || IsNil(o.Billed) {
		var ret int32
		return ret
	}
	return *o.Billed
}

// GetBilledOk returns a tuple with the Billed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetBilledOk() (*int32, bool) {
	if o == nil || IsNil(o.Billed) {
		return nil, false
	}
	return o.Billed, true
}

// HasBilled returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasBilled() bool {
	if o != nil && !IsNil(o.Billed) {
		return true
	}

	return false
}

// SetBilled gets a reference to the given int32 and assigns it to the Billed field.
func (o *CateringEventsAttendeesType) SetBilled(v int32) {
	o.Billed = &v
}

// GetPackageExpected returns the PackageExpected field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetPackageExpected() int32 {
	if o == nil || IsNil(o.PackageExpected) {
		var ret int32
		return ret
	}
	return *o.PackageExpected
}

// GetPackageExpectedOk returns a tuple with the PackageExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetPackageExpectedOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageExpected) {
		return nil, false
	}
	return o.PackageExpected, true
}

// HasPackageExpected returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasPackageExpected() bool {
	if o != nil && !IsNil(o.PackageExpected) {
		return true
	}

	return false
}

// SetPackageExpected gets a reference to the given int32 and assigns it to the PackageExpected field.
func (o *CateringEventsAttendeesType) SetPackageExpected(v int32) {
	o.PackageExpected = &v
}

// GetPackageGuaranteed returns the PackageGuaranteed field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetPackageGuaranteed() int32 {
	if o == nil || IsNil(o.PackageGuaranteed) {
		var ret int32
		return ret
	}
	return *o.PackageGuaranteed
}

// GetPackageGuaranteedOk returns a tuple with the PackageGuaranteed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetPackageGuaranteedOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageGuaranteed) {
		return nil, false
	}
	return o.PackageGuaranteed, true
}

// HasPackageGuaranteed returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasPackageGuaranteed() bool {
	if o != nil && !IsNil(o.PackageGuaranteed) {
		return true
	}

	return false
}

// SetPackageGuaranteed gets a reference to the given int32 and assigns it to the PackageGuaranteed field.
func (o *CateringEventsAttendeesType) SetPackageGuaranteed(v int32) {
	o.PackageGuaranteed = &v
}

// GetPackageActual returns the PackageActual field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetPackageActual() int32 {
	if o == nil || IsNil(o.PackageActual) {
		var ret int32
		return ret
	}
	return *o.PackageActual
}

// GetPackageActualOk returns a tuple with the PackageActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetPackageActualOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageActual) {
		return nil, false
	}
	return o.PackageActual, true
}

// HasPackageActual returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasPackageActual() bool {
	if o != nil && !IsNil(o.PackageActual) {
		return true
	}

	return false
}

// SetPackageActual gets a reference to the given int32 and assigns it to the PackageActual field.
func (o *CateringEventsAttendeesType) SetPackageActual(v int32) {
	o.PackageActual = &v
}

// GetPackageBilled returns the PackageBilled field value if set, zero value otherwise.
func (o *CateringEventsAttendeesType) GetPackageBilled() int32 {
	if o == nil || IsNil(o.PackageBilled) {
		var ret int32
		return ret
	}
	return *o.PackageBilled
}

// GetPackageBilledOk returns a tuple with the PackageBilled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CateringEventsAttendeesType) GetPackageBilledOk() (*int32, bool) {
	if o == nil || IsNil(o.PackageBilled) {
		return nil, false
	}
	return o.PackageBilled, true
}

// HasPackageBilled returns a boolean if a field has been set.
func (o *CateringEventsAttendeesType) HasPackageBilled() bool {
	if o != nil && !IsNil(o.PackageBilled) {
		return true
	}

	return false
}

// SetPackageBilled gets a reference to the given int32 and assigns it to the PackageBilled field.
func (o *CateringEventsAttendeesType) SetPackageBilled(v int32) {
	o.PackageBilled = &v
}

func (o CateringEventsAttendeesType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CateringEventsAttendeesType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expected) {
		toSerialize["expected"] = o.Expected
	}
	if !IsNil(o.Guaranteed) {
		toSerialize["guaranteed"] = o.Guaranteed
	}
	if !IsNil(o.Actual) {
		toSerialize["actual"] = o.Actual
	}
	if !IsNil(o.Set) {
		toSerialize["set"] = o.Set
	}
	if !IsNil(o.Billed) {
		toSerialize["billed"] = o.Billed
	}
	if !IsNil(o.PackageExpected) {
		toSerialize["packageExpected"] = o.PackageExpected
	}
	if !IsNil(o.PackageGuaranteed) {
		toSerialize["packageGuaranteed"] = o.PackageGuaranteed
	}
	if !IsNil(o.PackageActual) {
		toSerialize["packageActual"] = o.PackageActual
	}
	if !IsNil(o.PackageBilled) {
		toSerialize["packageBilled"] = o.PackageBilled
	}
	return toSerialize, nil
}

type NullableCateringEventsAttendeesType struct {
	value *CateringEventsAttendeesType
	isSet bool
}

func (v NullableCateringEventsAttendeesType) Get() *CateringEventsAttendeesType {
	return v.value
}

func (v *NullableCateringEventsAttendeesType) Set(val *CateringEventsAttendeesType) {
	v.value = val
	v.isSet = true
}

func (v NullableCateringEventsAttendeesType) IsSet() bool {
	return v.isSet
}

func (v *NullableCateringEventsAttendeesType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCateringEventsAttendeesType(val *CateringEventsAttendeesType) *NullableCateringEventsAttendeesType {
	return &NullableCateringEventsAttendeesType{value: val, isSet: true}
}

func (v NullableCateringEventsAttendeesType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCateringEventsAttendeesType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


