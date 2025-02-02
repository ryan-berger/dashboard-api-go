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

// NetworksNetworkIdSwitchLinkAggregationsSwitchPorts struct for NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
type NetworksNetworkIdSwitchLinkAggregationsSwitchPorts struct {
	// Serial number of the switch.
	Serial string `json:"serial"`
	// Port identifier of switch port. For modules, the identifier is \"SlotNumber_ModuleType_PortNumber\" (Ex: \"1_8X10G_1\"), otherwise it is just the port number (Ex: \"8\").
	PortId string `json:"portId"`
}

// NewNetworksNetworkIdSwitchLinkAggregationsSwitchPorts instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchLinkAggregationsSwitchPorts(serial string, portId string) *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	this := NetworksNetworkIdSwitchLinkAggregationsSwitchPorts{}
	this.Serial = serial
	this.PortId = portId
	return &this
}

// NewNetworksNetworkIdSwitchLinkAggregationsSwitchPortsWithDefaults instantiates a new NetworksNetworkIdSwitchLinkAggregationsSwitchPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchLinkAggregationsSwitchPortsWithDefaults() *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	this := NetworksNetworkIdSwitchLinkAggregationsSwitchPorts{}
	return &this
}

// GetSerial returns the Serial field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) SetSerial(v string) {
	o.Serial = v
}

// GetPortId returns the PortId field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetPortId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) GetPortIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PortId, true
}

// SetPortId sets field value
func (o *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) SetPortId(v string) {
	o.PortId = v
}

func (o NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	if true {
		toSerialize["portId"] = o.PortId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts struct {
	value *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts) Get() *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts) Set(val *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts(val *NetworksNetworkIdSwitchLinkAggregationsSwitchPorts) *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts {
	return &NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchLinkAggregationsSwitchPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


