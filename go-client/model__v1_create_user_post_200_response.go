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

// checks if the V1CreateUserPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1CreateUserPost200Response{}

// V1CreateUserPost200Response struct for V1CreateUserPost200Response
type V1CreateUserPost200Response struct {
	// Internal ID to identify the user's membership within your organization on Uplimit.
	UplimitSubscriptionEnrollmentId string `json:"uplimitSubscriptionEnrollmentId"`
	// Internal ID to identify the user across the Uplimit platform.
	UplimitUserId string `json:"uplimitUserId"`
}

type _V1CreateUserPost200Response V1CreateUserPost200Response

// NewV1CreateUserPost200Response instantiates a new V1CreateUserPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateUserPost200Response(uplimitSubscriptionEnrollmentId string, uplimitUserId string) *V1CreateUserPost200Response {
	this := V1CreateUserPost200Response{}
	this.UplimitSubscriptionEnrollmentId = uplimitSubscriptionEnrollmentId
	this.UplimitUserId = uplimitUserId
	return &this
}

// NewV1CreateUserPost200ResponseWithDefaults instantiates a new V1CreateUserPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateUserPost200ResponseWithDefaults() *V1CreateUserPost200Response {
	this := V1CreateUserPost200Response{}
	return &this
}

// GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field value
func (o *V1CreateUserPost200Response) GetUplimitSubscriptionEnrollmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitSubscriptionEnrollmentId
}

// GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field value
// and a boolean to check if the value has been set.
func (o *V1CreateUserPost200Response) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitSubscriptionEnrollmentId, true
}

// SetUplimitSubscriptionEnrollmentId sets field value
func (o *V1CreateUserPost200Response) SetUplimitSubscriptionEnrollmentId(v string) {
	o.UplimitSubscriptionEnrollmentId = v
}

// GetUplimitUserId returns the UplimitUserId field value
func (o *V1CreateUserPost200Response) GetUplimitUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitUserId
}

// GetUplimitUserIdOk returns a tuple with the UplimitUserId field value
// and a boolean to check if the value has been set.
func (o *V1CreateUserPost200Response) GetUplimitUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitUserId, true
}

// SetUplimitUserId sets field value
func (o *V1CreateUserPost200Response) SetUplimitUserId(v string) {
	o.UplimitUserId = v
}

func (o V1CreateUserPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1CreateUserPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uplimitSubscriptionEnrollmentId"] = o.UplimitSubscriptionEnrollmentId
	toSerialize["uplimitUserId"] = o.UplimitUserId
	return toSerialize, nil
}

func (o *V1CreateUserPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uplimitSubscriptionEnrollmentId",
		"uplimitUserId",
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

	varV1CreateUserPost200Response := _V1CreateUserPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1CreateUserPost200Response)

	if err != nil {
		return err
	}

	*o = V1CreateUserPost200Response(varV1CreateUserPost200Response)

	return err
}

type NullableV1CreateUserPost200Response struct {
	value *V1CreateUserPost200Response
	isSet bool
}

func (v NullableV1CreateUserPost200Response) Get() *V1CreateUserPost200Response {
	return v.value
}

func (v *NullableV1CreateUserPost200Response) Set(val *V1CreateUserPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateUserPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateUserPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateUserPost200Response(val *V1CreateUserPost200Response) *NullableV1CreateUserPost200Response {
	return &NullableV1CreateUserPost200Response{value: val, isSet: true}
}

func (v NullableV1CreateUserPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateUserPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


