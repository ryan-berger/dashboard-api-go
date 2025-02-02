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

// InlineObject209 struct for InlineObject209
type InlineObject209 struct {
	Destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination `json:"destination"`
	// The list of licenses to move
	Licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses `json:"licenses"`
}

// NewInlineObject209 instantiates a new InlineObject209 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject209(destination OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, licenses []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) *InlineObject209 {
	this := InlineObject209{}
	this.Destination = destination
	this.Licenses = licenses
	return &this
}

// NewInlineObject209WithDefaults instantiates a new InlineObject209 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject209WithDefaults() *InlineObject209 {
	this := InlineObject209{}
	return &this
}

// GetDestination returns the Destination field value
func (o *InlineObject209) GetDestination() OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination {
	if o == nil {
		var ret OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetDestinationOk() (*OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *InlineObject209) SetDestination(v OrganizationsOrganizationIdLicensingCotermLicensesMoveDestination) {
	o.Destination = v
}

// GetLicenses returns the Licenses field value
func (o *InlineObject209) GetLicenses() []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses {
	if o == nil {
		var ret []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses
		return ret
	}

	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value
// and a boolean to check if the value has been set.
func (o *InlineObject209) GetLicensesOk() ([]OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses, bool) {
	if o == nil {
    return nil, false
	}
	return o.Licenses, true
}

// SetLicenses sets field value
func (o *InlineObject209) SetLicenses(v []OrganizationsOrganizationIdLicensingCotermLicensesMoveLicenses) {
	o.Licenses = v
}

func (o InlineObject209) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["destination"] = o.Destination
	}
	if true {
		toSerialize["licenses"] = o.Licenses
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject209 struct {
	value *InlineObject209
	isSet bool
}

func (v NullableInlineObject209) Get() *InlineObject209 {
	return v.value
}

func (v *NullableInlineObject209) Set(val *InlineObject209) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject209) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject209) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject209(val *InlineObject209) *NullableInlineObject209 {
	return &NullableInlineObject209{value: val, isSet: true}
}

func (v NullableInlineObject209) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject209) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


