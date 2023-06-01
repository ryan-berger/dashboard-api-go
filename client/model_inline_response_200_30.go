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

// InlineResponse20030 struct for InlineResponse20030
type InlineResponse20030 struct {
	// Id of staged upgrade group
	GroupId *string `json:"groupId,omitempty"`
	// Name of the Staged Upgrade Group
	Name *string `json:"name,omitempty"`
	// Description of the Staged Upgrade Group
	Description *string `json:"description,omitempty"`
	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault *bool `json:"isDefault,omitempty"`
	AssignedDevices *NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices `json:"assignedDevices,omitempty"`
}

// NewInlineResponse20030 instantiates a new InlineResponse20030 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20030() *InlineResponse20030 {
	this := InlineResponse20030{}
	return &this
}

// NewInlineResponse20030WithDefaults instantiates a new InlineResponse20030 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20030WithDefaults() *InlineResponse20030 {
	this := InlineResponse20030{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *InlineResponse20030) GetGroupId() string {
	if o == nil || isNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetGroupIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupId) {
    return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineResponse20030) HasGroupId() bool {
	if o != nil && !isNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *InlineResponse20030) SetGroupId(v string) {
	o.GroupId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20030) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20030) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20030) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20030) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20030) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20030) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *InlineResponse20030) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *InlineResponse20030) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *InlineResponse20030) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetAssignedDevices returns the AssignedDevices field value if set, zero value otherwise.
func (o *InlineResponse20030) GetAssignedDevices() NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices {
	if o == nil || isNil(o.AssignedDevices) {
		var ret NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices
		return ret
	}
	return *o.AssignedDevices
}

// GetAssignedDevicesOk returns a tuple with the AssignedDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20030) GetAssignedDevicesOk() (*NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices, bool) {
	if o == nil || isNil(o.AssignedDevices) {
    return nil, false
	}
	return o.AssignedDevices, true
}

// HasAssignedDevices returns a boolean if a field has been set.
func (o *InlineResponse20030) HasAssignedDevices() bool {
	if o != nil && !isNil(o.AssignedDevices) {
		return true
	}

	return false
}

// SetAssignedDevices gets a reference to the given NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices and assigns it to the AssignedDevices field.
func (o *InlineResponse20030) SetAssignedDevices(v NetworksNetworkIdFirmwareUpgradesStagedGroupsAssignedDevices) {
	o.AssignedDevices = &v
}

func (o InlineResponse20030) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.AssignedDevices) {
		toSerialize["assignedDevices"] = o.AssignedDevices
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20030 struct {
	value *InlineResponse20030
	isSet bool
}

func (v NullableInlineResponse20030) Get() *InlineResponse20030 {
	return v.value
}

func (v *NullableInlineResponse20030) Set(val *InlineResponse20030) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20030) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20030) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20030(val *InlineResponse20030) *NullableInlineResponse20030 {
	return &NullableInlineResponse20030{value: val, isSet: true}
}

func (v NullableInlineResponse20030) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20030) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


