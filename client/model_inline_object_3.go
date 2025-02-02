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

// InlineObject3 struct for InlineObject3
type InlineObject3 struct {
	// Enable custom analytics
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the custom analytics artifact
	ArtifactId *string `json:"artifactId,omitempty"`
	// Parameters for the custom analytics workload
	Parameters []DevicesSerialCameraCustomAnalyticsParameters `json:"parameters,omitempty"`
}

// NewInlineObject3 instantiates a new InlineObject3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject3() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// NewInlineObject3WithDefaults instantiates a new InlineObject3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject3WithDefaults() *InlineObject3 {
	this := InlineObject3{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineObject3) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineObject3) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineObject3) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetArtifactId returns the ArtifactId field value if set, zero value otherwise.
func (o *InlineObject3) GetArtifactId() string {
	if o == nil || isNil(o.ArtifactId) {
		var ret string
		return ret
	}
	return *o.ArtifactId
}

// GetArtifactIdOk returns a tuple with the ArtifactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetArtifactIdOk() (*string, bool) {
	if o == nil || isNil(o.ArtifactId) {
    return nil, false
	}
	return o.ArtifactId, true
}

// HasArtifactId returns a boolean if a field has been set.
func (o *InlineObject3) HasArtifactId() bool {
	if o != nil && !isNil(o.ArtifactId) {
		return true
	}

	return false
}

// SetArtifactId gets a reference to the given string and assigns it to the ArtifactId field.
func (o *InlineObject3) SetArtifactId(v string) {
	o.ArtifactId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *InlineObject3) GetParameters() []DevicesSerialCameraCustomAnalyticsParameters {
	if o == nil || isNil(o.Parameters) {
		var ret []DevicesSerialCameraCustomAnalyticsParameters
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject3) GetParametersOk() ([]DevicesSerialCameraCustomAnalyticsParameters, bool) {
	if o == nil || isNil(o.Parameters) {
    return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *InlineObject3) HasParameters() bool {
	if o != nil && !isNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []DevicesSerialCameraCustomAnalyticsParameters and assigns it to the Parameters field.
func (o *InlineObject3) SetParameters(v []DevicesSerialCameraCustomAnalyticsParameters) {
	o.Parameters = v
}

func (o InlineObject3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.ArtifactId) {
		toSerialize["artifactId"] = o.ArtifactId
	}
	if !isNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject3 struct {
	value *InlineObject3
	isSet bool
}

func (v NullableInlineObject3) Get() *InlineObject3 {
	return v.value
}

func (v *NullableInlineObject3) Set(val *InlineObject3) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject3) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject3(val *InlineObject3) *NullableInlineObject3 {
	return &NullableInlineObject3{value: val, isSet: true}
}

func (v NullableInlineObject3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


