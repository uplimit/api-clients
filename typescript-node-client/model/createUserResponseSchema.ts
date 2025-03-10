/**
 * Uplimit Organization API
 * This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.
 *
 * The version of the OpenAPI document: 2025-03-07
 * Contact: hello@uplimit.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class CreateUserResponseSchema {
    /**
    * Internal ID to identify the user\'s membership within your organization on Uplimit.
    */
    'uplimitSubscriptionEnrollmentId': string;
    /**
    * Internal ID to identify the user across the Uplimit platform.
    */
    'uplimitUserId': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "uplimitSubscriptionEnrollmentId",
            "baseName": "uplimitSubscriptionEnrollmentId",
            "type": "string"
        },
        {
            "name": "uplimitUserId",
            "baseName": "uplimitUserId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return CreateUserResponseSchema.attributeTypeMap;
    }
}

