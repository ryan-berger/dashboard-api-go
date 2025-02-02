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

// InlineResponse20021Subnets struct for InlineResponse20021Subnets
type InlineResponse20021Subnets struct {
	// The CIDR notation subnet used within the VPN
	LocalSubnet *string `json:"localSubnet,omitempty"`
	// Indicates the presence of the subnet in the VPN
	UseVpn *bool `json:"useVpn,omitempty"`
}

// NewInlineResponse20021Subnets instantiates a new InlineResponse20021Subnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20021Subnets() *InlineResponse20021Subnets {
	this := InlineResponse20021Subnets{}
	return &this
}

// NewInlineResponse20021SubnetsWithDefaults instantiates a new InlineResponse20021Subnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20021SubnetsWithDefaults() *InlineResponse20021Subnets {
	this := InlineResponse20021Subnets{}
	return &this
}

// GetLocalSubnet returns the LocalSubnet field value if set, zero value otherwise.
func (o *InlineResponse20021Subnets) GetLocalSubnet() string {
	if o == nil || isNil(o.LocalSubnet) {
		var ret string
		return ret
	}
	return *o.LocalSubnet
}

// GetLocalSubnetOk returns a tuple with the LocalSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Subnets) GetLocalSubnetOk() (*string, bool) {
	if o == nil || isNil(o.LocalSubnet) {
    return nil, false
	}
	return o.LocalSubnet, true
}

// HasLocalSubnet returns a boolean if a field has been set.
func (o *InlineResponse20021Subnets) HasLocalSubnet() bool {
	if o != nil && !isNil(o.LocalSubnet) {
		return true
	}

	return false
}

// SetLocalSubnet gets a reference to the given string and assigns it to the LocalSubnet field.
func (o *InlineResponse20021Subnets) SetLocalSubnet(v string) {
	o.LocalSubnet = &v
}

// GetUseVpn returns the UseVpn field value if set, zero value otherwise.
func (o *InlineResponse20021Subnets) GetUseVpn() bool {
	if o == nil || isNil(o.UseVpn) {
		var ret bool
		return ret
	}
	return *o.UseVpn
}

// GetUseVpnOk returns a tuple with the UseVpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Subnets) GetUseVpnOk() (*bool, bool) {
	if o == nil || isNil(o.UseVpn) {
    return nil, false
	}
	return o.UseVpn, true
}

// HasUseVpn returns a boolean if a field has been set.
func (o *InlineResponse20021Subnets) HasUseVpn() bool {
	if o != nil && !isNil(o.UseVpn) {
		return true
	}

	return false
}

// SetUseVpn gets a reference to the given bool and assigns it to the UseVpn field.
func (o *InlineResponse20021Subnets) SetUseVpn(v bool) {
	o.UseVpn = &v
}

func (o InlineResponse20021Subnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.LocalSubnet) {
		toSerialize["localSubnet"] = o.LocalSubnet
	}
	if !isNil(o.UseVpn) {
		toSerialize["useVpn"] = o.UseVpn
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20021Subnets struct {
	value *InlineResponse20021Subnets
	isSet bool
}

func (v NullableInlineResponse20021Subnets) Get() *InlineResponse20021Subnets {
	return v.value
}

func (v *NullableInlineResponse20021Subnets) Set(val *InlineResponse20021Subnets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20021Subnets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20021Subnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20021Subnets(val *InlineResponse20021Subnets) *NullableInlineResponse20021Subnets {
	return &NullableInlineResponse20021Subnets{value: val, isSet: true}
}

func (v NullableInlineResponse20021Subnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20021Subnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


