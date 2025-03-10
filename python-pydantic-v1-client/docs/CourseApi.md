# openapi_client.CourseApi

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_list_courses_get**](CourseApi.md#v1_list_courses_get) | **GET** /v1/ListCourses | 
[**v1_list_sessions_in_course_get**](CourseApi.md#v1_list_sessions_in_course_get) | **GET** /v1/ListSessionsInCourse | 


# **v1_list_courses_get**
> V1ListCoursesGet200Response v1_list_courses_get(skip=skip, take=take)



This API allows developers to list all courses in an organization.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_list_courses_get200_response import V1ListCoursesGet200Response
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
    api_instance = openapi_client.CourseApi(api_client)
    skip = 3.4 # float |  (optional)
    take = 3.4 # float |  (optional)

    try:
        api_response = api_instance.v1_list_courses_get(skip=skip, take=take)
        print("The response of CourseApi->v1_list_courses_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CourseApi->v1_list_courses_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **float**|  | [optional] 
 **take** | **float**|  | [optional] 

### Return type

[**V1ListCoursesGet200Response**](V1ListCoursesGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The list of active users is returned successfully. |  -  |

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
    api_instance = openapi_client.CourseApi(api_client)
    uplimit_course_id = 'uplimit_course_id_example' # str | 
    skip = 3.4 # float |  (optional)
    take = 3.4 # float |  (optional)

    try:
        api_response = api_instance.v1_list_sessions_in_course_get(uplimit_course_id, skip=skip, take=take)
        print("The response of CourseApi->v1_list_sessions_in_course_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CourseApi->v1_list_sessions_in_course_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uplimit_course_id** | **str**|  | 
 **skip** | **float**|  | [optional] 
 **take** | **float**|  | [optional] 

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

