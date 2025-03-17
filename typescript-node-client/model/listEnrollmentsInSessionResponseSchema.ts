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
import { V1ListEnrollmentsInSessionGet200ResponseUsersInner } from './v1ListEnrollmentsInSessionGet200ResponseUsersInner';

export class ListEnrollmentsInSessionResponseSchema {
    'users': Array<V1ListEnrollmentsInSessionGet200ResponseUsersInner>;
    'totalCount': number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "users",
            "baseName": "users",
            "type": "Array<V1ListEnrollmentsInSessionGet200ResponseUsersInner>"
        },
        {
            "name": "totalCount",
            "baseName": "totalCount",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return ListEnrollmentsInSessionResponseSchema.attributeTypeMap;
    }
}

