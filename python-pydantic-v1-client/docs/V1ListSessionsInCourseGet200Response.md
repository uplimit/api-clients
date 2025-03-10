# V1ListSessionsInCourseGet200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sessions** | [**List[V1ListSessionsInCourseGet200ResponseSessionsInner]**](V1ListSessionsInCourseGet200ResponseSessionsInner.md) |  | 
**total_count** | **float** |  | 

## Example

```python
from openapi_client.models.v1_list_sessions_in_course_get200_response import V1ListSessionsInCourseGet200Response

# TODO update the JSON string below
json = "{}"
# create an instance of V1ListSessionsInCourseGet200Response from a JSON string
v1_list_sessions_in_course_get200_response_instance = V1ListSessionsInCourseGet200Response.from_json(json)
# print the JSON string representation of the object
print V1ListSessionsInCourseGet200Response.to_json()

# convert the object into a dict
v1_list_sessions_in_course_get200_response_dict = v1_list_sessions_in_course_get200_response_instance.to_dict()
# create an instance of V1ListSessionsInCourseGet200Response from a dict
v1_list_sessions_in_course_get200_response_from_dict = V1ListSessionsInCourseGet200Response.from_dict(v1_list_sessions_in_course_get200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


