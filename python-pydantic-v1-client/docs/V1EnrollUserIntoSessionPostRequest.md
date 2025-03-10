# V1EnrollUserIntoSessionPostRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**session_id** | **str** | The ID of the session to enroll the user into. You must provide either this or uplimitSessionId. | [optional] 
**uplimit_session_id** | **str** | Internal ID to identify the session across the Uplimit platform. | [optional] 
**subscription_commitment_id** | **str** | Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Example

```python
from openapi_client.models.v1_enroll_user_into_session_post_request import V1EnrollUserIntoSessionPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1EnrollUserIntoSessionPostRequest from a JSON string
v1_enroll_user_into_session_post_request_instance = V1EnrollUserIntoSessionPostRequest.from_json(json)
# print the JSON string representation of the object
print V1EnrollUserIntoSessionPostRequest.to_json()

# convert the object into a dict
v1_enroll_user_into_session_post_request_dict = v1_enroll_user_into_session_post_request_instance.to_dict()
# create an instance of V1EnrollUserIntoSessionPostRequest from a dict
v1_enroll_user_into_session_post_request_from_dict = V1EnrollUserIntoSessionPostRequest.from_dict(v1_enroll_user_into_session_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


