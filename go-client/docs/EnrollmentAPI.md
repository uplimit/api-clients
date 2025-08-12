# \EnrollmentAPI

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnrollUserIntoCoursePost**](EnrollmentAPI.md#V1EnrollUserIntoCoursePost) | **Post** /v1/EnrollUserIntoCourse | 
[**V1EnrollUserIntoSessionPost**](EnrollmentAPI.md#V1EnrollUserIntoSessionPost) | **Post** /v1/EnrollUserIntoSession | 
[**V1ListEnrollmentsInSessionGet**](EnrollmentAPI.md#V1ListEnrollmentsInSessionGet) | **Get** /v1/ListEnrollmentsInSession | 
[**V1UnenrollUserFromSessionPost**](EnrollmentAPI.md#V1UnenrollUserFromSessionPost) | **Post** /v1/UnenrollUserFromSession | 



## V1EnrollUserIntoCoursePost

> V1EnrollUserIntoSessionPost200Response V1EnrollUserIntoCoursePost(ctx).V1EnrollUserIntoCoursePostRequest(v1EnrollUserIntoCoursePostRequest).Execute()





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
	v1EnrollUserIntoCoursePostRequest := *openapiclient.NewV1EnrollUserIntoCoursePostRequest("EmailAddress_example", "UplimitCourseId_example", "UplimitEnrollUserIntoCourseSessionSelectionPolicy_example") // V1EnrollUserIntoCoursePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnrollmentAPI.V1EnrollUserIntoCoursePost(context.Background()).V1EnrollUserIntoCoursePostRequest(v1EnrollUserIntoCoursePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V1EnrollUserIntoCoursePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollUserIntoCoursePost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V1EnrollUserIntoCoursePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1EnrollUserIntoCoursePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1EnrollUserIntoCoursePostRequest** | [**V1EnrollUserIntoCoursePostRequest**](V1EnrollUserIntoCoursePostRequest.md) |  | 

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
	resp, r, err := apiClient.EnrollmentAPI.V1EnrollUserIntoSessionPost(context.Background()).V1EnrollUserIntoSessionPostRequest(v1EnrollUserIntoSessionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V1EnrollUserIntoSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollUserIntoSessionPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V1EnrollUserIntoSessionPost`: %v\n", resp)
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
	resp, r, err := apiClient.EnrollmentAPI.V1ListEnrollmentsInSessionGet(context.Background()).UplimitSessionId(uplimitSessionId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V1ListEnrollmentsInSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListEnrollmentsInSessionGet`: V1ListEnrollmentsInSessionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V1ListEnrollmentsInSessionGet`: %v\n", resp)
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
	resp, r, err := apiClient.EnrollmentAPI.V1UnenrollUserFromSessionPost(context.Background()).V1UnenrollUserFromSessionPostRequest(v1UnenrollUserFromSessionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnrollmentAPI.V1UnenrollUserFromSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UnenrollUserFromSessionPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EnrollmentAPI.V1UnenrollUserFromSessionPost`: %v\n", resp)
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

