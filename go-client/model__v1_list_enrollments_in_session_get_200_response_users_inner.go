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

// checks if the V1ListEnrollmentsInSessionGet200ResponseUsersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1ListEnrollmentsInSessionGet200ResponseUsersInner{}

// V1ListEnrollmentsInSessionGet200ResponseUsersInner struct for V1ListEnrollmentsInSessionGet200ResponseUsersInner
type V1ListEnrollmentsInSessionGet200ResponseUsersInner struct {
	// The email address of the user.
	EmailAddress string `json:"emailAddress"`
	// The first name of the user.
	FirstName string `json:"firstName"`
	// The last name of the user.
	LastName string `json:"lastName"`
	// Whether the user is allowed to access the Uplimit platform.
	UserAccountIsActive bool `json:"userAccountIsActive"`
	// Whether the user is activated in your organization.
	UserHasValidSubscriptionEnrollment bool `json:"userHasValidSubscriptionEnrollment"`
	// Internal ID to identify the user's membership within your organization on Uplimit.
	UplimitSubscriptionEnrollmentId string `json:"uplimitSubscriptionEnrollmentId"`
	// Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group.
	UplimitSubscriptionCommitmentId string `json:"uplimitSubscriptionCommitmentId"`
	// Internal ID to identify the user across the Uplimit platform.
	UplimitUserId string `json:"uplimitUserId"`
	// Whether the user has completed the session according to pre-defined completion criteria.
	SessionCompletionStatus string `json:"sessionCompletionStatus"`
}

type _V1ListEnrollmentsInSessionGet200ResponseUsersInner V1ListEnrollmentsInSessionGet200ResponseUsersInner

// NewV1ListEnrollmentsInSessionGet200ResponseUsersInner instantiates a new V1ListEnrollmentsInSessionGet200ResponseUsersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1ListEnrollmentsInSessionGet200ResponseUsersInner(emailAddress string, firstName string, lastName string, userAccountIsActive bool, userHasValidSubscriptionEnrollment bool, uplimitSubscriptionEnrollmentId string, uplimitSubscriptionCommitmentId string, uplimitUserId string, sessionCompletionStatus string) *V1ListEnrollmentsInSessionGet200ResponseUsersInner {
	this := V1ListEnrollmentsInSessionGet200ResponseUsersInner{}
	this.EmailAddress = emailAddress
	this.FirstName = firstName
	this.LastName = lastName
	this.UserAccountIsActive = userAccountIsActive
	this.UserHasValidSubscriptionEnrollment = userHasValidSubscriptionEnrollment
	this.UplimitSubscriptionEnrollmentId = uplimitSubscriptionEnrollmentId
	this.UplimitSubscriptionCommitmentId = uplimitSubscriptionCommitmentId
	this.UplimitUserId = uplimitUserId
	this.SessionCompletionStatus = sessionCompletionStatus
	return &this
}

// NewV1ListEnrollmentsInSessionGet200ResponseUsersInnerWithDefaults instantiates a new V1ListEnrollmentsInSessionGet200ResponseUsersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1ListEnrollmentsInSessionGet200ResponseUsersInnerWithDefaults() *V1ListEnrollmentsInSessionGet200ResponseUsersInner {
	this := V1ListEnrollmentsInSessionGet200ResponseUsersInner{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetFirstName returns the FirstName field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetLastName(v string) {
	o.LastName = v
}

// GetUserAccountIsActive returns the UserAccountIsActive field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserAccountIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserAccountIsActive
}

// GetUserAccountIsActiveOk returns a tuple with the UserAccountIsActive field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserAccountIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAccountIsActive, true
}

// SetUserAccountIsActive sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUserAccountIsActive(v bool) {
	o.UserAccountIsActive = v
}

// GetUserHasValidSubscriptionEnrollment returns the UserHasValidSubscriptionEnrollment field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserHasValidSubscriptionEnrollment() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserHasValidSubscriptionEnrollment
}

