/*
OPERA Cloud Channel Configuration API

APIs to cater for Channel Management functionality in OPERA Cloud. <br /><br /> Channel Management allows a property to configure and administer channels such as OTAs, and web channels, covering functionality such as channel configuration, availability, inventory and restrictions.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chl

import (
	"encoding/json"
)

// checks if the ChannelParameterType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelParameterType{}

// ChannelParameterType Parameters details for the Channel.
type ChannelParameterType struct {
	// Name of the parameter in the channel.
	ParameterName *string `json:"parameterName,omitempty"`
	// The display name for the parameter in the channel.
	ParameterDisplay *string `json:"parameterDisplay,omitempty"`
	// Short description for the parameter.
	ShortDescription *string `json:"shortDescription,omitempty"`
	// The Description for the parameter.
	ParameterDescription *string `json:"parameterDescription,omitempty"`
	// The Value of the parameter.
	ParameterValue *string `json:"parameterValue,omitempty"`
	// Hotel code identifying the parameters related to the Hotel.
	HotelId *string `json:"hotelId,omitempty"`
	// The type of the parameter, the possible types are P (Parameter), Setting (S).
	ParameterType *string `json:"parameterType,omitempty"`
	// Whether this is a global or property level parameter.
	Scope *string `json:"scope,omitempty"`
	ValueType *ChannelParameterValueType `json:"valueType,omitempty"`
	// Sequence number for displaying the parameter in the channel.
	Sequence *int32 `json:"sequence,omitempty"`
	Parameters []BaseChannelParameterType `json:"parameters,omitempty"`
	LevelType *ChannelParameterLevelType `json:"levelType,omitempty"`
	LevelCode *string `json:"levelCode,omitempty"`
}

// NewChannelParameterType instantiates a new ChannelParameterType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelParameterType() *ChannelParameterType {
	this := ChannelParameterType{}
	return &this
}

// NewChannelParameterTypeWithDefaults instantiates a new ChannelParameterType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelParameterTypeWithDefaults() *ChannelParameterType {
	this := ChannelParameterType{}
	return &this
}

// GetParameterName returns the ParameterName field value if set, zero value otherwise.
func (o *ChannelParameterType) GetParameterName() string {
	if o == nil || IsNil(o.ParameterName) {
		var ret string
		return ret
	}
	return *o.ParameterName
}

// GetParameterNameOk returns a tuple with the ParameterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetParameterNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterName) {
		return nil, false
	}
	return o.ParameterName, true
}

// HasParameterName returns a boolean if a field has been set.
func (o *ChannelParameterType) HasParameterName() bool {
	if o != nil && !IsNil(o.ParameterName) {
		return true
	}

	return false
}

// SetParameterName gets a reference to the given string and assigns it to the ParameterName field.
func (o *ChannelParameterType) SetParameterName(v string) {
	o.ParameterName = &v
}

// GetParameterDisplay returns the ParameterDisplay field value if set, zero value otherwise.
func (o *ChannelParameterType) GetParameterDisplay() string {
	if o == nil || IsNil(o.ParameterDisplay) {
		var ret string
		return ret
	}
	return *o.ParameterDisplay
}

// GetParameterDisplayOk returns a tuple with the ParameterDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetParameterDisplayOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterDisplay) {
		return nil, false
	}
	return o.ParameterDisplay, true
}

// HasParameterDisplay returns a boolean if a field has been set.
func (o *ChannelParameterType) HasParameterDisplay() bool {
	if o != nil && !IsNil(o.ParameterDisplay) {
		return true
	}

	return false
}

// SetParameterDisplay gets a reference to the given string and assigns it to the ParameterDisplay field.
func (o *ChannelParameterType) SetParameterDisplay(v string) {
	o.ParameterDisplay = &v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise.
func (o *ChannelParameterType) GetShortDescription() string {
	if o == nil || IsNil(o.ShortDescription) {
		var ret string
		return ret
	}
	return *o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetShortDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ShortDescription) {
		return nil, false
	}
	return o.ShortDescription, true
}

// HasShortDescription returns a boolean if a field has been set.
func (o *ChannelParameterType) HasShortDescription() bool {
	if o != nil && !IsNil(o.ShortDescription) {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given string and assigns it to the ShortDescription field.
func (o *ChannelParameterType) SetShortDescription(v string) {
	o.ShortDescription = &v
}

// GetParameterDescription returns the ParameterDescription field value if set, zero value otherwise.
func (o *ChannelParameterType) GetParameterDescription() string {
	if o == nil || IsNil(o.ParameterDescription) {
		var ret string
		return ret
	}
	return *o.ParameterDescription
}

// GetParameterDescriptionOk returns a tuple with the ParameterDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetParameterDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterDescription) {
		return nil, false
	}
	return o.ParameterDescription, true
}

// HasParameterDescription returns a boolean if a field has been set.
func (o *ChannelParameterType) HasParameterDescription() bool {
	if o != nil && !IsNil(o.ParameterDescription) {
		return true
	}

	return false
}

// SetParameterDescription gets a reference to the given string and assigns it to the ParameterDescription field.
func (o *ChannelParameterType) SetParameterDescription(v string) {
	o.ParameterDescription = &v
}

// GetParameterValue returns the ParameterValue field value if set, zero value otherwise.
func (o *ChannelParameterType) GetParameterValue() string {
	if o == nil || IsNil(o.ParameterValue) {
		var ret string
		return ret
	}
	return *o.ParameterValue
}

// GetParameterValueOk returns a tuple with the ParameterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetParameterValueOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterValue) {
		return nil, false
	}
	return o.ParameterValue, true
}

// HasParameterValue returns a boolean if a field has been set.
func (o *ChannelParameterType) HasParameterValue() bool {
	if o != nil && !IsNil(o.ParameterValue) {
		return true
	}

	return false
}

// SetParameterValue gets a reference to the given string and assigns it to the ParameterValue field.
func (o *ChannelParameterType) SetParameterValue(v string) {
	o.ParameterValue = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *ChannelParameterType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *ChannelParameterType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *ChannelParameterType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetParameterType returns the ParameterType field value if set, zero value otherwise.
func (o *ChannelParameterType) GetParameterType() string {
	if o == nil || IsNil(o.ParameterType) {
		var ret string
		return ret
	}
	return *o.ParameterType
}

// GetParameterTypeOk returns a tuple with the ParameterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetParameterTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParameterType) {
		return nil, false
	}
	return o.ParameterType, true
}

// HasParameterType returns a boolean if a field has been set.
func (o *ChannelParameterType) HasParameterType() bool {
	if o != nil && !IsNil(o.ParameterType) {
		return true
	}

	return false
}

// SetParameterType gets a reference to the given string and assigns it to the ParameterType field.
func (o *ChannelParameterType) SetParameterType(v string) {
	o.ParameterType = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ChannelParameterType) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ChannelParameterType) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ChannelParameterType) SetScope(v string) {
	o.Scope = &v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *ChannelParameterType) GetValueType() ChannelParameterValueType {
	if o == nil || IsNil(o.ValueType) {
		var ret ChannelParameterValueType
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetValueTypeOk() (*ChannelParameterValueType, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *ChannelParameterType) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given ChannelParameterValueType and assigns it to the ValueType field.
func (o *ChannelParameterType) SetValueType(v ChannelParameterValueType) {
	o.ValueType = &v
}

// GetSequence returns the Sequence field value if set, zero value otherwise.
func (o *ChannelParameterType) GetSequence() int32 {
	if o == nil || IsNil(o.Sequence) {
		var ret int32
		return ret
	}
	return *o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetSequenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Sequence) {
		return nil, false
	}
	return o.Sequence, true
}

// HasSequence returns a boolean if a field has been set.
func (o *ChannelParameterType) HasSequence() bool {
	if o != nil && !IsNil(o.Sequence) {
		return true
	}

	return false
}

// SetSequence gets a reference to the given int32 and assigns it to the Sequence field.
func (o *ChannelParameterType) SetSequence(v int32) {
	o.Sequence = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ChannelParameterType) GetParameters() []BaseChannelParameterType {
	if o == nil || IsNil(o.Parameters) {
		var ret []BaseChannelParameterType
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetParametersOk() ([]BaseChannelParameterType, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ChannelParameterType) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BaseChannelParameterType and assigns it to the Parameters field.
func (o *ChannelParameterType) SetParameters(v []BaseChannelParameterType) {
	o.Parameters = v
}

// GetLevelType returns the LevelType field value if set, zero value otherwise.
func (o *ChannelParameterType) GetLevelType() ChannelParameterLevelType {
	if o == nil || IsNil(o.LevelType) {
		var ret ChannelParameterLevelType
		return ret
	}
	return *o.LevelType
}

// GetLevelTypeOk returns a tuple with the LevelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetLevelTypeOk() (*ChannelParameterLevelType, bool) {
	if o == nil || IsNil(o.LevelType) {
		return nil, false
	}
	return o.LevelType, true
}

// HasLevelType returns a boolean if a field has been set.
func (o *ChannelParameterType) HasLevelType() bool {
	if o != nil && !IsNil(o.LevelType) {
		return true
	}

	return false
}

// SetLevelType gets a reference to the given ChannelParameterLevelType and assigns it to the LevelType field.
func (o *ChannelParameterType) SetLevelType(v ChannelParameterLevelType) {
	o.LevelType = &v
}

// GetLevelCode returns the LevelCode field value if set, zero value otherwise.
func (o *ChannelParameterType) GetLevelCode() string {
	if o == nil || IsNil(o.LevelCode) {
		var ret string
		return ret
	}
	return *o.LevelCode
}

// GetLevelCodeOk returns a tuple with the LevelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelParameterType) GetLevelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.LevelCode) {
		return nil, false
	}
	return o.LevelCode, true
}

// HasLevelCode returns a boolean if a field has been set.
func (o *ChannelParameterType) HasLevelCode() bool {
	if o != nil && !IsNil(o.LevelCode) {
		return true
	}

	return false
}

// SetLevelCode gets a reference to the given string and assigns it to the LevelCode field.
func (o *ChannelParameterType) SetLevelCode(v string) {
	o.LevelCode = &v
}

func (o ChannelParameterType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelParameterType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParameterName) {
		toSerialize["parameterName"] = o.ParameterName
	}
	if !IsNil(o.ParameterDisplay) {
		toSerialize["parameterDisplay"] = o.ParameterDisplay
	}
	if !IsNil(o.ShortDescription) {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if !IsNil(o.ParameterDescription) {
		toSerialize["parameterDescription"] = o.ParameterDescription
	}
	if !IsNil(o.ParameterValue) {
		toSerialize["parameterValue"] = o.ParameterValue
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.ParameterType) {
		toSerialize["parameterType"] = o.ParameterType
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.Sequence) {
		toSerialize["sequence"] = o.Sequence
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.LevelType) {
		toSerialize["levelType"] = o.LevelType
	}
	if !IsNil(o.LevelCode) {
		toSerialize["levelCode"] = o.LevelCode
	}
	return toSerialize, nil
}

type NullableChannelParameterType struct {
	value *ChannelParameterType
	isSet bool
}

func (v NullableChannelParameterType) Get() *ChannelParameterType {
	return v.value
}

func (v *NullableChannelParameterType) Set(val *ChannelParameterType) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelParameterType) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelParameterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelParameterType(val *ChannelParameterType) *NullableChannelParameterType {
	return &NullableChannelParameterType{value: val, isSet: true}
}

func (v NullableChannelParameterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelParameterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


