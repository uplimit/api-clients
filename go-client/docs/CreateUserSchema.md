# CreateUserSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**FirstName** | **string** | The first name of the user. | 
**LastName** | **string** | The last name of the user. | 
**SubscriptionCommitmentId** | Pointer to **string** | Internal ID to identify the \&quot;group\&quot; the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 
**DoNotSendWelcomeEmail** | Pointer to **bool** | Whether to send the welcome email to the user. If not provided, the welcome email will be sent. | [optional] 

## Methods

### NewCreateUserSchema

`func NewCreateUserSchema(emailAddress string, firstName string, lastName string, ) *CreateUserSchema`

NewCreateUserSchema instantiates a new CreateUserSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserSchemaWithDefaults

`func NewCreateUserSchemaWithDefaults() *CreateUserSchema`

NewCreateUserSchemaWithDefaults instantiates a new CreateUserSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *CreateUserSchema) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *CreateUserSchema) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *CreateUserSchema) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetFirstName

`func (o *CreateUserSchema) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateUserSchema) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateUserSchema) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *CreateUserSchema) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateUserSchema) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateUserSchema) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetSubscriptionCommitmentId

`func (o *CreateUserSchema) GetSubscriptionCommitmentId() string`

GetSubscriptionCommitmentId returns the SubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetSubscriptionCommitmentIdOk

`func (o *CreateUserSchema) GetSubscriptionCommitmentIdOk() (*string, bool)`

GetSubscriptionCommitmentIdOk returns a tuple with the SubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCommitmentId

`func (o *CreateUserSchema) SetSubscriptionCommitmentId(v string)`

SetSubscriptionCommitmentId sets SubscriptionCommitmentId field to given value.

### HasSubscriptionCommitmentId

`func (o *CreateUserSchema) HasSubscriptionCommitmentId() bool`

HasSubscriptionCommitmentId returns a boolean if a field has been set.

### GetDoNotSendWelcomeEmail

`func (o *CreateUserSchema) GetDoNotSendWelcomeEmail() bool`

GetDoNotSendWelcomeEmail returns the DoNotSendWelcomeEmail field if non-nil, zero value otherwise.

### GetDoNotSendWelcomeEmailOk

`func (o *CreateUserSchema) GetDoNotSendWelcomeEmailOk() (*bool, bool)`

GetDoNotSendWelcomeEmailOk returns a tuple with the DoNotSendWelcomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotSendWelcomeEmail

`func (o *CreateUserSchema) SetDoNotSendWelcomeEmail(v bool)`

SetDoNotSendWelcomeEmail sets DoNotSendWelcomeEmail field to given value.

### HasDoNotSendWelcomeEmail

`func (o *CreateUserSchema) HasDoNotSendWelcomeEmail() bool`

HasDoNotSendWelcomeEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


