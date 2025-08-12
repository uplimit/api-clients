# openapi_client.SessionApi

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_enroll_user_into_session_post**](SessionApi.md#v1_enroll_user_into_session_post) | **POST** /v1/EnrollUserIntoSession | 
[**v1_list_enrollments_in_session_get**](SessionApi.md#v1_list_enrollments_in_session_get) | **GET** /v1/ListEnrollmentsInSession | 
[**v1_list_sessions_in_course_get**](SessionApi.md#v1_list_sessions_in_course_get) | **GET** /v1/ListSessionsInCourse | 
[**v1_unenroll_user_from_session_post**](SessionApi.md#v1_unenroll_user_from_session_post) | **POST** /v1/UnenrollUserFromSession | 


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
    api_instance = openapi_client.SessionApi(api_client)
    v1_enroll_user_into_session_post_request = openapi_client.V1EnrollUserIntoSessionPostRequest() # V1EnrollUserIntoSessionPostRequest |  (optional)

    try:
        api_response = api_instance.v1_enroll_user_into_session_post(v1_enroll_user_into_session_post_request=v1_enroll_user_into_session_post_request)
        print("The response of SessionApi->v1_enroll_user_into_session_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SessionApi->v1_enroll_user_into_session_post: %s\n" % e)
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
    api_instance = openapi_client.SessionApi(api_client)
    uplimit_session_id = 'uplimit_session_id_example' # str | 
    skip = 56 # int |  (optional)
    take = 56 # int |  (optional)

    try:
        api_response = api_instance.v1_list_enrollments_in_session_get(uplimit_session_id, skip=skip, take=take)
        print("The response of SessionApi->v1_list_enrollments_in_session_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SessionApi->v1_list_enrollments_in_session_get: %s\n" % e)
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

# **v1_list_sessions_in_course_get**
> V1ListSessionsInCourseGet200Response v1_list_sessions_in_course_get(uplimit_course_id, skip=skip, take=take)



This API allows developers to list all sessions of a course.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_list_sessions_in_course_get200_response import V1ListSessionsInCourseGet200Response
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
    api_instance = openapi_client.SessionApi(api_client)
    uplimit_course_id = 'uplimit_course_id_example' # str | 
    skip = 56 # int |  (optional)
    take = 56 # int |  (optional)

    try:
        api_response = api_instance.v1_list_sessions_in_course_get(uplimit_course_id, skip=skip, take=take)
        print("The response of SessionApi->v1_list_sessions_in_course_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SessionApi->v1_list_sessions_in_course_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uplimit_course_id** | **str**|  | 
 **skip** | **int**|  | [optional] 
 **take** | **int**|  | [optional] 

### Return type

[**V1ListSessionsInCourseGet200Response**](V1ListSessionsInCourseGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The list of sessions in the course is returned successfully. |  -  |
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
    api_instance = openapi_client.SessionApi(api_client)
    v1_unenroll_user_from_session_post_request = openapi_client.V1UnenrollUserFromSessionPostRequest() # V1UnenrollUserFromSessionPostRequest |  (optional)

    try:
        api_response = api_instance.v1_unenroll_user_from_session_post(v1_unenroll_user_from_session_post_request=v1_unenroll_user_from_session_post_request)
        print("The response of SessionApi->v1_unenroll_user_from_session_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SessionApi->v1_unenroll_user_from_session_post: %s\n" % e)
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

