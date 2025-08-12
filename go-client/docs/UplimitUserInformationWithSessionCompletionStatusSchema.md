# UplimitUserInformationWithSessionCompletionStatusSchema

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

### NewUplimitUserInformationWithSessionCompletionStatusSchema

`func NewUplimitUserInformationWithSessionCompletionStatusSchema(emailAddress string, firstName string, lastName string, userAccountIsActive bool, userHasValidSubscriptionEnrollment bool, uplimitSubscriptionEnrollmentId string, uplimitSubscriptionCommitmentId string, uplimitActiveSubscriptionCommitmentIds []string, uplimitUserId string, sessionCompletionStatus string, ) *UplimitUserInformationWithSessionCompletionStatusSchema`

NewUplimitUserInformationWithSessionCompletionStatusSchema instantiates a new UplimitUserInformationWithSessionCompletionStatusSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplimitUserInformationWithSessionCompletionStatusSchemaWithDefaults

`func NewUplimitUserInformationWithSessionCompletionStatusSchemaWithDefaults() *UplimitUserInformationWithSessionCompletionStatusSchema`

NewUplimitUserInformationWithSessionCompletionStatusSchemaWithDefaults instantiates a new UplimitUserInformationWithSessionCompletionStatusSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserAccountIsActive

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUserAccountIsActive() bool`

GetUserAccountIsActive returns the UserAccountIsActive field if non-nil, zero value otherwise.

### GetUserAccountIsActiveOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUserAccountIsActiveOk() (*bool, bool)`

GetUserAccountIsActiveOk returns a tuple with the UserAccountIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountIsActive

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetUserAccountIsActive(v bool)`

SetUserAccountIsActive sets UserAccountIsActive field to given value.


### GetUserHasValidSubscriptionEnrollment

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUserHasValidSubscriptionEnrollment() bool`

GetUserHasValidSubscriptionEnrollment returns the UserHasValidSubscriptionEnrollment field if non-nil, zero value otherwise.

### GetUserHasValidSubscriptionEnrollmentOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUserHasValidSubscriptionEnrollmentOk() (*bool, bool)`

GetUserHasValidSubscriptionEnrollmentOk returns a tuple with the UserHasValidSubscriptionEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHasValidSubscriptionEnrollment

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetUserHasValidSubscriptionEnrollment(v bool)`

SetUserHasValidSubscriptionEnrollment sets UserHasValidSubscriptionEnrollment field to given value.


### GetUplimitSubscriptionEnrollmentId

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitSubscriptionEnrollmentId() string`

GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionEnrollmentIdOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool)`

GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionEnrollmentId

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetUplimitSubscriptionEnrollmentId(v string)`

SetUplimitSubscriptionEnrollmentId sets UplimitSubscriptionEnrollmentId field to given value.


### GetUplimitSubscriptionCommitmentId

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitSubscriptionCommitmentId() string`

GetUplimitSubscriptionCommitmentId returns the UplimitSubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionCommitmentIdOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitSubscriptionCommitmentIdOk() (*string, bool)`

GetUplimitSubscriptionCommitmentIdOk returns a tuple with the UplimitSubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionCommitmentId

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetUplimitSubscriptionCommitmentId(v string)`

SetUplimitSubscriptionCommitmentId sets UplimitSubscriptionCommitmentId field to given value.


### GetUplimitActiveSubscriptionCommitmentIds

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitActiveSubscriptionCommitmentIds() []string`

GetUplimitActiveSubscriptionCommitmentIds returns the UplimitActiveSubscriptionCommitmentIds field if non-nil, zero value otherwise.

### GetUplimitActiveSubscriptionCommitmentIdsOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitActiveSubscriptionCommitmentIdsOk() (*[]string, bool)`

GetUplimitActiveSubscriptionCommitmentIdsOk returns a tuple with the UplimitActiveSubscriptionCommitmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitActiveSubscriptionCommitmentIds

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetUplimitActiveSubscriptionCommitmentIds(v []string)`

SetUplimitActiveSubscriptionCommitmentIds sets UplimitActiveSubscriptionCommitmentIds field to given value.


### GetUplimitUserId

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitUserId() string`

GetUplimitUserId returns the UplimitUserId field if non-nil, zero value otherwise.

### GetUplimitUserIdOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetUplimitUserIdOk() (*string, bool)`

GetUplimitUserIdOk returns a tuple with the UplimitUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitUserId

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetUplimitUserId(v string)`

SetUplimitUserId sets UplimitUserId field to given value.


### GetSessionCompletionStatus

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetSessionCompletionStatus() string`

GetSessionCompletionStatus returns the SessionCompletionStatus field if non-nil, zero value otherwise.

### GetSessionCompletionStatusOk

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) GetSessionCompletionStatusOk() (*string, bool)`

GetSessionCompletionStatusOk returns a tuple with the SessionCompletionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCompletionStatus

`func (o *UplimitUserInformationWithSessionCompletionStatusSchema) SetSessionCompletionStatus(v string)`

SetSessionCompletionStatus sets SessionCompletionStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


