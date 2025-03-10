# ListCoursesResponseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**courses** | [**List[V1ListCoursesGet200ResponseCoursesInner]**](V1ListCoursesGet200ResponseCoursesInner.md) |  | 
**total_count** | **float** |  | 

## Example

```python
from openapi_client.models.list_courses_response_schema import ListCoursesResponseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of ListCoursesResponseSchema from a JSON string
list_courses_response_schema_instance = ListCoursesResponseSchema.from_json(json)
# print the JSON string representation of the object
print ListCoursesResponseSchema.to_json()

# convert the object into a dict
list_courses_response_schema_dict = list_courses_response_schema_instance.to_dict()
# create an instance of ListCoursesResponseSchema from a dict
list_courses_response_schema_from_dict = ListCoursesResponseSchema.from_dict(list_courses_response_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


