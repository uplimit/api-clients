# HealthcheckSuccessResponseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**success** | **bool** |  | 
**uplimit_organization_id** | **str** | The Uplimit ID of the organization. | 
**uplimit_organization_name** | **str** | The name of the organization as it appears in Uplimit. | 

## Example

```python
from openapi_client.models.healthcheck_success_response_schema import HealthcheckSuccessResponseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of HealthcheckSuccessResponseSchema from a JSON string
healthcheck_success_response_schema_instance = HealthcheckSuccessResponseSchema.from_json(json)
# print the JSON string representation of the object
print HealthcheckSuccessResponseSchema.to_json()

# convert the object into a dict
healthcheck_success_response_schema_dict = healthcheck_success_response_schema_instance.to_dict()
# create an instance of HealthcheckSuccessResponseSchema from a dict
healthcheck_success_response_schema_from_dict = HealthcheckSuccessResponseSchema.from_dict(healthcheck_success_response_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


