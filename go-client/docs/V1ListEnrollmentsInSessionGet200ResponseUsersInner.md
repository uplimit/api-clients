# V1ListEnrollmentsInSessionGet200ResponseUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**FirstName** | **string** | The first name of the user. | 
**LastName** | **string** | The last name of the user. | 
**UserAccountIsActive** | **bool** | Whether the user is allowed to access the Uplimit platform. | 
**UserHasValidSubscriptionEnrollment** | **bool** | Whether the user is activated in your organization. | 
**UplimitSubscriptionEnrollmentId** | **string** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**UplimitSubscriptionCommitmentId** | **string** | Internal ID to identify the \&quot;group\&quot; the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | 
**UplimitActiveSubscriptionCommitmentIds** | **[]string** | All the active subscription commitment ids for this user within this organization. | 
**UplimitUserId** | **string** | Internal ID to identify the user across the Uplimit platform. | 
**SessionCompletionStatus** | **string** | Whether the user has completed the session according to pre-defined completion criteria. | 

## Methods

### NewV1ListEnrollmentsInSessionGet200ResponseUsersInner

`func NewV1ListEnrollmentsInSessionGet200ResponseUsersInner(emailAddress string, firstName string, lastName string, userAccountIsActive bool, userHasValidSubscriptionEnrollment bool, uplimitSubscriptionEnrollmentId string, uplimitSubscriptionCommitmentId string, uplimitActiveSubscriptionCommitmentIds []string, uplimitUserId string, sessionCompletionStatus string, ) *V1ListEnrollmentsInSessionGet200ResponseUsersInner`

NewV1ListEnrollmentsInSessionGet200ResponseUsersInner instantiates a new V1ListEnrollmentsInSessionGet200ResponseUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListEnrollmentsInSessionGet200ResponseUsersInnerWithDefaults

`func NewV1ListEnrollmentsInSessionGet200ResponseUsersInnerWithDefaults() *V1ListEnrollmentsInSessionGet200ResponseUsersInner`

NewV1ListEnrollmentsInSessionGet200ResponseUsersInnerWithDefaults instantiates a new V1ListEnrollmentsInSessionGet200ResponseUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserAccountIsActive

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserAccountIsActive() bool`

GetUserAccountIsActive returns the UserAccountIsActive field if non-nil, zero value otherwise.

### GetUserAccountIsActiveOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserAccountIsActiveOk() (*bool, bool)`

GetUserAccountIsActiveOk returns a tuple with the UserAccountIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountIsActive

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUserAccountIsActive(v bool)`

SetUserAccountIsActive sets UserAccountIsActive field to given value.


### GetUserHasValidSubscriptionEnrollment

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserHasValidSubscriptionEnrollment() bool`

GetUserHasValidSubscriptionEnrollment returns the UserHasValidSubscriptionEnrollment field if non-nil, zero value otherwise.

### GetUserHasValidSubscriptionEnrollmentOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUserHasValidSubscriptionEnrollmentOk() (*bool, bool)`

GetUserHasValidSubscriptionEnrollmentOk returns a tuple with the UserHasValidSubscriptionEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHasValidSubscriptionEnrollment

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUserHasValidSubscriptionEnrollment(v bool)`

SetUserHasValidSubscriptionEnrollment sets UserHasValidSubscriptionEnrollment field to given value.


### GetUplimitSubscriptionEnrollmentId

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionEnrollmentId() string`

GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionEnrollmentIdOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool)`

GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionEnrollmentId

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitSubscriptionEnrollmentId(v string)`

SetUplimitSubscriptionEnrollmentId sets UplimitSubscriptionEnrollmentId field to given value.


### GetUplimitSubscriptionCommitmentId

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionCommitmentId() string`

GetUplimitSubscriptionCommitmentId returns the UplimitSubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionCommitmentIdOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitSubscriptionCommitmentIdOk() (*string, bool)`

GetUplimitSubscriptionCommitmentIdOk returns a tuple with the UplimitSubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionCommitmentId

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitSubscriptionCommitmentId(v string)`

SetUplimitSubscriptionCommitmentId sets UplimitSubscriptionCommitmentId field to given value.


### GetUplimitActiveSubscriptionCommitmentIds

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitActiveSubscriptionCommitmentIds() []string`

GetUplimitActiveSubscriptionCommitmentIds returns the UplimitActiveSubscriptionCommitmentIds field if non-nil, zero value otherwise.

### GetUplimitActiveSubscriptionCommitmentIdsOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitActiveSubscriptionCommitmentIdsOk() (*[]string, bool)`

GetUplimitActiveSubscriptionCommitmentIdsOk returns a tuple with the UplimitActiveSubscriptionCommitmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitActiveSubscriptionCommitmentIds

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitActiveSubscriptionCommitmentIds(v []string)`

SetUplimitActiveSubscriptionCommitmentIds sets UplimitActiveSubscriptionCommitmentIds field to given value.


### GetUplimitUserId

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitUserId() string`

GetUplimitUserId returns the UplimitUserId field if non-nil, zero value otherwise.

### GetUplimitUserIdOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetUplimitUserIdOk() (*string, bool)`

GetUplimitUserIdOk returns a tuple with the UplimitUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitUserId

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetUplimitUserId(v string)`

SetUplimitUserId sets UplimitUserId field to given value.


### GetSessionCompletionStatus

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetSessionCompletionStatus() string`

GetSessionCompletionStatus returns the SessionCompletionStatus field if non-nil, zero value otherwise.

### GetSessionCompletionStatusOk

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) GetSessionCompletionStatusOk() (*string, bool)`

GetSessionCompletionStatusOk returns a tuple with the SessionCompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCompletionStatus

`func (o *V1ListEnrollmentsInSessionGet200ResponseUsersInner) SetSessionCompletionStatus(v string)`

SetSessionCompletionStatus sets SessionCompletionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


