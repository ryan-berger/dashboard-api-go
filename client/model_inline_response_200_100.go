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

// InlineResponse200100 struct for InlineResponse200100
type InlineResponse200100 struct {
	Usage *InlineResponse200100Usage `json:"usage,omitempty"`
	Counts *InlineResponse200100Counts `json:"counts,omitempty"`
}

// NewInlineResponse200100 instantiates a new InlineResponse200100 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200100() *InlineResponse200100 {
	this := InlineResponse200100{}
	return &this
}

// NewInlineResponse200100WithDefaults instantiates a new InlineResponse200100 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200100WithDefaults() *InlineResponse200100 {
	this := InlineResponse200100{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *InlineResponse200100) GetUsage() InlineResponse200100Usage {
	if o == nil || isNil(o.Usage) {
		var ret InlineResponse200100Usage
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetUsageOk() (*InlineResponse200100Usage, bool) {
	if o == nil || isNil(o.Usage) {
    return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *InlineResponse200100) HasUsage() bool {
	if o != nil && !isNil(o.Usage) {
		return true
	}

	return false
}

// SetUsage gets a reference to the given InlineResponse200100Usage and assigns it to the Usage field.
func (o *InlineResponse200100) SetUsage(v InlineResponse200100Usage) {
	o.Usage = &v
}

// GetCounts returns the Counts field value if set, zero value otherwise.
func (o *InlineResponse200100) GetCounts() InlineResponse200100Counts {
	if o == nil || isNil(o.Counts) {
		var ret InlineResponse200100Counts
		return ret
	}
	return *o.Counts
}

// GetCountsOk returns a tuple with the Counts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200100) GetCountsOk() (*InlineResponse200100Counts, bool) {
	if o == nil || isNil(o.Counts) {
    return nil, false
	}
	return o.Counts, true
}

// HasCounts returns a boolean if a field has been set.
func (o *InlineResponse200100) HasCounts() bool {
	if o != nil && !isNil(o.Counts) {
		return true
	}

	return false
}

// SetCounts gets a reference to the given InlineResponse200100Counts and assigns it to the Counts field.
func (o *InlineResponse200100) SetCounts(v InlineResponse200100Counts) {
	o.Counts = &v
}

func (o InlineResponse200100) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Usage) {
		toSerialize["usage"] = o.Usage
	}
	if !isNil(o.Counts) {
		toSerialize["counts"] = o.Counts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200100 struct {
	value *InlineResponse200100
	isSet bool
}

func (v NullableInlineResponse200100) Get() *InlineResponse200100 {
	return v.value
}

func (v *NullableInlineResponse200100) Set(val *InlineResponse200100) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200100) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200100) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200100(val *InlineResponse200100) *NullableInlineResponse200100 {
	return &NullableInlineResponse200100{value: val, isSet: true}
}

func (v NullableInlineResponse200100) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200100) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


