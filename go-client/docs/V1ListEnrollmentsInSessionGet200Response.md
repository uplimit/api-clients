# V1ListEnrollmentsInSessionGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]V1ListEnrollmentsInSessionGet200ResponseUsersInner**](V1ListEnrollmentsInSessionGet200ResponseUsersInner.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewV1ListEnrollmentsInSessionGet200Response

`func NewV1ListEnrollmentsInSessionGet200Response(users []V1ListEnrollmentsInSessionGet200ResponseUsersInner, totalCount float32, ) *V1ListEnrollmentsInSessionGet200Response`

NewV1ListEnrollmentsInSessionGet200Response instantiates a new V1ListEnrollmentsInSessionGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListEnrollmentsInSessionGet200ResponseWithDefaults

`func NewV1ListEnrollmentsInSessionGet200ResponseWithDefaults() *V1ListEnrollmentsInSessionGet200Response`

NewV1ListEnrollmentsInSessionGet200ResponseWithDefaults instantiates a new V1ListEnrollmentsInSessionGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *V1ListEnrollmentsInSessionGet200Response) GetUsers() []V1ListEnrollmentsInSessionGet200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *V1ListEnrollmentsInSessionGet200Response) GetUsersOk() (*[]V1ListEnrollmentsInSessionGet200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *V1ListEnrollmentsInSessionGet200Response) SetUsers(v []V1ListEnrollmentsInSessionGet200ResponseUsersInner)`

SetUsers sets Users field to given value.


### GetTotalCount

`func (o *V1ListEnrollmentsInSessionGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1ListEnrollmentsInSessionGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1ListEnrollmentsInSessionGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


