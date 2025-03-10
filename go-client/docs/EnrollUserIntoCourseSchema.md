# EnrollUserIntoCourseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**SessionId** | Pointer to **string** | The ID of the session to enroll the user into. You must provide either this or uplimitSessionId. | [optional] 
**UplimitSessionId** | Pointer to **string** | Internal ID to identify the session across the Uplimit platform. | [optional] 
**SubscriptionCommitmentId** | Pointer to **string** | Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Methods

### NewEnrollUserIntoCourseSchema

`func NewEnrollUserIntoCourseSchema(emailAddress string, ) *EnrollUserIntoCourseSchema`

NewEnrollUserIntoCourseSchema instantiates a new EnrollUserIntoCourseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollUserIntoCourseSchemaWithDefaults

`func NewEnrollUserIntoCourseSchemaWithDefaults() *EnrollUserIntoCourseSchema`

NewEnrollUserIntoCourseSchemaWithDefaults instantiates a new EnrollUserIntoCourseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *EnrollUserIntoCourseSchema) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *EnrollUserIntoCourseSchema) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *EnrollUserIntoCourseSchema) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetSessionId

`func (o *EnrollUserIntoCourseSchema) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *EnrollUserIntoCourseSchema) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *EnrollUserIntoCourseSchema) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *EnrollUserIntoCourseSchema) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetUplimitSessionId

`func (o *EnrollUserIntoCourseSchema) GetUplimitSessionId() string`

GetUplimitSessionId returns the UplimitSessionId field if non-nil, zero value otherwise.

### GetUplimitSessionIdOk

`func (o *EnrollUserIntoCourseSchema) GetUplimitSessionIdOk() (*string, bool)`

GetUplimitSessionIdOk returns a tuple with the UplimitSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSessionId

`func (o *EnrollUserIntoCourseSchema) SetUplimitSessionId(v string)`

SetUplimitSessionId sets UplimitSessionId field to given value.

### HasUplimitSessionId

`func (o *EnrollUserIntoCourseSchema) HasUplimitSessionId() bool`

HasUplimitSessionId returns a boolean if a field has been set.

### GetSubscriptionCommitmentId

`func (o *EnrollUserIntoCourseSchema) GetSubscriptionCommitmentId() string`

GetSubscriptionCommitmentId returns the SubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetSubscriptionCommitmentIdOk

`func (o *EnrollUserIntoCourseSchema) GetSubscriptionCommitmentIdOk() (*string, bool)`

GetSubscriptionCommitmentIdOk returns a tuple with the SubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCommitmentId

`func (o *EnrollUserIntoCourseSchema) SetSubscriptionCommitmentId(v string)`

SetSubscriptionCommitmentId sets SubscriptionCommitmentId field to given value.

### HasSubscriptionCommitmentId

`func (o *EnrollUserIntoCourseSchema) HasSubscriptionCommitmentId() bool`

HasSubscriptionCommitmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


