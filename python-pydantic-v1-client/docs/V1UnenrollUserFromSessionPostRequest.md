# V1UnenrollUserFromSessionPostRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**uplimit_session_id** | **str** | Internal ID to identify the session across the Uplimit platform. | 

## Example

```python
from openapi_client.models.v1_unenroll_user_from_session_post_request import V1UnenrollUserFromSessionPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1UnenrollUserFromSessionPostRequest from a JSON string
v1_unenroll_user_from_session_post_request_instance = V1UnenrollUserFromSessionPostRequest.from_json(json)
# print the JSON string representation of the object
print V1UnenrollUserFromSessionPostRequest.to_json()

# convert the object into a dict
v1_unenroll_user_from_session_post_request_dict = v1_unenroll_user_from_session_post_request_instance.to_dict()
# create an instance of V1UnenrollUserFromSessionPostRequest from a dict
v1_unenroll_user_from_session_post_request_from_dict = V1UnenrollUserFromSessionPostRequest.from_dict(v1_unenroll_user_from_session_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


