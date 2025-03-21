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

// checks if the V1EnrollUserIntoSessionPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EnrollUserIntoSessionPostRequest{}

// V1EnrollUserIntoSessionPostRequest struct for V1EnrollUserIntoSessionPostRequest
type V1EnrollUserIntoSessionPostRequest struct {
	// The email address of the user.
	EmailAddress string `json:"emailAddress"`
	// The ID of the session to enroll the user into. You must provide either this or uplimitSessionId.
	SessionId *string `json:"sessionId,omitempty"`
	// Internal ID to identify the session across the Uplimit platform.
	UplimitSessionId *string `json:"uplimitSessionId,omitempty"`
	// Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group.
	SubscriptionCommitmentId *string `json:"subscriptionCommitmentId,omitempty"`
}

type _V1EnrollUserIntoSessionPostRequest V1EnrollUserIntoSessionPostRequest

// NewV1EnrollUserIntoSessionPostRequest instantiates a new V1EnrollUserIntoSessionPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EnrollUserIntoSessionPostRequest(emailAddress string) *V1EnrollUserIntoSessionPostRequest {
	this := V1EnrollUserIntoSessionPostRequest{}
	this.EmailAddress = emailAddress
	return &this
}

// NewV1EnrollUserIntoSessionPostRequestWithDefaults instantiates a new V1EnrollUserIntoSessionPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EnrollUserIntoSessionPostRequestWithDefaults() *V1EnrollUserIntoSessionPostRequest {
	this := V1EnrollUserIntoSessionPostRequest{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *V1EnrollUserIntoSessionPostRequest) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *V1EnrollUserIntoSessionPostRequest) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *V1EnrollUserIntoSessionPostRequest) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *V1EnrollUserIntoSessionPostRequest) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnrollUserIntoSessionPostRequest) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *V1EnrollUserIntoSessionPostRequest) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *V1EnrollUserIntoSessionPostRequest) SetSessionId(v string) {
	o.SessionId = &v
}

// GetUplimitSessionId returns the UplimitSessionId field value if set, zero value otherwise.
func (o *V1EnrollUserIntoSessionPostRequest) GetUplimitSessionId() string {
	if o == nil || IsNil(o.UplimitSessionId) {
		var ret string
		return ret
	}
	return *o.UplimitSessionId
}

// GetUplimitSessionIdOk returns a tuple with the UplimitSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnrollUserIntoSessionPostRequest) GetUplimitSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.UplimitSessionId) {
		return nil, false
	}
	return o.UplimitSessionId, true
}

// HasUplimitSessionId returns a boolean if a field has been set.
func (o *V1EnrollUserIntoSessionPostRequest) HasUplimitSessionId() bool {
	if o != nil && !IsNil(o.UplimitSessionId) {
		return true
	}

	return false
}

// SetUplimitSessionId gets a reference to the given string and assigns it to the UplimitSessionId field.
func (o *V1EnrollUserIntoSessionPostRequest) SetUplimitSessionId(v string) {
	o.UplimitSessionId = &v
}

// GetSubscriptionCommitmentId returns the SubscriptionCommitmentId field value if set, zero value otherwise.
func (o *V1EnrollUserIntoSessionPostRequest) GetSubscriptionCommitmentId() string {
	if o == nil || IsNil(o.SubscriptionCommitmentId) {
		var ret string
		return ret
	}
	return *o.SubscriptionCommitmentId
}

// GetSubscriptionCommitmentIdOk returns a tuple with the SubscriptionCommitmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EnrollUserIntoSessionPostRequest) GetSubscriptionCommitmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionCommitmentId) {
		return nil, false
	}
	return o.SubscriptionCommitmentId, true
}

// HasSubscriptionCommitmentId returns a boolean if a field has been set.
func (o *V1EnrollUserIntoSessionPostRequest) HasSubscriptionCommitmentId() bool {
	if o != nil && !IsNil(o.SubscriptionCommitmentId) {
		return true
	}

	return false
}

// SetSubscriptionCommitmentId gets a reference to the given string and assigns it to the SubscriptionCommitmentId field.
func (o *V1EnrollUserIntoSessionPostRequest) SetSubscriptionCommitmentId(v string) {
	o.SubscriptionCommitmentId = &v
}

func (o V1EnrollUserIntoSessionPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EnrollUserIntoSessionPostRequest) ToMap() (map[string]interface{}, error) {
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

func (o *V1EnrollUserIntoSessionPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varV1EnrollUserIntoSessionPostRequest := _V1EnrollUserIntoSessionPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1EnrollUserIntoSessionPostRequest)

	if err != nil {
		return err
	}

	*o = V1EnrollUserIntoSessionPostRequest(varV1EnrollUserIntoSessionPostRequest)

	return err
}

type NullableV1EnrollUserIntoSessionPostRequest struct {
	value *V1EnrollUserIntoSessionPostRequest
	isSet bool
}

func (v NullableV1EnrollUserIntoSessionPostRequest) Get() *V1EnrollUserIntoSessionPostRequest {
	return v.value
}

func (v *NullableV1EnrollUserIntoSessionPostRequest) Set(val *V1EnrollUserIntoSessionPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EnrollUserIntoSessionPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EnrollUserIntoSessionPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EnrollUserIntoSessionPostRequest(val *V1EnrollUserIntoSessionPostRequest) *NullableV1EnrollUserIntoSessionPostRequest {
	return &NullableV1EnrollUserIntoSessionPostRequest{value: val, isSet: true}
}

func (v NullableV1EnrollUserIntoSessionPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EnrollUserIntoSessionPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


