/*
OPERA Cloud Cashiering API

APIs to cater for Cashiering related functionality in OPERA Cloud.<br /><br /> Cashiering provides access to a guest folio, posting journals, receipt histories, currency calculations, credit card settlements, and check a guest out.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 22.3.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package csh

import (
	"encoding/json"
)

// checks if the VaultHTTPTransactionMessageType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VaultHTTPTransactionMessageType{}

// VaultHTTPTransactionMessageType struct for VaultHTTPTransactionMessageType
type VaultHTTPTransactionMessageType struct {
	// The hotel context of the transaction.
	HotelId *string `json:"hotelId,omitempty"`
	// The HTTP request entity content. The needs to use escape characters.
	EscapedRequestContent *string `json:"escapedRequestContent,omitempty"`
	// The HTTP response entity content. The needs to use escape characters.
	EscapedResponseContent *string `json:"escapedResponseContent,omitempty"`
	HTTPTransactionDuration *DateRangeType `json:"hTTPTransactionDuration,omitempty"`
	HTTPError *ErrorType `json:"hTTPError,omitempty"`
	AuthorizationApproval *VaultHTTPTransactionMessageTypeAuthorizationApproval `json:"authorizationApproval,omitempty"`
	Type *VaultHTTPTransactionType `json:"type,omitempty"`
}

// NewVaultHTTPTransactionMessageType instantiates a new VaultHTTPTransactionMessageType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultHTTPTransactionMessageType() *VaultHTTPTransactionMessageType {
	this := VaultHTTPTransactionMessageType{}
	return &this
}

// NewVaultHTTPTransactionMessageTypeWithDefaults instantiates a new VaultHTTPTransactionMessageType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultHTTPTransactionMessageTypeWithDefaults() *VaultHTTPTransactionMessageType {
	this := VaultHTTPTransactionMessageType{}
	return &this
}

// GetHotelId returns the HotelId field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetHotelId() string {
	if o == nil || IsNil(o.HotelId) {
		var ret string
		return ret
	}
	return *o.HotelId
}

// GetHotelIdOk returns a tuple with the HotelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetHotelIdOk() (*string, bool) {
	if o == nil || IsNil(o.HotelId) {
		return nil, false
	}
	return o.HotelId, true
}

// HasHotelId returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasHotelId() bool {
	if o != nil && !IsNil(o.HotelId) {
		return true
	}

	return false
}

// SetHotelId gets a reference to the given string and assigns it to the HotelId field.
func (o *VaultHTTPTransactionMessageType) SetHotelId(v string) {
	o.HotelId = &v
}

// GetEscapedRequestContent returns the EscapedRequestContent field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetEscapedRequestContent() string {
	if o == nil || IsNil(o.EscapedRequestContent) {
		var ret string
		return ret
	}
	return *o.EscapedRequestContent
}

// GetEscapedRequestContentOk returns a tuple with the EscapedRequestContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetEscapedRequestContentOk() (*string, bool) {
	if o == nil || IsNil(o.EscapedRequestContent) {
		return nil, false
	}
	return o.EscapedRequestContent, true
}

// HasEscapedRequestContent returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasEscapedRequestContent() bool {
	if o != nil && !IsNil(o.EscapedRequestContent) {
		return true
	}

	return false
}

// SetEscapedRequestContent gets a reference to the given string and assigns it to the EscapedRequestContent field.
func (o *VaultHTTPTransactionMessageType) SetEscapedRequestContent(v string) {
	o.EscapedRequestContent = &v
}

// GetEscapedResponseContent returns the EscapedResponseContent field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetEscapedResponseContent() string {
	if o == nil || IsNil(o.EscapedResponseContent) {
		var ret string
		return ret
	}
	return *o.EscapedResponseContent
}

// GetEscapedResponseContentOk returns a tuple with the EscapedResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetEscapedResponseContentOk() (*string, bool) {
	if o == nil || IsNil(o.EscapedResponseContent) {
		return nil, false
	}
	return o.EscapedResponseContent, true
}

// HasEscapedResponseContent returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasEscapedResponseContent() bool {
	if o != nil && !IsNil(o.EscapedResponseContent) {
		return true
	}

	return false
}

// SetEscapedResponseContent gets a reference to the given string and assigns it to the EscapedResponseContent field.
func (o *VaultHTTPTransactionMessageType) SetEscapedResponseContent(v string) {
	o.EscapedResponseContent = &v
}

// GetHTTPTransactionDuration returns the HTTPTransactionDuration field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetHTTPTransactionDuration() DateRangeType {
	if o == nil || IsNil(o.HTTPTransactionDuration) {
		var ret DateRangeType
		return ret
	}
	return *o.HTTPTransactionDuration
}

// GetHTTPTransactionDurationOk returns a tuple with the HTTPTransactionDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetHTTPTransactionDurationOk() (*DateRangeType, bool) {
	if o == nil || IsNil(o.HTTPTransactionDuration) {
		return nil, false
	}
	return o.HTTPTransactionDuration, true
}

// HasHTTPTransactionDuration returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasHTTPTransactionDuration() bool {
	if o != nil && !IsNil(o.HTTPTransactionDuration) {
		return true
	}

	return false
}

// SetHTTPTransactionDuration gets a reference to the given DateRangeType and assigns it to the HTTPTransactionDuration field.
func (o *VaultHTTPTransactionMessageType) SetHTTPTransactionDuration(v DateRangeType) {
	o.HTTPTransactionDuration = &v
}

// GetHTTPError returns the HTTPError field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetHTTPError() ErrorType {
	if o == nil || IsNil(o.HTTPError) {
		var ret ErrorType
		return ret
	}
	return *o.HTTPError
}

// GetHTTPErrorOk returns a tuple with the HTTPError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetHTTPErrorOk() (*ErrorType, bool) {
	if o == nil || IsNil(o.HTTPError) {
		return nil, false
	}
	return o.HTTPError, true
}

// HasHTTPError returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasHTTPError() bool {
	if o != nil && !IsNil(o.HTTPError) {
		return true
	}

	return false
}

// SetHTTPError gets a reference to the given ErrorType and assigns it to the HTTPError field.
func (o *VaultHTTPTransactionMessageType) SetHTTPError(v ErrorType) {
	o.HTTPError = &v
}

// GetAuthorizationApproval returns the AuthorizationApproval field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetAuthorizationApproval() VaultHTTPTransactionMessageTypeAuthorizationApproval {
	if o == nil || IsNil(o.AuthorizationApproval) {
		var ret VaultHTTPTransactionMessageTypeAuthorizationApproval
		return ret
	}
	return *o.AuthorizationApproval
}

// GetAuthorizationApprovalOk returns a tuple with the AuthorizationApproval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetAuthorizationApprovalOk() (*VaultHTTPTransactionMessageTypeAuthorizationApproval, bool) {
	if o == nil || IsNil(o.AuthorizationApproval) {
		return nil, false
	}
	return o.AuthorizationApproval, true
}

// HasAuthorizationApproval returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasAuthorizationApproval() bool {
	if o != nil && !IsNil(o.AuthorizationApproval) {
		return true
	}

	return false
}

// SetAuthorizationApproval gets a reference to the given VaultHTTPTransactionMessageTypeAuthorizationApproval and assigns it to the AuthorizationApproval field.
func (o *VaultHTTPTransactionMessageType) SetAuthorizationApproval(v VaultHTTPTransactionMessageTypeAuthorizationApproval) {
	o.AuthorizationApproval = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VaultHTTPTransactionMessageType) GetType() VaultHTTPTransactionType {
	if o == nil || IsNil(o.Type) {
		var ret VaultHTTPTransactionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultHTTPTransactionMessageType) GetTypeOk() (*VaultHTTPTransactionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VaultHTTPTransactionMessageType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given VaultHTTPTransactionType and assigns it to the Type field.
func (o *VaultHTTPTransactionMessageType) SetType(v VaultHTTPTransactionType) {
	o.Type = &v
}

func (o VaultHTTPTransactionMessageType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VaultHTTPTransactionMessageType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HotelId) {
		toSerialize["hotelId"] = o.HotelId
	}
	if !IsNil(o.EscapedRequestContent) {
		toSerialize["escapedRequestContent"] = o.EscapedRequestContent
	}
	if !IsNil(o.EscapedResponseContent) {
		toSerialize["escapedResponseContent"] = o.EscapedResponseContent
	}
	if !IsNil(o.HTTPTransactionDuration) {
		toSerialize["hTTPTransactionDuration"] = o.HTTPTransactionDuration
	}
	if !IsNil(o.HTTPError) {
		toSerialize["hTTPError"] = o.HTTPError
	}
	if !IsNil(o.AuthorizationApproval) {
		toSerialize["authorizationApproval"] = o.AuthorizationApproval
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableVaultHTTPTransactionMessageType struct {
	value *VaultHTTPTransactionMessageType
	isSet bool
}

func (v NullableVaultHTTPTransactionMessageType) Get() *VaultHTTPTransactionMessageType {
	return v.value
}

func (v *NullableVaultHTTPTransactionMessageType) Set(val *VaultHTTPTransactionMessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultHTTPTransactionMessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultHTTPTransactionMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultHTTPTransactionMessageType(val *VaultHTTPTransactionMessageType) *NullableVaultHTTPTransactionMessageType {
	return &NullableVaultHTTPTransactionMessageType{value: val, isSet: true}
}

func (v NullableVaultHTTPTransactionMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultHTTPTransactionMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


