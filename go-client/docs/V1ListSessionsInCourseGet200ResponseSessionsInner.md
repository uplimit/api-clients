# V1ListSessionsInCourseGet200ResponseSessionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplimitSessionId** | **string** | Internal ID to identify the session across the Uplimit platform. | 
**Name** | **string** | The name of the session. | 
**StartsAt** | **time.Time** | The start date of the session. | 
**EnrollmentEnabled** | **bool** | Whether the session allows enrollments | 

## Methods

### NewV1ListSessionsInCourseGet200ResponseSessionsInner

`func NewV1ListSessionsInCourseGet200ResponseSessionsInner(uplimitSessionId string, name string, startsAt time.Time, enrollmentEnabled bool, ) *V1ListSessionsInCourseGet200ResponseSessionsInner`

NewV1ListSessionsInCourseGet200ResponseSessionsInner instantiates a new V1ListSessionsInCourseGet200ResponseSessionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListSessionsInCourseGet200ResponseSessionsInnerWithDefaults

`func NewV1ListSessionsInCourseGet200ResponseSessionsInnerWithDefaults() *V1ListSessionsInCourseGet200ResponseSessionsInner`

NewV1ListSessionsInCourseGet200ResponseSessionsInnerWithDefaults instantiates a new V1ListSessionsInCourseGet200ResponseSessionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplimitSessionId

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetUplimitSessionId() string`

GetUplimitSessionId returns the UplimitSessionId field if non-nil, zero value otherwise.

### GetUplimitSessionIdOk

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetUplimitSessionIdOk() (*string, bool)`

GetUplimitSessionIdOk returns a tuple with the UplimitSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSessionId

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) SetUplimitSessionId(v string)`

SetUplimitSessionId sets UplimitSessionId field to given value.


### GetName

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) SetName(v string)`

SetName sets Name field to given value.


### GetStartsAt

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEnrollmentEnabled

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetEnrollmentEnabled() bool`

GetEnrollmentEnabled returns the EnrollmentEnabled field if non-nil, zero value otherwise.

### GetEnrollmentEnabledOk

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) GetEnrollmentEnabledOk() (*bool, bool)`

GetEnrollmentEnabledOk returns a tuple with the EnrollmentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentEnabled

`func (o *V1ListSessionsInCourseGet200ResponseSessionsInner) SetEnrollmentEnabled(v bool)`

SetEnrollmentEnabled sets EnrollmentEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


