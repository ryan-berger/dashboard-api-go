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

// InlineObject46 struct for InlineObject46
type InlineObject46 struct {
	// The subnet of the single LAN configuration
	Subnet *string `json:"subnet,omitempty"`
	// The appliance IP address of the single LAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
	Ipv6 *NetworksNetworkIdApplianceSingleLanIpv6 `json:"ipv6,omitempty"`
	MandatoryDhcp *NetworksNetworkIdApplianceSingleLanMandatoryDhcp `json:"mandatoryDhcp,omitempty"`
}

// NewInlineObject46 instantiates a new InlineObject46 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject46() *InlineObject46 {
	this := InlineObject46{}
	return &this
}

// NewInlineObject46WithDefaults instantiates a new InlineObject46 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject46WithDefaults() *InlineObject46 {
	this := InlineObject46{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject46) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject46) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject46) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *InlineObject46) GetApplianceIp() string {
	if o == nil || isNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetApplianceIpOk() (*string, bool) {
	if o == nil || isNil(o.ApplianceIp) {
    return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *InlineObject46) HasApplianceIp() bool {
	if o != nil && !isNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *InlineObject46) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineObject46) GetIpv6() NetworksNetworkIdApplianceSingleLanIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret NetworksNetworkIdApplianceSingleLanIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetIpv6Ok() (*NetworksNetworkIdApplianceSingleLanIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineObject46) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NetworksNetworkIdApplianceSingleLanIpv6 and assigns it to the Ipv6 field.
func (o *InlineObject46) SetIpv6(v NetworksNetworkIdApplianceSingleLanIpv6) {
	o.Ipv6 = &v
}

// GetMandatoryDhcp returns the MandatoryDhcp field value if set, zero value otherwise.
func (o *InlineObject46) GetMandatoryDhcp() NetworksNetworkIdApplianceSingleLanMandatoryDhcp {
	if o == nil || isNil(o.MandatoryDhcp) {
		var ret NetworksNetworkIdApplianceSingleLanMandatoryDhcp
		return ret
	}
	return *o.MandatoryDhcp
}

// GetMandatoryDhcpOk returns a tuple with the MandatoryDhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject46) GetMandatoryDhcpOk() (*NetworksNetworkIdApplianceSingleLanMandatoryDhcp, bool) {
	if o == nil || isNil(o.MandatoryDhcp) {
    return nil, false
	}
	return o.MandatoryDhcp, true
}

// HasMandatoryDhcp returns a boolean if a field has been set.
func (o *InlineObject46) HasMandatoryDhcp() bool {
	if o != nil && !isNil(o.MandatoryDhcp) {
		return true
	}

	return false
}

// SetMandatoryDhcp gets a reference to the given NetworksNetworkIdApplianceSingleLanMandatoryDhcp and assigns it to the MandatoryDhcp field.
func (o *InlineObject46) SetMandatoryDhcp(v NetworksNetworkIdApplianceSingleLanMandatoryDhcp) {
	o.MandatoryDhcp = &v
}

func (o InlineObject46) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !isNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	if !isNil(o.MandatoryDhcp) {
		toSerialize["mandatoryDhcp"] = o.MandatoryDhcp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject46 struct {
	value *InlineObject46
	isSet bool
}

func (v NullableInlineObject46) Get() *InlineObject46 {
	return v.value
}

func (v *NullableInlineObject46) Set(val *InlineObject46) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject46) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject46) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject46(val *InlineObject46) *NullableInlineObject46 {
	return &NullableInlineObject46{value: val, isSet: true}
}

func (v NullableInlineObject46) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject46) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


