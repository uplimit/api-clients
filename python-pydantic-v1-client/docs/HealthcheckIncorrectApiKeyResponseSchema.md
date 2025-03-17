# HealthcheckIncorrectApiKeyResponseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** |  | 
**error** | **str** | The error message. | 

## Example

```python
from openapi_client.models.healthcheck_incorrect_api_key_response_schema import HealthcheckIncorrectApiKeyResponseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of HealthcheckIncorrectApiKeyResponseSchema from a JSON string
healthcheck_incorrect_api_key_response_schema_instance = HealthcheckIncorrectApiKeyResponseSchema.from_json(json)
# print the JSON string representation of the object
print HealthcheckIncorrectApiKeyResponseSchema.to_json()

# convert the object into a dict
healthcheck_incorrect_api_key_response_schema_dict = healthcheck_incorrect_api_key_response_schema_instance.to_dict()
# create an instance of HealthcheckIncorrectApiKeyResponseSchema from a dict
healthcheck_incorrect_api_key_response_schema_from_dict = HealthcheckIncorrectApiKeyResponseSchema.from_dict(healthcheck_incorrect_api_key_response_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


