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

export class V1HealthcheckGet200Response {
    'success': boolean;
    /**
    * The Uplimit ID of the organization.
    */
    'uplimitOrganizationId': string;
    /**
    * The name of the organization as it appears in Uplimit.
    */
    'uplimitOrganizationName': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "success",
            "baseName": "success",
            "type": "boolean"
        },
        {
            "name": "uplimitOrganizationId",
            "baseName": "uplimitOrganizationId",
            "type": "string"
        },
        {
            "name": "uplimitOrganizationName",
            "baseName": "uplimitOrganizationName",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return V1HealthcheckGet200Response.attributeTypeMap;
    }
}

