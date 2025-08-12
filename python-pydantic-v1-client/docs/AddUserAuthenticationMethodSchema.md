# AddUserAuthenticationMethodSchema


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**email_address** | **str** | The email address of the user. | 
**authentication_method** | **str** | The authentication method to add to the user. | 
**custom_authentication_method_provider_id** | **str** | For SAML and OAUTH2, this is the provider ID (will be supplied by Uplimit to the developer). | 
**authentication_secret** | **str** | The unique identity for the authentication method. For SAML, this is the SAML entity ID. For OAUTH2, this is the subject identifier (or sub). | 

## Example

```python
from openapi_client.models.add_user_authentication_method_schema import AddUserAuthenticationMethodSchema

# TODO update the JSON string below
json = "{}"
# create an instance of AddUserAuthenticationMethodSchema from a JSON string
add_user_authentication_method_schema_instance = AddUserAuthenticationMethodSchema.from_json(json)
# print the JSON string representation of the object
print AddUserAuthenticationMethodSchema.to_json()

# convert the object into a dict
add_user_authentication_method_schema_dict = add_user_authentication_method_schema_instance.to_dict()
# create an instance of AddUserAuthenticationMethodSchema from a dict
add_user_authentication_method_schema_from_dict = AddUserAuthenticationMethodSchema.from_dict(add_user_authentication_method_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


