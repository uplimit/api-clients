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

from openapi_client.models.v1_list_active_users_get200_response import V1ListActiveUsersGet200Response  # noqa: E501

class TestV1ListActiveUsersGet200Response(unittest.TestCase):
    """V1ListActiveUsersGet200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> V1ListActiveUsersGet200Response:
        """Test V1ListActiveUsersGet200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `V1ListActiveUsersGet200Response`
        """
        model = V1ListActiveUsersGet200Response()  # noqa: E501
        if include_optional:
            return V1ListActiveUsersGet200Response(
                users = [
                    openapi_client.models._v1_get_user_information__email_address_or_user_id__get_200_response._v1_GetUserInformation__emailAddressOrUserId__get_200_response(
                        email_address = '', 
                        first_name = '', 
                        last_name = '', 
                        user_account_is_active = True, 
                        user_has_valid_subscription_enrollment = True, 
                        uplimit_subscription_enrollment_id = '', 
                        uplimit_subscription_commitment_id = '', 
                        uplimit_user_id = '', )
                    ],
                total_count = 1.337
            )
        else:
            return V1ListActiveUsersGet200Response(
                users = [
                    openapi_client.models._v1_get_user_information__email_address_or_user_id__get_200_response._v1_GetUserInformation__emailAddressOrUserId__get_200_response(
                        email_address = '', 
                        first_name = '', 
                        last_name = '', 
                        user_account_is_active = True, 
                        user_has_valid_subscription_enrollment = True, 
                        uplimit_subscription_enrollment_id = '', 
                        uplimit_subscription_commitment_id = '', 
                        uplimit_user_id = '', )
                    ],
                total_count = 1.337,
        )
        """

    def testV1ListActiveUsersGet200Response(self):
        """Test V1ListActiveUsersGet200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
