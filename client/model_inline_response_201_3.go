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

// InlineResponse2013 struct for InlineResponse2013
type InlineResponse2013 struct {
	// Webhook delivery identifier
	Id *string `json:"id,omitempty"`
	// URL where the webhook was delivered
	Url *string `json:"url,omitempty"`
	// Current status of the webhook delivery
	Status *string `json:"status,omitempty"`
}

// NewInlineResponse2013 instantiates a new InlineResponse2013 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2013() *InlineResponse2013 {
	this := InlineResponse2013{}
	return &this
}

// NewInlineResponse2013WithDefaults instantiates a new InlineResponse2013 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2013WithDefaults() *InlineResponse2013 {
	this := InlineResponse2013{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2013) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2013) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2013) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse2013) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse2013) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse2013) SetUrl(v string) {
	o.Url = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2013) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2013) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2013) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse2013) SetStatus(v string) {
	o.Status = &v
}

func (o InlineResponse2013) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2013 struct {
	value *InlineResponse2013
	isSet bool
}

func (v NullableInlineResponse2013) Get() *InlineResponse2013 {
	return v.value
}

func (v *NullableInlineResponse2013) Set(val *InlineResponse2013) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2013) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2013) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2013(val *InlineResponse2013) *NullableInlineResponse2013 {
	return &NullableInlineResponse2013{value: val, isSet: true}
}

func (v NullableInlineResponse2013) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2013) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


