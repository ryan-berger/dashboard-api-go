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

// NetworksNetworkIdMqttBrokersSecuritySecurity TLS settings of the MQTT broker.
type NetworksNetworkIdMqttBrokersSecuritySecurity struct {
	// CA Certificate of the MQTT broker.
	CaCertificate *string `json:"caCertificate,omitempty"`
	// Whether the TLS hostname verification is enabled for the MQTT broker.
	VerifyHostnames *bool `json:"verifyHostnames,omitempty"`
}

// NewNetworksNetworkIdMqttBrokersSecuritySecurity instantiates a new NetworksNetworkIdMqttBrokersSecuritySecurity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdMqttBrokersSecuritySecurity() *NetworksNetworkIdMqttBrokersSecuritySecurity {
	this := NetworksNetworkIdMqttBrokersSecuritySecurity{}
	return &this
}

// NewNetworksNetworkIdMqttBrokersSecuritySecurityWithDefaults instantiates a new NetworksNetworkIdMqttBrokersSecuritySecurity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdMqttBrokersSecuritySecurityWithDefaults() *NetworksNetworkIdMqttBrokersSecuritySecurity {
	this := NetworksNetworkIdMqttBrokersSecuritySecurity{}
	return &this
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) GetCaCertificate() string {
	if o == nil || isNil(o.CaCertificate) {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) GetCaCertificateOk() (*string, bool) {
	if o == nil || isNil(o.CaCertificate) {
    return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) HasCaCertificate() bool {
	if o != nil && !isNil(o.CaCertificate) {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetVerifyHostnames returns the VerifyHostnames field value if set, zero value otherwise.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) GetVerifyHostnames() bool {
	if o == nil || isNil(o.VerifyHostnames) {
		var ret bool
		return ret
	}
	return *o.VerifyHostnames
}

// GetVerifyHostnamesOk returns a tuple with the VerifyHostnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) GetVerifyHostnamesOk() (*bool, bool) {
	if o == nil || isNil(o.VerifyHostnames) {
    return nil, false
	}
	return o.VerifyHostnames, true
}

// HasVerifyHostnames returns a boolean if a field has been set.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) HasVerifyHostnames() bool {
	if o != nil && !isNil(o.VerifyHostnames) {
		return true
	}

	return false
}

// SetVerifyHostnames gets a reference to the given bool and assigns it to the VerifyHostnames field.
func (o *NetworksNetworkIdMqttBrokersSecuritySecurity) SetVerifyHostnames(v bool) {
	o.VerifyHostnames = &v
}

func (o NetworksNetworkIdMqttBrokersSecuritySecurity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CaCertificate) {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	if !isNil(o.VerifyHostnames) {
		toSerialize["verifyHostnames"] = o.VerifyHostnames
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdMqttBrokersSecuritySecurity struct {
	value *NetworksNetworkIdMqttBrokersSecuritySecurity
	isSet bool
}

func (v NullableNetworksNetworkIdMqttBrokersSecuritySecurity) Get() *NetworksNetworkIdMqttBrokersSecuritySecurity {
	return v.value
}

func (v *NullableNetworksNetworkIdMqttBrokersSecuritySecurity) Set(val *NetworksNetworkIdMqttBrokersSecuritySecurity) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdMqttBrokersSecuritySecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdMqttBrokersSecuritySecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdMqttBrokersSecuritySecurity(val *NetworksNetworkIdMqttBrokersSecuritySecurity) *NullableNetworksNetworkIdMqttBrokersSecuritySecurity {
	return &NullableNetworksNetworkIdMqttBrokersSecuritySecurity{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdMqttBrokersSecuritySecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdMqttBrokersSecuritySecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


