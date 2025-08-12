# EnrollUserIntoSessionSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**session_id** | **str** | The ID of the session to enroll the user into. You must provide either this or uplimitSessionId. | [optional] 
**uplimit_session_id** | **str** | Internal ID to identify the session across the Uplimit platform. | [optional] 
**subscription_commitment_id** | **str** | Internal ID to identify the \&quot;group\&quot; the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Example

```python
from openapi_client.models.enroll_user_into_session_schema import EnrollUserIntoSessionSchema

# TODO update the JSON string below
json = "{}"
# create an instance of EnrollUserIntoSessionSchema from a JSON string
enroll_user_into_session_schema_instance = EnrollUserIntoSessionSchema.from_json(json)
# print the JSON string representation of the object
print EnrollUserIntoSessionSchema.to_json()

# convert the object into a dict
enroll_user_into_session_schema_dict = enroll_user_into_session_schema_instance.to_dict()
# create an instance of EnrollUserIntoSessionSchema from a dict
enroll_user_into_session_schema_from_dict = EnrollUserIntoSessionSchema.from_dict(enroll_user_into_session_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


