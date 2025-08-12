# ToggleUserActivationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**SetIsActive** | **bool** | Whether to set the user as active or inactive. | 
**SubscriptionCommitmentId** | Pointer to **string** | (optional) The subscription commitment id to target. If not provided, the user will be activated on the default subscription commitment, or deactivated across all their subscription commitments. | [optional] 
**DoNotSendWelcomeEmail** | Pointer to **bool** | (optional) Whether to send the welcome email to the user when reactivating them. If not provided, the welcome email will be sent. This option is ignored when deactivating the user. | [optional] 

## Methods

### NewToggleUserActivationSchema

`func NewToggleUserActivationSchema(emailAddress string, setIsActive bool, ) *ToggleUserActivationSchema`

NewToggleUserActivationSchema instantiates a new ToggleUserActivationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToggleUserActivationSchemaWithDefaults

`func NewToggleUserActivationSchemaWithDefaults() *ToggleUserActivationSchema`

NewToggleUserActivationSchemaWithDefaults instantiates a new ToggleUserActivationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *ToggleUserActivationSchema) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ToggleUserActivationSchema) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ToggleUserActivationSchema) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetSetIsActive

`func (o *ToggleUserActivationSchema) GetSetIsActive() bool`

GetSetIsActive returns the SetIsActive field if non-nil, zero value otherwise.

### GetSetIsActiveOk

`func (o *ToggleUserActivationSchema) GetSetIsActiveOk() (*bool, bool)`

GetSetIsActiveOk returns a tuple with the SetIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetIsActive

`func (o *ToggleUserActivationSchema) SetSetIsActive(v bool)`

SetSetIsActive sets SetIsActive field to given value.


### GetSubscriptionCommitmentId

`func (o *ToggleUserActivationSchema) GetSubscriptionCommitmentId() string`

GetSubscriptionCommitmentId returns the SubscriptionCommitmentId field if non-nil, zero value otherwise.

### GetSubscriptionCommitmentIdOk

`func (o *ToggleUserActivationSchema) GetSubscriptionCommitmentIdOk() (*string, bool)`

GetSubscriptionCommitmentIdOk returns a tuple with the SubscriptionCommitmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCommitmentId

`func (o *ToggleUserActivationSchema) SetSubscriptionCommitmentId(v string)`

SetSubscriptionCommitmentId sets SubscriptionCommitmentId field to given value.

### HasSubscriptionCommitmentId

`func (o *ToggleUserActivationSchema) HasSubscriptionCommitmentId() bool`

HasSubscriptionCommitmentId returns a boolean if a field has been set.

### GetDoNotSendWelcomeEmail

`func (o *ToggleUserActivationSchema) GetDoNotSendWelcomeEmail() bool`

GetDoNotSendWelcomeEmail returns the DoNotSendWelcomeEmail field if non-nil, zero value otherwise.

### GetDoNotSendWelcomeEmailOk

`func (o *ToggleUserActivationSchema) GetDoNotSendWelcomeEmailOk() (*bool, bool)`

GetDoNotSendWelcomeEmailOk returns a tuple with the DoNotSendWelcomeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotSendWelcomeEmail

`func (o *ToggleUserActivationSchema) SetDoNotSendWelcomeEmail(v bool)`

SetDoNotSendWelcomeEmail sets DoNotSendWelcomeEmail field to given value.

### HasDoNotSendWelcomeEmail

`func (o *ToggleUserActivationSchema) HasDoNotSendWelcomeEmail() bool`

HasDoNotSendWelcomeEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


