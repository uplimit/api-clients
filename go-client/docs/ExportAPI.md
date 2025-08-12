# \ExportAPI

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1GetLearnerActivityGetGet**](ExportAPI.md#V1GetLearnerActivityGetGet) | **Get** /v1/GetLearnerActivity/get | 
[**V1GetLearnerActivityStartPost**](ExportAPI.md#V1GetLearnerActivityStartPost) | **Post** /v1/GetLearnerActivity/start | 



## V1GetLearnerActivityGetGet

> V1GetLearnerActivityGetGet200Response V1GetLearnerActivityGetGet(ctx).JobId(jobId).Execute()





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
	jobId := "jobId_example" // string | The ID of the job to get the status of.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.V1GetLearnerActivityGetGet(context.Background()).JobId(jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.V1GetLearnerActivityGetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetLearnerActivityGetGet`: V1GetLearnerActivityGetGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.V1GetLearnerActivityGetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetLearnerActivityGetGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobId** | **string** | The ID of the job to get the status of. | 

### Return type

[**V1GetLearnerActivityGetGet200Response**](V1GetLearnerActivityGetGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1GetLearnerActivityStartPost

> V1GetLearnerActivityStartPost200Response V1GetLearnerActivityStartPost(ctx).V1GetLearnerActivityStartPostRequest(v1GetLearnerActivityStartPostRequest).Execute()





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
	v1GetLearnerActivityStartPostRequest := *openapiclient.NewV1GetLearnerActivityStartPostRequest() // V1GetLearnerActivityStartPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ExportAPI.V1GetLearnerActivityStartPost(context.Background()).V1GetLearnerActivityStartPostRequest(v1GetLearnerActivityStartPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.V1GetLearnerActivityStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetLearnerActivityStartPost`: V1GetLearnerActivityStartPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ExportAPI.V1GetLearnerActivityStartPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1GetLearnerActivityStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1GetLearnerActivityStartPostRequest** | [**V1GetLearnerActivityStartPostRequest**](V1GetLearnerActivityStartPostRequest.md) |  | 

### Return type

[**V1GetLearnerActivityStartPost200Response**](V1GetLearnerActivityStartPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

