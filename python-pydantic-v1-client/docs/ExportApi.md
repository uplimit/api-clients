# openapi_client.ExportApi

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_get_learner_activity_get_get**](ExportApi.md#v1_get_learner_activity_get_get) | **GET** /v1/GetLearnerActivity/get | 
[**v1_get_learner_activity_start_post**](ExportApi.md#v1_get_learner_activity_start_post) | **POST** /v1/GetLearnerActivity/start | 


# **v1_get_learner_activity_get_get**
> V1GetLearnerActivityGetGet200Response v1_get_learner_activity_get_get(job_id)



This API gets the status of a job to export learner activity for a given session.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_get_learner_activity_get_get200_response import V1GetLearnerActivityGetGet200Response
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
    api_instance = openapi_client.ExportApi(api_client)
    job_id = 'job_id_example' # str | The ID of the job to get the status of.

    try:
        api_response = api_instance.v1_get_learner_activity_get_get(job_id)
        print("The response of ExportApi->v1_get_learner_activity_get_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExportApi->v1_get_learner_activity_get_get: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_id** | **str**| The ID of the job to get the status of. | 

### Return type

[**V1GetLearnerActivityGetGet200Response**](V1GetLearnerActivityGetGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The job status was retrieved successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **v1_get_learner_activity_start_post**
> V1GetLearnerActivityStartPost200Response v1_get_learner_activity_start_post(v1_get_learner_activity_start_post_request=v1_get_learner_activity_start_post_request)



This API starts a job to export learner activity for a given sessions. If no sessions are provided, all sessions in the course will be exported.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_get_learner_activity_start_post200_response import V1GetLearnerActivityStartPost200Response
from openapi_client.models.v1_get_learner_activity_start_post_request import V1GetLearnerActivityStartPostRequest
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
    api_instance = openapi_client.ExportApi(api_client)
    v1_get_learner_activity_start_post_request = openapi_client.V1GetLearnerActivityStartPostRequest() # V1GetLearnerActivityStartPostRequest |  (optional)

    try:
        api_response = api_instance.v1_get_learner_activity_start_post(v1_get_learner_activity_start_post_request=v1_get_learner_activity_start_post_request)
        print("The response of ExportApi->v1_get_learner_activity_start_post:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ExportApi->v1_get_learner_activity_start_post: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_get_learner_activity_start_post_request** | [**V1GetLearnerActivityStartPostRequest**](V1GetLearnerActivityStartPostRequest.md)|  | [optional] 

### Return type

[**V1GetLearnerActivityStartPost200Response**](V1GetLearnerActivityStartPost200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The job was started successfully. |  -  |
**400** | The request is invalid. |  -  |
**401** | The request is unauthorized. |  -  |
**404** | One or more of the resources required to fulfill the request were not found. |  -  |
**405** | The request method is not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

