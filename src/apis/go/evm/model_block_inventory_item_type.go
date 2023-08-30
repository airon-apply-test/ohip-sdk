/*
OPERA Cloud Sales Event Management API

APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>

API version: 21.5.0.0
Contact: hospitality_apis_ww_grp@oracle.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evm

import (
	"encoding/json"
)

// checks if the BlockInventoryItemType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockInventoryItemType{}

// BlockInventoryItemType An identifier used to uniquely reference an object in a system (e.g. an airline reservation reference, customer profile reference, booking confirmation number, or a reference to a previous availability quote).
type BlockInventoryItemType struct {
	// URL that identifies the location associated with the record identified by the UniqueID.
	Url *string `json:"url,omitempty"`
	// A reference to the type of object defined by the UniqueID elementSpace.
	Type *string `json:"type,omitempty"`
	// The identification of a record as it exists at a point in time. An instance is used in update messages where the sender must assure the server that the update sent refers to the most recent modification level of the object being updated.
	Instance *string `json:"instance,omitempty"`
	// Used to identify the source of the identifier (e.g., IATA, ABTA).
	IdContext *string `json:"idContext,omitempty"`
	// A unique identifying value assigned by the creating system. The ID attribute may be used to reference a primary-key value within a database or in a particular implementation.
	Id *string `json:"id,omitempty"`
	// Additional identifying value assigned by the creating system.
	IdExtension *int32 `json:"idExtension,omitempty"`
	Item *ItemInfoType `json:"item,omitempty"`
	// Number of items booked.
	Quantity *int32 `json:"quantity,omitempty"`
	BlockDates *DateRangeType `json:"blockDates,omitempty"`
	Source *BlockInventoryItemSourceType `json:"source,omitempty"`
}

// NewBlockInventoryItemType instantiates a new BlockInventoryItemType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockInventoryItemType() *BlockInventoryItemType {
	this := BlockInventoryItemType{}
	return &this
}

// NewBlockInventoryItemTypeWithDefaults instantiates a new BlockInventoryItemType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockInventoryItemTypeWithDefaults() *BlockInventoryItemType {
	this := BlockInventoryItemType{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BlockInventoryItemType) SetUrl(v string) {
	o.Url = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BlockInventoryItemType) SetType(v string) {
	o.Type = &v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *BlockInventoryItemType) SetInstance(v string) {
	o.Instance = &v
}

// GetIdContext returns the IdContext field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetIdContext() string {
	if o == nil || IsNil(o.IdContext) {
		var ret string
		return ret
	}
	return *o.IdContext
}

// GetIdContextOk returns a tuple with the IdContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetIdContextOk() (*string, bool) {
	if o == nil || IsNil(o.IdContext) {
		return nil, false
	}
	return o.IdContext, true
}

// HasIdContext returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasIdContext() bool {
	if o != nil && !IsNil(o.IdContext) {
		return true
	}

	return false
}

// SetIdContext gets a reference to the given string and assigns it to the IdContext field.
func (o *BlockInventoryItemType) SetIdContext(v string) {
	o.IdContext = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BlockInventoryItemType) SetId(v string) {
	o.Id = &v
}

// GetIdExtension returns the IdExtension field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetIdExtension() int32 {
	if o == nil || IsNil(o.IdExtension) {
		var ret int32
		return ret
	}
	return *o.IdExtension
}

// GetIdExtensionOk returns a tuple with the IdExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetIdExtensionOk() (*int32, bool) {
	if o == nil || IsNil(o.IdExtension) {
		return nil, false
	}
	return o.IdExtension, true
}

// HasIdExtension returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasIdExtension() bool {
	if o != nil && !IsNil(o.IdExtension) {
		return true
	}

	return false
}

// SetIdExtension gets a reference to the given int32 and assigns it to the IdExtension field.
func (o *BlockInventoryItemType) SetIdExtension(v int32) {
	o.IdExtension = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetItem() ItemInfoType {
	if o == nil || IsNil(o.Item) {
		var ret ItemInfoType
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetItemOk() (*ItemInfoType, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given ItemInfoType and assigns it to the Item field.
func (o *BlockInventoryItemType) SetItem(v ItemInfoType) {
	o.Item = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity) {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *BlockInventoryItemType) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetBlockDates returns the BlockDates field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetBlockDates() DateRangeType {
	if o == nil || IsNil(o.BlockDates) {
		var ret DateRangeType
		return ret
	}
	return *o.BlockDates
}

// GetBlockDatesOk returns a tuple with the BlockDates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetBlockDatesOk() (*DateRangeType, bool) {
	if o == nil || IsNil(o.BlockDates) {
		return nil, false
	}
	return o.BlockDates, true
}

// HasBlockDates returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasBlockDates() bool {
	if o != nil && !IsNil(o.BlockDates) {
		return true
	}

	return false
}

// SetBlockDates gets a reference to the given DateRangeType and assigns it to the BlockDates field.
func (o *BlockInventoryItemType) SetBlockDates(v DateRangeType) {
	o.BlockDates = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BlockInventoryItemType) GetSource() BlockInventoryItemSourceType {
	if o == nil || IsNil(o.Source) {
		var ret BlockInventoryItemSourceType
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockInventoryItemType) GetSourceOk() (*BlockInventoryItemSourceType, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BlockInventoryItemType) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given BlockInventoryItemSourceType and assigns it to the Source field.
func (o *BlockInventoryItemType) SetSource(v BlockInventoryItemSourceType) {
	o.Source = &v
}

func (o BlockInventoryItemType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockInventoryItemType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !IsNil(o.IdContext) {
		toSerialize["idContext"] = o.IdContext
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IdExtension) {
		toSerialize["idExtension"] = o.IdExtension
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.BlockDates) {
		toSerialize["blockDates"] = o.BlockDates
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	return toSerialize, nil
}

type NullableBlockInventoryItemType struct {
	value *BlockInventoryItemType
	isSet bool
}

func (v NullableBlockInventoryItemType) Get() *BlockInventoryItemType {
	return v.value
}

func (v *NullableBlockInventoryItemType) Set(val *BlockInventoryItemType) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockInventoryItemType) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockInventoryItemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockInventoryItemType(val *BlockInventoryItemType) *NullableBlockInventoryItemType {
	return &NullableBlockInventoryItemType{value: val, isSet: true}
}

func (v NullableBlockInventoryItemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockInventoryItemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


