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

// InlineResponse200128Tags struct for InlineResponse200128Tags
type InlineResponse200128Tags struct {
	// The name of the tag
	Tag *string `json:"tag,omitempty"`
	// The privilege of the SAML administrator on the tag
	Access *string `json:"access,omitempty"`
}

// NewInlineResponse200128Tags instantiates a new InlineResponse200128Tags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200128Tags() *InlineResponse200128Tags {
	this := InlineResponse200128Tags{}
	return &this
}

// NewInlineResponse200128TagsWithDefaults instantiates a new InlineResponse200128Tags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200128TagsWithDefaults() *InlineResponse200128Tags {
	this := InlineResponse200128Tags{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *InlineResponse200128Tags) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128Tags) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *InlineResponse200128Tags) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *InlineResponse200128Tags) SetTag(v string) {
	o.Tag = &v
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *InlineResponse200128Tags) GetAccess() string {
	if o == nil || isNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200128Tags) GetAccessOk() (*string, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *InlineResponse200128Tags) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *InlineResponse200128Tags) SetAccess(v string) {
	o.Access = &v
}

func (o InlineResponse200128Tags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200128Tags struct {
	value *InlineResponse200128Tags
	isSet bool
}

func (v NullableInlineResponse200128Tags) Get() *InlineResponse200128Tags {
	return v.value
}

func (v *NullableInlineResponse200128Tags) Set(val *InlineResponse200128Tags) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200128Tags) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200128Tags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200128Tags(val *InlineResponse200128Tags) *NullableInlineResponse200128Tags {
	return &NullableInlineResponse200128Tags{value: val, isSet: true}
}

func (v NullableInlineResponse200128Tags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200128Tags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


