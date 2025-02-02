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

// InlineResponse20041NamedVlans A hash of Named VLANs options applied to the Network.
type InlineResponse20041NamedVlans struct {
	// Enables / disables Named VLANs on the Network.
	Enabled bool `json:"enabled"`
}

// NewInlineResponse20041NamedVlans instantiates a new InlineResponse20041NamedVlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20041NamedVlans(enabled bool) *InlineResponse20041NamedVlans {
	this := InlineResponse20041NamedVlans{}
	this.Enabled = enabled
	return &this
}

// NewInlineResponse20041NamedVlansWithDefaults instantiates a new InlineResponse20041NamedVlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20041NamedVlansWithDefaults() *InlineResponse20041NamedVlans {
	this := InlineResponse20041NamedVlans{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *InlineResponse20041NamedVlans) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *InlineResponse20041NamedVlans) GetEnabledOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *InlineResponse20041NamedVlans) SetEnabled(v bool) {
	o.Enabled = v
}

func (o InlineResponse20041NamedVlans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20041NamedVlans struct {
	value *InlineResponse20041NamedVlans
	isSet bool
}

func (v NullableInlineResponse20041NamedVlans) Get() *InlineResponse20041NamedVlans {
	return v.value
}

func (v *NullableInlineResponse20041NamedVlans) Set(val *InlineResponse20041NamedVlans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20041NamedVlans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20041NamedVlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20041NamedVlans(val *InlineResponse20041NamedVlans) *NullableInlineResponse20041NamedVlans {
	return &NullableInlineResponse20041NamedVlans{value: val, isSet: true}
}

func (v NullableInlineResponse20041NamedVlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20041NamedVlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


