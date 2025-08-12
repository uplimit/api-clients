# V1UnenrollUserFromSessionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**UplimitSessionId** | **string** | Internal ID to identify the session across the Uplimit platform. | 

## Methods

### NewV1UnenrollUserFromSessionPostRequest

`func NewV1UnenrollUserFromSessionPostRequest(emailAddress string, uplimitSessionId string, ) *V1UnenrollUserFromSessionPostRequest`

NewV1UnenrollUserFromSessionPostRequest instantiates a new V1UnenrollUserFromSessionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1UnenrollUserFromSessionPostRequestWithDefaults

`func NewV1UnenrollUserFromSessionPostRequestWithDefaults() *V1UnenrollUserFromSessionPostRequest`

NewV1UnenrollUserFromSessionPostRequestWithDefaults instantiates a new V1UnenrollUserFromSessionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V1UnenrollUserFromSessionPostRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V1UnenrollUserFromSessionPostRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V1UnenrollUserFromSessionPostRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetUplimitSessionId

`func (o *V1UnenrollUserFromSessionPostRequest) GetUplimitSessionId() string`

GetUplimitSessionId returns the UplimitSessionId field if non-nil, zero value otherwise.

### GetUplimitSessionIdOk

`func (o *V1UnenrollUserFromSessionPostRequest) GetUplimitSessionIdOk() (*string, bool)`

GetUplimitSessionIdOk returns a tuple with the UplimitSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSessionId

`func (o *V1UnenrollUserFromSessionPostRequest) SetUplimitSessionId(v string)`

SetUplimitSessionId sets UplimitSessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


