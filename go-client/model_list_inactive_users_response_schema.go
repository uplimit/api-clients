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

// checks if the ListInactiveUsersResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInactiveUsersResponseSchema{}

// ListInactiveUsersResponseSchema struct for ListInactiveUsersResponseSchema
type ListInactiveUsersResponseSchema struct {
	Users []V1GetUserInformationEmailAddressOrUserIdGet200Response `json:"users"`
	TotalCount float32 `json:"totalCount"`
}

type _ListInactiveUsersResponseSchema ListInactiveUsersResponseSchema

// NewListInactiveUsersResponseSchema instantiates a new ListInactiveUsersResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInactiveUsersResponseSchema(users []V1GetUserInformationEmailAddressOrUserIdGet200Response, totalCount float32) *ListInactiveUsersResponseSchema {
	this := ListInactiveUsersResponseSchema{}
	this.Users = users
	this.TotalCount = totalCount
	return &this
}

// NewListInactiveUsersResponseSchemaWithDefaults instantiates a new ListInactiveUsersResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInactiveUsersResponseSchemaWithDefaults() *ListInactiveUsersResponseSchema {
	this := ListInactiveUsersResponseSchema{}
	return &this
}

// GetUsers returns the Users field value
func (o *ListInactiveUsersResponseSchema) GetUsers() []V1GetUserInformationEmailAddressOrUserIdGet200Response {
	if o == nil {
		var ret []V1GetUserInformationEmailAddressOrUserIdGet200Response
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *ListInactiveUsersResponseSchema) GetUsersOk() ([]V1GetUserInformationEmailAddressOrUserIdGet200Response, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *ListInactiveUsersResponseSchema) SetUsers(v []V1GetUserInformationEmailAddressOrUserIdGet200Response) {
	o.Users = v
}

// GetTotalCount returns the TotalCount field value
func (o *ListInactiveUsersResponseSchema) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ListInactiveUsersResponseSchema) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ListInactiveUsersResponseSchema) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o ListInactiveUsersResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInactiveUsersResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *ListInactiveUsersResponseSchema) UnmarshalJSON(data []byte) (err error) {
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

	varListInactiveUsersResponseSchema := _ListInactiveUsersResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListInactiveUsersResponseSchema)

	if err != nil {
		return err
	}

	*o = ListInactiveUsersResponseSchema(varListInactiveUsersResponseSchema)

	return err
}

type NullableListInactiveUsersResponseSchema struct {
	value *ListInactiveUsersResponseSchema
	isSet bool
}

func (v NullableListInactiveUsersResponseSchema) Get() *ListInactiveUsersResponseSchema {
	return v.value
}

func (v *NullableListInactiveUsersResponseSchema) Set(val *ListInactiveUsersResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableListInactiveUsersResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableListInactiveUsersResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInactiveUsersResponseSchema(val *ListInactiveUsersResponseSchema) *NullableListInactiveUsersResponseSchema {
	return &NullableListInactiveUsersResponseSchema{value: val, isSet: true}
}

func (v NullableListInactiveUsersResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInactiveUsersResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


