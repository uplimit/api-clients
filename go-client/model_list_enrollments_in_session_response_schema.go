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

// checks if the ListEnrollmentsInSessionResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEnrollmentsInSessionResponseSchema{}

// ListEnrollmentsInSessionResponseSchema struct for ListEnrollmentsInSessionResponseSchema
type ListEnrollmentsInSessionResponseSchema struct {
	Users []V1ListEnrollmentsInSessionGet200ResponseUsersInner `json:"users"`
	TotalCount float32 `json:"totalCount"`
}

type _ListEnrollmentsInSessionResponseSchema ListEnrollmentsInSessionResponseSchema

// NewListEnrollmentsInSessionResponseSchema instantiates a new ListEnrollmentsInSessionResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEnrollmentsInSessionResponseSchema(users []V1ListEnrollmentsInSessionGet200ResponseUsersInner, totalCount float32) *ListEnrollmentsInSessionResponseSchema {
	this := ListEnrollmentsInSessionResponseSchema{}
	this.Users = users
	this.TotalCount = totalCount
	return &this
}

// NewListEnrollmentsInSessionResponseSchemaWithDefaults instantiates a new ListEnrollmentsInSessionResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEnrollmentsInSessionResponseSchemaWithDefaults() *ListEnrollmentsInSessionResponseSchema {
	this := ListEnrollmentsInSessionResponseSchema{}
	return &this
}

// GetUsers returns the Users field value
func (o *ListEnrollmentsInSessionResponseSchema) GetUsers() []V1ListEnrollmentsInSessionGet200ResponseUsersInner {
	if o == nil {
		var ret []V1ListEnrollmentsInSessionGet200ResponseUsersInner
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ListEnrollmentsInSessionResponseSchema) GetUsersOk() ([]V1ListEnrollmentsInSessionGet200ResponseUsersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ListEnrollmentsInSessionResponseSchema) SetUsers(v []V1ListEnrollmentsInSessionGet200ResponseUsersInner) {
	o.Users = v
}

// GetTotalCount returns the TotalCount field value
func (o *ListEnrollmentsInSessionResponseSchema) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ListEnrollmentsInSessionResponseSchema) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ListEnrollmentsInSessionResponseSchema) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o ListEnrollmentsInSessionResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEnrollmentsInSessionResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *ListEnrollmentsInSessionResponseSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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

	varListEnrollmentsInSessionResponseSchema := _ListEnrollmentsInSessionResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListEnrollmentsInSessionResponseSchema)

	if err != nil {
		return err
	}

	*o = ListEnrollmentsInSessionResponseSchema(varListEnrollmentsInSessionResponseSchema)

	return err
}

type NullableListEnrollmentsInSessionResponseSchema struct {
	value *ListEnrollmentsInSessionResponseSchema
	isSet bool
}

func (v NullableListEnrollmentsInSessionResponseSchema) Get() *ListEnrollmentsInSessionResponseSchema {
	return v.value
}

func (v *NullableListEnrollmentsInSessionResponseSchema) Set(val *ListEnrollmentsInSessionResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableListEnrollmentsInSessionResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableListEnrollmentsInSessionResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEnrollmentsInSessionResponseSchema(val *ListEnrollmentsInSessionResponseSchema) *NullableListEnrollmentsInSessionResponseSchema {
	return &NullableListEnrollmentsInSessionResponseSchema{value: val, isSet: true}
}

func (v NullableListEnrollmentsInSessionResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEnrollmentsInSessionResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


