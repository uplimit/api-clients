# ListCoursesResponseSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Courses** | [**[]V1ListCoursesGet200ResponseCoursesInner**](V1ListCoursesGet200ResponseCoursesInner.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewListCoursesResponseSchema

`func NewListCoursesResponseSchema(courses []V1ListCoursesGet200ResponseCoursesInner, totalCount float32, ) *ListCoursesResponseSchema`

NewListCoursesResponseSchema instantiates a new ListCoursesResponseSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCoursesResponseSchemaWithDefaults

`func NewListCoursesResponseSchemaWithDefaults() *ListCoursesResponseSchema`

NewListCoursesResponseSchemaWithDefaults instantiates a new ListCoursesResponseSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourses

`func (o *ListCoursesResponseSchema) GetCourses() []V1ListCoursesGet200ResponseCoursesInner`

GetCourses returns the Courses field if non-nil, zero value otherwise.

### GetCoursesOk

`func (o *ListCoursesResponseSchema) GetCoursesOk() (*[]V1ListCoursesGet200ResponseCoursesInner, bool)`

GetCoursesOk returns a tuple with the Courses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourses

`func (o *ListCoursesResponseSchema) SetCourses(v []V1ListCoursesGet200ResponseCoursesInner)`

SetCourses sets Courses field to given value.


### GetTotalCount

`func (o *ListCoursesResponseSchema) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListCoursesResponseSchema) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListCoursesResponseSchema) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


