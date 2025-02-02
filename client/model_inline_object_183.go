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

// InlineObject183 struct for InlineObject183
type InlineObject183 struct {
	// The list of VPN peers
	Peers []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers `json:"peers"`
}

// NewInlineObject183 instantiates a new InlineObject183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject183(peers []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) *InlineObject183 {
	this := InlineObject183{}
	this.Peers = peers
	return &this
}

// NewInlineObject183WithDefaults instantiates a new InlineObject183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject183WithDefaults() *InlineObject183 {
	this := InlineObject183{}
	return &this
}

// GetPeers returns the Peers field value
func (o *InlineObject183) GetPeers() []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers {
	if o == nil {
		var ret []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *InlineObject183) GetPeersOk() ([]OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers, bool) {
	if o == nil {
    return nil, false
	}
	return o.Peers, true
}

// SetPeers sets field value
func (o *InlineObject183) SetPeers(v []OrganizationsOrganizationIdApplianceVpnThirdPartyVPNPeersPeers) {
	o.Peers = v
}

func (o InlineObject183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject183 struct {
	value *InlineObject183
	isSet bool
}

func (v NullableInlineObject183) Get() *InlineObject183 {
	return v.value
}

func (v *NullableInlineObject183) Set(val *InlineObject183) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject183) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject183(val *InlineObject183) *NullableInlineObject183 {
	return &NullableInlineObject183{value: val, isSet: true}
}

func (v NullableInlineObject183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


