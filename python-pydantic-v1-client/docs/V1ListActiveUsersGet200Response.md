# V1ListActiveUsersGet200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**users** | [**List[V1GetUserInformationEmailAddressOrUserIdGet200Response]**](V1GetUserInformationEmailAddressOrUserIdGet200Response.md) |  | 
**total_count** | **float** |  | 

## Example

```python
from openapi_client.models.v1_list_active_users_get200_response import V1ListActiveUsersGet200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1ListActiveUsersGet200Response from a JSON string
v1_list_active_users_get200_response_instance = V1ListActiveUsersGet200Response.from_json(json)
# print the JSON string representation of the object
print V1ListActiveUsersGet200Response.to_json()

# convert the object into a dict
v1_list_active_users_get200_response_dict = v1_list_active_users_get200_response_instance.to_dict()
# create an instance of V1ListActiveUsersGet200Response from a dict
v1_list_active_users_get200_response_from_dict = V1ListActiveUsersGet200Response.from_dict(v1_list_active_users_get200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


