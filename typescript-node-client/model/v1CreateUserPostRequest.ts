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

export class V1CreateUserPostRequest {
    /**
    * The email address of the user.
    */
    'emailAddress': string;
    /**
    * The first name of the user.
    */
    'firstName': string;
    /**
    * The last name of the user.
    */
    'lastName': string;
    /**
    * Internal ID to identify the “group” the user belongs to within your organization. Leaving this blank will enroll the user into the default group.
    */
    'subscriptionCommitmentId'?: string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "emailAddress",
            "baseName": "emailAddress",
            "type": "string"
        },
        {
            "name": "firstName",
            "baseName": "firstName",
            "type": "string"
        },
        {
            "name": "lastName",
            "baseName": "lastName",
            "type": "string"
        },
        {
            "name": "subscriptionCommitmentId",
            "baseName": "subscriptionCommitmentId",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return V1CreateUserPostRequest.attributeTypeMap;
    }
}

