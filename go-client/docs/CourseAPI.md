# \CourseAPI

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ListCoursesGet**](CourseAPI.md#V1ListCoursesGet) | **Get** /v1/ListCourses | 
[**V1ListSessionsInCourseGet**](CourseAPI.md#V1ListSessionsInCourseGet) | **Get** /v1/ListSessionsInCourse | 



## V1ListCoursesGet

> V1ListCoursesGet200Response V1ListCoursesGet(ctx).Skip(skip).Take(take).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	skip := int32(56) // int32 |  (optional)
	take := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseAPI.V1ListCoursesGet(context.Background()).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAPI.V1ListCoursesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListCoursesGet`: V1ListCoursesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CourseAPI.V1ListCoursesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ListCoursesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | 
 **take** | **int32** |  | 

### Return type

[**V1ListCoursesGet200Response**](V1ListCoursesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListSessionsInCourseGet

> V1ListSessionsInCourseGet200Response V1ListSessionsInCourseGet(ctx).UplimitCourseId(uplimitCourseId).Skip(skip).Take(take).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	uplimitCourseId := "uplimitCourseId_example" // string | 
	skip := int32(56) // int32 |  (optional)
	take := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseAPI.V1ListSessionsInCourseGet(context.Background()).UplimitCourseId(uplimitCourseId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAPI.V1ListSessionsInCourseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListSessionsInCourseGet`: V1ListSessionsInCourseGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CourseAPI.V1ListSessionsInCourseGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ListSessionsInCourseGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uplimitCourseId** | **string** |  | 
 **skip** | **int32** |  | 
 **take** | **int32** |  | 

### Return type

[**V1ListSessionsInCourseGet200Response**](V1ListSessionsInCourseGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

