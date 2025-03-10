# V1CreateUserPost200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uplimit_subscription_enrollment_id** | **str** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**uplimit_user_id** | **str** | Internal ID to identify the user across the Uplimit platform. | 

## Example

```python
from openapi_client.models.v1_create_user_post200_response import V1CreateUserPost200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1CreateUserPost200Response from a JSON string
v1_create_user_post200_response_instance = V1CreateUserPost200Response.from_json(json)
# print the JSON string representation of the object
print V1CreateUserPost200Response.to_json()

# convert the object into a dict
v1_create_user_post200_response_dict = v1_create_user_post200_response_instance.to_dict()
# create an instance of V1CreateUserPost200Response from a dict
v1_create_user_post200_response_from_dict = V1CreateUserPost200Response.from_dict(v1_create_user_post200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


