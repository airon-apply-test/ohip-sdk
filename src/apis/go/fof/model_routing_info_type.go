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

// checks if the RoutingInfoType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingInfoType{}

// RoutingInfoType A routing info object can either be of type Folio OR of type Room with its corresponding object.
type RoutingInfoType struct {
	Folio *RoutingInfoTypeFolio `json:"folio,omitempty"`
	Room *RoutingInfoTypeRoom `json:"room,omitempty"`
	Comp *RoutingInfoTypeComp `json:"comp,omitempty"`
	Request *RoutingInfoTypeRequest `json:"request,omitempty"`
}

// NewRoutingInfoType instantiates a new RoutingInfoType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingInfoType() *RoutingInfoType {
	this := RoutingInfoType{}
	return &this
}

// NewRoutingInfoTypeWithDefaults instantiates a new RoutingInfoType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingInfoTypeWithDefaults() *RoutingInfoType {
	this := RoutingInfoType{}
	return &this
}

// GetFolio returns the Folio field value if set, zero value otherwise.
func (o *RoutingInfoType) GetFolio() RoutingInfoTypeFolio {
	if o == nil || IsNil(o.Folio) {
		var ret RoutingInfoTypeFolio
		return ret
	}
	return *o.Folio
}

// GetFolioOk returns a tuple with the Folio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoType) GetFolioOk() (*RoutingInfoTypeFolio, bool) {
	if o == nil || IsNil(o.Folio) {
		return nil, false
	}
	return o.Folio, true
}

// HasFolio returns a boolean if a field has been set.
func (o *RoutingInfoType) HasFolio() bool {
	if o != nil && !IsNil(o.Folio) {
		return true
	}

	return false
}

// SetFolio gets a reference to the given RoutingInfoTypeFolio and assigns it to the Folio field.
func (o *RoutingInfoType) SetFolio(v RoutingInfoTypeFolio) {
	o.Folio = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *RoutingInfoType) GetRoom() RoutingInfoTypeRoom {
	if o == nil || IsNil(o.Room) {
		var ret RoutingInfoTypeRoom
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoType) GetRoomOk() (*RoutingInfoTypeRoom, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *RoutingInfoType) HasRoom() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given RoutingInfoTypeRoom and assigns it to the Room field.
func (o *RoutingInfoType) SetRoom(v RoutingInfoTypeRoom) {
	o.Room = &v
}

// GetComp returns the Comp field value if set, zero value otherwise.
func (o *RoutingInfoType) GetComp() RoutingInfoTypeComp {
	if o == nil || IsNil(o.Comp) {
		var ret RoutingInfoTypeComp
		return ret
	}
	return *o.Comp
}

// GetCompOk returns a tuple with the Comp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoType) GetCompOk() (*RoutingInfoTypeComp, bool) {
	if o == nil || IsNil(o.Comp) {
		return nil, false
	}
	return o.Comp, true
}

// HasComp returns a boolean if a field has been set.
func (o *RoutingInfoType) HasComp() bool {
	if o != nil && !IsNil(o.Comp) {
		return true
	}

	return false
}

// SetComp gets a reference to the given RoutingInfoTypeComp and assigns it to the Comp field.
func (o *RoutingInfoType) SetComp(v RoutingInfoTypeComp) {
	o.Comp = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *RoutingInfoType) GetRequest() RoutingInfoTypeRequest {
	if o == nil || IsNil(o.Request) {
		var ret RoutingInfoTypeRequest
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingInfoType) GetRequestOk() (*RoutingInfoTypeRequest, bool) {
	if o == nil || IsNil(o.Request) {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *RoutingInfoType) HasRequest() bool {
	if o != nil && !IsNil(o.Request) {
		return true
	}

	return false
}

// SetRequest gets a reference to the given RoutingInfoTypeRequest and assigns it to the Request field.
func (o *RoutingInfoType) SetRequest(v RoutingInfoTypeRequest) {
	o.Request = &v
}

func (o RoutingInfoType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingInfoType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Folio) {
		toSerialize["folio"] = o.Folio
	}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if !IsNil(o.Comp) {
		toSerialize["comp"] = o.Comp
	}
	if !IsNil(o.Request) {
		toSerialize["request"] = o.Request
	}
	return toSerialize, nil
}

type NullableRoutingInfoType struct {
	value *RoutingInfoType
	isSet bool
}

func (v NullableRoutingInfoType) Get() *RoutingInfoType {
	return v.value
}

func (v *NullableRoutingInfoType) Set(val *RoutingInfoType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingInfoType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingInfoType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingInfoType(val *RoutingInfoType) *NullableRoutingInfoType {
	return &NullableRoutingInfoType{value: val, isSet: true}
}

func (v NullableRoutingInfoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingInfoType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


