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

// InlineObject157 struct for InlineObject157
type InlineObject157 struct {
	// General EAP timeout in seconds.
	Timeout *int32 `json:"timeout,omitempty"`
	Identity *InlineResponse20085Identity `json:"identity,omitempty"`
	// Maximum number of general EAP retries.
	MaxRetries *int32 `json:"maxRetries,omitempty"`
	EapolKey *InlineResponse20085EapolKey `json:"eapolKey,omitempty"`
}

// NewInlineObject157 instantiates a new InlineObject157 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject157() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// NewInlineObject157WithDefaults instantiates a new InlineObject157 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject157WithDefaults() *InlineObject157 {
	this := InlineObject157{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineObject157) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineObject157) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineObject157) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *InlineObject157) GetIdentity() InlineResponse20085Identity {
	if o == nil || isNil(o.Identity) {
		var ret InlineResponse20085Identity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetIdentityOk() (*InlineResponse20085Identity, bool) {
	if o == nil || isNil(o.Identity) {
    return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *InlineObject157) HasIdentity() bool {
	if o != nil && !isNil(o.Identity) {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given InlineResponse20085Identity and assigns it to the Identity field.
func (o *InlineObject157) SetIdentity(v InlineResponse20085Identity) {
	o.Identity = &v
}

// GetMaxRetries returns the MaxRetries field value if set, zero value otherwise.
func (o *InlineObject157) GetMaxRetries() int32 {
	if o == nil || isNil(o.MaxRetries) {
		var ret int32
		return ret
	}
	return *o.MaxRetries
}

// GetMaxRetriesOk returns a tuple with the MaxRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetMaxRetriesOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRetries) {
    return nil, false
	}
	return o.MaxRetries, true
}

// HasMaxRetries returns a boolean if a field has been set.
func (o *InlineObject157) HasMaxRetries() bool {
	if o != nil && !isNil(o.MaxRetries) {
		return true
	}

	return false
}

// SetMaxRetries gets a reference to the given int32 and assigns it to the MaxRetries field.
func (o *InlineObject157) SetMaxRetries(v int32) {
	o.MaxRetries = &v
}

// GetEapolKey returns the EapolKey field value if set, zero value otherwise.
func (o *InlineObject157) GetEapolKey() InlineResponse20085EapolKey {
	if o == nil || isNil(o.EapolKey) {
		var ret InlineResponse20085EapolKey
		return ret
	}
	return *o.EapolKey
}

// GetEapolKeyOk returns a tuple with the EapolKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject157) GetEapolKeyOk() (*InlineResponse20085EapolKey, bool) {
	if o == nil || isNil(o.EapolKey) {
    return nil, false
	}
	return o.EapolKey, true
}

// HasEapolKey returns a boolean if a field has been set.
func (o *InlineObject157) HasEapolKey() bool {
	if o != nil && !isNil(o.EapolKey) {
		return true
	}

	return false
}

// SetEapolKey gets a reference to the given InlineResponse20085EapolKey and assigns it to the EapolKey field.
func (o *InlineObject157) SetEapolKey(v InlineResponse20085EapolKey) {
	o.EapolKey = &v
}

func (o InlineObject157) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !isNil(o.Identity) {
		toSerialize["identity"] = o.Identity
	}
	if !isNil(o.MaxRetries) {
		toSerialize["maxRetries"] = o.MaxRetries
	}
	if !isNil(o.EapolKey) {
		toSerialize["eapolKey"] = o.EapolKey
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject157 struct {
	value *InlineObject157
	isSet bool
}

func (v NullableInlineObject157) Get() *InlineObject157 {
	return v.value
}

func (v *NullableInlineObject157) Set(val *InlineObject157) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject157) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject157) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject157(val *InlineObject157) *NullableInlineObject157 {
	return &NullableInlineObject157{value: val, isSet: true}
}

func (v NullableInlineObject157) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject157) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


