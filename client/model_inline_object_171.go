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

// InlineObject171 struct for InlineObject171
type InlineObject171 struct {
	// Name of the adaptive policy ACL
	Name string `json:"name"`
	// Description of the adaptive policy ACL
	Description *string `json:"description,omitempty"`
	// An ordered array of the adaptive policy ACL rules.
	Rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 `json:"rules"`
	// IP version of adpative policy ACL. One of: 'any', 'ipv4' or 'ipv6'
	IpVersion string `json:"ipVersion"`
}

// NewInlineObject171 instantiates a new InlineObject171 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject171(name string, rules []OrganizationsOrganizationIdAdaptivePolicyAclsRules1, ipVersion string) *InlineObject171 {
	this := InlineObject171{}
	this.Name = name
	var description string = ""
	this.Description = &description
	this.Rules = rules
	this.IpVersion = ipVersion
	return &this
}

// NewInlineObject171WithDefaults instantiates a new InlineObject171 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject171WithDefaults() *InlineObject171 {
	this := InlineObject171{}
	var description string = ""
	this.Description = &description
	return &this
}

// GetName returns the Name field value
func (o *InlineObject171) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject171) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineObject171) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineObject171) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineObject171) SetDescription(v string) {
	o.Description = &v
}

// GetRules returns the Rules field value
func (o *InlineObject171) GetRules() []OrganizationsOrganizationIdAdaptivePolicyAclsRules1 {
	if o == nil {
		var ret []OrganizationsOrganizationIdAdaptivePolicyAclsRules1
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetRulesOk() ([]OrganizationsOrganizationIdAdaptivePolicyAclsRules1, bool) {
	if o == nil {
    return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *InlineObject171) SetRules(v []OrganizationsOrganizationIdAdaptivePolicyAclsRules1) {
	o.Rules = v
}

// GetIpVersion returns the IpVersion field value
func (o *InlineObject171) GetIpVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value
// and a boolean to check if the value has been set.
func (o *InlineObject171) GetIpVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.IpVersion, true
}

// SetIpVersion sets field value
func (o *InlineObject171) SetIpVersion(v string) {
	o.IpVersion = v
}

func (o InlineObject171) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["rules"] = o.Rules
	}
	if true {
		toSerialize["ipVersion"] = o.IpVersion
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject171 struct {
	value *InlineObject171
	isSet bool
}

func (v NullableInlineObject171) Get() *InlineObject171 {
	return v.value
}

func (v *NullableInlineObject171) Set(val *InlineObject171) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject171) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject171) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject171(val *InlineObject171) *NullableInlineObject171 {
	return &NullableInlineObject171{value: val, isSet: true}
}

func (v NullableInlineObject171) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject171) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