// GetUserHasValidSubscriptionEnrollmentOk returns a tuple with the UserHasValidSubscriptionEnrollment field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserHasValidSubscriptionEnrollmentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserHasValidSubscriptionEnrollment, true
}

// SetUserHasValidSubscriptionEnrollment sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUserHasValidSubscriptionEnrollment(v bool) {
	o.UserHasValidSubscriptionEnrollment = v
}

// GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionEnrollmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitSubscriptionEnrollmentId
}

// GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitSubscriptionEnrollmentId, true
}

// SetUplimitSubscriptionEnrollmentId sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitSubscriptionEnrollmentId(v string) {
	o.UplimitSubscriptionEnrollmentId = v
}

// GetUplimitSubscriptionCommitmentId returns the UplimitSubscriptionCommitmentId field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionCommitmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitSubscriptionCommitmentId
}

// GetUplimitSubscriptionCommitmentIdOk returns a tuple with the UplimitSubscriptionCommitmentId field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionCommitmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitSubscriptionCommitmentId, true
}

// SetUplimitSubscriptionCommitmentId sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitSubscriptionCommitmentId(v string) {
	o.UplimitSubscriptionCommitmentId = v
}

// GetUplimitUserId returns the UplimitUserId field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UplimitUserId
}

// GetUplimitUserIdOk returns a tuple with the UplimitUserId field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplimitUserId, true
}

// SetUplimitUserId sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitUserId(v string) {
	o.UplimitUserId = v
}

// GetSessionCompletionStatus returns the SessionCompletionStatus field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetSessionCompletionStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionCompletionStatus
}

// GetSessionCompletionStatusOk returns a tuple with the SessionCompletionStatus field value
// and a boolean to check if the value has been set.
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetSessionCompletionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionCompletionStatus, true
}

// SetSessionCompletionStatus sets field value
func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetSessionCompletionStatus(v string) {
	o.SessionCompletionStatus = v
}

func (o V1ListEnrollmentsInSessionGet200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1ListEnrollmentsInSessionGet200ResponseUsersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["emailAddress"] = o.EmailAddress
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["userAccountIsActive"] = o.UserAccountIsActive
	toSerialize["userHasValidSubscriptionEnrollment"] = o.UserHasValidSubscriptionEnrollment
	toSerialize["uplimitSubscriptionEnrollmentId"] = o.UplimitSubscriptionEnrollmentId
	toSerialize["uplimitSubscriptionCommitmentId"] = o.UplimitSubscriptionCommitmentId
	toSerialize["uplimitUserId"] = o.UplimitUserId
	toSerialize["sessionCompletionStatus"] = o.SessionCompletionStatus
	return toSerialize, nil
}

func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"emailAddress",
		"firstName",
		"lastName",
		"userAccountIsActive",
		"userHasValidSubscriptionEnrollment",
		"uplimitSubscriptionEnrollmentId",
		"uplimitSubscriptionCommitmentId",
		"uplimitUserId",
		"sessionCompletionStatus",
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

	varV1ListEnrollmentsInSessionGet200ResponseUsersInner := _V1ListEnrollmentsInSessionGet200ResponseUsersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1ListEnrollmentsInSessionGet200ResponseUsersInner)

	if err != nil {
		return err
	}

	*o = V1ListEnrollmentsInSessionGet200ResponseUsersInner(varV1ListEnrollmentsInSessionGet200ResponseUsersInner)

	return err
}

type NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner struct {
	value *V1ListEnrollmentsInSessionGet200ResponseUsersInner
	isSet bool
}

func (v NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner) Get() *V1ListEnrollmentsInSessionGet200ResponseUsersInner {
	return v.value
}

func (v *NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner) Set(val *V1ListEnrollmentsInSessionGet200ResponseUsersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1ListEnrollmentsInSessionGet200ResponseUsersInner(val *V1ListEnrollmentsInSessionGet200ResponseUsersInner) *NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner {
	return &NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner{value: val, isSet: true}
}

func (v NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1ListEnrollmentsInSessionGet200ResponseUsersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


