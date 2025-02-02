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

// InlineObject37 struct for InlineObject37
type InlineObject37 struct {
	// An array of 1:1 nat rules
	Rules []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules `json:"rules"`
}

// NewInlineObject37 instantiates a new InlineObject37 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject37(rules []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) *InlineObject37 {
	this := InlineObject37{}
	this.Rules = rules
	return &this
}

// NewInlineObject37WithDefaults instantiates a new InlineObject37 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject37WithDefaults() *InlineObject37 {
	this := InlineObject37{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject37) GetRules() []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject37) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject37) SetRules(v []NetworksNetworkIdApplianceFirewallOneToOneNatRulesRules) {
	o.Rules = v
}

func (o InlineObject37) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject37 struct {
	value *InlineObject37
	isSet bool
}

func (v NullableInlineObject37) Get() *InlineObject37 {
	return v.value
}

func (v *NullableInlineObject37) Set(val *InlineObject37) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject37) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject37) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject37(val *InlineObject37) *NullableInlineObject37 {
	return &NullableInlineObject37{value: val, isSet: true}
}

func (v NullableInlineObject37) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject37) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


