# V1GetUserInformationEmailAddressOrUserIdGet200Response

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

### NewV1GetUserInformationEmailAddressOrUserIdGet200Response

`func NewV1GetUserInformationEmailAddressOrUserIdGet200Response(emailAddress string, firstName string, lastName string, userAccountIsActive bool, userHasValidSubscriptionEnrollment bool, uplimitSubscriptionEnrollmentId string, uplimitSubscriptionCommitmentId string, uplimitUserId string, ) *V1GetUserInformationEmailAddressOrUserIdGet200Response`

NewV1GetUserInformationEmailAddressOrUserIdGet200Response instantiates a new V1GetUserInformationEmailAddressOrUserIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GetUserInformationEmailAddressOrUserIdGet200ResponseWithDefaults

`func NewV1GetUserInformationEmailAddressOrUserIdGet200ResponseWithDefaults() *V1GetUserInformationEmailAddressOrUserIdGet200Response`

NewV1GetUserInformationEmailAddressOrUserIdGet200ResponseWithDefaults instantiates a new V1GetUserInformationEmailAddressOrUserIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetUserAccountIsActive

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUserAccountIsActive() bool`

GetUserAccountIsActive returns the UserAccountIsActive field if non-nil, zero value otherwise.

### GetUserAccountIsActiveOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUserAccountIsActiveOk() (*bool, bool)`

GetUserAccountIsActiveOk returns a tuple with the UserAccountIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccountIsActive

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetUserAccountIsActive(v bool)`

SetUserAccountIsActive sets UserAccountIsActive field to given value.


### GetUserHasValidSubscriptionEnrollment

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUserHasValidSubscriptionEnrollment() bool`

GetUserHasValidSubscriptionEnrollment returns the UserHasValidSubscriptionEnrollment field if non-nil, zero value otherwise.

### GetUserHasValidSubscriptionEnrollmentOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUserHasValidSubscriptionEnrollmentOk() (*bool, bool)`

GetUserHasValidSubscriptionEnrollmentOk returns a tuple with the UserHasValidSubscriptionEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHasValidSubscriptionEnrollment

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetUserHasValidSubscriptionEnrollment(v bool)`

SetUserHasValidSubscriptionEnrollment sets UserHasValidSubscriptionEnrollment field to given value.


### GetUplimitSubscriptionEnrollmentId

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUplimitSubscriptionEnrollmentId() string`

GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionEnrollmentIdOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool)`

GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionEnrollmentId

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetUplimitSubscriptionEnrollmentId(v string)`

SetUplimitSubscriptionEnrollmentId sets UplimitSubscriptionEnrollmentId field to given value.


### GetUplimitSubscriptionCommitmentId

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUplimitSubscriptionCommitmentId() string`

GetUplimitSubscriptionCommitmentId returns the UplimitSubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionCommitmentIdOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUplimitSubscriptionCommitmentIdOk() (*string, bool)`

GetUplimitSubscriptionCommitmentIdOk returns a tuple with the UplimitSubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionCommitmentId

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetUplimitSubscriptionCommitmentId(v string)`

SetUplimitSubscriptionCommitmentId sets UplimitSubscriptionCommitmentId field to given value.


### GetUplimitUserId

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUplimitUserId() string`

GetUplimitUserId returns the UplimitUserId field if non-nil, zero value otherwise.

### GetUplimitUserIdOk

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) GetUplimitUserIdOk() (*string, bool)`

GetUplimitUserIdOk returns a tuple with the UplimitUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitUserId

`func (o *V1GetUserInformationEmailAddressOrUserIdGet200Response) SetUplimitUserId(v string)`

SetUplimitUserId sets UplimitUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


