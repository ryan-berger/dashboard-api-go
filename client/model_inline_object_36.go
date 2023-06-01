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

// InlineObject36 struct for InlineObject36
type InlineObject36 struct {
	// An array of 1:Many nat rules
	Rules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules `json:"rules"`
}

// NewInlineObject36 instantiates a new InlineObject36 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject36(rules []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) *InlineObject36 {
	this := InlineObject36{}
	this.Rules = rules
	return &this
}

// NewInlineObject36WithDefaults instantiates a new InlineObject36 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject36WithDefaults() *InlineObject36 {
	this := InlineObject36{}
	return &this
}

// GetRules returns the Rules field value
func (o *InlineObject36) GetRules() []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules {
	if o == nil {
		var ret []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject36) GetRulesOk() ([]NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject36) SetRules(v []NetworksNetworkIdApplianceFirewallOneToManyNatRulesRules) {
	o.Rules = v
}

func (o InlineObject36) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject36 struct {
	value *InlineObject36
	isSet bool
}

func (v NullableInlineObject36) Get() *InlineObject36 {
	return v.value
}

func (v *NullableInlineObject36) Set(val *InlineObject36) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject36) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject36) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject36(val *InlineObject36) *NullableInlineObject36 {
	return &NullableInlineObject36{value: val, isSet: true}
}

func (v NullableInlineObject36) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject36) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


