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

// InlineResponse2015 struct for InlineResponse2015
type InlineResponse2015 struct {
	// ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
	Id *string `json:"id,omitempty"`
	// ID of the organization this action batch belongs to
	OrganizationId *string `json:"organizationId,omitempty"`
	// Flag describing whether the action should be previewed before executing or not
	Confirmed *bool `json:"confirmed,omitempty"`
	// Flag describing whether actions should run synchronously or asynchronously
	Synchronous *bool `json:"synchronous,omitempty"`
	Status *InlineResponse2015Status `json:"status,omitempty"`
	// A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	Actions []InlineResponse2015Actions `json:"actions"`
}

// NewInlineResponse2015 instantiates a new InlineResponse2015 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2015(actions []InlineResponse2015Actions) *InlineResponse2015 {
	this := InlineResponse2015{}
	this.Actions = actions
	return &this
}

// NewInlineResponse2015WithDefaults instantiates a new InlineResponse2015 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2015WithDefaults() *InlineResponse2015 {
	this := InlineResponse2015{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse2015) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse2015) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse2015) SetId(v string) {
	o.Id = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *InlineResponse2015) GetOrganizationId() string {
	if o == nil || isNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetOrganizationIdOk() (*string, bool) {
	if o == nil || isNil(o.OrganizationId) {
    return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *InlineResponse2015) HasOrganizationId() bool {
	if o != nil && !isNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *InlineResponse2015) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetConfirmed returns the Confirmed field value if set, zero value otherwise.
func (o *InlineResponse2015) GetConfirmed() bool {
	if o == nil || isNil(o.Confirmed) {
		var ret bool
		return ret
	}
	return *o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetConfirmedOk() (*bool, bool) {
	if o == nil || isNil(o.Confirmed) {
    return nil, false
	}
	return o.Confirmed, true
}

// HasConfirmed returns a boolean if a field has been set.
func (o *InlineResponse2015) HasConfirmed() bool {
	if o != nil && !isNil(o.Confirmed) {
		return true
	}

	return false
}

// SetConfirmed gets a reference to the given bool and assigns it to the Confirmed field.
func (o *InlineResponse2015) SetConfirmed(v bool) {
	o.Confirmed = &v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *InlineResponse2015) GetSynchronous() bool {
	if o == nil || isNil(o.Synchronous) {
		var ret bool
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetSynchronousOk() (*bool, bool) {
	if o == nil || isNil(o.Synchronous) {
    return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *InlineResponse2015) HasSynchronous() bool {
	if o != nil && !isNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given bool and assigns it to the Synchronous field.
func (o *InlineResponse2015) SetSynchronous(v bool) {
	o.Synchronous = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse2015) GetStatus() InlineResponse2015Status {
	if o == nil || isNil(o.Status) {
		var ret InlineResponse2015Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetStatusOk() (*InlineResponse2015Status, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse2015) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given InlineResponse2015Status and assigns it to the Status field.
func (o *InlineResponse2015) SetStatus(v InlineResponse2015Status) {
	o.Status = &v
}

// GetActions returns the Actions field value
func (o *InlineResponse2015) GetActions() []InlineResponse2015Actions {
	if o == nil {
		var ret []InlineResponse2015Actions
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2015) GetActionsOk() ([]InlineResponse2015Actions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *InlineResponse2015) SetActions(v []InlineResponse2015Actions) {
	o.Actions = v
}

func (o InlineResponse2015) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !isNil(o.Confirmed) {
		toSerialize["confirmed"] = o.Confirmed
	}
	if !isNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2015 struct {
	value *InlineResponse2015
	isSet bool
}

func (v NullableInlineResponse2015) Get() *InlineResponse2015 {
	return v.value
}

func (v *NullableInlineResponse2015) Set(val *InlineResponse2015) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2015) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2015) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2015(val *InlineResponse2015) *NullableInlineResponse2015 {
	return &NullableInlineResponse2015{value: val, isSet: true}
}

func (v NullableInlineResponse2015) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2015) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


