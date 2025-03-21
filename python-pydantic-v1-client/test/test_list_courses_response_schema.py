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

from openapi_client.models.list_courses_response_schema import ListCoursesResponseSchema  # noqa: E501

class TestListCoursesResponseSchema(unittest.TestCase):
    """ListCoursesResponseSchema unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ListCoursesResponseSchema:
        """Test ListCoursesResponseSchema
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ListCoursesResponseSchema`
        """
        model = ListCoursesResponseSchema()  # noqa: E501
        if include_optional:
            return ListCoursesResponseSchema(
                courses = [
                    openapi_client.models._v1_list_courses_get_200_response_courses_inner._v1_ListCourses_get_200_response_courses_inner(
                        uplimit_course_id = '', 
                        uplimit_course_slug = '', 
                        name = '', )
                    ],
                total_count = 1.337
            )
        else:
            return ListCoursesResponseSchema(
                courses = [
                    openapi_client.models._v1_list_courses_get_200_response_courses_inner._v1_ListCourses_get_200_response_courses_inner(
                        uplimit_course_id = '', 
                        uplimit_course_slug = '', 
                        name = '', )
                    ],
                total_count = 1.337,
        )
        """

    def testListCoursesResponseSchema(self):
        """Test ListCoursesResponseSchema"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
