# V1ListEnrollmentsInSessionGet200ResponseUsersInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**first_name** | **str** | The first name of the user. | 
**last_name** | **str** | The last name of the user. | 
**user_account_is_active** | **bool** | Whether the user is allowed to access the Uplimit platform. | 
**user_has_valid_subscription_enrollment** | **bool** | Whether the user is activated in your organization. | 
**uplimit_subscription_enrollment_id** | **str** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**uplimit_subscription_commitment_id** | **str** | Internal ID to identify the \&quot;group\&quot; the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | 
**uplimit_active_subscription_commitment_ids** | **List[str]** | All the active subscription commitment ids for this user within this organization. | 
**uplimit_user_id** | **str** | Internal ID to identify the user across the Uplimit platform. | 
**session_completion_status** | **str** | Whether the user has completed the session according to pre-defined completion criteria. | 
**uplimit_session_enrollment_id** | **str** | The uplimit internal ID for the user enrollment in session. | 

## Example

```python
from openapi_client.models.v1_list_enrollments_in_session_get200_response_users_inner import V1ListEnrollmentsInSessionGet200ResponseUsersInner

# TODO update the JSON string below
json = "{}"
# create an instance of V1ListEnrollmentsInSessionGet200ResponseUsersInner from a JSON string
v1_list_enrollments_in_session_get200_response_users_inner_instance = V1ListEnrollmentsInSessionGet200ResponseUsersInner.from_json(json)
# print the JSON string representation of the object
print V1ListEnrollmentsInSessionGet200ResponseUsersInner.to_json()

# convert the object into a dict
v1_list_enrollments_in_session_get200_response_users_inner_dict = v1_list_enrollments_in_session_get200_response_users_inner_instance.to_dict()
# create an instance of V1ListEnrollmentsInSessionGet200ResponseUsersInner from a dict
v1_list_enrollments_in_session_get200_response_users_inner_from_dict = V1ListEnrollmentsInSessionGet200ResponseUsersInner.from_dict(v1_list_enrollments_in_session_get200_response_users_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


