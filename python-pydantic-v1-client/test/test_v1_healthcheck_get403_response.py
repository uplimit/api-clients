# coding: utf-8

"""
    Uplimit Organization API

    This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.

    The version of the OpenAPI document: 2025-03-17
    Contact: hello@uplimit.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from openapi_client.models.v1_healthcheck_get403_response import V1HealthcheckGet403Response  # noqa: E501

class TestV1HealthcheckGet403Response(unittest.TestCase):
    """V1HealthcheckGet403Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> V1HealthcheckGet403Response:
        """Test V1HealthcheckGet403Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `V1HealthcheckGet403Response`
        """
        model = V1HealthcheckGet403Response()  # noqa: E501
        if include_optional:
            return V1HealthcheckGet403Response(
                success = True,
                error = ''
            )
        else:
            return V1HealthcheckGet403Response(
                success = True,
                error = '',
        )
        """

    def testV1HealthcheckGet403Response(self):
        """Test V1HealthcheckGet403Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
