# V1ToggleUserActivationPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**SetIsActive** | **bool** | Whether to set the user as active or inactive. | 

## Methods

### NewV1ToggleUserActivationPostRequest

`func NewV1ToggleUserActivationPostRequest(emailAddress string, setIsActive bool, ) *V1ToggleUserActivationPostRequest`

NewV1ToggleUserActivationPostRequest instantiates a new V1ToggleUserActivationPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ToggleUserActivationPostRequestWithDefaults

`func NewV1ToggleUserActivationPostRequestWithDefaults() *V1ToggleUserActivationPostRequest`

NewV1ToggleUserActivationPostRequestWithDefaults instantiates a new V1ToggleUserActivationPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V1ToggleUserActivationPostRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V1ToggleUserActivationPostRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V1ToggleUserActivationPostRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetSetIsActive

`func (o *V1ToggleUserActivationPostRequest) GetSetIsActive() bool`

GetSetIsActive returns the SetIsActive field if non-nil, zero value otherwise.

### GetSetIsActiveOk

`func (o *V1ToggleUserActivationPostRequest) GetSetIsActiveOk() (*bool, bool)`

GetSetIsActiveOk returns a tuple with the SetIsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetIsActive

`func (o *V1ToggleUserActivationPostRequest) SetSetIsActive(v bool)`

SetSetIsActive sets SetIsActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


