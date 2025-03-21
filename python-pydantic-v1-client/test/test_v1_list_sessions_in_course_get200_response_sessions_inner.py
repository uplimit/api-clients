# coding: utf-8

"""
    Uplimit Organization API

    This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

    The version of the OpenAPI document: 2025-03-07
    Contact: hello@uplimit.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from openapi_client.models.v1_list_sessions_in_course_get200_response_sessions_inner import V1ListSessionsInCourseGet200ResponseSessionsInner  # noqa: E501

class TestV1ListSessionsInCourseGet200ResponseSessionsInner(unittest.TestCase):
    """V1ListSessionsInCourseGet200ResponseSessionsInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> V1ListSessionsInCourseGet200ResponseSessionsInner:
        """Test V1ListSessionsInCourseGet200ResponseSessionsInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `V1ListSessionsInCourseGet200ResponseSessionsInner`
        """
        model = V1ListSessionsInCourseGet200ResponseSessionsInner()  # noqa: E501
        if include_optional:
            return V1ListSessionsInCourseGet200ResponseSessionsInner(
                uplimit_session_id = '',
                name = '',
                starts_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f')
            )
        else:
            return V1ListSessionsInCourseGet200ResponseSessionsInner(
                uplimit_session_id = '',
                name = '',
                starts_at = datetime.datetime.strptime('2013-10-20 19:20:30.00', '%Y-%m-%d %H:%M:%S.%f'),
        )
        """

    def testV1ListSessionsInCourseGet200ResponseSessionsInner(self):
        """Test V1ListSessionsInCourseGet200ResponseSessionsInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
