/**
 * Uplimit Organization API
 * This API is used to manage organizations within the Uplimit platform. For more information, please reach out to your Uplimit Enterprise contact.
 *
 * The version of the OpenAPI document: 2025-03-17
 * Contact: hello@uplimit.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';

export class ToggleUserActivationSchema {
    /**
    * The email address of the user.
    */
    'emailAddress': string;
    /**
    * Whether to set the user as active or inactive.
    */
    'setIsActive': boolean;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "emailAddress",
            "baseName": "emailAddress",
            "type": "string"
        },
        {
            "name": "setIsActive",
            "baseName": "setIsActive",
            "type": "boolean"
        }    ];

    static getAttributeTypeMap() {
        return ToggleUserActivationSchema.attributeTypeMap;
    }
}

