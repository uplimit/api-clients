# ToggleUserActivationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**SetIsActive** | **bool** | Whether to set the user as active or inactive. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


