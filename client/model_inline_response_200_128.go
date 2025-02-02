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

// InlineResponse200128 struct for InlineResponse200128
type InlineResponse200128 struct {
	// ID associated with the SAML role
	Id *string `json:"id,omitempty"`
	// The role of the SAML administrator
	Role *string `json:"role,omitempty"`
	// The privilege of the SAML administrator on the organization
	OrgAccess *string `json:"orgAccess,omitempty"`
	// The list of networks that the SAML administrator has privileges on
	Networks []InlineResponse200128Networks `json:"networks,omitempty"`
	// The list of tags that the SAML administrator has privleges on
	Tags []InlineResponse200128Tags `json:"tags,omitempty"`
}

// NewInlineResponse200128 instantiates a new InlineResponse200128 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200128() *InlineResponse200128 {
	this := InlineResponse200128{}
	return &this
}

// NewInlineResponse200128WithDefaults instantiates a new InlineResponse200128 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200128WithDefaults() *InlineResponse200128 {
	this := InlineResponse200128{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200128) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200128) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200128) SetId(v string) {
	o.Id = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *InlineResponse200128) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *InlineResponse200128) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *InlineResponse200128) SetRole(v string) {
	o.Role = &v
}

// GetOrgAccess returns the OrgAccess field value if set, zero value otherwise.
func (o *InlineResponse200128) GetOrgAccess() string {
	if o == nil || isNil(o.OrgAccess) {
		var ret string
		return ret
	}
	return *o.OrgAccess
}

// GetOrgAccessOk returns a tuple with the OrgAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetOrgAccessOk() (*string, bool) {
	if o == nil || isNil(o.OrgAccess) {
    return nil, false
	}
	return o.OrgAccess, true
}

// HasOrgAccess returns a boolean if a field has been set.
func (o *InlineResponse200128) HasOrgAccess() bool {
	if o != nil && !isNil(o.OrgAccess) {
		return true
	}

	return false
}

// SetOrgAccess gets a reference to the given string and assigns it to the OrgAccess field.
func (o *InlineResponse200128) SetOrgAccess(v string) {
	o.OrgAccess = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *InlineResponse200128) GetNetworks() []InlineResponse200128Networks {
	if o == nil || isNil(o.Networks) {
		var ret []InlineResponse200128Networks
		return ret
	}
	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetNetworksOk() ([]InlineResponse200128Networks, bool) {
	if o == nil || isNil(o.Networks) {
    return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *InlineResponse200128) HasNetworks() bool {
	if o != nil && !isNil(o.Networks) {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []InlineResponse200128Networks and assigns it to the Networks field.
func (o *InlineResponse200128) SetNetworks(v []InlineResponse200128Networks) {
	o.Networks = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200128) GetTags() []InlineResponse200128Tags {
	if o == nil || isNil(o.Tags) {
		var ret []InlineResponse200128Tags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128) GetTagsOk() ([]InlineResponse200128Tags, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200128) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []InlineResponse200128Tags and assigns it to the Tags field.
func (o *InlineResponse200128) SetTags(v []InlineResponse200128Tags) {
	o.Tags = v
}

func (o InlineResponse200128) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.OrgAccess) {
		toSerialize["orgAccess"] = o.OrgAccess
	}
	if !isNil(o.Networks) {
		toSerialize["networks"] = o.Networks
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200128 struct {
	value *InlineResponse200128
	isSet bool
}

func (v NullableInlineResponse200128) Get() *InlineResponse200128 {
	return v.value
}

func (v *NullableInlineResponse200128) Set(val *InlineResponse200128) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200128) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200128) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200128(val *InlineResponse200128) *NullableInlineResponse200128 {
	return &NullableInlineResponse200128{value: val, isSet: true}
}

func (v NullableInlineResponse200128) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200128) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


