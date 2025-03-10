# V1CreateUserPostRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**first_name** | **str** | The first name of the user. | 
**last_name** | **str** | The last name of the user. | 
**subscription_commitment_id** | **str** | Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Example

```python
from openapi_client.models.v1_create_user_post_request import V1CreateUserPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1CreateUserPostRequest from a JSON string
v1_create_user_post_request_instance = V1CreateUserPostRequest.from_json(json)
# print the JSON string representation of the object
print V1CreateUserPostRequest.to_json()

# convert the object into a dict
v1_create_user_post_request_dict = v1_create_user_post_request_instance.to_dict()
# create an instance of V1CreateUserPostRequest from a dict
v1_create_user_post_request_from_dict = V1CreateUserPostRequest.from_dict(v1_create_user_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


