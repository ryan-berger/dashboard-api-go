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

// NetworksNetworkIdWirelessSsidsNumberSchedulesRanges struct for NetworksNetworkIdWirelessSsidsNumberSchedulesRanges
type NetworksNetworkIdWirelessSsidsNumberSchedulesRanges struct {
	// Day of when the outage starts. Can be either full day name, or three letter abbreviation.
	StartDay string `json:"startDay"`
	// 24 hour time when the outage starts.
	StartTime string `json:"startTime"`
	// Day of when the outage ends. Can be either full day name, or three letter abbreviation
	EndDay string `json:"endDay"`
	// 24 hour time when the outage ends.
	EndTime string `json:"endTime"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSchedulesRanges instantiates a new NetworksNetworkIdWirelessSsidsNumberSchedulesRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSchedulesRanges(startDay string, startTime string, endDay string, endTime string) *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges {
	this := NetworksNetworkIdWirelessSsidsNumberSchedulesRanges{}
	this.StartDay = startDay
	this.StartTime = startTime
	this.EndDay = endDay
	this.EndTime = endTime
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSchedulesRangesWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSchedulesRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSchedulesRangesWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges {
	this := NetworksNetworkIdWirelessSsidsNumberSchedulesRanges{}
	return &this
}

// GetStartDay returns the StartDay field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetStartDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDay
}

// GetStartDayOk returns a tuple with the StartDay field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetStartDayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StartDay, true
}

// SetStartDay sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) SetStartDay(v string) {
	o.StartDay = v
}

// GetStartTime returns the StartTime field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetStartTimeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) SetStartTime(v string) {
	o.StartTime = v
}

// GetEndDay returns the EndDay field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetEndDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDay
}

// GetEndDayOk returns a tuple with the EndDay field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetEndDayOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EndDay, true
}

// SetEndDay sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) SetEndDay(v string) {
	o.EndDay = v
}

// GetEndTime returns the EndTime field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) GetEndTimeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) SetEndTime(v string) {
	o.EndTime = v
}

func (o NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startDay"] = o.StartDay
	}
	if true {
		toSerialize["startTime"] = o.StartTime
	}
	if true {
		toSerialize["endDay"] = o.EndDay
	}
	if true {
		toSerialize["endTime"] = o.EndTime
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges struct {
	value *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges) Get() *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges) Set(val *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges(val *NetworksNetworkIdWirelessSsidsNumberSchedulesRanges) *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSchedulesRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


