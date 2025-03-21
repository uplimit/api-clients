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

// checks if the ListCoursesResponseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCoursesResponseSchema{}

// ListCoursesResponseSchema struct for ListCoursesResponseSchema
type ListCoursesResponseSchema struct {
	Courses []V1ListCoursesGet200ResponseCoursesInner `json:"courses"`
	TotalCount float32 `json:"totalCount"`
}

type _ListCoursesResponseSchema ListCoursesResponseSchema

// NewListCoursesResponseSchema instantiates a new ListCoursesResponseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCoursesResponseSchema(courses []V1ListCoursesGet200ResponseCoursesInner, totalCount float32) *ListCoursesResponseSchema {
	this := ListCoursesResponseSchema{}
	this.Courses = courses
	this.TotalCount = totalCount
	return &this
}

// NewListCoursesResponseSchemaWithDefaults instantiates a new ListCoursesResponseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCoursesResponseSchemaWithDefaults() *ListCoursesResponseSchema {
	this := ListCoursesResponseSchema{}
	return &this
}

// GetCourses returns the Courses field value
func (o *ListCoursesResponseSchema) GetCourses() []V1ListCoursesGet200ResponseCoursesInner {
	if o == nil {
		var ret []V1ListCoursesGet200ResponseCoursesInner
		return ret
	}

	return o.Courses
}

// GetCoursesOk returns a tuple with the Courses field value
// and a boolean to check if the value has been set.
func (o *ListCoursesResponseSchema) GetCoursesOk() ([]V1ListCoursesGet200ResponseCoursesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Courses, true
}

// SetCourses sets field value
func (o *ListCoursesResponseSchema) SetCourses(v []V1ListCoursesGet200ResponseCoursesInner) {
	o.Courses = v
}

// GetTotalCount returns the TotalCount field value
func (o *ListCoursesResponseSchema) GetTotalCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *ListCoursesResponseSchema) GetTotalCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *ListCoursesResponseSchema) SetTotalCount(v float32) {
	o.TotalCount = v
}

func (o ListCoursesResponseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCoursesResponseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["courses"] = o.Courses
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

func (o *ListCoursesResponseSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"courses",
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

	varListCoursesResponseSchema := _ListCoursesResponseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListCoursesResponseSchema)

	if err != nil {
		return err
	}

	*o = ListCoursesResponseSchema(varListCoursesResponseSchema)

	return err
}

type NullableListCoursesResponseSchema struct {
	value *ListCoursesResponseSchema
	isSet bool
}

func (v NullableListCoursesResponseSchema) Get() *ListCoursesResponseSchema {
	return v.value
}

func (v *NullableListCoursesResponseSchema) Set(val *ListCoursesResponseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableListCoursesResponseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableListCoursesResponseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCoursesResponseSchema(val *ListCoursesResponseSchema) *NullableListCoursesResponseSchema {
	return &NullableListCoursesResponseSchema{value: val, isSet: true}
}

func (v NullableListCoursesResponseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCoursesResponseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


