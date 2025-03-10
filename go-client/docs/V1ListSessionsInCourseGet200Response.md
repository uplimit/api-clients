# V1ListSessionsInCourseGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | [**[]V1ListSessionsInCourseGet200ResponseSessionsInner**](V1ListSessionsInCourseGet200ResponseSessionsInner.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewV1ListSessionsInCourseGet200Response

`func NewV1ListSessionsInCourseGet200Response(sessions []V1ListSessionsInCourseGet200ResponseSessionsInner, totalCount float32, ) *V1ListSessionsInCourseGet200Response`

NewV1ListSessionsInCourseGet200Response instantiates a new V1ListSessionsInCourseGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListSessionsInCourseGet200ResponseWithDefaults

`func NewV1ListSessionsInCourseGet200ResponseWithDefaults() *V1ListSessionsInCourseGet200Response`

NewV1ListSessionsInCourseGet200ResponseWithDefaults instantiates a new V1ListSessionsInCourseGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *V1ListSessionsInCourseGet200Response) GetSessions() []V1ListSessionsInCourseGet200ResponseSessionsInner`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *V1ListSessionsInCourseGet200Response) GetSessionsOk() (*[]V1ListSessionsInCourseGet200ResponseSessionsInner, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *V1ListSessionsInCourseGet200Response) SetSessions(v []V1ListSessionsInCourseGet200ResponseSessionsInner)`

SetSessions sets Sessions field to given value.


### GetTotalCount

`func (o *V1ListSessionsInCourseGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1ListSessionsInCourseGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1ListSessionsInCourseGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


