# V1AddUserAuthenticationMethodPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The email address of the user. | 
**AuthenticationMethod** | **string** | The authentication method to add to the user. | 
**CustomAuthenticationMethodProviderId** | **string** | For SAML and OAUTH2, this is the provider ID (will be supplied by Uplimit to the developer). | 
**AuthenticationSecret** | **string** | The unique identity for the authentication method. For SAML, this is the SAML entity ID. For OAUTH2, this is the subject identifier (or sub). | 

## Methods

### NewV1AddUserAuthenticationMethodPostRequest

`func NewV1AddUserAuthenticationMethodPostRequest(emailAddress string, authenticationMethod string, customAuthenticationMethodProviderId string, authenticationSecret string, ) *V1AddUserAuthenticationMethodPostRequest`

NewV1AddUserAuthenticationMethodPostRequest instantiates a new V1AddUserAuthenticationMethodPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AddUserAuthenticationMethodPostRequestWithDefaults

`func NewV1AddUserAuthenticationMethodPostRequestWithDefaults() *V1AddUserAuthenticationMethodPostRequest`

NewV1AddUserAuthenticationMethodPostRequestWithDefaults instantiates a new V1AddUserAuthenticationMethodPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *V1AddUserAuthenticationMethodPostRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *V1AddUserAuthenticationMethodPostRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *V1AddUserAuthenticationMethodPostRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetAuthenticationMethod

`func (o *V1AddUserAuthenticationMethodPostRequest) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *V1AddUserAuthenticationMethodPostRequest) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *V1AddUserAuthenticationMethodPostRequest) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetCustomAuthenticationMethodProviderId

`func (o *V1AddUserAuthenticationMethodPostRequest) GetCustomAuthenticationMethodProviderId() string`

GetCustomAuthenticationMethodProviderId returns the CustomAuthenticationMethodProviderId field if non-nil, zero value otherwise.

### GetCustomAuthenticationMethodProviderIdOk

`func (o *V1AddUserAuthenticationMethodPostRequest) GetCustomAuthenticationMethodProviderIdOk() (*string, bool)`

GetCustomAuthenticationMethodProviderIdOk returns a tuple with the CustomAuthenticationMethodProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAuthenticationMethodProviderId

`func (o *V1AddUserAuthenticationMethodPostRequest) SetCustomAuthenticationMethodProviderId(v string)`

SetCustomAuthenticationMethodProviderId sets CustomAuthenticationMethodProviderId field to given value.


### GetAuthenticationSecret

`func (o *V1AddUserAuthenticationMethodPostRequest) GetAuthenticationSecret() string`

GetAuthenticationSecret returns the AuthenticationSecret field if non-nil, zero value otherwise.

### GetAuthenticationSecretOk

`func (o *V1AddUserAuthenticationMethodPostRequest) GetAuthenticationSecretOk() (*string, bool)`

GetAuthenticationSecretOk returns a tuple with the AuthenticationSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSecret

`func (o *V1AddUserAuthenticationMethodPostRequest) SetAuthenticationSecret(v string)`

SetAuthenticationSecret sets AuthenticationSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


