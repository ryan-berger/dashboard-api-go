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

// InlineResponse200130 struct for InlineResponse200130
type InlineResponse200130 struct {
	// Serial number of the sensor that took the readings.
	Serial *string `json:"serial,omitempty"`
	Network *OrganizationsOrganizationIdSensorReadingsHistoryNetwork `json:"network,omitempty"`
	// Array of latest readings from the sensor. Each object represents a single reading for a single metric.
	Readings []OrganizationsOrganizationIdSensorReadingsLatestReadings `json:"readings,omitempty"`
}

// NewInlineResponse200130 instantiates a new InlineResponse200130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200130() *InlineResponse200130 {
	this := InlineResponse200130{}
	return &this
}

// NewInlineResponse200130WithDefaults instantiates a new InlineResponse200130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200130WithDefaults() *InlineResponse200130 {
	this := InlineResponse200130{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200130) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200130) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200130) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200130) SetSerial(v string) {
	o.Serial = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200130) GetNetwork() OrganizationsOrganizationIdSensorReadingsHistoryNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSensorReadingsHistoryNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200130) GetNetworkOk() (*OrganizationsOrganizationIdSensorReadingsHistoryNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200130) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSensorReadingsHistoryNetwork and assigns it to the Network field.
func (o *InlineResponse200130) SetNetwork(v OrganizationsOrganizationIdSensorReadingsHistoryNetwork) {
	o.Network = &v
}

// GetReadings returns the Readings field value if set, zero value otherwise.
func (o *InlineResponse200130) GetReadings() []OrganizationsOrganizationIdSensorReadingsLatestReadings {
	if o == nil || isNil(o.Readings) {
		var ret []OrganizationsOrganizationIdSensorReadingsLatestReadings
		return ret
	}
	return o.Readings
}

// GetReadingsOk returns a tuple with the Readings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200130) GetReadingsOk() ([]OrganizationsOrganizationIdSensorReadingsLatestReadings, bool) {
	if o == nil || isNil(o.Readings) {
    return nil, false
	}
	return o.Readings, true
}

// HasReadings returns a boolean if a field has been set.
func (o *InlineResponse200130) HasReadings() bool {
	if o != nil && !isNil(o.Readings) {
		return true
	}

	return false
}

// SetReadings gets a reference to the given []OrganizationsOrganizationIdSensorReadingsLatestReadings and assigns it to the Readings field.
func (o *InlineResponse200130) SetReadings(v []OrganizationsOrganizationIdSensorReadingsLatestReadings) {
	o.Readings = v
}

func (o InlineResponse200130) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Readings) {
		toSerialize["readings"] = o.Readings
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200130 struct {
	value *InlineResponse200130
	isSet bool
}

func (v NullableInlineResponse200130) Get() *InlineResponse200130 {
	return v.value
}

func (v *NullableInlineResponse200130) Set(val *InlineResponse200130) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200130) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200130(val *InlineResponse200130) *NullableInlineResponse200130 {
	return &NullableInlineResponse200130{value: val, isSet: true}
}

func (v NullableInlineResponse200130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


