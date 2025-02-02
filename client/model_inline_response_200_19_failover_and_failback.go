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

// InlineResponse20019FailoverAndFailback WAN failover and failback
type InlineResponse20019FailoverAndFailback struct {
	Immediate *InlineResponse20019FailoverAndFailbackImmediate `json:"immediate,omitempty"`
}

// NewInlineResponse20019FailoverAndFailback instantiates a new InlineResponse20019FailoverAndFailback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20019FailoverAndFailback() *InlineResponse20019FailoverAndFailback {
	this := InlineResponse20019FailoverAndFailback{}
	return &this
}

// NewInlineResponse20019FailoverAndFailbackWithDefaults instantiates a new InlineResponse20019FailoverAndFailback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20019FailoverAndFailbackWithDefaults() *InlineResponse20019FailoverAndFailback {
	this := InlineResponse20019FailoverAndFailback{}
	return &this
}

// GetImmediate returns the Immediate field value if set, zero value otherwise.
func (o *InlineResponse20019FailoverAndFailback) GetImmediate() InlineResponse20019FailoverAndFailbackImmediate {
	if o == nil || isNil(o.Immediate) {
		var ret InlineResponse20019FailoverAndFailbackImmediate
		return ret
	}
	return *o.Immediate
}

// GetImmediateOk returns a tuple with the Immediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20019FailoverAndFailback) GetImmediateOk() (*InlineResponse20019FailoverAndFailbackImmediate, bool) {
	if o == nil || isNil(o.Immediate) {
    return nil, false
	}
	return o.Immediate, true
}

// HasImmediate returns a boolean if a field has been set.
func (o *InlineResponse20019FailoverAndFailback) HasImmediate() bool {
	if o != nil && !isNil(o.Immediate) {
		return true
	}

	return false
}

// SetImmediate gets a reference to the given InlineResponse20019FailoverAndFailbackImmediate and assigns it to the Immediate field.
func (o *InlineResponse20019FailoverAndFailback) SetImmediate(v InlineResponse20019FailoverAndFailbackImmediate) {
	o.Immediate = &v
}

func (o InlineResponse20019FailoverAndFailback) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Immediate) {
		toSerialize["immediate"] = o.Immediate
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20019FailoverAndFailback struct {
	value *InlineResponse20019FailoverAndFailback
	isSet bool
}

func (v NullableInlineResponse20019FailoverAndFailback) Get() *InlineResponse20019FailoverAndFailback {
	return v.value
}

func (v *NullableInlineResponse20019FailoverAndFailback) Set(val *InlineResponse20019FailoverAndFailback) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20019FailoverAndFailback) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20019FailoverAndFailback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20019FailoverAndFailback(val *InlineResponse20019FailoverAndFailback) *NullableInlineResponse20019FailoverAndFailback {
	return &NullableInlineResponse20019FailoverAndFailback{value: val, isSet: true}
}

func (v NullableInlineResponse20019FailoverAndFailback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20019FailoverAndFailback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


