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

// checks if the V1ListSessionsInCourseGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ListSessionsInCourseGet200Response{}

// V1ListSessionsInCourseGet200Response struct for V1ListSessionsInCourseGet200Response
type V1ListSessionsInCourseGet200Response struct {
	Sessions []V1ListSessionsInCourseGet200ResponseSessionsInner `json:"sessions"`
	TotalCount float32 `json:"totalCount"`
}

type _V1ListSessionsInCourseGet200Response V1ListSessionsInCourseGet200Response

// NewV1ListSessionsInCourseGet200Response instantiates a new V1ListSessionsInCourseGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListSessionsInCourseGet200Response(sessions []V1ListSessionsInCourseGet200ResponseSessionsInner, totalCount float32) *V1ListSessionsInCourseGet200Response {
	this := V1ListSessionsInCourseGet200Response{}
	this.Sessions = sessions
	this.TotalCount = totalCount
	return &this
}

// NewV1ListSessionsInCourseGet200ResponseWithDefaults instantiates a new V1ListSessionsInCourseGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListSessionsInCourseGet200ResponseWithDefaults() *V1ListSessionsInCourseGet200Response {
	this := V1ListSessionsInCourseGet200Response{}
	return &this
}

// GetSessions returns the Sessions field value
func (o *V1ListSessionsInCourseGet200Response) GetSessions() []V1ListSessionsInCourseGet200ResponseSessionsInner {
	if o == nil {
		var ret []V1ListSessionsInCourseGet200ResponseSessionsInner
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *V1ListSessionsInCourseGet200Response) GetSessionsOk() ([]V1ListSessionsInCourseGet200ResponseSessionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *V1ListSessionsInCourseGet200Response) SetSessions(v []V1ListSessionsInCourseGet200ResponseSessionsInner) {
	o.Sessions = v
}

// GetTotalCount returns the TotalCount field value
func (o *V1ListSessionsInCourseGet200Response) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *V1ListSessionsInCourseGet200Response) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *V1ListSessionsInCourseGet200Response) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o V1ListSessionsInCourseGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ListSessionsInCourseGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessions"] = o.Sessions
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *V1ListSessionsInCourseGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sessions",
		"totalCount",
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

	varV1ListSessionsInCourseGet200Response := _V1ListSessionsInCourseGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ListSessionsInCourseGet200Response)

	if err != nil {
		return err
	}

	*o = V1ListSessionsInCourseGet200Response(varV1ListSessionsInCourseGet200Response)

	return err
}

type NullableV1ListSessionsInCourseGet200Response struct {
	value *V1ListSessionsInCourseGet200Response
	isSet bool
}

func (v NullableV1ListSessionsInCourseGet200Response) Get() *V1ListSessionsInCourseGet200Response {
	return v.value
}

func (v *NullableV1ListSessionsInCourseGet200Response) Set(val *V1ListSessionsInCourseGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListSessionsInCourseGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListSessionsInCourseGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListSessionsInCourseGet200Response(val *V1ListSessionsInCourseGet200Response) *NullableV1ListSessionsInCourseGet200Response {
	return &NullableV1ListSessionsInCourseGet200Response{value: val, isSet: true}
}

func (v NullableV1ListSessionsInCourseGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListSessionsInCourseGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


