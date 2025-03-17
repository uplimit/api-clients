# V1HealthcheckGet200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** |  | 
**uplimit_organization_id** | **str** | The Uplimit ID of the organization. | 
**uplimit_organization_name** | **str** | The name of the organization as it appears in Uplimit. | 

## Example

```python
from openapi_client.models.v1_healthcheck_get200_response import V1HealthcheckGet200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthcheckGet200Response from a JSON string
v1_healthcheck_get200_response_instance = V1HealthcheckGet200Response.from_json(json)
# print the JSON string representation of the object
print V1HealthcheckGet200Response.to_json()

# convert the object into a dict
v1_healthcheck_get200_response_dict = v1_healthcheck_get200_response_instance.to_dict()
# create an instance of V1HealthcheckGet200Response from a dict
v1_healthcheck_get200_response_from_dict = V1HealthcheckGet200Response.from_dict(v1_healthcheck_get200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


