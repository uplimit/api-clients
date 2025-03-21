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

from openapi_client.models.v1_list_enrollments_in_session_get200_response_users_inner import V1ListEnrollmentsInSessionGet200ResponseUsersInner  # noqa: E501

class TestV1ListEnrollmentsInSessionGet200ResponseUsersInner(unittest.TestCase):
    """V1ListEnrollmentsInSessionGet200ResponseUsersInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> V1ListEnrollmentsInSessionGet200ResponseUsersInner:
        """Test V1ListEnrollmentsInSessionGet200ResponseUsersInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `V1ListEnrollmentsInSessionGet200ResponseUsersInner`
        """
        model = V1ListEnrollmentsInSessionGet200ResponseUsersInner()  # noqa: E501
        if include_optional:
            return V1ListEnrollmentsInSessionGet200ResponseUsersInner(
                email_address = '',
                first_name = '',
                last_name = '',
                user_account_is_active = True,
                user_has_valid_subscription_enrollment = True,
                uplimit_subscription_enrollment_id = '',
                uplimit_subscription_commitment_id = '',
                uplimit_user_id = '',
                session_completion_status = 'PENDING'
            )
        else:
            return V1ListEnrollmentsInSessionGet200ResponseUsersInner(
                email_address = '',
                first_name = '',
                last_name = '',
                user_account_is_active = True,
                user_has_valid_subscription_enrollment = True,
                uplimit_subscription_enrollment_id = '',
                uplimit_subscription_commitment_id = '',
                uplimit_user_id = '',
                session_completion_status = 'PENDING',
        )
        """

    def testV1ListEnrollmentsInSessionGet200ResponseUsersInner(self):
        """Test V1ListEnrollmentsInSessionGet200ResponseUsersInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
