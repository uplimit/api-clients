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

// checks if the ListSessionsInCourseResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListSessionsInCourseResponseSchema{}

// ListSessionsInCourseResponseSchema struct for ListSessionsInCourseResponseSchema
type ListSessionsInCourseResponseSchema struct {
	Sessions []V1ListSessionsInCourseGet200ResponseSessionsInner `json:"sessions"`
	TotalCount float32 `json:"totalCount"`
}

type _ListSessionsInCourseResponseSchema ListSessionsInCourseResponseSchema

// NewListSessionsInCourseResponseSchema instantiates a new ListSessionsInCourseResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSessionsInCourseResponseSchema(sessions []V1ListSessionsInCourseGet200ResponseSessionsInner, totalCount float32) *ListSessionsInCourseResponseSchema {
	this := ListSessionsInCourseResponseSchema{}
	this.Sessions = sessions
	this.TotalCount = totalCount
	return &this
}

// NewListSessionsInCourseResponseSchemaWithDefaults instantiates a new ListSessionsInCourseResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSessionsInCourseResponseSchemaWithDefaults() *ListSessionsInCourseResponseSchema {
	this := ListSessionsInCourseResponseSchema{}
	return &this
}

// GetSessions returns the Sessions field value
func (o *ListSessionsInCourseResponseSchema) GetSessions() []V1ListSessionsInCourseGet200ResponseSessionsInner {
	if o == nil {
		var ret []V1ListSessionsInCourseGet200ResponseSessionsInner
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *ListSessionsInCourseResponseSchema) GetSessionsOk() ([]V1ListSessionsInCourseGet200ResponseSessionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sessions, true
}

// SetSessions sets field value
func (o *ListSessionsInCourseResponseSchema) SetSessions(v []V1ListSessionsInCourseGet200ResponseSessionsInner) {
	o.Sessions = v
}

// GetTotalCount returns the TotalCount field value
func (o *ListSessionsInCourseResponseSchema) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ListSessionsInCourseResponseSchema) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ListSessionsInCourseResponseSchema) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o ListSessionsInCourseResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListSessionsInCourseResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessions"] = o.Sessions
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *ListSessionsInCourseResponseSchema) UnmarshalJSON(data []byte) (err error) {
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

	varListSessionsInCourseResponseSchema := _ListSessionsInCourseResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListSessionsInCourseResponseSchema)

	if err != nil {
		return err
	}

	*o = ListSessionsInCourseResponseSchema(varListSessionsInCourseResponseSchema)

	return err
}

type NullableListSessionsInCourseResponseSchema struct {
	value *ListSessionsInCourseResponseSchema
	isSet bool
}

func (v NullableListSessionsInCourseResponseSchema) Get() *ListSessionsInCourseResponseSchema {
	return v.value
}

func (v *NullableListSessionsInCourseResponseSchema) Set(val *ListSessionsInCourseResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableListSessionsInCourseResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableListSessionsInCourseResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSessionsInCourseResponseSchema(val *ListSessionsInCourseResponseSchema) *NullableListSessionsInCourseResponseSchema {
	return &NullableListSessionsInCourseResponseSchema{value: val, isSet: true}
}

func (v NullableListSessionsInCourseResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSessionsInCourseResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


