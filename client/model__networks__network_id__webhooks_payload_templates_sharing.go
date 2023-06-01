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

// NetworksNetworkIdWebhooksPayloadTemplatesSharing Information on which entities have access to the template
type NetworksNetworkIdWebhooksPayloadTemplatesSharing struct {
	ByNetwork *NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork `json:"byNetwork,omitempty"`
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesSharing instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesSharing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWebhooksPayloadTemplatesSharing() *NetworksNetworkIdWebhooksPayloadTemplatesSharing {
	this := NetworksNetworkIdWebhooksPayloadTemplatesSharing{}
	return &this
}

// NewNetworksNetworkIdWebhooksPayloadTemplatesSharingWithDefaults instantiates a new NetworksNetworkIdWebhooksPayloadTemplatesSharing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWebhooksPayloadTemplatesSharingWithDefaults() *NetworksNetworkIdWebhooksPayloadTemplatesSharing {
	this := NetworksNetworkIdWebhooksPayloadTemplatesSharing{}
	return &this
}

// GetByNetwork returns the ByNetwork field value if set, zero value otherwise.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharing) GetByNetwork() NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork {
	if o == nil || isNil(o.ByNetwork) {
		var ret NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork
		return ret
	}
	return *o.ByNetwork
}

// GetByNetworkOk returns a tuple with the ByNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharing) GetByNetworkOk() (*NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork, bool) {
	if o == nil || isNil(o.ByNetwork) {
    return nil, false
	}
	return o.ByNetwork, true
}

// HasByNetwork returns a boolean if a field has been set.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharing) HasByNetwork() bool {
	if o != nil && !isNil(o.ByNetwork) {
		return true
	}

	return false
}

// SetByNetwork gets a reference to the given NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork and assigns it to the ByNetwork field.
func (o *NetworksNetworkIdWebhooksPayloadTemplatesSharing) SetByNetwork(v NetworksNetworkIdWebhooksPayloadTemplatesSharingByNetwork) {
	o.ByNetwork = &v
}

func (o NetworksNetworkIdWebhooksPayloadTemplatesSharing) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ByNetwork) {
		toSerialize["byNetwork"] = o.ByNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing struct {
	value *NetworksNetworkIdWebhooksPayloadTemplatesSharing
	isSet bool
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing) Get() *NetworksNetworkIdWebhooksPayloadTemplatesSharing {
	return v.value
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing) Set(val *NetworksNetworkIdWebhooksPayloadTemplatesSharing) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWebhooksPayloadTemplatesSharing(val *NetworksNetworkIdWebhooksPayloadTemplatesSharing) *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing {
	return &NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWebhooksPayloadTemplatesSharing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


