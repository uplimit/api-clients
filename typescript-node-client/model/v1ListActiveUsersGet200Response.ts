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
import { V1GetUserInformationEmailAddressOrUserIdGet200Response } from './v1GetUserInformationEmailAddressOrUserIdGet200Response';

export class V1ListActiveUsersGet200Response {
    'users': Array<V1GetUserInformationEmailAddressOrUserIdGet200Response>;
    'totalCount': number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "users",
            "baseName": "users",
            "type": "Array<V1GetUserInformationEmailAddressOrUserIdGet200Response>"
        },
        {
            "name": "totalCount",
            "baseName": "totalCount",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return V1ListActiveUsersGet200Response.attributeTypeMap;
    }
}

