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

export class V1ListCoursesGet200ResponseCoursesInner {
    /**
    * Internal ID to identify the course across the Uplimit platform.
    */
    'uplimitCourseId': string;
    /**
    * The slug (i.e. short name) of the course across the Uplimit platform.
    */
    'uplimitCourseSlug': string;
    /**
    * The name of the course.
    */
    'name': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "uplimitCourseId",
            "baseName": "uplimitCourseId",
            "type": "string"
        },
        {
            "name": "uplimitCourseSlug",
            "baseName": "uplimitCourseSlug",
            "type": "string"
        },
        {
            "name": "name",
            "baseName": "name",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return V1ListCoursesGet200ResponseCoursesInner.attributeTypeMap;
    }
}

