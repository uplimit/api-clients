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

export class GetUplimitSessionIdUsingDoceboSessionCodeResponseSchema {
    'uplimitSessionId': string;
    'name': string;
    'startsAt': Date;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "uplimitSessionId",
            "baseName": "uplimitSessionId",
            "type": "string"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string"
        },
        {
            "name": "startsAt",
            "baseName": "startsAt",
            "type": "Date"
        }    ];

    static getAttributeTypeMap() {
        return GetUplimitSessionIdUsingDoceboSessionCodeResponseSchema.attributeTypeMap;
    }
}

