# UplimitCourseInformationSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uplimit_course_id** | **str** | Internal ID to identify the course across the Uplimit platform. | 
**uplimit_course_slug** | **str** | The slug (i.e. short name) of the course across the Uplimit platform. | 
**uplimit_course_description** | **str** | The description of the course (may be empty). | 
**name** | **str** | The name of the course. | 

## Example

```python
from openapi_client.models.uplimit_course_information_schema import UplimitCourseInformationSchema

# TODO update the JSON string below
json = "{}"
# create an instance of UplimitCourseInformationSchema from a JSON string
uplimit_course_information_schema_instance = UplimitCourseInformationSchema.from_json(json)
# print the JSON string representation of the object
print UplimitCourseInformationSchema.to_json()

# convert the object into a dict
uplimit_course_information_schema_dict = uplimit_course_information_schema_instance.to_dict()
# create an instance of UplimitCourseInformationSchema from a dict
uplimit_course_information_schema_from_dict = UplimitCourseInformationSchema.from_dict(uplimit_course_information_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


