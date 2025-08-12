# \UserAPI

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1AddUserAuthenticationMethodPost**](UserAPI.md#V1AddUserAuthenticationMethodPost) | **Post** /v1/AddUserAuthenticationMethod | 
[**V1CreateUserPost**](UserAPI.md#V1CreateUserPost) | **Post** /v1/CreateUser | 
[**V1EnrollUserIntoCoursePost**](UserAPI.md#V1EnrollUserIntoCoursePost) | **Post** /v1/EnrollUserIntoCourse | 
[**V1EnrollUserIntoSessionPost**](UserAPI.md#V1EnrollUserIntoSessionPost) | **Post** /v1/EnrollUserIntoSession | 
[**V1GetUserInformationEmailAddressOrUserIdGet**](UserAPI.md#V1GetUserInformationEmailAddressOrUserIdGet) | **Get** /v1/GetUserInformation/{emailAddressOrUserId} | 
[**V1ListActiveUsersGet**](UserAPI.md#V1ListActiveUsersGet) | **Get** /v1/ListActiveUsers | 
[**V1ListEnrollmentsInSessionGet**](UserAPI.md#V1ListEnrollmentsInSessionGet) | **Get** /v1/ListEnrollmentsInSession | 
[**V1ListInactiveUsersGet**](UserAPI.md#V1ListInactiveUsersGet) | **Get** /v1/ListInactiveUsers | 
[**V1ToggleUserActivationPost**](UserAPI.md#V1ToggleUserActivationPost) | **Post** /v1/ToggleUserActivation | 
[**V1UnenrollUserFromSessionPost**](UserAPI.md#V1UnenrollUserFromSessionPost) | **Post** /v1/UnenrollUserFromSession | 



## V1AddUserAuthenticationMethodPost

> V1EnrollUserIntoSessionPost200Response V1AddUserAuthenticationMethodPost(ctx).V1AddUserAuthenticationMethodPostRequest(v1AddUserAuthenticationMethodPostRequest).Execute()





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
	v1AddUserAuthenticationMethodPostRequest := *openapiclient.NewV1AddUserAuthenticationMethodPostRequest("EmailAddress_example", "AuthenticationMethod_example", "CustomAuthenticationMethodProviderId_example", "AuthenticationSecret_example") // V1AddUserAuthenticationMethodPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1AddUserAuthenticationMethodPost(context.Background()).V1AddUserAuthenticationMethodPostRequest(v1AddUserAuthenticationMethodPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1AddUserAuthenticationMethodPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1AddUserAuthenticationMethodPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1AddUserAuthenticationMethodPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1AddUserAuthenticationMethodPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AddUserAuthenticationMethodPostRequest** | [**V1AddUserAuthenticationMethodPostRequest**](V1AddUserAuthenticationMethodPostRequest.md) |  | 

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


## V1CreateUserPost

> V1CreateUserPost200Response V1CreateUserPost(ctx).V1CreateUserPostRequest(v1CreateUserPostRequest).Execute()





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
	v1CreateUserPostRequest := *openapiclient.NewV1CreateUserPostRequest("EmailAddress_example", "FirstName_example", "LastName_example") // V1CreateUserPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1CreateUserPost(context.Background()).V1CreateUserPostRequest(v1CreateUserPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1CreateUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1CreateUserPost`: V1CreateUserPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1CreateUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1CreateUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1CreateUserPostRequest** | [**V1CreateUserPostRequest**](V1CreateUserPostRequest.md) |  | 

### Return type

[**V1CreateUserPost200Response**](V1CreateUserPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.UserAPI.V1EnrollUserIntoCoursePost(context.Background()).V1EnrollUserIntoCoursePostRequest(v1EnrollUserIntoCoursePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1EnrollUserIntoCoursePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollUserIntoCoursePost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1EnrollUserIntoCoursePost`: %v\n", resp)
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
	resp, r, err := apiClient.UserAPI.V1EnrollUserIntoSessionPost(context.Background()).V1EnrollUserIntoSessionPostRequest(v1EnrollUserIntoSessionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1EnrollUserIntoSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1EnrollUserIntoSessionPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1EnrollUserIntoSessionPost`: %v\n", resp)
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


## V1GetUserInformationEmailAddressOrUserIdGet

> V1GetUserInformationEmailAddressOrUserIdGet200Response V1GetUserInformationEmailAddressOrUserIdGet(ctx, emailAddressOrUserId).Execute()





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
	emailAddressOrUserId := "emailAddressOrUserId_example" // string | The email address or uplimit User ID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1GetUserInformationEmailAddressOrUserIdGet(context.Background(), emailAddressOrUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1GetUserInformationEmailAddressOrUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1GetUserInformationEmailAddressOrUserIdGet`: V1GetUserInformationEmailAddressOrUserIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1GetUserInformationEmailAddressOrUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddressOrUserId** | **string** | The email address or uplimit User ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1GetUserInformationEmailAddressOrUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1GetUserInformationEmailAddressOrUserIdGet200Response**](V1GetUserInformationEmailAddressOrUserIdGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ListActiveUsersGet

> V1ListActiveUsersGet200Response V1ListActiveUsersGet(ctx).Skip(skip).Take(take).Execute()





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
	resp, r, err := apiClient.UserAPI.V1ListActiveUsersGet(context.Background()).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1ListActiveUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListActiveUsersGet`: V1ListActiveUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1ListActiveUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ListActiveUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | 
 **take** | **int32** |  | 

### Return type

[**V1ListActiveUsersGet200Response**](V1ListActiveUsersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
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
	resp, r, err := apiClient.UserAPI.V1ListEnrollmentsInSessionGet(context.Background()).UplimitSessionId(uplimitSessionId).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1ListEnrollmentsInSessionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListEnrollmentsInSessionGet`: V1ListEnrollmentsInSessionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1ListEnrollmentsInSessionGet`: %v\n", resp)
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


## V1ListInactiveUsersGet

> V1ListActiveUsersGet200Response V1ListInactiveUsersGet(ctx).Skip(skip).Take(take).Execute()





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
	resp, r, err := apiClient.UserAPI.V1ListInactiveUsersGet(context.Background()).Skip(skip).Take(take).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1ListInactiveUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ListInactiveUsersGet`: V1ListActiveUsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1ListInactiveUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ListInactiveUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** |  | 
 **take** | **int32** |  | 

### Return type

[**V1ListActiveUsersGet200Response**](V1ListActiveUsersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ToggleUserActivationPost

> V1EnrollUserIntoSessionPost200Response V1ToggleUserActivationPost(ctx).V1ToggleUserActivationPostRequest(v1ToggleUserActivationPostRequest).Execute()





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
	v1ToggleUserActivationPostRequest := *openapiclient.NewV1ToggleUserActivationPostRequest("EmailAddress_example", false) // V1ToggleUserActivationPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.V1ToggleUserActivationPost(context.Background()).V1ToggleUserActivationPostRequest(v1ToggleUserActivationPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1ToggleUserActivationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ToggleUserActivationPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1ToggleUserActivationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ToggleUserActivationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1ToggleUserActivationPostRequest** | [**V1ToggleUserActivationPostRequest**](V1ToggleUserActivationPostRequest.md) |  | 

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
	resp, r, err := apiClient.UserAPI.V1UnenrollUserFromSessionPost(context.Background()).V1UnenrollUserFromSessionPostRequest(v1UnenrollUserFromSessionPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.V1UnenrollUserFromSessionPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1UnenrollUserFromSessionPost`: V1EnrollUserIntoSessionPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.V1UnenrollUserFromSessionPost`: %v\n", resp)
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

