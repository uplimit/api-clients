# CreateUserSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**first_name** | **str** | The first name of the user. | 
**last_name** | **str** | The last name of the user. | 
**subscription_commitment_id** | **str** | Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Example

```python
from openapi_client.models.create_user_schema import CreateUserSchema

# TODO update the JSON string below
json = "{}"
# create an instance of CreateUserSchema from a JSON string
create_user_schema_instance = CreateUserSchema.from_json(json)
# print the JSON string representation of the object
print CreateUserSchema.to_json()

# convert the object into a dict
create_user_schema_dict = create_user_schema_instance.to_dict()
# create an instance of CreateUserSchema from a dict
create_user_schema_from_dict = CreateUserSchema.from_dict(create_user_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


