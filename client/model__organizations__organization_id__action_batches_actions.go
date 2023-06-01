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

// OrganizationsOrganizationIdActionBatchesActions struct for OrganizationsOrganizationIdActionBatchesActions
type OrganizationsOrganizationIdActionBatchesActions struct {
	// Unique identifier for the resource to be acted on
	Resource string `json:"resource"`
	// The operation to be used
	Operation string `json:"operation"`
	// The body of the action
	Body map[string]interface{} `json:"body,omitempty"`
}

// NewOrganizationsOrganizationIdActionBatchesActions instantiates a new OrganizationsOrganizationIdActionBatchesActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdActionBatchesActions(resource string, operation string) *OrganizationsOrganizationIdActionBatchesActions {
	this := OrganizationsOrganizationIdActionBatchesActions{}
	this.Resource = resource
	this.Operation = operation
	return &this
}

// NewOrganizationsOrganizationIdActionBatchesActionsWithDefaults instantiates a new OrganizationsOrganizationIdActionBatchesActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdActionBatchesActionsWithDefaults() *OrganizationsOrganizationIdActionBatchesActions {
	this := OrganizationsOrganizationIdActionBatchesActions{}
	return &this
}

// GetResource returns the Resource field value
func (o *OrganizationsOrganizationIdActionBatchesActions) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions) GetResourceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *OrganizationsOrganizationIdActionBatchesActions) SetResource(v string) {
	o.Resource = v
}

// GetOperation returns the Operation field value
func (o *OrganizationsOrganizationIdActionBatchesActions) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions) GetOperationOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *OrganizationsOrganizationIdActionBatchesActions) SetOperation(v string) {
	o.Operation = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdActionBatchesActions) GetBody() map[string]interface{} {
	if o == nil || isNil(o.Body) {
		var ret map[string]interface{}
		return ret
	}
	return o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions) GetBodyOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Body) {
    return map[string]interface{}{}, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdActionBatchesActions) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given map[string]interface{} and assigns it to the Body field.
func (o *OrganizationsOrganizationIdActionBatchesActions) SetBody(v map[string]interface{}) {
	o.Body = v
}

func (o OrganizationsOrganizationIdActionBatchesActions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resource"] = o.Resource
	}
	if true {
		toSerialize["operation"] = o.Operation
	}
	if !isNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdActionBatchesActions struct {
	value *OrganizationsOrganizationIdActionBatchesActions
	isSet bool
}

func (v NullableOrganizationsOrganizationIdActionBatchesActions) Get() *OrganizationsOrganizationIdActionBatchesActions {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdActionBatchesActions) Set(val *OrganizationsOrganizationIdActionBatchesActions) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdActionBatchesActions) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdActionBatchesActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdActionBatchesActions(val *OrganizationsOrganizationIdActionBatchesActions) *NullableOrganizationsOrganizationIdActionBatchesActions {
	return &NullableOrganizationsOrganizationIdActionBatchesActions{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdActionBatchesActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdActionBatchesActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


