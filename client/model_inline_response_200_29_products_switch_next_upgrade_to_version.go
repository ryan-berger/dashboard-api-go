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

// InlineResponse20029ProductsSwitchNextUpgradeToVersion Details of the version the device will upgrade to
type InlineResponse20029ProductsSwitchNextUpgradeToVersion struct {
	// Id of the Version being upgraded to
	Id *string `json:"id,omitempty"`
	// Firmware version short name
	ShortName *string `json:"shortName,omitempty"`
}

// NewInlineResponse20029ProductsSwitchNextUpgradeToVersion instantiates a new InlineResponse20029ProductsSwitchNextUpgradeToVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20029ProductsSwitchNextUpgradeToVersion() *InlineResponse20029ProductsSwitchNextUpgradeToVersion {
	this := InlineResponse20029ProductsSwitchNextUpgradeToVersion{}
	return &this
}

// NewInlineResponse20029ProductsSwitchNextUpgradeToVersionWithDefaults instantiates a new InlineResponse20029ProductsSwitchNextUpgradeToVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20029ProductsSwitchNextUpgradeToVersionWithDefaults() *InlineResponse20029ProductsSwitchNextUpgradeToVersion {
	this := InlineResponse20029ProductsSwitchNextUpgradeToVersion{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse20029ProductsSwitchNextUpgradeToVersion) SetShortName(v string) {
	o.ShortName = &v
}

func (o InlineResponse20029ProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion struct {
	value *InlineResponse20029ProductsSwitchNextUpgradeToVersion
	isSet bool
}

func (v NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion) Get() *InlineResponse20029ProductsSwitchNextUpgradeToVersion {
	return v.value
}

func (v *NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion) Set(val *InlineResponse20029ProductsSwitchNextUpgradeToVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20029ProductsSwitchNextUpgradeToVersion(val *InlineResponse20029ProductsSwitchNextUpgradeToVersion) *NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion {
	return &NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion{value: val, isSet: true}
}

func (v NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20029ProductsSwitchNextUpgradeToVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


