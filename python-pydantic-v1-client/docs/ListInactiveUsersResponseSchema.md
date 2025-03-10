# ListInactiveUsersResponseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**users** | [**List[V1GetUserInformationEmailAddressOrUserIdGet200Response]**](V1GetUserInformationEmailAddressOrUserIdGet200Response.md) |  | 
**total_count** | **float** |  | 

## Example

```python
from openapi_client.models.list_inactive_users_response_schema import ListInactiveUsersResponseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of ListInactiveUsersResponseSchema from a JSON string
list_inactive_users_response_schema_instance = ListInactiveUsersResponseSchema.from_json(json)
# print the JSON string representation of the object
print ListInactiveUsersResponseSchema.to_json()

# convert the object into a dict
list_inactive_users_response_schema_dict = list_inactive_users_response_schema_instance.to_dict()
# create an instance of ListInactiveUsersResponseSchema from a dict
list_inactive_users_response_schema_from_dict = ListInactiveUsersResponseSchema.from_dict(list_inactive_users_response_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


