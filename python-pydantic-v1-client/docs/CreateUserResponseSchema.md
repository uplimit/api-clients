# CreateUserResponseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uplimit_subscription_enrollment_id** | **str** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**uplimit_user_id** | **str** | Internal ID to identify the user across the Uplimit platform. | 

## Example

```python
from openapi_client.models.create_user_response_schema import CreateUserResponseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of CreateUserResponseSchema from a JSON string
create_user_response_schema_instance = CreateUserResponseSchema.from_json(json)
# print the JSON string representation of the object
print CreateUserResponseSchema.to_json()

# convert the object into a dict
create_user_response_schema_dict = create_user_response_schema_instance.to_dict()
# create an instance of CreateUserResponseSchema from a dict
create_user_response_schema_from_dict = CreateUserResponseSchema.from_dict(create_user_response_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


