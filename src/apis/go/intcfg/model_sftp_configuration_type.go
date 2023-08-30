/*
OPERA Cloud Integration Configuration API

APIs catering to Integration Configuration in OPERA Cloud.  Operations such as get Hotel Interface Types, or get UDF mappings can be found in this module.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intcfg

import (
	"encoding/json"
)

// checks if the SftpConfigurationType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SftpConfigurationType{}

// SftpConfigurationType Information which uniquely identifies SFTP Configuration
type SftpConfigurationType struct {
	// Unique id associated with this configuration
	ConfigurationId *int32 `json:"configurationId,omitempty"`
	// SFTP destination
	Destination *string `json:"destination,omitempty"`
	// Description of the destination, such as Shift Reports.
	Description *string `json:"description,omitempty"`
	// Indicates whether the configuration is inactive or not.
	Inactive *bool `json:"inactive,omitempty"`
	// Hotel code
	HotelId *string `json:"hotelId,omitempty"`
}

// NewSftpConfigurationType instantiates a new SftpConfigurationType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSftpConfigurationType() *SftpConfigurationType {
	this := SftpConfigurationType{}
	return &this
}

// NewSftpConfigurationTypeWithDefaults instantiates a new SftpConfigurationType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSftpConfigurationTypeWithDefaults() *SftpConfigurationType {
	this := SftpConfigurationType{}
	return &this
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *SftpConfigurationType) GetConfigurationId() int32 {
	if o == nil || IsNil(o.ConfigurationId) {
		var ret int32
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SftpConfigurationType) GetConfigurationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfigurationId) {
		return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *SftpConfigurationType) HasConfigurationId() bool {
	if o != nil && !IsNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given int32 and assigns it to the ConfigurationId field.
func (o *SftpConfigurationType) SetConfigurationId(v int32) {
	o.ConfigurationId = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *SftpConfigurationType) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SftpConfigurationType) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *SftpConfigurationType) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *SftpConfigurationType) SetDestination(v string) {
	o.Destination = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SftpConfigurationType) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SftpConfigurationType) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SftpConfigurationType) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SftpConfigurationType) SetDescription(v string) {
	o.Description = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *SftpConfigurationType) GetInactive() bool {
	if o == nil || IsNil(o.Inactive) {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SftpConfigurationType) GetInactiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *SftpConfigurationType) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *SftpConfigurationType) SetInactive(v bool) {
	o.Inactive = &v
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *SftpConfigurationType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SftpConfigurationType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *SftpConfigurationType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *SftpConfigurationType) SetHotelId(v string) {
	o.HotelId = &v
}

func (o SftpConfigurationType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SftpConfigurationType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigurationId) {
		toSerialize["configurationId"] = o.ConfigurationId
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	return toSerialize, nil
}

type NullableSftpConfigurationType struct {
	value *SftpConfigurationType
	isSet bool
}

func (v NullableSftpConfigurationType) Get() *SftpConfigurationType {
	return v.value
}

func (v *NullableSftpConfigurationType) Set(val *SftpConfigurationType) {
	v.value = val
	v.isSet = true
}

func (v NullableSftpConfigurationType) IsSet() bool {
	return v.isSet
}

func (v *NullableSftpConfigurationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSftpConfigurationType(val *SftpConfigurationType) *NullableSftpConfigurationType {
	return &NullableSftpConfigurationType{value: val, isSet: true}
}

func (v NullableSftpConfigurationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSftpConfigurationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


