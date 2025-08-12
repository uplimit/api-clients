# UnenrollUserFromSessionSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**uplimit_session_id** | **str** | Internal ID to identify the session across the Uplimit platform. | 

## Example

```python
from openapi_client.models.unenroll_user_from_session_schema import UnenrollUserFromSessionSchema

# TODO update the JSON string below
json = "{}"
# create an instance of UnenrollUserFromSessionSchema from a JSON string
unenroll_user_from_session_schema_instance = UnenrollUserFromSessionSchema.from_json(json)
# print the JSON string representation of the object
print UnenrollUserFromSessionSchema.to_json()

# convert the object into a dict
unenroll_user_from_session_schema_dict = unenroll_user_from_session_schema_instance.to_dict()
# create an instance of UnenrollUserFromSessionSchema from a dict
unenroll_user_from_session_schema_from_dict = UnenrollUserFromSessionSchema.from_dict(unenroll_user_from_session_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


