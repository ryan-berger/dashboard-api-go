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

// InlineObject struct for InlineObject
type InlineObject struct {
	// The name of a device
	Name *string `json:"name,omitempty"`
	// The list of tags of a device
	Tags []string `json:"tags,omitempty"`
	// The latitude of a device
	Lat *float32 `json:"lat,omitempty"`
	// The longitude of a device
	Lng *float32 `json:"lng,omitempty"`
	// The address of a device
	Address *string `json:"address,omitempty"`
	// The notes for the device. String. Limited to 255 characters.
	Notes *string `json:"notes,omitempty"`
	// Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	MoveMapMarker *bool `json:"moveMapMarker,omitempty"`
	// The ID of a switch profile to bind to the device (for available switch profiles, see the 'Switch Profiles' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch profile, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	SwitchProfileId *string `json:"switchProfileId,omitempty"`
	// The floor plan to associate to this device. null disassociates the device from the floorplan.
	FloorPlanId *string `json:"floorPlanId,omitempty"`
}

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject() *InlineObject {
	this := InlineObject{}
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineObject) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineObject) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineObject) SetTags(v []string) {
	o.Tags = v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *InlineObject) GetLat() float32 {
	if o == nil || isNil(o.Lat) {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetLatOk() (*float32, bool) {
	if o == nil || isNil(o.Lat) {
    return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *InlineObject) HasLat() bool {
	if o != nil && !isNil(o.Lat) {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *InlineObject) SetLat(v float32) {
	o.Lat = &v
}

// GetLng returns the Lng field value if set, zero value otherwise.
func (o *InlineObject) GetLng() float32 {
	if o == nil || isNil(o.Lng) {
		var ret float32
		return ret
	}
	return *o.Lng
}

// GetLngOk returns a tuple with the Lng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetLngOk() (*float32, bool) {
	if o == nil || isNil(o.Lng) {
    return nil, false
	}
	return o.Lng, true
}

// HasLng returns a boolean if a field has been set.
func (o *InlineObject) HasLng() bool {
	if o != nil && !isNil(o.Lng) {
		return true
	}

	return false
}

// SetLng gets a reference to the given float32 and assigns it to the Lng field.
func (o *InlineObject) SetLng(v float32) {
	o.Lng = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineObject) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineObject) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineObject) SetAddress(v string) {
	o.Address = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InlineObject) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InlineObject) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InlineObject) SetNotes(v string) {
	o.Notes = &v
}

// GetMoveMapMarker returns the MoveMapMarker field value if set, zero value otherwise.
func (o *InlineObject) GetMoveMapMarker() bool {
	if o == nil || isNil(o.MoveMapMarker) {
		var ret bool
		return ret
	}
	return *o.MoveMapMarker
}

// GetMoveMapMarkerOk returns a tuple with the MoveMapMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetMoveMapMarkerOk() (*bool, bool) {
	if o == nil || isNil(o.MoveMapMarker) {
    return nil, false
	}
	return o.MoveMapMarker, true
}

// HasMoveMapMarker returns a boolean if a field has been set.
func (o *InlineObject) HasMoveMapMarker() bool {
	if o != nil && !isNil(o.MoveMapMarker) {
		return true
	}

	return false
}

// SetMoveMapMarker gets a reference to the given bool and assigns it to the MoveMapMarker field.
func (o *InlineObject) SetMoveMapMarker(v bool) {
	o.MoveMapMarker = &v
}

// GetSwitchProfileId returns the SwitchProfileId field value if set, zero value otherwise.
func (o *InlineObject) GetSwitchProfileId() string {
	if o == nil || isNil(o.SwitchProfileId) {
		var ret string
		return ret
	}
	return *o.SwitchProfileId
}

// GetSwitchProfileIdOk returns a tuple with the SwitchProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetSwitchProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.SwitchProfileId) {
    return nil, false
	}
	return o.SwitchProfileId, true
}

// HasSwitchProfileId returns a boolean if a field has been set.
func (o *InlineObject) HasSwitchProfileId() bool {
	if o != nil && !isNil(o.SwitchProfileId) {
		return true
	}

	return false
}

// SetSwitchProfileId gets a reference to the given string and assigns it to the SwitchProfileId field.
func (o *InlineObject) SetSwitchProfileId(v string) {
	o.SwitchProfileId = &v
}

// GetFloorPlanId returns the FloorPlanId field value if set, zero value otherwise.
func (o *InlineObject) GetFloorPlanId() string {
	if o == nil || isNil(o.FloorPlanId) {
		var ret string
		return ret
	}
	return *o.FloorPlanId
}

// GetFloorPlanIdOk returns a tuple with the FloorPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject) GetFloorPlanIdOk() (*string, bool) {
	if o == nil || isNil(o.FloorPlanId) {
    return nil, false
	}
	return o.FloorPlanId, true
}

// HasFloorPlanId returns a boolean if a field has been set.
func (o *InlineObject) HasFloorPlanId() bool {
	if o != nil && !isNil(o.FloorPlanId) {
		return true
	}

	return false
}

// SetFloorPlanId gets a reference to the given string and assigns it to the FloorPlanId field.
func (o *InlineObject) SetFloorPlanId(v string) {
	o.FloorPlanId = &v
}

func (o InlineObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Lat) {
		toSerialize["lat"] = o.Lat
	}
	if !isNil(o.Lng) {
		toSerialize["lng"] = o.Lng
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.MoveMapMarker) {
		toSerialize["moveMapMarker"] = o.MoveMapMarker
	}
	if !isNil(o.SwitchProfileId) {
		toSerialize["switchProfileId"] = o.SwitchProfileId
	}
	if !isNil(o.FloorPlanId) {
		toSerialize["floorPlanId"] = o.FloorPlanId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


