# UplimitUserInformationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**FirstName** | **string** | The first name of the user. | 
**LastName** | **string** | The last name of the user. | 
**UserAccountIsActive** | **bool** | Whether the user is allowed to access the Uplimit platform. | 
**UserHasValidSubscriptionEnrollment** | **bool** | Whether the user is activated in your organization. | 
**UplimitSubscriptionEnrollmentId** | **string** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**UplimitSubscriptionCommitmentId** | **string** | Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | 
**UplimitUserId** | **string** | Internal ID to identify the user across the Uplimit platform. | 

## Methods

### NewUplimitUserInformationSchema

`func NewUplimitUserInformationSchema(emailAddress string, firstName string, lastName string, userAccountIsActive bool, userHasValidSubscriptionEnrollment bool, uplimitSubscriptionEnrollmentId string, uplimitSubscriptionCommitmentId string, uplimitUserId string, ) *UplimitUserInformationSchema`

NewUplimitUserInformationSchema instantiates a new UplimitUserInformationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplimitUserInformationSchemaWithDefaults

`func NewUplimitUserInformationSchemaWithDefaults() *UplimitUserInformationSchema`

NewUplimitUserInformationSchemaWithDefaults instantiates a new UplimitUserInformationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *UplimitUserInformationSchema) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UplimitUserInformationSchema) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UplimitUserInformationSchema) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *UplimitUserInformationSchema) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UplimitUserInformationSchema) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UplimitUserInformationSchema) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UplimitUserInformationSchema) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UplimitUserInformationSchema) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UplimitUserInformationSchema) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserAccountIsActive

`func (o *UplimitUserInformationSchema) GetUserAccountIsActive() bool`

GetUserAccountIsActive returns the UserAccountIsActive field if non-nil, zero value otherwise.

### GetUserAccountIsActiveOk

`func (o *UplimitUserInformationSchema) GetUserAccountIsActiveOk() (*bool, bool)`

GetUserAccountIsActiveOk returns a tuple with the UserAccountIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountIsActive

`func (o *UplimitUserInformationSchema) SetUserAccountIsActive(v bool)`

SetUserAccountIsActive sets UserAccountIsActive field to given value.


### GetUserHasValidSubscriptionEnrollment

`func (o *UplimitUserInformationSchema) GetUserHasValidSubscriptionEnrollment() bool`

GetUserHasValidSubscriptionEnrollment returns the UserHasValidSubscriptionEnrollment field if non-nil, zero value otherwise.

### GetUserHasValidSubscriptionEnrollmentOk

`func (o *UplimitUserInformationSchema) GetUserHasValidSubscriptionEnrollmentOk() (*bool, bool)`

GetUserHasValidSubscriptionEnrollmentOk returns a tuple with the UserHasValidSubscriptionEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHasValidSubscriptionEnrollment

`func (o *UplimitUserInformationSchema) SetUserHasValidSubscriptionEnrollment(v bool)`

SetUserHasValidSubscriptionEnrollment sets UserHasValidSubscriptionEnrollment field to given value.


### GetUplimitSubscriptionEnrollmentId

`func (o *UplimitUserInformationSchema) GetUplimitSubscriptionEnrollmentId() string`

GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionEnrollmentIdOk

`func (o *UplimitUserInformationSchema) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool)`

GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionEnrollmentId

`func (o *UplimitUserInformationSchema) SetUplimitSubscriptionEnrollmentId(v string)`

SetUplimitSubscriptionEnrollmentId sets UplimitSubscriptionEnrollmentId field to given value.


### GetUplimitSubscriptionCommitmentId

`func (o *UplimitUserInformationSchema) GetUplimitSubscriptionCommitmentId() string`

GetUplimitSubscriptionCommitmentId returns the UplimitSubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionCommitmentIdOk

`func (o *UplimitUserInformationSchema) GetUplimitSubscriptionCommitmentIdOk() (*string, bool)`

GetUplimitSubscriptionCommitmentIdOk returns a tuple with the UplimitSubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionCommitmentId

`func (o *UplimitUserInformationSchema) SetUplimitSubscriptionCommitmentId(v string)`

SetUplimitSubscriptionCommitmentId sets UplimitSubscriptionCommitmentId field to given value.


### GetUplimitUserId

`func (o *UplimitUserInformationSchema) GetUplimitUserId() string`

GetUplimitUserId returns the UplimitUserId field if non-nil, zero value otherwise.

### GetUplimitUserIdOk

`func (o *UplimitUserInformationSchema) GetUplimitUserIdOk() (*string, bool)`

GetUplimitUserIdOk returns a tuple with the UplimitUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitUserId

`func (o *UplimitUserInformationSchema) SetUplimitUserId(v string)`

SetUplimitUserId sets UplimitUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


