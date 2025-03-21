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

// checks if the V1ListCoursesGet200ResponseCoursesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ListCoursesGet200ResponseCoursesInner{}

// V1ListCoursesGet200ResponseCoursesInner struct for V1ListCoursesGet200ResponseCoursesInner
type V1ListCoursesGet200ResponseCoursesInner struct {
	// Internal ID to identify the course across the Uplimit platform.
	UplimitCourseId string `json:"uplimitCourseId"`
	// The slug (i.e. short name) of the course across the Uplimit platform.
	UplimitCourseSlug string `json:"uplimitCourseSlug"`
	// The name of the course.
	Name string `json:"name"`
}

type _V1ListCoursesGet200ResponseCoursesInner V1ListCoursesGet200ResponseCoursesInner

// NewV1ListCoursesGet200ResponseCoursesInner instantiates a new V1ListCoursesGet200ResponseCoursesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListCoursesGet200ResponseCoursesInner(uplimitCourseId string, uplimitCourseSlug string, name string) *V1ListCoursesGet200ResponseCoursesInner {
	this := V1ListCoursesGet200ResponseCoursesInner{}
	this.UplimitCourseId = uplimitCourseId
	this.UplimitCourseSlug = uplimitCourseSlug
	this.Name = name
	return &this
}

// NewV1ListCoursesGet200ResponseCoursesInnerWithDefaults instantiates a new V1ListCoursesGet200ResponseCoursesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListCoursesGet200ResponseCoursesInnerWithDefaults() *V1ListCoursesGet200ResponseCoursesInner {
	this := V1ListCoursesGet200ResponseCoursesInner{}
	return &this
}

// GetUplimitCourseId returns the UplimitCourseId field value
func (o *V1ListCoursesGet200ResponseCoursesInner) GetUplimitCourseId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitCourseId
}

// GetUplimitCourseIdOk returns a tuple with the UplimitCourseId field value
// and a boolean to check if the value has been set.
func (o *V1ListCoursesGet200ResponseCoursesInner) GetUplimitCourseIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitCourseId, true
}

// SetUplimitCourseId sets field value
func (o *V1ListCoursesGet200ResponseCoursesInner) SetUplimitCourseId(v string) {
	o.UplimitCourseId = v
}

// GetUplimitCourseSlug returns the UplimitCourseSlug field value
func (o *V1ListCoursesGet200ResponseCoursesInner) GetUplimitCourseSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitCourseSlug
}

// GetUplimitCourseSlugOk returns a tuple with the UplimitCourseSlug field value
// and a boolean to check if the value has been set.
func (o *V1ListCoursesGet200ResponseCoursesInner) GetUplimitCourseSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitCourseSlug, true
}

// SetUplimitCourseSlug sets field value
func (o *V1ListCoursesGet200ResponseCoursesInner) SetUplimitCourseSlug(v string) {
	o.UplimitCourseSlug = v
}

// GetName returns the Name field value
func (o *V1ListCoursesGet200ResponseCoursesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1ListCoursesGet200ResponseCoursesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1ListCoursesGet200ResponseCoursesInner) SetName(v string) {
	o.Name = v
}

func (o V1ListCoursesGet200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ListCoursesGet200ResponseCoursesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uplimitCourseId"] = o.UplimitCourseId
	toSerialize["uplimitCourseSlug"] = o.UplimitCourseSlug
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *V1ListCoursesGet200ResponseCoursesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uplimitCourseId",
		"uplimitCourseSlug",
		"name",
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

	varV1ListCoursesGet200ResponseCoursesInner := _V1ListCoursesGet200ResponseCoursesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ListCoursesGet200ResponseCoursesInner)

	if err != nil {
		return err
	}

	*o = V1ListCoursesGet200ResponseCoursesInner(varV1ListCoursesGet200ResponseCoursesInner)

	return err
}

type NullableV1ListCoursesGet200ResponseCoursesInner struct {
	value *V1ListCoursesGet200ResponseCoursesInner
	isSet bool
}

func (v NullableV1ListCoursesGet200ResponseCoursesInner) Get() *V1ListCoursesGet200ResponseCoursesInner {
	return v.value
}

func (v *NullableV1ListCoursesGet200ResponseCoursesInner) Set(val *V1ListCoursesGet200ResponseCoursesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListCoursesGet200ResponseCoursesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListCoursesGet200ResponseCoursesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListCoursesGet200ResponseCoursesInner(val *V1ListCoursesGet200ResponseCoursesInner) *NullableV1ListCoursesGet200ResponseCoursesInner {
	return &NullableV1ListCoursesGet200ResponseCoursesInner{value: val, isSet: true}
}

func (v NullableV1ListCoursesGet200ResponseCoursesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListCoursesGet200ResponseCoursesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


