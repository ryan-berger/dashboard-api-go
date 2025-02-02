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

// InlineResponse20015DynamicDns Dynamic DNS settings for a network
type InlineResponse20015DynamicDns struct {
	// Dynamic DNS enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Dynamic DNS url prefix. DDNS must be enabled to update
	Prefix *string `json:"prefix,omitempty"`
	// Dynamic DNS url. DDNS must be enabled to update
	Url *string `json:"url,omitempty"`
}

// NewInlineResponse20015DynamicDns instantiates a new InlineResponse20015DynamicDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20015DynamicDns() *InlineResponse20015DynamicDns {
	this := InlineResponse20015DynamicDns{}
	return &this
}

// NewInlineResponse20015DynamicDnsWithDefaults instantiates a new InlineResponse20015DynamicDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20015DynamicDnsWithDefaults() *InlineResponse20015DynamicDns {
	this := InlineResponse20015DynamicDns{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20015DynamicDns) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DynamicDns) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20015DynamicDns) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20015DynamicDns) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InlineResponse20015DynamicDns) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DynamicDns) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InlineResponse20015DynamicDns) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InlineResponse20015DynamicDns) SetPrefix(v string) {
	o.Prefix = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20015DynamicDns) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20015DynamicDns) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20015DynamicDns) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20015DynamicDns) SetUrl(v string) {
	o.Url = &v
}

func (o InlineResponse20015DynamicDns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20015DynamicDns struct {
	value *InlineResponse20015DynamicDns
	isSet bool
}

func (v NullableInlineResponse20015DynamicDns) Get() *InlineResponse20015DynamicDns {
	return v.value
}

func (v *NullableInlineResponse20015DynamicDns) Set(val *InlineResponse20015DynamicDns) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20015DynamicDns) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20015DynamicDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20015DynamicDns(val *InlineResponse20015DynamicDns) *NullableInlineResponse20015DynamicDns {
	return &NullableInlineResponse20015DynamicDns{value: val, isSet: true}
}

func (v NullableInlineResponse20015DynamicDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20015DynamicDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


