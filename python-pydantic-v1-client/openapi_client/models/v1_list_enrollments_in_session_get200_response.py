# coding: utf-8

"""
    Uplimit Organization API

    This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

    The version of the OpenAPI document: 2025-03-07
    Contact: hello@uplimit.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import List, Union
from pydantic import BaseModel, Field, StrictFloat, StrictInt, conlist
from openapi_client.models.v1_list_enrollments_in_session_get200_response_users_inner import V1ListEnrollmentsInSessionGet200ResponseUsersInner

class V1ListEnrollmentsInSessionGet200Response(BaseModel):
    """
    V1ListEnrollmentsInSessionGet200Response
    """
    users: conlist(V1ListEnrollmentsInSessionGet200ResponseUsersInner) = Field(...)
    total_count: Union[StrictFloat, StrictInt] = Field(default=..., alias="totalCount")
    __properties = ["users", "totalCount"]

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
    def from_json(cls, json_str: str) -> V1ListEnrollmentsInSessionGet200Response:
        """Create an instance of V1ListEnrollmentsInSessionGet200Response from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self):
        """Returns the dictionary representation of the model using alias"""
        _dict = self.dict(by_alias=True,
                          exclude={
                          },
                          exclude_none=True)
        # override the default output from pydantic by calling `to_dict()` of each item in users (list)
        _items = []
        if self.users:
            for _item in self.users:
                if _item:
                    _items.append(_item.to_dict())
            _dict['users'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: dict) -> V1ListEnrollmentsInSessionGet200Response:
        """Create an instance of V1ListEnrollmentsInSessionGet200Response from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return V1ListEnrollmentsInSessionGet200Response.parse_obj(obj)

        _obj = V1ListEnrollmentsInSessionGet200Response.parse_obj({
            "users": [V1ListEnrollmentsInSessionGet200ResponseUsersInner.from_dict(_item) for _item in obj.get("users")] if obj.get("users") is not None else None,
            "total_count": obj.get("totalCount")
        })
        return _obj


