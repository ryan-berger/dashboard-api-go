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

// InlineResponse20041SecurePort A hash of SecureConnect options applied to the Network.
type InlineResponse20041SecurePort struct {
	// Enables / disables SecureConnect on the network. Optional.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewInlineResponse20041SecurePort instantiates a new InlineResponse20041SecurePort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20041SecurePort() *InlineResponse20041SecurePort {
	this := InlineResponse20041SecurePort{}
	return &this
}

// NewInlineResponse20041SecurePortWithDefaults instantiates a new InlineResponse20041SecurePort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20041SecurePortWithDefaults() *InlineResponse20041SecurePort {
	this := InlineResponse20041SecurePort{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20041SecurePort) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20041SecurePort) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20041SecurePort) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20041SecurePort) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o InlineResponse20041SecurePort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20041SecurePort struct {
	value *InlineResponse20041SecurePort
	isSet bool
}

func (v NullableInlineResponse20041SecurePort) Get() *InlineResponse20041SecurePort {
	return v.value
}

func (v *NullableInlineResponse20041SecurePort) Set(val *InlineResponse20041SecurePort) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20041SecurePort) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20041SecurePort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20041SecurePort(val *InlineResponse20041SecurePort) *NullableInlineResponse20041SecurePort {
	return &NullableInlineResponse20041SecurePort{value: val, isSet: true}
}

func (v NullableInlineResponse20041SecurePort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20041SecurePort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


