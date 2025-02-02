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

// InlineObject156 struct for InlineObject156
type InlineObject156 struct {
	// If true, the SSID device type group policies are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of device type policies.
	DeviceTypePolicies []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies `json:"deviceTypePolicies,omitempty"`
}

// NewInlineObject156 instantiates a new InlineObject156 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject156() *InlineObject156 {
	this := InlineObject156{}
	return &this
}

// NewInlineObject156WithDefaults instantiates a new InlineObject156 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject156WithDefaults() *InlineObject156 {
	this := InlineObject156{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject156) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject156) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject156) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeviceTypePolicies returns the DeviceTypePolicies field value if set, zero value otherwise.
func (o *InlineObject156) GetDeviceTypePolicies() []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies {
	if o == nil || isNil(o.DeviceTypePolicies) {
		var ret []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies
		return ret
	}
	return o.DeviceTypePolicies
}

// GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject156) GetDeviceTypePoliciesOk() ([]NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies, bool) {
	if o == nil || isNil(o.DeviceTypePolicies) {
    return nil, false
	}
	return o.DeviceTypePolicies, true
}

// HasDeviceTypePolicies returns a boolean if a field has been set.
func (o *InlineObject156) HasDeviceTypePolicies() bool {
	if o != nil && !isNil(o.DeviceTypePolicies) {
		return true
	}

	return false
}

// SetDeviceTypePolicies gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies and assigns it to the DeviceTypePolicies field.
func (o *InlineObject156) SetDeviceTypePolicies(v []NetworksNetworkIdWirelessSsidsNumberDeviceTypeGroupPoliciesDeviceTypePolicies) {
	o.DeviceTypePolicies = v
}

func (o InlineObject156) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.DeviceTypePolicies) {
		toSerialize["deviceTypePolicies"] = o.DeviceTypePolicies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject156 struct {
	value *InlineObject156
	isSet bool
}

func (v NullableInlineObject156) Get() *InlineObject156 {
	return v.value
}

func (v *NullableInlineObject156) Set(val *InlineObject156) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject156) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject156) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject156(val *InlineObject156) *NullableInlineObject156 {
	return &NullableInlineObject156{value: val, isSet: true}
}

func (v NullableInlineObject156) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject156) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


