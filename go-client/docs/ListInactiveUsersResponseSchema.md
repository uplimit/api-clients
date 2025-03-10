# ListInactiveUsersResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]V1GetUserInformationEmailAddressOrUserIdGet200Response**](V1GetUserInformationEmailAddressOrUserIdGet200Response.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewListInactiveUsersResponseSchema

`func NewListInactiveUsersResponseSchema(users []V1GetUserInformationEmailAddressOrUserIdGet200Response, totalCount float32, ) *ListInactiveUsersResponseSchema`

NewListInactiveUsersResponseSchema instantiates a new ListInactiveUsersResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInactiveUsersResponseSchemaWithDefaults

`func NewListInactiveUsersResponseSchemaWithDefaults() *ListInactiveUsersResponseSchema`

NewListInactiveUsersResponseSchemaWithDefaults instantiates a new ListInactiveUsersResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListInactiveUsersResponseSchema) GetUsers() []V1GetUserInformationEmailAddressOrUserIdGet200Response`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListInactiveUsersResponseSchema) GetUsersOk() (*[]V1GetUserInformationEmailAddressOrUserIdGet200Response, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListInactiveUsersResponseSchema) SetUsers(v []V1GetUserInformationEmailAddressOrUserIdGet200Response)`

SetUsers sets Users field to given value.


### GetTotalCount

`func (o *ListInactiveUsersResponseSchema) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListInactiveUsersResponseSchema) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListInactiveUsersResponseSchema) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


