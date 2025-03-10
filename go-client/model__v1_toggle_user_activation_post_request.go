/*
Uplimit Organization API

This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

API version: 2025-03-07
Contact: hello@uplimit.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V1ToggleUserActivationPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ToggleUserActivationPostRequest{}

// V1ToggleUserActivationPostRequest struct for V1ToggleUserActivationPostRequest
type V1ToggleUserActivationPostRequest struct {
	// The email address of the user.
	EmailAddress string `json:"emailAddress"`
	// Whether to set the user as active or inactive.
	SetIsActive bool `json:"setIsActive"`
}

type _V1ToggleUserActivationPostRequest V1ToggleUserActivationPostRequest

// NewV1ToggleUserActivationPostRequest instantiates a new V1ToggleUserActivationPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ToggleUserActivationPostRequest(emailAddress string, setIsActive bool) *V1ToggleUserActivationPostRequest {
	this := V1ToggleUserActivationPostRequest{}
	this.EmailAddress = emailAddress
	this.SetIsActive = setIsActive
	return &this
}

// NewV1ToggleUserActivationPostRequestWithDefaults instantiates a new V1ToggleUserActivationPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ToggleUserActivationPostRequestWithDefaults() *V1ToggleUserActivationPostRequest {
	this := V1ToggleUserActivationPostRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *V1ToggleUserActivationPostRequest) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *V1ToggleUserActivationPostRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *V1ToggleUserActivationPostRequest) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetSetIsActive returns the SetIsActive field value
func (o *V1ToggleUserActivationPostRequest) GetSetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SetIsActive
}

// GetSetIsActiveOk returns a tuple with the SetIsActive field value
// and a boolean to check if the value has been set.
func (o *V1ToggleUserActivationPostRequest) GetSetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SetIsActive, true
}

// SetSetIsActive sets field value
func (o *V1ToggleUserActivationPostRequest) SetSetIsActive(v bool) {
	o.SetIsActive = v
}

func (o V1ToggleUserActivationPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ToggleUserActivationPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	toSerialize["setIsActive"] = o.SetIsActive
	return toSerialize, nil
}

func (o *V1ToggleUserActivationPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailAddress",
		"setIsActive",
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

	varV1ToggleUserActivationPostRequest := _V1ToggleUserActivationPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ToggleUserActivationPostRequest)

	if err != nil {
		return err
	}

	*o = V1ToggleUserActivationPostRequest(varV1ToggleUserActivationPostRequest)

	return err
}

type NullableV1ToggleUserActivationPostRequest struct {
	value *V1ToggleUserActivationPostRequest
	isSet bool
}

func (v NullableV1ToggleUserActivationPostRequest) Get() *V1ToggleUserActivationPostRequest {
	return v.value
}

func (v *NullableV1ToggleUserActivationPostRequest) Set(val *V1ToggleUserActivationPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ToggleUserActivationPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ToggleUserActivationPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ToggleUserActivationPostRequest(val *V1ToggleUserActivationPostRequest) *NullableV1ToggleUserActivationPostRequest {
	return &NullableV1ToggleUserActivationPostRequest{value: val, isSet: true}
}

func (v NullableV1ToggleUserActivationPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ToggleUserActivationPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


