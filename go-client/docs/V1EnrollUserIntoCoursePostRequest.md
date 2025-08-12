# V1EnrollUserIntoCoursePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**UplimitCourseId** | **string** | Internal ID to identify the course across the Uplimit platform. | 
**UplimitEnrollUserIntoCourseSessionSelectionPolicy** | **string** | The policy to decide which session to enroll a user into when enrolling the user into a course. | 
**SubscriptionCommitmentId** | Pointer to **string** | Internal ID to identify the \&quot;group\&quot; the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Methods

### NewV1EnrollUserIntoCoursePostRequest

`func NewV1EnrollUserIntoCoursePostRequest(emailAddress string, uplimitCourseId string, uplimitEnrollUserIntoCourseSessionSelectionPolicy string, ) *V1EnrollUserIntoCoursePostRequest`

NewV1EnrollUserIntoCoursePostRequest instantiates a new V1EnrollUserIntoCoursePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EnrollUserIntoCoursePostRequestWithDefaults

`func NewV1EnrollUserIntoCoursePostRequestWithDefaults() *V1EnrollUserIntoCoursePostRequest`

NewV1EnrollUserIntoCoursePostRequestWithDefaults instantiates a new V1EnrollUserIntoCoursePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V1EnrollUserIntoCoursePostRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V1EnrollUserIntoCoursePostRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V1EnrollUserIntoCoursePostRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetUplimitCourseId

`func (o *V1EnrollUserIntoCoursePostRequest) GetUplimitCourseId() string`

GetUplimitCourseId returns the UplimitCourseId field if non-nil, zero value otherwise.

### GetUplimitCourseIdOk

`func (o *V1EnrollUserIntoCoursePostRequest) GetUplimitCourseIdOk() (*string, bool)`

GetUplimitCourseIdOk returns a tuple with the UplimitCourseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitCourseId

`func (o *V1EnrollUserIntoCoursePostRequest) SetUplimitCourseId(v string)`

SetUplimitCourseId sets UplimitCourseId field to given value.


### GetUplimitEnrollUserIntoCourseSessionSelectionPolicy

`func (o *V1EnrollUserIntoCoursePostRequest) GetUplimitEnrollUserIntoCourseSessionSelectionPolicy() string`

GetUplimitEnrollUserIntoCourseSessionSelectionPolicy returns the UplimitEnrollUserIntoCourseSessionSelectionPolicy field if non-nil, zero value otherwise.

### GetUplimitEnrollUserIntoCourseSessionSelectionPolicyOk

`func (o *V1EnrollUserIntoCoursePostRequest) GetUplimitEnrollUserIntoCourseSessionSelectionPolicyOk() (*string, bool)`

GetUplimitEnrollUserIntoCourseSessionSelectionPolicyOk returns a tuple with the UplimitEnrollUserIntoCourseSessionSelectionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitEnrollUserIntoCourseSessionSelectionPolicy

`func (o *V1EnrollUserIntoCoursePostRequest) SetUplimitEnrollUserIntoCourseSessionSelectionPolicy(v string)`

SetUplimitEnrollUserIntoCourseSessionSelectionPolicy sets UplimitEnrollUserIntoCourseSessionSelectionPolicy field to given value.


### GetSubscriptionCommitmentId

`func (o *V1EnrollUserIntoCoursePostRequest) GetSubscriptionCommitmentId() string`

GetSubscriptionCommitmentId returns the SubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetSubscriptionCommitmentIdOk

`func (o *V1EnrollUserIntoCoursePostRequest) GetSubscriptionCommitmentIdOk() (*string, bool)`

GetSubscriptionCommitmentIdOk returns a tuple with the SubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCommitmentId

`func (o *V1EnrollUserIntoCoursePostRequest) SetSubscriptionCommitmentId(v string)`

SetSubscriptionCommitmentId sets SubscriptionCommitmentId field to given value.

### HasSubscriptionCommitmentId

`func (o *V1EnrollUserIntoCoursePostRequest) HasSubscriptionCommitmentId() bool`

HasSubscriptionCommitmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


