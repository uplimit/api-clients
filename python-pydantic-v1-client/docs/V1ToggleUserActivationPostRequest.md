# V1ToggleUserActivationPostRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**set_is_active** | **bool** | Whether to set the user as active or inactive. | 

## Example

```python
from openapi_client.models.v1_toggle_user_activation_post_request import V1ToggleUserActivationPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1ToggleUserActivationPostRequest from a JSON string
v1_toggle_user_activation_post_request_instance = V1ToggleUserActivationPostRequest.from_json(json)
# print the JSON string representation of the object
print V1ToggleUserActivationPostRequest.to_json()

# convert the object into a dict
v1_toggle_user_activation_post_request_dict = v1_toggle_user_activation_post_request_instance.to_dict()
# create an instance of V1ToggleUserActivationPostRequest from a dict
v1_toggle_user_activation_post_request_from_dict = V1ToggleUserActivationPostRequest.from_dict(v1_toggle_user_activation_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


