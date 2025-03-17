# coding: utf-8

"""
    Uplimit Organization API

    This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

    The version of the OpenAPI document: 2025-03-17
    Contact: hello@uplimit.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json



from pydantic import BaseModel, Field, StrictBool, StrictStr, validator

class UplimitUserInformationWithSessionCompletionStatusSchema(BaseModel):
    """
    UplimitUserInformationWithSessionCompletionStatusSchema
    """
    email_address: StrictStr = Field(default=..., alias="emailAddress", description="The email address of the user.")
    first_name: StrictStr = Field(default=..., alias="firstName", description="The first name of the user.")
    last_name: StrictStr = Field(default=..., alias="lastName", description="The last name of the user.")
    user_account_is_active: StrictBool = Field(default=..., alias="userAccountIsActive", description="Whether the user is allowed to access the Uplimit platform.")
    user_has_valid_subscription_enrollment: StrictBool = Field(default=..., alias="userHasValidSubscriptionEnrollment", description="Whether the user is activated in your organization.")
    uplimit_subscription_enrollment_id: StrictStr = Field(default=..., alias="uplimitSubscriptionEnrollmentId", description="Internal ID to identify the user's membership within your organization on Uplimit.")
    uplimit_subscription_commitment_id: StrictStr = Field(default=..., alias="uplimitSubscriptionCommitmentId", description="Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group.")
    uplimit_user_id: StrictStr = Field(default=..., alias="uplimitUserId", description="Internal ID to identify the user across the Uplimit platform.")
    session_completion_status: StrictStr = Field(default=..., alias="sessionCompletionStatus", description="Whether the user has completed the session according to pre-defined completion criteria.")
    __properties = ["emailAddress", "firstName", "lastName", "userAccountIsActive", "userHasValidSubscriptionEnrollment", "uplimitSubscriptionEnrollmentId", "uplimitSubscriptionCommitmentId", "uplimitUserId", "sessionCompletionStatus"]

    @validator('session_completion_status')
    def session_completion_status_validate_enum(cls, value):
        """Validates the enum"""
        if value not in ('PENDING', 'COMPLETED',):
            raise ValueError("must be one of enum values ('PENDING', 'COMPLETED')")
        return value

    class Config:
        """Pydantic configuration"""
        allow_population_by_field_name = True
        validate_assignment = True

    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.dict(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> UplimitUserInformationWithSessionCompletionStatusSchema:
        """Create an instance of UplimitUserInformationWithSessionCompletionStatusSchema from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> UplimitUserInformationWithSessionCompletionStatusSchema:
        """Create an instance of UplimitUserInformationWithSessionCompletionStatusSchema from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return UplimitUserInformationWithSessionCompletionStatusSchema.parse_obj(obj)

        _obj = UplimitUserInformationWithSessionCompletionStatusSchema.parse_obj({
            "email_address": obj.get("emailAddress"),
            "first_name": obj.get("firstName"),
            "last_name": obj.get("lastName"),
            "user_account_is_active": obj.get("userAccountIsActive"),
            "user_has_valid_subscription_enrollment": obj.get("userHasValidSubscriptionEnrollment"),
            "uplimit_subscription_enrollment_id": obj.get("uplimitSubscriptionEnrollmentId"),
            "uplimit_subscription_commitment_id": obj.get("uplimitSubscriptionCommitmentId"),
            "uplimit_user_id": obj.get("uplimitUserId"),
            "session_completion_status": obj.get("sessionCompletionStatus")
        })
        return _obj


