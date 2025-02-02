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

// InlineObject42 struct for InlineObject42
type InlineObject42 struct {
	// A static IPv6 prefix
	Prefix *string `json:"prefix,omitempty"`
	Origin *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 `json:"origin,omitempty"`
	// A name or description for the prefix
	Description *string `json:"description,omitempty"`
}

// NewInlineObject42 instantiates a new InlineObject42 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject42() *InlineObject42 {
	this := InlineObject42{}
	return &this
}

// NewInlineObject42WithDefaults instantiates a new InlineObject42 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject42WithDefaults() *InlineObject42 {
	this := InlineObject42{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InlineObject42) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject42) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InlineObject42) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InlineObject42) SetPrefix(v string) {
	o.Prefix = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *InlineObject42) GetOrigin() NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 {
	if o == nil || isNil(o.Origin) {
		var ret NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject42) GetOriginOk() (*NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1, bool) {
	if o == nil || isNil(o.Origin) {
    return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *InlineObject42) HasOrigin() bool {
	if o != nil && !isNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1 and assigns it to the Origin field.
func (o *InlineObject42) SetOrigin(v NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1) {
	o.Origin = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject42) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject42) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject42) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject42) SetDescription(v string) {
	o.Description = &v
}

func (o InlineObject42) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject42 struct {
	value *InlineObject42
	isSet bool
}

func (v NullableInlineObject42) Get() *InlineObject42 {
	return v.value
}

func (v *NullableInlineObject42) Set(val *InlineObject42) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject42) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject42) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject42(val *InlineObject42) *NullableInlineObject42 {
	return &NullableInlineObject42{value: val, isSet: true}
}

func (v NullableInlineObject42) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject42) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


