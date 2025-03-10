# CreateUserResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplimitSubscriptionEnrollmentId** | **string** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**UplimitUserId** | **string** | Internal ID to identify the user across the Uplimit platform. | 

## Methods

### NewCreateUserResponseSchema

`func NewCreateUserResponseSchema(uplimitSubscriptionEnrollmentId string, uplimitUserId string, ) *CreateUserResponseSchema`

NewCreateUserResponseSchema instantiates a new CreateUserResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserResponseSchemaWithDefaults

`func NewCreateUserResponseSchemaWithDefaults() *CreateUserResponseSchema`

NewCreateUserResponseSchemaWithDefaults instantiates a new CreateUserResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplimitSubscriptionEnrollmentId

`func (o *CreateUserResponseSchema) GetUplimitSubscriptionEnrollmentId() string`

GetUplimitSubscriptionEnrollmentId returns the UplimitSubscriptionEnrollmentId field if non-nil, zero value otherwise.

### GetUplimitSubscriptionEnrollmentIdOk

`func (o *CreateUserResponseSchema) GetUplimitSubscriptionEnrollmentIdOk() (*string, bool)`

GetUplimitSubscriptionEnrollmentIdOk returns a tuple with the UplimitSubscriptionEnrollmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSubscriptionEnrollmentId

`func (o *CreateUserResponseSchema) SetUplimitSubscriptionEnrollmentId(v string)`

SetUplimitSubscriptionEnrollmentId sets UplimitSubscriptionEnrollmentId field to given value.


### GetUplimitUserId

`func (o *CreateUserResponseSchema) GetUplimitUserId() string`

GetUplimitUserId returns the UplimitUserId field if non-nil, zero value otherwise.

### GetUplimitUserIdOk

`func (o *CreateUserResponseSchema) GetUplimitUserIdOk() (*string, bool)`

GetUplimitUserIdOk returns a tuple with the UplimitUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitUserId

`func (o *CreateUserResponseSchema) SetUplimitUserId(v string)`

SetUplimitUserId sets UplimitUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


