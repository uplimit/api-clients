# UplimitUserInformationSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**first_name** | **str** | The first name of the user. | 
**last_name** | **str** | The last name of the user. | 
**user_account_is_active** | **bool** | Whether the user is allowed to access the Uplimit platform. | 
**user_has_valid_subscription_enrollment** | **bool** | Whether the user is activated in your organization. | 
**uplimit_subscription_enrollment_id** | **str** | Internal ID to identify the user&#39;s membership within your organization on Uplimit. | 
**uplimit_subscription_commitment_id** | **str** | Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group. | 
**uplimit_user_id** | **str** | Internal ID to identify the user across the Uplimit platform. | 

## Example

```python
from openapi_client.models.uplimit_user_information_schema import UplimitUserInformationSchema

# TODO update the JSON string below
json = "{}"
# create an instance of UplimitUserInformationSchema from a JSON string
uplimit_user_information_schema_instance = UplimitUserInformationSchema.from_json(json)
# print the JSON string representation of the object
print UplimitUserInformationSchema.to_json()

# convert the object into a dict
uplimit_user_information_schema_dict = uplimit_user_information_schema_instance.to_dict()
# create an instance of UplimitUserInformationSchema from a dict
uplimit_user_information_schema_from_dict = UplimitUserInformationSchema.from_dict(uplimit_user_information_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


