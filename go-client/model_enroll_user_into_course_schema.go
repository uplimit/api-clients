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

// checks if the EnrollUserIntoCourseSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnrollUserIntoCourseSchema{}

// EnrollUserIntoCourseSchema struct for EnrollUserIntoCourseSchema
type EnrollUserIntoCourseSchema struct {
	// The email address of the user.
	EmailAddress string `json:"emailAddress"`
	// The ID of the session to enroll the user into. You must provide either this or uplimitSessionId.
	SessionId *string `json:"sessionId,omitempty"`
	// Internal ID to identify the session across the Uplimit platform.
	UplimitSessionId *string `json:"uplimitSessionId,omitempty"`
	// Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group.
	SubscriptionCommitmentId *string `json:"subscriptionCommitmentId,omitempty"`
}

type _EnrollUserIntoCourseSchema EnrollUserIntoCourseSchema

// NewEnrollUserIntoCourseSchema instantiates a new EnrollUserIntoCourseSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnrollUserIntoCourseSchema(emailAddress string) *EnrollUserIntoCourseSchema {
	this := EnrollUserIntoCourseSchema{}
	this.EmailAddress = emailAddress
	return &this
}

// NewEnrollUserIntoCourseSchemaWithDefaults instantiates a new EnrollUserIntoCourseSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnrollUserIntoCourseSchemaWithDefaults() *EnrollUserIntoCourseSchema {
	this := EnrollUserIntoCourseSchema{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *EnrollUserIntoCourseSchema) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *EnrollUserIntoCourseSchema) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *EnrollUserIntoCourseSchema) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *EnrollUserIntoCourseSchema) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollUserIntoCourseSchema) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *EnrollUserIntoCourseSchema) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *EnrollUserIntoCourseSchema) SetSessionId(v string) {
	o.SessionId = &v
}

// GetUplimitSessionId returns the UplimitSessionId field value if set, zero value otherwise.
func (o *EnrollUserIntoCourseSchema) GetUplimitSessionId() string {
	if o == nil || IsNil(o.UplimitSessionId) {
		var ret string
		return ret
	}
	return *o.UplimitSessionId
}

// GetUplimitSessionIdOk returns a tuple with the UplimitSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollUserIntoCourseSchema) GetUplimitSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.UplimitSessionId) {
		return nil, false
	}
	return o.UplimitSessionId, true
}

// HasUplimitSessionId returns a boolean if a field has been set.
func (o *EnrollUserIntoCourseSchema) HasUplimitSessionId() bool {
	if o != nil && !IsNil(o.UplimitSessionId) {
		return true
	}

	return false
}

// SetUplimitSessionId gets a reference to the given string and assigns it to the UplimitSessionId field.
func (o *EnrollUserIntoCourseSchema) SetUplimitSessionId(v string) {
	o.UplimitSessionId = &v
}

// GetSubscriptionCommitmentId returns the SubscriptionCommitmentId field value if set, zero value otherwise.
func (o *EnrollUserIntoCourseSchema) GetSubscriptionCommitmentId() string {
	if o == nil || IsNil(o.SubscriptionCommitmentId) {
		var ret string
		return ret
	}
	return *o.SubscriptionCommitmentId
}

// GetSubscriptionCommitmentIdOk returns a tuple with the SubscriptionCommitmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnrollUserIntoCourseSchema) GetSubscriptionCommitmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionCommitmentId) {
		return nil, false
	}
	return o.SubscriptionCommitmentId, true
}

// HasSubscriptionCommitmentId returns a boolean if a field has been set.
func (o *EnrollUserIntoCourseSchema) HasSubscriptionCommitmentId() bool {
	if o != nil && !IsNil(o.SubscriptionCommitmentId) {
		return true
	}

	return false
}

// SetSubscriptionCommitmentId gets a reference to the given string and assigns it to the SubscriptionCommitmentId field.
func (o *EnrollUserIntoCourseSchema) SetSubscriptionCommitmentId(v string) {
	o.SubscriptionCommitmentId = &v
}

func (o EnrollUserIntoCourseSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnrollUserIntoCourseSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	if !IsNil(o.SessionId) {
		toSerialize["sessionId"] = o.SessionId
	}
	if !IsNil(o.UplimitSessionId) {
		toSerialize["uplimitSessionId"] = o.UplimitSessionId
	}
	if !IsNil(o.SubscriptionCommitmentId) {
		toSerialize["subscriptionCommitmentId"] = o.SubscriptionCommitmentId
	}
	return toSerialize, nil
}

func (o *EnrollUserIntoCourseSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailAddress",
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

	varEnrollUserIntoCourseSchema := _EnrollUserIntoCourseSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEnrollUserIntoCourseSchema)

	if err != nil {
		return err
	}

	*o = EnrollUserIntoCourseSchema(varEnrollUserIntoCourseSchema)

	return err
}

type NullableEnrollUserIntoCourseSchema struct {
	value *EnrollUserIntoCourseSchema
	isSet bool
}

func (v NullableEnrollUserIntoCourseSchema) Get() *EnrollUserIntoCourseSchema {
	return v.value
}

func (v *NullableEnrollUserIntoCourseSchema) Set(val *EnrollUserIntoCourseSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEnrollUserIntoCourseSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEnrollUserIntoCourseSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnrollUserIntoCourseSchema(val *EnrollUserIntoCourseSchema) *NullableEnrollUserIntoCourseSchema {
	return &NullableEnrollUserIntoCourseSchema{value: val, isSet: true}
}

func (v NullableEnrollUserIntoCourseSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnrollUserIntoCourseSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


