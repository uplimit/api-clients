# V1HealthcheckGet403Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** |  | 
**error** | **str** | The error message. | 

## Example

```python
from openapi_client.models.v1_healthcheck_get403_response import V1HealthcheckGet403Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1HealthcheckGet403Response from a JSON string
v1_healthcheck_get403_response_instance = V1HealthcheckGet403Response.from_json(json)
# print the JSON string representation of the object
print V1HealthcheckGet403Response.to_json()

# convert the object into a dict
v1_healthcheck_get403_response_dict = v1_healthcheck_get403_response_instance.to_dict()
# create an instance of V1HealthcheckGet403Response from a dict
v1_healthcheck_get403_response_from_dict = V1HealthcheckGet403Response.from_dict(v1_healthcheck_get403_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


