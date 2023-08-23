/*
OPERA Cloud Front Desk Operations Service

APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ExceptionDetailType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionDetailType{}

// ExceptionDetailType Complex type that contains error details for a REST call.
type ExceptionDetailType struct {
	// Absolute URI [RFC3986] that identifies the problem type.  When dereferenced, it SHOULD provide a human-readable summary of the problem (for example, using HTML).
	Type string `json:"type"`
	// Short, human-readable summary of the problem.  The summary SHOULD NOT change for subsequent occurrences of the problem, except for purposes of localization.
	Title string `json:"title"`
	// HTTP status code for this occurrence of the problem, set by the origin server.
	Status *int32 `json:"status,omitempty"`
	// Human-readable description specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`
	// Absolute URI that identifies the specific occurrence of the problem.  It may or may not provide additional information if dereferenced.
	Instance *string `json:"instance,omitempty"`
	// Application error code, which is different from HTTP error code.
	OerrorCode *string `json:"o:errorCode,omitempty"`
	// Path to the problem at the resource or property level.
	OerrorPath *string `json:"o:errorPath,omitempty"`
	// Details of the error message, consisting of a hierarchical tree structure.
	OerrorDetails []ErrorInstance `json:"o:errorDetails,omitempty"`
	Links []InstanceLink `json:"links,omitempty"`
}

// NewExceptionDetailType instantiates a new ExceptionDetailType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionDetailType(type_ string, title string) *ExceptionDetailType {
	this := ExceptionDetailType{}
	this.Type = type_
	this.Title = title
	return &this
}

// NewExceptionDetailTypeWithDefaults instantiates a new ExceptionDetailType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionDetailTypeWithDefaults() *ExceptionDetailType {
	this := ExceptionDetailType{}
	return &this
}

// GetType returns the Type field value
func (o *ExceptionDetailType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExceptionDetailType) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ExceptionDetailType) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExceptionDetailType) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetStatus() int32 {
	if o == nil || IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetStatusOk() (*int32, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ExceptionDetailType) SetStatus(v int32) {
	o.Status = &v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ExceptionDetailType) SetDetail(v string) {
	o.Detail = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ExceptionDetailType) SetInstance(v string) {
	o.Instance = &v
}

// GetOerrorCode returns the OerrorCode field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetOerrorCode() string {
	if o == nil || IsNil(o.OerrorCode) {
		var ret string
		return ret
	}
	return *o.OerrorCode
}

// GetOerrorCodeOk returns a tuple with the OerrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetOerrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorCode) {
		return nil, false
	}
	return o.OerrorCode, true
}

// HasOerrorCode returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasOerrorCode() bool {
	if o != nil && !IsNil(o.OerrorCode) {
		return true
	}

	return false
}

// SetOerrorCode gets a reference to the given string and assigns it to the OerrorCode field.
func (o *ExceptionDetailType) SetOerrorCode(v string) {
	o.OerrorCode = &v
}

// GetOerrorPath returns the OerrorPath field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetOerrorPath() string {
	if o == nil || IsNil(o.OerrorPath) {
		var ret string
		return ret
	}
	return *o.OerrorPath
}

// GetOerrorPathOk returns a tuple with the OerrorPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetOerrorPathOk() (*string, bool) {
	if o == nil || IsNil(o.OerrorPath) {
		return nil, false
	}
	return o.OerrorPath, true
}

// HasOerrorPath returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasOerrorPath() bool {
	if o != nil && !IsNil(o.OerrorPath) {
		return true
	}

	return false
}

// SetOerrorPath gets a reference to the given string and assigns it to the OerrorPath field.
func (o *ExceptionDetailType) SetOerrorPath(v string) {
	o.OerrorPath = &v
}

// GetOerrorDetails returns the OerrorDetails field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetOerrorDetails() []ErrorInstance {
	if o == nil || IsNil(o.OerrorDetails) {
		var ret []ErrorInstance
		return ret
	}
	return o.OerrorDetails
}

// GetOerrorDetailsOk returns a tuple with the OerrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetOerrorDetailsOk() ([]ErrorInstance, bool) {
	if o == nil || IsNil(o.OerrorDetails) {
		return nil, false
	}
	return o.OerrorDetails, true
}

// HasOerrorDetails returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasOerrorDetails() bool {
	if o != nil && !IsNil(o.OerrorDetails) {
		return true
	}

	return false
}

// SetOerrorDetails gets a reference to the given []ErrorInstance and assigns it to the OerrorDetails field.
func (o *ExceptionDetailType) SetOerrorDetails(v []ErrorInstance) {
	o.OerrorDetails = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExceptionDetailType) GetLinks() []InstanceLink {
	if o == nil || IsNil(o.Links) {
		var ret []InstanceLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionDetailType) GetLinksOk() ([]InstanceLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExceptionDetailType) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []InstanceLink and assigns it to the Links field.
func (o *ExceptionDetailType) SetLinks(v []InstanceLink) {
	o.Links = v
}

func (o ExceptionDetailType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionDetailType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.OerrorCode) {
		toSerialize["o:errorCode"] = o.OerrorCode
	}
	if !IsNil(o.OerrorPath) {
		toSerialize["o:errorPath"] = o.OerrorPath
	}
	if !IsNil(o.OerrorDetails) {
		toSerialize["o:errorDetails"] = o.OerrorDetails
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableExceptionDetailType struct {
	value *ExceptionDetailType
	isSet bool
}

func (v NullableExceptionDetailType) Get() *ExceptionDetailType {
	return v.value
}

func (v *NullableExceptionDetailType) Set(val *ExceptionDetailType) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionDetailType) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionDetailType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionDetailType(val *ExceptionDetailType) *NullableExceptionDetailType {
	return &NullableExceptionDetailType{value: val, isSet: true}
}

func (v NullableExceptionDetailType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionDetailType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


