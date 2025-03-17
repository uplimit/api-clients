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

// checks if the V1HealthcheckGet403Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1HealthcheckGet403Response{}

// V1HealthcheckGet403Response struct for V1HealthcheckGet403Response
type V1HealthcheckGet403Response struct {
	Success bool `json:"success"`
	// The error message.
	Error string `json:"error"`
}

type _V1HealthcheckGet403Response V1HealthcheckGet403Response

// NewV1HealthcheckGet403Response instantiates a new V1HealthcheckGet403Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1HealthcheckGet403Response(success bool, error_ string) *V1HealthcheckGet403Response {
	this := V1HealthcheckGet403Response{}
	this.Success = success
	this.Error = error_
	return &this
}

// NewV1HealthcheckGet403ResponseWithDefaults instantiates a new V1HealthcheckGet403Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1HealthcheckGet403ResponseWithDefaults() *V1HealthcheckGet403Response {
	this := V1HealthcheckGet403Response{}
	return &this
}

// GetSuccess returns the Success field value
func (o *V1HealthcheckGet403Response) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *V1HealthcheckGet403Response) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *V1HealthcheckGet403Response) SetSuccess(v bool) {
	o.Success = v
}

// GetError returns the Error field value
func (o *V1HealthcheckGet403Response) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *V1HealthcheckGet403Response) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *V1HealthcheckGet403Response) SetError(v string) {
	o.Error = v
}

func (o V1HealthcheckGet403Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1HealthcheckGet403Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["success"] = o.Success
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

func (o *V1HealthcheckGet403Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"success",
		"error",
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

	varV1HealthcheckGet403Response := _V1HealthcheckGet403Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1HealthcheckGet403Response)

	if err != nil {
		return err
	}

	*o = V1HealthcheckGet403Response(varV1HealthcheckGet403Response)

	return err
}

type NullableV1HealthcheckGet403Response struct {
	value *V1HealthcheckGet403Response
	isSet bool
}

func (v NullableV1HealthcheckGet403Response) Get() *V1HealthcheckGet403Response {
	return v.value
}

func (v *NullableV1HealthcheckGet403Response) Set(val *V1HealthcheckGet403Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1HealthcheckGet403Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1HealthcheckGet403Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1HealthcheckGet403Response(val *V1HealthcheckGet403Response) *NullableV1HealthcheckGet403Response {
	return &NullableV1HealthcheckGet403Response{value: val, isSet: true}
}

func (v NullableV1HealthcheckGet403Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1HealthcheckGet403Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


