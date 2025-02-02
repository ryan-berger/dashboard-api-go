/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20028 struct for InlineResponse20028
type InlineResponse20028 struct {
	// Product type to rollback (if the network is a combined network)
	Product *string `json:"product,omitempty"`
	// Status of the rollback
	Status *string `json:"status,omitempty"`
	// Batch ID of the firmware rollback
	UpgradeBatchId *string `json:"upgradeBatchId,omitempty"`
	// Scheduled time for the rollback
	Time *time.Time `json:"time,omitempty"`
	ToVersion *InlineResponse20028ToVersion `json:"toVersion,omitempty"`
	// Reasons for the rollback
	Reasons []InlineResponse20028Reasons `json:"reasons,omitempty"`
}

// NewInlineResponse20028 instantiates a new InlineResponse20028 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20028() *InlineResponse20028 {
	this := InlineResponse20028{}
	return &this
}

// NewInlineResponse20028WithDefaults instantiates a new InlineResponse20028 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20028WithDefaults() *InlineResponse20028 {
	this := InlineResponse20028{}
	return &this
}

// GetProduct returns the Product field value if set, zero value otherwise.
func (o *InlineResponse20028) GetProduct() string {
	if o == nil || isNil(o.Product) {
		var ret string
		return ret
	}
	return *o.Product
}

// GetProductOk returns a tuple with the Product field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetProductOk() (*string, bool) {
	if o == nil || isNil(o.Product) {
    return nil, false
	}
	return o.Product, true
}

// HasProduct returns a boolean if a field has been set.
func (o *InlineResponse20028) HasProduct() bool {
	if o != nil && !isNil(o.Product) {
		return true
	}

	return false
}

// SetProduct gets a reference to the given string and assigns it to the Product field.
func (o *InlineResponse20028) SetProduct(v string) {
	o.Product = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20028) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20028) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20028) SetStatus(v string) {
	o.Status = &v
}

// GetUpgradeBatchId returns the UpgradeBatchId field value if set, zero value otherwise.
func (o *InlineResponse20028) GetUpgradeBatchId() string {
	if o == nil || isNil(o.UpgradeBatchId) {
		var ret string
		return ret
	}
	return *o.UpgradeBatchId
}

// GetUpgradeBatchIdOk returns a tuple with the UpgradeBatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetUpgradeBatchIdOk() (*string, bool) {
	if o == nil || isNil(o.UpgradeBatchId) {
    return nil, false
	}
	return o.UpgradeBatchId, true
}

// HasUpgradeBatchId returns a boolean if a field has been set.
func (o *InlineResponse20028) HasUpgradeBatchId() bool {
	if o != nil && !isNil(o.UpgradeBatchId) {
		return true
	}

	return false
}

// SetUpgradeBatchId gets a reference to the given string and assigns it to the UpgradeBatchId field.
func (o *InlineResponse20028) SetUpgradeBatchId(v string) {
	o.UpgradeBatchId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *InlineResponse20028) GetTime() time.Time {
	if o == nil || isNil(o.Time) {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.Time) {
    return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *InlineResponse20028) HasTime() bool {
	if o != nil && !isNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *InlineResponse20028) SetTime(v time.Time) {
	o.Time = &v
}

// GetToVersion returns the ToVersion field value if set, zero value otherwise.
func (o *InlineResponse20028) GetToVersion() InlineResponse20028ToVersion {
	if o == nil || isNil(o.ToVersion) {
		var ret InlineResponse20028ToVersion
		return ret
	}
	return *o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetToVersionOk() (*InlineResponse20028ToVersion, bool) {
	if o == nil || isNil(o.ToVersion) {
    return nil, false
	}
	return o.ToVersion, true
}

// HasToVersion returns a boolean if a field has been set.
func (o *InlineResponse20028) HasToVersion() bool {
	if o != nil && !isNil(o.ToVersion) {
		return true
	}

	return false
}

// SetToVersion gets a reference to the given InlineResponse20028ToVersion and assigns it to the ToVersion field.
func (o *InlineResponse20028) SetToVersion(v InlineResponse20028ToVersion) {
	o.ToVersion = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *InlineResponse20028) GetReasons() []InlineResponse20028Reasons {
	if o == nil || isNil(o.Reasons) {
		var ret []InlineResponse20028Reasons
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20028) GetReasonsOk() ([]InlineResponse20028Reasons, bool) {
	if o == nil || isNil(o.Reasons) {
    return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *InlineResponse20028) HasReasons() bool {
	if o != nil && !isNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []InlineResponse20028Reasons and assigns it to the Reasons field.
func (o *InlineResponse20028) SetReasons(v []InlineResponse20028Reasons) {
	o.Reasons = v
}

func (o InlineResponse20028) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Product) {
		toSerialize["product"] = o.Product
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.UpgradeBatchId) {
		toSerialize["upgradeBatchId"] = o.UpgradeBatchId
	}
	if !isNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !isNil(o.ToVersion) {
		toSerialize["toVersion"] = o.ToVersion
	}
	if !isNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20028 struct {
	value *InlineResponse20028
	isSet bool
}

func (v NullableInlineResponse20028) Get() *InlineResponse20028 {
	return v.value
}

func (v *NullableInlineResponse20028) Set(val *InlineResponse20028) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20028) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20028) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20028(val *InlineResponse20028) *NullableInlineResponse20028 {
	return &NullableInlineResponse20028{value: val, isSet: true}
}

func (v NullableInlineResponse20028) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20028) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


