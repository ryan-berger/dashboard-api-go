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

// OrganizationsOrganizationIdSwitchPortsBySwitchNetwork Identifying information of the switch's network.
type OrganizationsOrganizationIdSwitchPortsBySwitchNetwork struct {
	// The name of the network.
	Name *string `json:"name,omitempty"`
	// The ID of the network.
	Id *string `json:"id,omitempty"`
}

// NewOrganizationsOrganizationIdSwitchPortsBySwitchNetwork instantiates a new OrganizationsOrganizationIdSwitchPortsBySwitchNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSwitchPortsBySwitchNetwork() *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork {
	this := OrganizationsOrganizationIdSwitchPortsBySwitchNetwork{}
	return &this
}

// NewOrganizationsOrganizationIdSwitchPortsBySwitchNetworkWithDefaults instantiates a new OrganizationsOrganizationIdSwitchPortsBySwitchNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSwitchPortsBySwitchNetworkWithDefaults() *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork {
	this := OrganizationsOrganizationIdSwitchPortsBySwitchNetwork{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) SetId(v string) {
	o.Id = &v
}

func (o OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork struct {
	value *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork) Get() *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork) Set(val *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork(val *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) *NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork {
	return &NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSwitchPortsBySwitchNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


