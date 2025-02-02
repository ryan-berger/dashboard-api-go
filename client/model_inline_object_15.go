/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject15 struct for InlineObject15
type InlineObject15 struct {
	Livestream *DevicesSerialSensorRelationshipsLivestream1 `json:"livestream,omitempty"`
}

// NewInlineObject15 instantiates a new InlineObject15 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject15() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// NewInlineObject15WithDefaults instantiates a new InlineObject15 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject15WithDefaults() *InlineObject15 {
	this := InlineObject15{}
	return &this
}

// GetLivestream returns the Livestream field value if set, zero value otherwise.
func (o *InlineObject15) GetLivestream() DevicesSerialSensorRelationshipsLivestream1 {
	if o == nil || isNil(o.Livestream) {
		var ret DevicesSerialSensorRelationshipsLivestream1
		return ret
	}
	return *o.Livestream
}

// GetLivestreamOk returns a tuple with the Livestream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject15) GetLivestreamOk() (*DevicesSerialSensorRelationshipsLivestream1, bool) {
	if o == nil || isNil(o.Livestream) {
    return nil, false
	}
	return o.Livestream, true
}

// HasLivestream returns a boolean if a field has been set.
func (o *InlineObject15) HasLivestream() bool {
	if o != nil && !isNil(o.Livestream) {
		return true
	}

	return false
}

// SetLivestream gets a reference to the given DevicesSerialSensorRelationshipsLivestream1 and assigns it to the Livestream field.
func (o *InlineObject15) SetLivestream(v DevicesSerialSensorRelationshipsLivestream1) {
	o.Livestream = &v
}

func (o InlineObject15) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Livestream) {
		toSerialize["livestream"] = o.Livestream
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject15 struct {
	value *InlineObject15
	isSet bool
}

func (v NullableInlineObject15) Get() *InlineObject15 {
	return v.value
}

func (v *NullableInlineObject15) Set(val *InlineObject15) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject15) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject15) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject15(val *InlineObject15) *NullableInlineObject15 {
	return &NullableInlineObject15{value: val, isSet: true}
}

func (v NullableInlineObject15) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject15) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


