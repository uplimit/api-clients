/*
Uplimit Organization API

This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

API version: 2025-03-17
Contact: hello@uplimit.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HealthcheckSuccessResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthcheckSuccessResponseSchema{}

// HealthcheckSuccessResponseSchema struct for HealthcheckSuccessResponseSchema
type HealthcheckSuccessResponseSchema struct {
	Success bool `json:"success"`
	// The Uplimit ID of the organization.
	UplimitOrganizationId string `json:"uplimitOrganizationId"`
	// The name of the organization as it appears in Uplimit.
	UplimitOrganizationName string `json:"uplimitOrganizationName"`
}

type _HealthcheckSuccessResponseSchema HealthcheckSuccessResponseSchema

// NewHealthcheckSuccessResponseSchema instantiates a new HealthcheckSuccessResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthcheckSuccessResponseSchema(success bool, uplimitOrganizationId string, uplimitOrganizationName string) *HealthcheckSuccessResponseSchema {
	this := HealthcheckSuccessResponseSchema{}
	this.Success = success
	this.UplimitOrganizationId = uplimitOrganizationId
	this.UplimitOrganizationName = uplimitOrganizationName
	return &this
}

// NewHealthcheckSuccessResponseSchemaWithDefaults instantiates a new HealthcheckSuccessResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthcheckSuccessResponseSchemaWithDefaults() *HealthcheckSuccessResponseSchema {
	this := HealthcheckSuccessResponseSchema{}
	return &this
}

// GetSuccess returns the Success field value
func (o *HealthcheckSuccessResponseSchema) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *HealthcheckSuccessResponseSchema) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *HealthcheckSuccessResponseSchema) SetSuccess(v bool) {
	o.Success = v
}

// GetUplimitOrganizationId returns the UplimitOrganizationId field value
func (o *HealthcheckSuccessResponseSchema) GetUplimitOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitOrganizationId
}

// GetUplimitOrganizationIdOk returns a tuple with the UplimitOrganizationId field value
// and a boolean to check if the value has been set.
func (o *HealthcheckSuccessResponseSchema) GetUplimitOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitOrganizationId, true
}

// SetUplimitOrganizationId sets field value
func (o *HealthcheckSuccessResponseSchema) SetUplimitOrganizationId(v string) {
	o.UplimitOrganizationId = v
}

// GetUplimitOrganizationName returns the UplimitOrganizationName field value
func (o *HealthcheckSuccessResponseSchema) GetUplimitOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitOrganizationName
}

// GetUplimitOrganizationNameOk returns a tuple with the UplimitOrganizationName field value
// and a boolean to check if the value has been set.
func (o *HealthcheckSuccessResponseSchema) GetUplimitOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitOrganizationName, true
}

// SetUplimitOrganizationName sets field value
func (o *HealthcheckSuccessResponseSchema) SetUplimitOrganizationName(v string) {
	o.UplimitOrganizationName = v
}

func (o HealthcheckSuccessResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthcheckSuccessResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["uplimitOrganizationId"] = o.UplimitOrganizationId
	toSerialize["uplimitOrganizationName"] = o.UplimitOrganizationName
	return toSerialize, nil
}

func (o *HealthcheckSuccessResponseSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"uplimitOrganizationId",
		"uplimitOrganizationName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHealthcheckSuccessResponseSchema := _HealthcheckSuccessResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHealthcheckSuccessResponseSchema)

	if err != nil {
		return err
	}

	*o = HealthcheckSuccessResponseSchema(varHealthcheckSuccessResponseSchema)

	return err
}

type NullableHealthcheckSuccessResponseSchema struct {
	value *HealthcheckSuccessResponseSchema
	isSet bool
}

func (v NullableHealthcheckSuccessResponseSchema) Get() *HealthcheckSuccessResponseSchema {
	return v.value
}

func (v *NullableHealthcheckSuccessResponseSchema) Set(val *HealthcheckSuccessResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthcheckSuccessResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthcheckSuccessResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthcheckSuccessResponseSchema(val *HealthcheckSuccessResponseSchema) *NullableHealthcheckSuccessResponseSchema {
	return &NullableHealthcheckSuccessResponseSchema{value: val, isSet: true}
}

func (v NullableHealthcheckSuccessResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthcheckSuccessResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


