/*
OPERA Cloud Leisure Management API

APIs to cater for external Leisure Management functionality integrated with OPERA Cloud.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lms

import (
	"encoding/json"
	"time"
)

// checks if the ChangeActivityTypesRS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeActivityTypesRS{}

// ChangeActivityTypesRS Existing Operations Responses will eventually be modified to be extended from this type.
type ChangeActivityTypesRS struct {
	// Returning an empty element of this type indicates the successful processing of an message. This is used in conjunction with the Warning Type to report any warnings or business errors.
	Success map[string]interface{} `json:"success,omitempty"`
	// Used in conjunction with the Success element to define a business error.
	Warnings []WarningType `json:"warnings,omitempty"`
	// An error that occurred during the processing of a message.
	Errors []ErrorType `json:"errors,omitempty"`
	// A reference for additional message identification, assigned by the requesting host system. When a request message includes an echo token the corresponding response message MUST include an echo token with an identical value.
	EchoToken *string `json:"echoToken,omitempty"`
	// Indicates the creation date and time of the message in UTC using the following format specified by ISO 8601; YYYY-MM-DDThh:mm:ssZ with time values using the 24 hour clock (e.g. 20 November 2003, 1:59:38 pm UTC becomes 2003-11-20T13:59:38Z).
	TimeStamp *time.Time `json:"timeStamp,omitempty"`
	// For all Opera versioned messages, the version of the message is indicated by a Opera Version value.
	Version *string `json:"version,omitempty"`
	// Allow end-to-end correlation of log messages with the corresponding Web service message throughout the processing of the Web service message.
	CorrelationId *string `json:"correlationId,omitempty"`
	// Indicates if the operation supports the ability to retry the request.
	RetryAllowed *bool `json:"retryAllowed,omitempty"`
	// Indicates if the operation supports the ability to force the retry request through OPERA services in the case where the external system continues to fail.
	EnforceAllowed *bool `json:"enforceAllowed,omitempty"`
	// This attribute carries the user selected confirmation value on confirmation popup.
	UseLocalAllowed *bool `json:"useLocalAllowed,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
}

// NewChangeActivityTypesRS instantiates a new ChangeActivityTypesRS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeActivityTypesRS() *ChangeActivityTypesRS {
	this := ChangeActivityTypesRS{}
	return &this
}

// NewChangeActivityTypesRSWithDefaults instantiates a new ChangeActivityTypesRS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeActivityTypesRSWithDefaults() *ChangeActivityTypesRS {
	this := ChangeActivityTypesRS{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetSuccess() map[string]interface{} {
	if o == nil || IsNil(o.Success) {
		var ret map[string]interface{}
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetSuccessOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Success) {
		return map[string]interface{}{}, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given map[string]interface{} and assigns it to the Success field.
func (o *ChangeActivityTypesRS) SetSuccess(v map[string]interface{}) {
	o.Success = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetWarnings() []WarningType {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningType
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetWarningsOk() ([]WarningType, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningType and assigns it to the Warnings field.
func (o *ChangeActivityTypesRS) SetWarnings(v []WarningType) {
	o.Warnings = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetErrors() []ErrorType {
	if o == nil || IsNil(o.Errors) {
		var ret []ErrorType
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetErrorsOk() ([]ErrorType, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorType and assigns it to the Errors field.
func (o *ChangeActivityTypesRS) SetErrors(v []ErrorType) {
	o.Errors = v
}

// GetEchoToken returns the EchoToken field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetEchoToken() string {
	if o == nil || IsNil(o.EchoToken) {
		var ret string
		return ret
	}
	return *o.EchoToken
}

// GetEchoTokenOk returns a tuple with the EchoToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetEchoTokenOk() (*string, bool) {
	if o == nil || IsNil(o.EchoToken) {
		return nil, false
	}
	return o.EchoToken, true
}

// HasEchoToken returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasEchoToken() bool {
	if o != nil && !IsNil(o.EchoToken) {
		return true
	}

	return false
}

// SetEchoToken gets a reference to the given string and assigns it to the EchoToken field.
func (o *ChangeActivityTypesRS) SetEchoToken(v string) {
	o.EchoToken = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetTimeStamp() time.Time {
	if o == nil || IsNil(o.TimeStamp) {
		var ret time.Time
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetTimeStampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given time.Time and assigns it to the TimeStamp field.
func (o *ChangeActivityTypesRS) SetTimeStamp(v time.Time) {
	o.TimeStamp = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ChangeActivityTypesRS) SetVersion(v string) {
	o.Version = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *ChangeActivityTypesRS) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetRetryAllowed returns the RetryAllowed field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetRetryAllowed() bool {
	if o == nil || IsNil(o.RetryAllowed) {
		var ret bool
		return ret
	}
	return *o.RetryAllowed
}

// GetRetryAllowedOk returns a tuple with the RetryAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetRetryAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.RetryAllowed) {
		return nil, false
	}
	return o.RetryAllowed, true
}

// HasRetryAllowed returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasRetryAllowed() bool {
	if o != nil && !IsNil(o.RetryAllowed) {
		return true
	}

	return false
}

// SetRetryAllowed gets a reference to the given bool and assigns it to the RetryAllowed field.
func (o *ChangeActivityTypesRS) SetRetryAllowed(v bool) {
	o.RetryAllowed = &v
}

// GetEnforceAllowed returns the EnforceAllowed field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetEnforceAllowed() bool {
	if o == nil || IsNil(o.EnforceAllowed) {
		var ret bool
		return ret
	}
	return *o.EnforceAllowed
}

// GetEnforceAllowedOk returns a tuple with the EnforceAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetEnforceAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.EnforceAllowed) {
		return nil, false
	}
	return o.EnforceAllowed, true
}

// HasEnforceAllowed returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasEnforceAllowed() bool {
	if o != nil && !IsNil(o.EnforceAllowed) {
		return true
	}

	return false
}

// SetEnforceAllowed gets a reference to the given bool and assigns it to the EnforceAllowed field.
func (o *ChangeActivityTypesRS) SetEnforceAllowed(v bool) {
	o.EnforceAllowed = &v
}

// GetUseLocalAllowed returns the UseLocalAllowed field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetUseLocalAllowed() bool {
	if o == nil || IsNil(o.UseLocalAllowed) {
		var ret bool
		return ret
	}
	return *o.UseLocalAllowed
}

// GetUseLocalAllowedOk returns a tuple with the UseLocalAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetUseLocalAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.UseLocalAllowed) {
		return nil, false
	}
	return o.UseLocalAllowed, true
}

// HasUseLocalAllowed returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasUseLocalAllowed() bool {
	if o != nil && !IsNil(o.UseLocalAllowed) {
		return true
	}

	return false
}

// SetUseLocalAllowed gets a reference to the given bool and assigns it to the UseLocalAllowed field.
func (o *ChangeActivityTypesRS) SetUseLocalAllowed(v bool) {
	o.UseLocalAllowed = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ChangeActivityTypesRS) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeActivityTypesRS) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ChangeActivityTypesRS) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ChangeActivityTypesRS) SetLinks(v []InstanceLink) {
	o.Links = v
}

func (o ChangeActivityTypesRS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeActivityTypesRS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.EchoToken) {
		toSerialize["echoToken"] = o.EchoToken
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.RetryAllowed) {
		toSerialize["retryAllowed"] = o.RetryAllowed
	}
	if !IsNil(o.EnforceAllowed) {
		toSerialize["enforceAllowed"] = o.EnforceAllowed
	}
	if !IsNil(o.UseLocalAllowed) {
		toSerialize["useLocalAllowed"] = o.UseLocalAllowed
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableChangeActivityTypesRS struct {
	value *ChangeActivityTypesRS
	isSet bool
}

func (v NullableChangeActivityTypesRS) Get() *ChangeActivityTypesRS {
	return v.value
}

func (v *NullableChangeActivityTypesRS) Set(val *ChangeActivityTypesRS) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeActivityTypesRS) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeActivityTypesRS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeActivityTypesRS(val *ChangeActivityTypesRS) *NullableChangeActivityTypesRS {
	return &NullableChangeActivityTypesRS{value: val, isSet: true}
}

func (v NullableChangeActivityTypesRS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeActivityTypesRS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


