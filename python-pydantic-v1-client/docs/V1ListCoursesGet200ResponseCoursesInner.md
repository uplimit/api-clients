# V1ListCoursesGet200ResponseCoursesInner


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uplimit_course_id** | **str** | Internal ID to identify the course across the Uplimit platform. | 
**uplimit_course_slug** | **str** | The slug (i.e. short name) of the course across the Uplimit platform. | 
**name** | **str** | The name of the course. | 

## Example

```python
from openapi_client.models.v1_list_courses_get200_response_courses_inner import V1ListCoursesGet200ResponseCoursesInner

# TODO update the JSON string below
json = "{}"
# create an instance of V1ListCoursesGet200ResponseCoursesInner from a JSON string
v1_list_courses_get200_response_courses_inner_instance = V1ListCoursesGet200ResponseCoursesInner.from_json(json)
# print the JSON string representation of the object
print V1ListCoursesGet200ResponseCoursesInner.to_json()

# convert the object into a dict
v1_list_courses_get200_response_courses_inner_dict = v1_list_courses_get200_response_courses_inner_instance.to_dict()
# create an instance of V1ListCoursesGet200ResponseCoursesInner from a dict
v1_list_courses_get200_response_courses_inner_from_dict = V1ListCoursesGet200ResponseCoursesInner.from_dict(v1_list_courses_get200_response_courses_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


