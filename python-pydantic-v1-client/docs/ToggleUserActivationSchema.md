# ToggleUserActivationSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**set_is_active** | **bool** | Whether to set the user as active or inactive. | 

## Example

```python
from openapi_client.models.toggle_user_activation_schema import ToggleUserActivationSchema

# TODO update the JSON string below
json = "{}"
# create an instance of ToggleUserActivationSchema from a JSON string
toggle_user_activation_schema_instance = ToggleUserActivationSchema.from_json(json)
# print the JSON string representation of the object
print ToggleUserActivationSchema.to_json()

# convert the object into a dict
toggle_user_activation_schema_dict = toggle_user_activation_schema_instance.to_dict()
# create an instance of ToggleUserActivationSchema from a dict
toggle_user_activation_schema_from_dict = ToggleUserActivationSchema.from_dict(toggle_user_activation_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


