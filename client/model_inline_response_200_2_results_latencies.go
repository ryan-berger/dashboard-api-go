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

// InlineResponse2002ResultsLatencies Packet latency stats
type InlineResponse2002ResultsLatencies struct {
	// Minimum latency
	Minimum *float32 `json:"minimum,omitempty"`
	// Average latency
	Average *float32 `json:"average,omitempty"`
	// Maximum latency
	Maximum *float32 `json:"maximum,omitempty"`
}

// NewInlineResponse2002ResultsLatencies instantiates a new InlineResponse2002ResultsLatencies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002ResultsLatencies() *InlineResponse2002ResultsLatencies {
	this := InlineResponse2002ResultsLatencies{}
	return &this
}

// NewInlineResponse2002ResultsLatenciesWithDefaults instantiates a new InlineResponse2002ResultsLatencies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002ResultsLatenciesWithDefaults() *InlineResponse2002ResultsLatencies {
	this := InlineResponse2002ResultsLatencies{}
	return &this
}

// GetMinimum returns the Minimum field value if set, zero value otherwise.
func (o *InlineResponse2002ResultsLatencies) GetMinimum() float32 {
	if o == nil || isNil(o.Minimum) {
		var ret float32
		return ret
	}
	return *o.Minimum
}

// GetMinimumOk returns a tuple with the Minimum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ResultsLatencies) GetMinimumOk() (*float32, bool) {
	if o == nil || isNil(o.Minimum) {
    return nil, false
	}
	return o.Minimum, true
}

// HasMinimum returns a boolean if a field has been set.
func (o *InlineResponse2002ResultsLatencies) HasMinimum() bool {
	if o != nil && !isNil(o.Minimum) {
		return true
	}

	return false
}

// SetMinimum gets a reference to the given float32 and assigns it to the Minimum field.
func (o *InlineResponse2002ResultsLatencies) SetMinimum(v float32) {
	o.Minimum = &v
}

// GetAverage returns the Average field value if set, zero value otherwise.
func (o *InlineResponse2002ResultsLatencies) GetAverage() float32 {
	if o == nil || isNil(o.Average) {
		var ret float32
		return ret
	}
	return *o.Average
}

// GetAverageOk returns a tuple with the Average field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ResultsLatencies) GetAverageOk() (*float32, bool) {
	if o == nil || isNil(o.Average) {
    return nil, false
	}
	return o.Average, true
}

// HasAverage returns a boolean if a field has been set.
func (o *InlineResponse2002ResultsLatencies) HasAverage() bool {
	if o != nil && !isNil(o.Average) {
		return true
	}

	return false
}

// SetAverage gets a reference to the given float32 and assigns it to the Average field.
func (o *InlineResponse2002ResultsLatencies) SetAverage(v float32) {
	o.Average = &v
}

// GetMaximum returns the Maximum field value if set, zero value otherwise.
func (o *InlineResponse2002ResultsLatencies) GetMaximum() float32 {
	if o == nil || isNil(o.Maximum) {
		var ret float32
		return ret
	}
	return *o.Maximum
}

// GetMaximumOk returns a tuple with the Maximum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002ResultsLatencies) GetMaximumOk() (*float32, bool) {
	if o == nil || isNil(o.Maximum) {
    return nil, false
	}
	return o.Maximum, true
}

// HasMaximum returns a boolean if a field has been set.
func (o *InlineResponse2002ResultsLatencies) HasMaximum() bool {
	if o != nil && !isNil(o.Maximum) {
		return true
	}

	return false
}

// SetMaximum gets a reference to the given float32 and assigns it to the Maximum field.
func (o *InlineResponse2002ResultsLatencies) SetMaximum(v float32) {
	o.Maximum = &v
}

func (o InlineResponse2002ResultsLatencies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Minimum) {
		toSerialize["minimum"] = o.Minimum
	}
	if !isNil(o.Average) {
		toSerialize["average"] = o.Average
	}
	if !isNil(o.Maximum) {
		toSerialize["maximum"] = o.Maximum
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002ResultsLatencies struct {
	value *InlineResponse2002ResultsLatencies
	isSet bool
}

func (v NullableInlineResponse2002ResultsLatencies) Get() *InlineResponse2002ResultsLatencies {
	return v.value
}

func (v *NullableInlineResponse2002ResultsLatencies) Set(val *InlineResponse2002ResultsLatencies) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002ResultsLatencies) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002ResultsLatencies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002ResultsLatencies(val *InlineResponse2002ResultsLatencies) *NullableInlineResponse2002ResultsLatencies {
	return &NullableInlineResponse2002ResultsLatencies{value: val, isSet: true}
}

func (v NullableInlineResponse2002ResultsLatencies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002ResultsLatencies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


