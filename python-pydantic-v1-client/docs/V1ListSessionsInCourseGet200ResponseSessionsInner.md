# V1ListSessionsInCourseGet200ResponseSessionsInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uplimit_session_id** | **str** | Internal ID to identify the session across the Uplimit platform. | 
**name** | **str** | The name of the session. | 
**starts_at** | **datetime** | The start date of the session. | 
**ends_at** | **datetime** | The end date of the session. | 
**enrollment_enabled** | **bool** | Whether the session allows enrollments | 

## Example

```python
from openapi_client.models.v1_list_sessions_in_course_get200_response_sessions_inner import V1ListSessionsInCourseGet200ResponseSessionsInner

# TODO update the JSON string below
json = "{}"
# create an instance of V1ListSessionsInCourseGet200ResponseSessionsInner from a JSON string
v1_list_sessions_in_course_get200_response_sessions_inner_instance = V1ListSessionsInCourseGet200ResponseSessionsInner.from_json(json)
# print the JSON string representation of the object
print V1ListSessionsInCourseGet200ResponseSessionsInner.to_json()

# convert the object into a dict
v1_list_sessions_in_course_get200_response_sessions_inner_dict = v1_list_sessions_in_course_get200_response_sessions_inner_instance.to_dict()
# create an instance of V1ListSessionsInCourseGet200ResponseSessionsInner from a dict
v1_list_sessions_in_course_get200_response_sessions_inner_from_dict = V1ListSessionsInCourseGet200ResponseSessionsInner.from_dict(v1_list_sessions_in_course_get200_response_sessions_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


