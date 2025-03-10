# ListSessionsInCourseResponseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**sessions** | [**List[V1ListSessionsInCourseGet200ResponseSessionsInner]**](V1ListSessionsInCourseGet200ResponseSessionsInner.md) |  | 
**total_count** | **float** |  | 

## Example

```python
from openapi_client.models.list_sessions_in_course_response_schema import ListSessionsInCourseResponseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of ListSessionsInCourseResponseSchema from a JSON string
list_sessions_in_course_response_schema_instance = ListSessionsInCourseResponseSchema.from_json(json)
# print the JSON string representation of the object
print ListSessionsInCourseResponseSchema.to_json()

# convert the object into a dict
list_sessions_in_course_response_schema_dict = list_sessions_in_course_response_schema_instance.to_dict()
# create an instance of ListSessionsInCourseResponseSchema from a dict
list_sessions_in_course_response_schema_from_dict = ListSessionsInCourseResponseSchema.from_dict(list_sessions_in_course_response_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


