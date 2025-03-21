# UplimitSessionInformationSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UplimitSessionId** | **string** | Internal ID to identify the session across the Uplimit platform. | 
**Name** | **string** | The name of the session. | 
**StartsAt** | **time.Time** | The start date of the session. | 

## Methods

### NewUplimitSessionInformationSchema

`func NewUplimitSessionInformationSchema(uplimitSessionId string, name string, startsAt time.Time, ) *UplimitSessionInformationSchema`

NewUplimitSessionInformationSchema instantiates a new UplimitSessionInformationSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUplimitSessionInformationSchemaWithDefaults

`func NewUplimitSessionInformationSchemaWithDefaults() *UplimitSessionInformationSchema`

NewUplimitSessionInformationSchemaWithDefaults instantiates a new UplimitSessionInformationSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUplimitSessionId

`func (o *UplimitSessionInformationSchema) GetUplimitSessionId() string`

GetUplimitSessionId returns the UplimitSessionId field if non-nil, zero value otherwise.

### GetUplimitSessionIdOk

`func (o *UplimitSessionInformationSchema) GetUplimitSessionIdOk() (*string, bool)`

GetUplimitSessionIdOk returns a tuple with the UplimitSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplimitSessionId

`func (o *UplimitSessionInformationSchema) SetUplimitSessionId(v string)`

SetUplimitSessionId sets UplimitSessionId field to given value.


### GetName

`func (o *UplimitSessionInformationSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UplimitSessionInformationSchema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UplimitSessionInformationSchema) SetName(v string)`

SetName sets Name field to given value.


### GetStartsAt

`func (o *UplimitSessionInformationSchema) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UplimitSessionInformationSchema) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UplimitSessionInformationSchema) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


