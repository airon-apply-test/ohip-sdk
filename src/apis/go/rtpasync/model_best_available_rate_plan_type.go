/*
Opera Cloud Rate Plan Asynchronous Service API

APIs catering to the Rate Plan asynchronous related functionality in a hotel.  This includes adding/updating daily rates&apos; pricing schedules and best available rates by day or length of stay. <p>This API follows an async pattern where</p><ul><li>You make an initial request, which returns a Location header</li><li>You poll HEAD on the Location header returned to obtain the status of the process</li><li>Once the process completes HEAD will return in the Location header the URL that must be called to obtain the results of the process</li><li>You call the URL to obtain the results of the process</li></ul><br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rtpasync

import (
	"encoding/json"
)

// checks if the BestAvailableRatePlanType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BestAvailableRatePlanType{}

// BestAvailableRatePlanType Defines best available rate plans.
type BestAvailableRatePlanType struct {
	// Collection of best available rate codes.
	RatePlanCodes []string `json:"ratePlanCodes,omitempty"`
	// Validity of best available rate plan.
	RateDate *string `json:"rateDate,omitempty"`
	// Indicates Length of Stay 1 configuration.
	Los1 *bool `json:"los1,omitempty"`
	// Indicates Length of Stay 2 configuration.
	Los2 *bool `json:"los2,omitempty"`
	// Indicates Length of Stay 3 configuration.
	Los3 *bool `json:"los3,omitempty"`
	// Indicates Length of Stay 4 configuration.
	Los4 *bool `json:"los4,omitempty"`
	// Indicates Length of Stay 5 configuration.
	Los5 *bool `json:"los5,omitempty"`
	// Indicates Length of Stay 6 configuration.
	Los6 *bool `json:"los6,omitempty"`
	// Indicates Length of Stay 7 configuration.
	Los7 *bool `json:"los7,omitempty"`
	// Indicates Length of Stay 8 configuration.
	Los8 *bool `json:"los8,omitempty"`
}

// NewBestAvailableRatePlanType instantiates a new BestAvailableRatePlanType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBestAvailableRatePlanType() *BestAvailableRatePlanType {
	this := BestAvailableRatePlanType{}
	return &this
}

// NewBestAvailableRatePlanTypeWithDefaults instantiates a new BestAvailableRatePlanType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBestAvailableRatePlanTypeWithDefaults() *BestAvailableRatePlanType {
	this := BestAvailableRatePlanType{}
	return &this
}

// GetRatePlanCodes returns the RatePlanCodes field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetRatePlanCodes() []string {
	if o == nil || IsNil(o.RatePlanCodes) {
		var ret []string
		return ret
	}
	return o.RatePlanCodes
}

// GetRatePlanCodesOk returns a tuple with the RatePlanCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetRatePlanCodesOk() ([]string, bool) {
	if o == nil || IsNil(o.RatePlanCodes) {
		return nil, false
	}
	return o.RatePlanCodes, true
}

// HasRatePlanCodes returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasRatePlanCodes() bool {
	if o != nil && !IsNil(o.RatePlanCodes) {
		return true
	}

	return false
}

// SetRatePlanCodes gets a reference to the given []string and assigns it to the RatePlanCodes field.
func (o *BestAvailableRatePlanType) SetRatePlanCodes(v []string) {
	o.RatePlanCodes = v
}

// GetRateDate returns the RateDate field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetRateDate() string {
	if o == nil || IsNil(o.RateDate) {
		var ret string
		return ret
	}
	return *o.RateDate
}

// GetRateDateOk returns a tuple with the RateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetRateDateOk() (*string, bool) {
	if o == nil || IsNil(o.RateDate) {
		return nil, false
	}
	return o.RateDate, true
}

// HasRateDate returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasRateDate() bool {
	if o != nil && !IsNil(o.RateDate) {
		return true
	}

	return false
}

// SetRateDate gets a reference to the given string and assigns it to the RateDate field.
func (o *BestAvailableRatePlanType) SetRateDate(v string) {
	o.RateDate = &v
}

// GetLos1 returns the Los1 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos1() bool {
	if o == nil || IsNil(o.Los1) {
		var ret bool
		return ret
	}
	return *o.Los1
}

// GetLos1Ok returns a tuple with the Los1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos1Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los1) {
		return nil, false
	}
	return o.Los1, true
}

// HasLos1 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos1() bool {
	if o != nil && !IsNil(o.Los1) {
		return true
	}

	return false
}

// SetLos1 gets a reference to the given bool and assigns it to the Los1 field.
func (o *BestAvailableRatePlanType) SetLos1(v bool) {
	o.Los1 = &v
}

// GetLos2 returns the Los2 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos2() bool {
	if o == nil || IsNil(o.Los2) {
		var ret bool
		return ret
	}
	return *o.Los2
}

// GetLos2Ok returns a tuple with the Los2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos2Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los2) {
		return nil, false
	}
	return o.Los2, true
}

// HasLos2 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos2() bool {
	if o != nil && !IsNil(o.Los2) {
		return true
	}

	return false
}

// SetLos2 gets a reference to the given bool and assigns it to the Los2 field.
func (o *BestAvailableRatePlanType) SetLos2(v bool) {
	o.Los2 = &v
}

// GetLos3 returns the Los3 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos3() bool {
	if o == nil || IsNil(o.Los3) {
		var ret bool
		return ret
	}
	return *o.Los3
}

// GetLos3Ok returns a tuple with the Los3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos3Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los3) {
		return nil, false
	}
	return o.Los3, true
}

// HasLos3 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos3() bool {
	if o != nil && !IsNil(o.Los3) {
		return true
	}

	return false
}

// SetLos3 gets a reference to the given bool and assigns it to the Los3 field.
func (o *BestAvailableRatePlanType) SetLos3(v bool) {
	o.Los3 = &v
}

// GetLos4 returns the Los4 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos4() bool {
	if o == nil || IsNil(o.Los4) {
		var ret bool
		return ret
	}
	return *o.Los4
}

// GetLos4Ok returns a tuple with the Los4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos4Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los4) {
		return nil, false
	}
	return o.Los4, true
}

// HasLos4 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos4() bool {
	if o != nil && !IsNil(o.Los4) {
		return true
	}

	return false
}

// SetLos4 gets a reference to the given bool and assigns it to the Los4 field.
func (o *BestAvailableRatePlanType) SetLos4(v bool) {
	o.Los4 = &v
}

// GetLos5 returns the Los5 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos5() bool {
	if o == nil || IsNil(o.Los5) {
		var ret bool
		return ret
	}
	return *o.Los5
}

// GetLos5Ok returns a tuple with the Los5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos5Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los5) {
		return nil, false
	}
	return o.Los5, true
}

// HasLos5 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos5() bool {
	if o != nil && !IsNil(o.Los5) {
		return true
	}

	return false
}

// SetLos5 gets a reference to the given bool and assigns it to the Los5 field.
func (o *BestAvailableRatePlanType) SetLos5(v bool) {
	o.Los5 = &v
}

// GetLos6 returns the Los6 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos6() bool {
	if o == nil || IsNil(o.Los6) {
		var ret bool
		return ret
	}
	return *o.Los6
}

// GetLos6Ok returns a tuple with the Los6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos6Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los6) {
		return nil, false
	}
	return o.Los6, true
}

// HasLos6 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos6() bool {
	if o != nil && !IsNil(o.Los6) {
		return true
	}

	return false
}

// SetLos6 gets a reference to the given bool and assigns it to the Los6 field.
func (o *BestAvailableRatePlanType) SetLos6(v bool) {
	o.Los6 = &v
}

// GetLos7 returns the Los7 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos7() bool {
	if o == nil || IsNil(o.Los7) {
		var ret bool
		return ret
	}
	return *o.Los7
}

// GetLos7Ok returns a tuple with the Los7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos7Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los7) {
		return nil, false
	}
	return o.Los7, true
}

// HasLos7 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos7() bool {
	if o != nil && !IsNil(o.Los7) {
		return true
	}

	return false
}

// SetLos7 gets a reference to the given bool and assigns it to the Los7 field.
func (o *BestAvailableRatePlanType) SetLos7(v bool) {
	o.Los7 = &v
}

// GetLos8 returns the Los8 field value if set, zero value otherwise.
func (o *BestAvailableRatePlanType) GetLos8() bool {
	if o == nil || IsNil(o.Los8) {
		var ret bool
		return ret
	}
	return *o.Los8
}

// GetLos8Ok returns a tuple with the Los8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BestAvailableRatePlanType) GetLos8Ok() (*bool, bool) {
	if o == nil || IsNil(o.Los8) {
		return nil, false
	}
	return o.Los8, true
}

// HasLos8 returns a boolean if a field has been set.
func (o *BestAvailableRatePlanType) HasLos8() bool {
	if o != nil && !IsNil(o.Los8) {
		return true
	}

	return false
}

// SetLos8 gets a reference to the given bool and assigns it to the Los8 field.
func (o *BestAvailableRatePlanType) SetLos8(v bool) {
	o.Los8 = &v
}

func (o BestAvailableRatePlanType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BestAvailableRatePlanType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RatePlanCodes) {
		toSerialize["ratePlanCodes"] = o.RatePlanCodes
	}
	if !IsNil(o.RateDate) {
		toSerialize["rateDate"] = o.RateDate
	}
	if !IsNil(o.Los1) {
		toSerialize["los1"] = o.Los1
	}
	if !IsNil(o.Los2) {
		toSerialize["los2"] = o.Los2
	}
	if !IsNil(o.Los3) {
		toSerialize["los3"] = o.Los3
	}
	if !IsNil(o.Los4) {
		toSerialize["los4"] = o.Los4
	}
	if !IsNil(o.Los5) {
		toSerialize["los5"] = o.Los5
	}
	if !IsNil(o.Los6) {
		toSerialize["los6"] = o.Los6
	}
	if !IsNil(o.Los7) {
		toSerialize["los7"] = o.Los7
	}
	if !IsNil(o.Los8) {
		toSerialize["los8"] = o.Los8
	}
	return toSerialize, nil
}

type NullableBestAvailableRatePlanType struct {
	value *BestAvailableRatePlanType
	isSet bool
}

func (v NullableBestAvailableRatePlanType) Get() *BestAvailableRatePlanType {
	return v.value
}

func (v *NullableBestAvailableRatePlanType) Set(val *BestAvailableRatePlanType) {
	v.value = val
	v.isSet = true
}

func (v NullableBestAvailableRatePlanType) IsSet() bool {
	return v.isSet
}

func (v *NullableBestAvailableRatePlanType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBestAvailableRatePlanType(val *BestAvailableRatePlanType) *NullableBestAvailableRatePlanType {
	return &NullableBestAvailableRatePlanType{value: val, isSet: true}
}

func (v NullableBestAvailableRatePlanType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBestAvailableRatePlanType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


