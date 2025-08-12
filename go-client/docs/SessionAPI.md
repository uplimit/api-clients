# \SessionAPI

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnrollUserIntoSessionPost**](SessionAPI.md#V1EnrollUserIntoSessionPost) | **Post** /v1/EnrollUserIntoSession | 
[**V1ListEnrollmentsInSessionGet**](SessionAPI.md#V1ListEnrollmentsInSessionGet) | **Get** /v1/ListEnrollmentsInSession | 
[**V1ListSessionsInCourseGet**](SessionAPI.md#V1ListSessionsInCourseGet) | **Get** /v1/ListSessionsInCourse | 
[**V1UnenrollUserFromSessionPost**](SessionAPI.md#V1UnenrollUserFromSessionPost) | **Post** /v1/UnenrollUserFromSession | 



## V1EnrollUserIntoSessionPost

> V1EnrollUserIntoSessionPost200Response V1EnrollUserIntoSessionPost(ctx).V1EnrollUserIntoSessionPostRequest(v1EnrollUserIntoSessionPostRequest).Execute()





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
	v1EnrollUserIntoSessionPostRequest := *openapiclient.NewV1EnrollUserIntoSessionPostRequest("EmailAddress_example") // V1EnrollUserIntoSessionPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionAPI.V1EnrollUserIntoSessionPost(context.Background()).V1EnrollUserIntoSessionPostRequest(v1EnrollUserIntoSessionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.V1EnrollUserIntoSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollUserIntoSessionPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.V1EnrollUserIntoSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollUserIntoSessionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1EnrollUserIntoSessionPostRequest** | [**V1EnrollUserIntoSessionPostRequest**](V1EnrollUserIntoSessionPostRequest.md) |  | 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListEnrollmentsInSessionGet

> V1ListEnrollmentsInSessionGet200Response V1ListEnrollmentsInSessionGet(ctx).UplimitSessionId(uplimitSessionId).Skip(skip).Take(take).Execute()





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
	uplimitSessionId := "uplimitSessionId_example" // string | 
	skip := int32(56) // int32 |  (optional)
	take := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionAPI.V1ListEnrollmentsInSessionGet(context.Background()).UplimitSessionId(uplimitSessionId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.V1ListEnrollmentsInSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListEnrollmentsInSessionGet`: V1ListEnrollmentsInSessionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.V1ListEnrollmentsInSessionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ListEnrollmentsInSessionGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uplimitSessionId** | **string** |  | 
 **skip** | **int32** |  | 
 **take** | **int32** |  | 

### Return type

[**V1ListEnrollmentsInSessionGet200Response**](V1ListEnrollmentsInSessionGet200Response.md)

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
	resp, r, err := apiClient.SessionAPI.V1ListSessionsInCourseGet(context.Background()).UplimitCourseId(uplimitCourseId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.V1ListSessionsInCourseGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListSessionsInCourseGet`: V1ListSessionsInCourseGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.V1ListSessionsInCourseGet`: %v\n", resp)
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


## V1UnenrollUserFromSessionPost

> V1EnrollUserIntoSessionPost200Response V1UnenrollUserFromSessionPost(ctx).V1UnenrollUserFromSessionPostRequest(v1UnenrollUserFromSessionPostRequest).Execute()





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
	v1UnenrollUserFromSessionPostRequest := *openapiclient.NewV1UnenrollUserFromSessionPostRequest("EmailAddress_example", "UplimitSessionId_example") // V1UnenrollUserFromSessionPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SessionAPI.V1UnenrollUserFromSessionPost(context.Background()).V1UnenrollUserFromSessionPostRequest(v1UnenrollUserFromSessionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SessionAPI.V1UnenrollUserFromSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UnenrollUserFromSessionPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SessionAPI.V1UnenrollUserFromSessionPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1UnenrollUserFromSessionPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1UnenrollUserFromSessionPostRequest** | [**V1UnenrollUserFromSessionPostRequest**](V1UnenrollUserFromSessionPostRequest.md) |  | 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

