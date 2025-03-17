# openapi_client.PlatformApi

All URIs are relative to *https://uplimit.com/api/organization*

Method | HTTP request | Description
------------- | ------------- | -------------
[**v1_healthcheck_get**](PlatformApi.md#v1_healthcheck_get) | **GET** /v1/Healthcheck | 


# **v1_healthcheck_get**
> V1HealthcheckGet200Response v1_healthcheck_get()



This API allows developers to check the health of the platform and verify that the API key provided is valid.

### Example

* Bearer Authentication (bearerAuth):
```python
import time
import os
import openapi_client
from openapi_client.models.v1_healthcheck_get200_response import V1HealthcheckGet200Response
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
    api_instance = openapi_client.PlatformApi(api_client)

    try:
        api_response = api_instance.v1_healthcheck_get()
        print("The response of PlatformApi->v1_healthcheck_get:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PlatformApi->v1_healthcheck_get: %s\n" % e)
```



### Parameters
This endpoint does not need any parameter.

### Return type

[**V1HealthcheckGet200Response**](V1HealthcheckGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The platform is up and running, and the API key provided is correct. |  -  |
**403** | The platform is up and running, but the API key provided is incorrect or missing. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

