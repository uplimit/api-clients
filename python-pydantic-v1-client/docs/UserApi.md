# openapi_client.UserApi

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_add_user_authentication_method_post**](UserApi.md#v1_add_user_authentication_method_post) | **POST** /v1/AddUserAuthenticationMethod | 
[**v1_create_user_post**](UserApi.md#v1_create_user_post) | **POST** /v1/CreateUser | 
[**v1_enroll_user_into_course_post**](UserApi.md#v1_enroll_user_into_course_post) | **POST** /v1/EnrollUserIntoCourse | 
[**v1_enroll_user_into_session_post**](UserApi.md#v1_enroll_user_into_session_post) | **POST** /v1/EnrollUserIntoSession | 
[**v1_get_user_information_email_address_or_user_id_get**](UserApi.md#v1_get_user_information_email_address_or_user_id_get) | **GET** /v1/GetUserInformation/{emailAddressOrUserId} | 
[**v1_list_active_users_get**](UserApi.md#v1_list_active_users_get) | **GET** /v1/ListActiveUsers | 
[**v1_list_enrollments_in_session_get**](UserApi.md#v1_list_enrollments_in_session_get) | **GET** /v1/ListEnrollmentsInSession | 
[**v1_list_inactive_users_get**](UserApi.md#v1_list_inactive_users_get) | **GET** /v1/ListInactiveUsers | 
[**v1_toggle_user_activation_post**](UserApi.md#v1_toggle_user_activation_post) | **POST** /v1/ToggleUserActivation | 
[**v1_unenroll_user_from_session_post**](UserApi.md#v1_unenroll_user_from_session_post) | **POST** /v1/UnenrollUserFromSession | 


# **v1_add_user_authentication_method_post**
> V1EnrollUserIntoSessionPost200Response v1_add_user_authentication_method_post(v1_add_user_authentication_method_post_request=v1_add_user_authentication_method_post_request)



This API creates a new user authentication method for a user within your organization.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_add_user_authentication_method_post_request import V1AddUserAuthenticationMethodPostRequest
from openapi_client.models.v1_enroll_user_into_session_post200_response import V1EnrollUserIntoSessionPost200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    v1_add_user_authentication_method_post_request = openapi_client.V1AddUserAuthenticationMethodPostRequest() # V1AddUserAuthenticationMethodPostRequest |  (optional)

    try:
        api_response = api_instance.v1_add_user_authentication_method_post(v1_add_user_authentication_method_post_request=v1_add_user_authentication_method_post_request)
        print("The response of UserApi->v1_add_user_authentication_method_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_add_user_authentication_method_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_add_user_authentication_method_post_request** | [**V1AddUserAuthenticationMethodPostRequest**](V1AddUserAuthenticationMethodPostRequest.md)|  | [optional] 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The user authentication method was added successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_create_user_post**
> V1CreateUserPost200Response v1_create_user_post(v1_create_user_post_request=v1_create_user_post_request)



This API creates a new user account on Uplimit and enrolls the user into your organization. If a user account with the same email already exists, we will just enroll that existing user into your organization. 

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_create_user_post200_response import V1CreateUserPost200Response
from openapi_client.models.v1_create_user_post_request import V1CreateUserPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    v1_create_user_post_request = openapi_client.V1CreateUserPostRequest() # V1CreateUserPostRequest |  (optional)

    try:
        api_response = api_instance.v1_create_user_post(v1_create_user_post_request=v1_create_user_post_request)
        print("The response of UserApi->v1_create_user_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_create_user_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_create_user_post_request** | [**V1CreateUserPostRequest**](V1CreateUserPostRequest.md)|  | [optional] 

### Return type

[**V1CreateUserPost200Response**](V1CreateUserPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The user was created successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_enroll_user_into_course_post**
> V1EnrollUserIntoSessionPost200Response v1_enroll_user_into_course_post(v1_enroll_user_into_course_post_request=v1_enroll_user_into_course_post_request)



This API allows developers to add a user into a course. The user must have already been created with the Create User API (see above) before you can add this user. In general, prefer to use the EnrollUserIntoSession API instead of this API, because that is more specific and you can specify the exact session you want to enroll the user into, instead of relying on a policy.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_enroll_user_into_course_post_request import V1EnrollUserIntoCoursePostRequest
from openapi_client.models.v1_enroll_user_into_session_post200_response import V1EnrollUserIntoSessionPost200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    v1_enroll_user_into_course_post_request = openapi_client.V1EnrollUserIntoCoursePostRequest() # V1EnrollUserIntoCoursePostRequest |  (optional)

    try:
        api_response = api_instance.v1_enroll_user_into_course_post(v1_enroll_user_into_course_post_request=v1_enroll_user_into_course_post_request)
        print("The response of UserApi->v1_enroll_user_into_course_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_enroll_user_into_course_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_enroll_user_into_course_post_request** | [**V1EnrollUserIntoCoursePostRequest**](V1EnrollUserIntoCoursePostRequest.md)|  | [optional] 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The user is enrolled into a session of the course successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_enroll_user_into_session_post**
> V1EnrollUserIntoSessionPost200Response v1_enroll_user_into_session_post(v1_enroll_user_into_session_post_request=v1_enroll_user_into_session_post_request)



This API allows developers to add a user into a session. The user must have already been created with the Create User API (see above) before you can add this user.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_enroll_user_into_session_post200_response import V1EnrollUserIntoSessionPost200Response
from openapi_client.models.v1_enroll_user_into_session_post_request import V1EnrollUserIntoSessionPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    v1_enroll_user_into_session_post_request = openapi_client.V1EnrollUserIntoSessionPostRequest() # V1EnrollUserIntoSessionPostRequest |  (optional)

    try:
        api_response = api_instance.v1_enroll_user_into_session_post(v1_enroll_user_into_session_post_request=v1_enroll_user_into_session_post_request)
        print("The response of UserApi->v1_enroll_user_into_session_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_enroll_user_into_session_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_enroll_user_into_session_post_request** | [**V1EnrollUserIntoSessionPostRequest**](V1EnrollUserIntoSessionPostRequest.md)|  | [optional] 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The user is enrolled into the session successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_get_user_information_email_address_or_user_id_get**
> V1GetUserInformationEmailAddressOrUserIdGet200Response v1_get_user_information_email_address_or_user_id_get(email_address_or_user_id)



This API allows developers to add a user into a session. The user must have already been created with the Create User API (see above) before you can add this user.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_get_user_information_email_address_or_user_id_get200_response import V1GetUserInformationEmailAddressOrUserIdGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    email_address_or_user_id = 'email_address_or_user_id_example' # str | The email address or uplimit User ID of the user.

    try:
        api_response = api_instance.v1_get_user_information_email_address_or_user_id_get(email_address_or_user_id)
        print("The response of UserApi->v1_get_user_information_email_address_or_user_id_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_get_user_information_email_address_or_user_id_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email_address_or_user_id** | **str**| The email address or uplimit User ID of the user. | 

### Return type

[**V1GetUserInformationEmailAddressOrUserIdGet200Response**](V1GetUserInformationEmailAddressOrUserIdGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Information about the user is returned successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_list_active_users_get**
> V1ListActiveUsersGet200Response v1_list_active_users_get(skip=skip, take=take)



This API lists all active users in your organization on the Uplimit platform.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_list_active_users_get200_response import V1ListActiveUsersGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    skip = 56 # int |  (optional)
    take = 56 # int |  (optional)

    try:
        api_response = api_instance.v1_list_active_users_get(skip=skip, take=take)
        print("The response of UserApi->v1_list_active_users_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_list_active_users_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int**|  | [optional] 
 **take** | **int**|  | [optional] 

### Return type

[**V1ListActiveUsersGet200Response**](V1ListActiveUsersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The list of active users is returned successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_list_enrollments_in_session_get**
> V1ListEnrollmentsInSessionGet200Response v1_list_enrollments_in_session_get(uplimit_session_id, skip=skip, take=take)



This API allows developers to list all active enrollments in a session.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_list_enrollments_in_session_get200_response import V1ListEnrollmentsInSessionGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    uplimit_session_id = 'uplimit_session_id_example' # str | 
    skip = 56 # int |  (optional)
    take = 56 # int |  (optional)

    try:
        api_response = api_instance.v1_list_enrollments_in_session_get(uplimit_session_id, skip=skip, take=take)
        print("The response of UserApi->v1_list_enrollments_in_session_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_list_enrollments_in_session_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uplimit_session_id** | **str**|  | 
 **skip** | **int**|  | [optional] 
 **take** | **int**|  | [optional] 

### Return type

[**V1ListEnrollmentsInSessionGet200Response**](V1ListEnrollmentsInSessionGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The list of enrollments in the session is returned successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_list_inactive_users_get**
> V1ListActiveUsersGet200Response v1_list_inactive_users_get(skip=skip, take=take)



This API lists all inactive users in your organization on the Uplimit platform.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_list_active_users_get200_response import V1ListActiveUsersGet200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    skip = 56 # int |  (optional)
    take = 56 # int |  (optional)

    try:
        api_response = api_instance.v1_list_inactive_users_get(skip=skip, take=take)
        print("The response of UserApi->v1_list_inactive_users_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_list_inactive_users_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int**|  | [optional] 
 **take** | **int**|  | [optional] 

### Return type

[**V1ListActiveUsersGet200Response**](V1ListActiveUsersGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The list of inactive users is returned successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_toggle_user_activation_post**
> V1EnrollUserIntoSessionPost200Response v1_toggle_user_activation_post(v1_toggle_user_activation_post_request=v1_toggle_user_activation_post_request)



This API changes whether a user is active in your organization (i.e. sets their state to activated or deactivated).

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_enroll_user_into_session_post200_response import V1EnrollUserIntoSessionPost200Response
from openapi_client.models.v1_toggle_user_activation_post_request import V1ToggleUserActivationPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    v1_toggle_user_activation_post_request = openapi_client.V1ToggleUserActivationPostRequest() # V1ToggleUserActivationPostRequest |  (optional)

    try:
        api_response = api_instance.v1_toggle_user_activation_post(v1_toggle_user_activation_post_request=v1_toggle_user_activation_post_request)
        print("The response of UserApi->v1_toggle_user_activation_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_toggle_user_activation_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_toggle_user_activation_post_request** | [**V1ToggleUserActivationPostRequest**](V1ToggleUserActivationPostRequest.md)|  | [optional] 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The user activation was toggled successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_unenroll_user_from_session_post**
> V1EnrollUserIntoSessionPost200Response v1_unenroll_user_from_session_post(v1_unenroll_user_from_session_post_request=v1_unenroll_user_from_session_post_request)



This API allows developers to remove a user from a session. The user must currently be enrolled in the session to be eligible for unenrollment.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_enroll_user_into_session_post200_response import V1EnrollUserIntoSessionPost200Response
from openapi_client.models.v1_unenroll_user_from_session_post_request import V1UnenrollUserFromSessionPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://uplimit.com/api/organization
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://uplimit.com/api/organization"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization: bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.UserApi(api_client)
    v1_unenroll_user_from_session_post_request = openapi_client.V1UnenrollUserFromSessionPostRequest() # V1UnenrollUserFromSessionPostRequest |  (optional)

    try:
        api_response = api_instance.v1_unenroll_user_from_session_post(v1_unenroll_user_from_session_post_request=v1_unenroll_user_from_session_post_request)
        print("The response of UserApi->v1_unenroll_user_from_session_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UserApi->v1_unenroll_user_from_session_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_unenroll_user_from_session_post_request** | [**V1UnenrollUserFromSessionPostRequest**](V1UnenrollUserFromSessionPostRequest.md)|  | [optional] 

### Return type

[**V1EnrollUserIntoSessionPost200Response**](V1EnrollUserIntoSessionPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The user is unenrolled from the session successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

