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

// NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 Splash authorization for SSID 13
type NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 struct {
	// New authorization status for the SSID (true, false).
	IsAuthorized *bool `json:"isAuthorized,omitempty"`
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13{}
	return &this
}

// NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13WithDefaults instantiates a new NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13WithDefaults() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 {
	this := NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13{}
	return &this
}

// GetIsAuthorized returns the IsAuthorized field value if set, zero value otherwise.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) GetIsAuthorized() bool {
	if o == nil || isNil(o.IsAuthorized) {
		var ret bool
		return ret
	}
	return *o.IsAuthorized
}

// GetIsAuthorizedOk returns a tuple with the IsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) GetIsAuthorizedOk() (*bool, bool) {
	if o == nil || isNil(o.IsAuthorized) {
    return nil, false
	}
	return o.IsAuthorized, true
}

// HasIsAuthorized returns a boolean if a field has been set.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) HasIsAuthorized() bool {
	if o != nil && !isNil(o.IsAuthorized) {
		return true
	}

	return false
}

// SetIsAuthorized gets a reference to the given bool and assigns it to the IsAuthorized field.
func (o *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) SetIsAuthorized(v bool) {
	o.IsAuthorized = &v
}

func (o NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IsAuthorized) {
		toSerialize["isAuthorized"] = o.IsAuthorized
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 struct {
	value *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13
	isSet bool
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) Get() *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 {
	return v.value
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) Set(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13(val *NetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13 {
	return &NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdClientsClientIdSplashAuthorizationStatusSsids13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


