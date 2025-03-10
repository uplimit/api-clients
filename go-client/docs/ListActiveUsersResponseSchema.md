# ListActiveUsersResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]V1GetUserInformationEmailAddressOrUserIdGet200Response**](V1GetUserInformationEmailAddressOrUserIdGet200Response.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewListActiveUsersResponseSchema

`func NewListActiveUsersResponseSchema(users []V1GetUserInformationEmailAddressOrUserIdGet200Response, totalCount float32, ) *ListActiveUsersResponseSchema`

NewListActiveUsersResponseSchema instantiates a new ListActiveUsersResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActiveUsersResponseSchemaWithDefaults

`func NewListActiveUsersResponseSchemaWithDefaults() *ListActiveUsersResponseSchema`

NewListActiveUsersResponseSchemaWithDefaults instantiates a new ListActiveUsersResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListActiveUsersResponseSchema) GetUsers() []V1GetUserInformationEmailAddressOrUserIdGet200Response`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListActiveUsersResponseSchema) GetUsersOk() (*[]V1GetUserInformationEmailAddressOrUserIdGet200Response, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListActiveUsersResponseSchema) SetUsers(v []V1GetUserInformationEmailAddressOrUserIdGet200Response)`

SetUsers sets Users field to given value.


### GetTotalCount

`func (o *ListActiveUsersResponseSchema) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListActiveUsersResponseSchema) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListActiveUsersResponseSchema) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


