# V1ListCoursesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Courses** | [**[]V1ListCoursesGet200ResponseCoursesInner**](V1ListCoursesGet200ResponseCoursesInner.md) |  | 
**TotalCount** | **float32** |  | 

## Methods

### NewV1ListCoursesGet200Response

`func NewV1ListCoursesGet200Response(courses []V1ListCoursesGet200ResponseCoursesInner, totalCount float32, ) *V1ListCoursesGet200Response`

NewV1ListCoursesGet200Response instantiates a new V1ListCoursesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListCoursesGet200ResponseWithDefaults

`func NewV1ListCoursesGet200ResponseWithDefaults() *V1ListCoursesGet200Response`

NewV1ListCoursesGet200ResponseWithDefaults instantiates a new V1ListCoursesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourses

`func (o *V1ListCoursesGet200Response) GetCourses() []V1ListCoursesGet200ResponseCoursesInner`

GetCourses returns the Courses field if non-nil, zero value otherwise.

### GetCoursesOk

`func (o *V1ListCoursesGet200Response) GetCoursesOk() (*[]V1ListCoursesGet200ResponseCoursesInner, bool)`

GetCoursesOk returns a tuple with the Courses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourses

`func (o *V1ListCoursesGet200Response) SetCourses(v []V1ListCoursesGet200ResponseCoursesInner)`

SetCourses sets Courses field to given value.


### GetTotalCount

`func (o *V1ListCoursesGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1ListCoursesGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1ListCoursesGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


