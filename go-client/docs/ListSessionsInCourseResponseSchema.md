# ListSessionsInCourseResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | [**[]V1ListSessionsInCourseGet200ResponseSessionsInner**](V1ListSessionsInCourseGet200ResponseSessionsInner.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewListSessionsInCourseResponseSchema

`func NewListSessionsInCourseResponseSchema(sessions []V1ListSessionsInCourseGet200ResponseSessionsInner, totalCount float32, ) *ListSessionsInCourseResponseSchema`

NewListSessionsInCourseResponseSchema instantiates a new ListSessionsInCourseResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSessionsInCourseResponseSchemaWithDefaults

`func NewListSessionsInCourseResponseSchemaWithDefaults() *ListSessionsInCourseResponseSchema`

NewListSessionsInCourseResponseSchemaWithDefaults instantiates a new ListSessionsInCourseResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *ListSessionsInCourseResponseSchema) GetSessions() []V1ListSessionsInCourseGet200ResponseSessionsInner`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListSessionsInCourseResponseSchema) GetSessionsOk() (*[]V1ListSessionsInCourseGet200ResponseSessionsInner, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListSessionsInCourseResponseSchema) SetSessions(v []V1ListSessionsInCourseGet200ResponseSessionsInner)`

SetSessions sets Sessions field to given value.


### GetTotalCount

`func (o *ListSessionsInCourseResponseSchema) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListSessionsInCourseResponseSchema) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListSessionsInCourseResponseSchema) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


