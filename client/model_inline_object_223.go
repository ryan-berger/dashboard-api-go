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

// InlineObject223 struct for InlineObject223
type InlineObject223 struct {
	// Serial number of the source switch (must be on a network not bound to a template)
	SourceSerial string `json:"sourceSerial"`
	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	TargetSerials []string `json:"targetSerials"`
}

// NewInlineObject223 instantiates a new InlineObject223 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject223(sourceSerial string, targetSerials []string) *InlineObject223 {
	this := InlineObject223{}
	this.SourceSerial = sourceSerial
	this.TargetSerials = targetSerials
	return &this
}

// NewInlineObject223WithDefaults instantiates a new InlineObject223 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject223WithDefaults() *InlineObject223 {
	this := InlineObject223{}
	return &this
}

// GetSourceSerial returns the SourceSerial field value
func (o *InlineObject223) GetSourceSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceSerial
}

// GetSourceSerialOk returns a tuple with the SourceSerial field value
// and a boolean to check if the value has been set.
func (o *InlineObject223) GetSourceSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SourceSerial, true
}

// SetSourceSerial sets field value
func (o *InlineObject223) SetSourceSerial(v string) {
	o.SourceSerial = v
}

// GetTargetSerials returns the TargetSerials field value
func (o *InlineObject223) GetTargetSerials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TargetSerials
}

// GetTargetSerialsOk returns a tuple with the TargetSerials field value
// and a boolean to check if the value has been set.
func (o *InlineObject223) GetTargetSerialsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.TargetSerials, true
}

// SetTargetSerials sets field value
func (o *InlineObject223) SetTargetSerials(v []string) {
	o.TargetSerials = v
}

func (o InlineObject223) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceSerial"] = o.SourceSerial
	}
	if true {
		toSerialize["targetSerials"] = o.TargetSerials
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject223 struct {
	value *InlineObject223
	isSet bool
}

func (v NullableInlineObject223) Get() *InlineObject223 {
	return v.value
}

func (v *NullableInlineObject223) Set(val *InlineObject223) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject223) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject223) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject223(val *InlineObject223) *NullableInlineObject223 {
	return &NullableInlineObject223{value: val, isSet: true}
}

func (v NullableInlineObject223) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject223) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


