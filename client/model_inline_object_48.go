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

// InlineObject48 struct for InlineObject48
type InlineObject48 struct {
	// The name of the new static route
	Name string `json:"name"`
	// The subnet of the static route
	Subnet string `json:"subnet"`
	// The gateway IP (next hop) of the static route
	GatewayIp string `json:"gatewayIp"`
	// The gateway IP (next hop) VLAN ID of the static route
	GatewayVlanId *string `json:"gatewayVlanId,omitempty"`
}

// NewInlineObject48 instantiates a new InlineObject48 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject48(name string, subnet string, gatewayIp string) *InlineObject48 {
	this := InlineObject48{}
	this.Name = name
	this.Subnet = subnet
	this.GatewayIp = gatewayIp
	return &this
}

// NewInlineObject48WithDefaults instantiates a new InlineObject48 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject48WithDefaults() *InlineObject48 {
	this := InlineObject48{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject48) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject48) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject48) SetName(v string) {
	o.Name = v
}

// GetSubnet returns the Subnet field value
func (o *InlineObject48) GetSubnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value
// and a boolean to check if the value has been set.
func (o *InlineObject48) GetSubnetOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Subnet, true
}

// SetSubnet sets field value
func (o *InlineObject48) SetSubnet(v string) {
	o.Subnet = v
}

// GetGatewayIp returns the GatewayIp field value
func (o *InlineObject48) GetGatewayIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GatewayIp
}

// GetGatewayIpOk returns a tuple with the GatewayIp field value
// and a boolean to check if the value has been set.
func (o *InlineObject48) GetGatewayIpOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GatewayIp, true
}

// SetGatewayIp sets field value
func (o *InlineObject48) SetGatewayIp(v string) {
	o.GatewayIp = v
}

// GetGatewayVlanId returns the GatewayVlanId field value if set, zero value otherwise.
func (o *InlineObject48) GetGatewayVlanId() string {
	if o == nil || isNil(o.GatewayVlanId) {
		var ret string
		return ret
	}
	return *o.GatewayVlanId
}

// GetGatewayVlanIdOk returns a tuple with the GatewayVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject48) GetGatewayVlanIdOk() (*string, bool) {
	if o == nil || isNil(o.GatewayVlanId) {
    return nil, false
	}
	return o.GatewayVlanId, true
}

// HasGatewayVlanId returns a boolean if a field has been set.
func (o *InlineObject48) HasGatewayVlanId() bool {
	if o != nil && !isNil(o.GatewayVlanId) {
		return true
	}

	return false
}

// SetGatewayVlanId gets a reference to the given string and assigns it to the GatewayVlanId field.
func (o *InlineObject48) SetGatewayVlanId(v string) {
	o.GatewayVlanId = &v
}

func (o InlineObject48) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["subnet"] = o.Subnet
	}
	if true {
		toSerialize["gatewayIp"] = o.GatewayIp
	}
	if !isNil(o.GatewayVlanId) {
		toSerialize["gatewayVlanId"] = o.GatewayVlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject48 struct {
	value *InlineObject48
	isSet bool
}

func (v NullableInlineObject48) Get() *InlineObject48 {
	return v.value
}

func (v *NullableInlineObject48) Set(val *InlineObject48) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject48) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject48) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject48(val *InlineObject48) *NullableInlineObject48 {
	return &NullableInlineObject48{value: val, isSet: true}
}

func (v NullableInlineObject48) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject48) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


