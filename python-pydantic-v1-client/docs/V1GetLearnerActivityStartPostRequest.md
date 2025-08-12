# V1GetLearnerActivityStartPostRequest


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**session_ids** | **List[str]** | The IDs of the sessions to export activity for. If not provided, all sessions will be exported. | [optional] 

## Example

```python
from openapi_client.models.v1_get_learner_activity_start_post_request import V1GetLearnerActivityStartPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of V1GetLearnerActivityStartPostRequest from a JSON string
v1_get_learner_activity_start_post_request_instance = V1GetLearnerActivityStartPostRequest.from_json(json)
# print the JSON string representation of the object
print V1GetLearnerActivityStartPostRequest.to_json()

# convert the object into a dict
v1_get_learner_activity_start_post_request_dict = v1_get_learner_activity_start_post_request_instance.to_dict()
# create an instance of V1GetLearnerActivityStartPostRequest from a dict
v1_get_learner_activity_start_post_request_from_dict = V1GetLearnerActivityStartPostRequest.from_dict(v1_get_learner_activity_start_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


