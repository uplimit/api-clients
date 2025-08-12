# EnrollUserIntoCourseSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**uplimit_course_id** | **str** | Internal ID to identify the course across the Uplimit platform. | 
**uplimit_enroll_user_into_course_session_selection_policy** | **str** | The policy to decide which session to enroll a user into when enrolling the user into a course. | 
**subscription_commitment_id** | **str** | Internal ID to identify the \&quot;group\&quot; the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | [optional] 

## Example

```python
from openapi_client.models.enroll_user_into_course_schema import EnrollUserIntoCourseSchema

# TODO update the JSON string below
json = "{}"
# create an instance of EnrollUserIntoCourseSchema from a JSON string
enroll_user_into_course_schema_instance = EnrollUserIntoCourseSchema.from_json(json)
# print the JSON string representation of the object
print EnrollUserIntoCourseSchema.to_json()

# convert the object into a dict
enroll_user_into_course_schema_dict = enroll_user_into_course_schema_instance.to_dict()
# create an instance of EnrollUserIntoCourseSchema from a dict
enroll_user_into_course_schema_from_dict = EnrollUserIntoCourseSchema.from_dict(enroll_user_into_course_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


