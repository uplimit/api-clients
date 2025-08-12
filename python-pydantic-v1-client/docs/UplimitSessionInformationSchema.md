# UplimitSessionInformationSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uplimit_session_id** | **str** | Internal ID to identify the session across the Uplimit platform. | 
**name** | **str** | The name of the session. | 
**starts_at** | **datetime** | The start date of the session. | 
**enrollment_enabled** | **bool** | Whether the session allows enrollments | 

## Example

```python
from openapi_client.models.uplimit_session_information_schema import UplimitSessionInformationSchema

# TODO update the JSON string below
json = "{}"
# create an instance of UplimitSessionInformationSchema from a JSON string
uplimit_session_information_schema_instance = UplimitSessionInformationSchema.from_json(json)
# print the JSON string representation of the object
print UplimitSessionInformationSchema.to_json()

# convert the object into a dict
uplimit_session_information_schema_dict = uplimit_session_information_schema_instance.to_dict()
# create an instance of UplimitSessionInformationSchema from a dict
uplimit_session_information_schema_from_dict = UplimitSessionInformationSchema.from_dict(uplimit_session_information_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


