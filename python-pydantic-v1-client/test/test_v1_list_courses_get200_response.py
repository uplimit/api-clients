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

from openapi_client.models.v1_list_courses_get200_response import V1ListCoursesGet200Response  # noqa: E501

class TestV1ListCoursesGet200Response(unittest.TestCase):
    """V1ListCoursesGet200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> V1ListCoursesGet200Response:
        """Test V1ListCoursesGet200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `V1ListCoursesGet200Response`
        """
        model = V1ListCoursesGet200Response()  # noqa: E501
        if include_optional:
            return V1ListCoursesGet200Response(
                courses = [
                    openapi_client.models._v1_list_courses_get_200_response_courses_inner._v1_ListCourses_get_200_response_courses_inner(
                        uplimit_course_id = '', 
                        uplimit_course_slug = '', 
                        name = '', )
                    ],
                total_count = 1.337
            )
        else:
            return V1ListCoursesGet200Response(
                courses = [
                    openapi_client.models._v1_list_courses_get_200_response_courses_inner._v1_ListCourses_get_200_response_courses_inner(
                        uplimit_course_id = '', 
                        uplimit_course_slug = '', 
                        name = '', )
                    ],
                total_count = 1.337,
        )
        """

    def testV1ListCoursesGet200Response(self):
        """Test V1ListCoursesGet200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
