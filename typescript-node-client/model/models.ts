import localVarRequest from 'request';

export * from './createUserResponseSchema';
export * from './createUserSchema';
export * from './enrollUserIntoCourseSchema';
export * from './healthcheckIncorrectApiKeyResponseSchema';
export * from './healthcheckSuccessResponseSchema';
export * from './listActiveUsersResponseSchema';
export * from './listCoursesResponseSchema';
export * from './listEnrollmentsInSessionResponseSchema';
export * from './listInactiveUsersResponseSchema';
export * from './listSessionsInCourseResponseSchema';
export * from './toggleUserActivationSchema';
export * from './uplimitCourseInformationSchema';
export * from './uplimitSessionInformationSchema';
export * from './uplimitUserInformationSchema';
export * from './uplimitUserInformationWithSessionCompletionStatusSchema';
export * from './v1CreateUserPost200Response';
export * from './v1CreateUserPost400Response';
export * from './v1CreateUserPostRequest';
export * from './v1EnrollUserIntoSessionPost200Response';
export * from './v1EnrollUserIntoSessionPostRequest';
export * from './v1GetUserInformationEmailAddressOrUserIdGet200Response';
export * from './v1HealthcheckGet200Response';
export * from './v1HealthcheckGet403Response';
export * from './v1ListActiveUsersGet200Response';
export * from './v1ListCoursesGet200Response';
export * from './v1ListCoursesGet200ResponseCoursesInner';
export * from './v1ListEnrollmentsInSessionGet200Response';
export * from './v1ListEnrollmentsInSessionGet200ResponseUsersInner';
export * from './v1ListSessionsInCourseGet200Response';
export * from './v1ListSessionsInCourseGet200ResponseSessionsInner';
export * from './v1ToggleUserActivationPostRequest';

import * as fs from 'fs';

export interface RequestDetailedFile {
    value: Buffer;
    options?: {
        filename?: string;
        contentType?: string;
    }
}

export type RequestFile = string | Buffer | fs.ReadStream | RequestDetailedFile;


import { CreateUserResponseSchema } from './createUserResponseSchema';
import { CreateUserSchema } from './createUserSchema';
import { EnrollUserIntoCourseSchema } from './enrollUserIntoCourseSchema';
import { HealthcheckIncorrectApiKeyResponseSchema } from './healthcheckIncorrectApiKeyResponseSchema';
import { HealthcheckSuccessResponseSchema } from './healthcheckSuccessResponseSchema';
import { ListActiveUsersResponseSchema } from './listActiveUsersResponseSchema';
import { ListCoursesResponseSchema } from './listCoursesResponseSchema';
import { ListEnrollmentsInSessionResponseSchema } from './listEnrollmentsInSessionResponseSchema';
import { ListInactiveUsersResponseSchema } from './listInactiveUsersResponseSchema';
import { ListSessionsInCourseResponseSchema } from './listSessionsInCourseResponseSchema';
import { ToggleUserActivationSchema } from './toggleUserActivationSchema';
import { UplimitCourseInformationSchema } from './uplimitCourseInformationSchema';
import { UplimitSessionInformationSchema } from './uplimitSessionInformationSchema';
import { UplimitUserInformationSchema } from './uplimitUserInformationSchema';
import { UplimitUserInformationWithSessionCompletionStatusSchema } from './uplimitUserInformationWithSessionCompletionStatusSchema';
import { V1CreateUserPost200Response } from './v1CreateUserPost200Response';
import { V1CreateUserPost400Response } from './v1CreateUserPost400Response';
import { V1CreateUserPostRequest } from './v1CreateUserPostRequest';
import { V1EnrollUserIntoSessionPost200Response } from './v1EnrollUserIntoSessionPost200Response';
import { V1EnrollUserIntoSessionPostRequest } from './v1EnrollUserIntoSessionPostRequest';
import { V1GetUserInformationEmailAddressOrUserIdGet200Response } from './v1GetUserInformationEmailAddressOrUserIdGet200Response';
import { V1HealthcheckGet200Response } from './v1HealthcheckGet200Response';
import { V1HealthcheckGet403Response } from './v1HealthcheckGet403Response';
import { V1ListActiveUsersGet200Response } from './v1ListActiveUsersGet200Response';
import { V1ListCoursesGet200Response } from './v1ListCoursesGet200Response';
import { V1ListCoursesGet200ResponseCoursesInner } from './v1ListCoursesGet200ResponseCoursesInner';
import { V1ListEnrollmentsInSessionGet200Response } from './v1ListEnrollmentsInSessionGet200Response';
import { V1ListEnrollmentsInSessionGet200ResponseUsersInner } from './v1ListEnrollmentsInSessionGet200ResponseUsersInner';
import { V1ListSessionsInCourseGet200Response } from './v1ListSessionsInCourseGet200Response';
import { V1ListSessionsInCourseGet200ResponseSessionsInner } from './v1ListSessionsInCourseGet200ResponseSessionsInner';
import { V1ToggleUserActivationPostRequest } from './v1ToggleUserActivationPostRequest';

/* tslint:disable:no-unused-variable */
let primitives = [
                    "string",
                    "boolean",
                    "double",
                    "integer",
                    "long",
                    "float",
                    "number",
                    "any"
                 ];

let enumsMap: {[index: string]: any} = {
        "UplimitUserInformationWithSessionCompletionStatusSchema.SessionCompletionStatusEnum": UplimitUserInformationWithSessionCompletionStatusSchema.SessionCompletionStatusEnum,
        "V1ListEnrollmentsInSessionGet200ResponseUsersInner.SessionCompletionStatusEnum": V1ListEnrollmentsInSessionGet200ResponseUsersInner.SessionCompletionStatusEnum,
}

let typeMap: {[index: string]: any} = {
    "CreateUserResponseSchema": CreateUserResponseSchema,
    "CreateUserSchema": CreateUserSchema,
    "EnrollUserIntoCourseSchema": EnrollUserIntoCourseSchema,
    "HealthcheckIncorrectApiKeyResponseSchema": HealthcheckIncorrectApiKeyResponseSchema,
    "HealthcheckSuccessResponseSchema": HealthcheckSuccessResponseSchema,
    "ListActiveUsersResponseSchema": ListActiveUsersResponseSchema,
    "ListCoursesResponseSchema": ListCoursesResponseSchema,
    "ListEnrollmentsInSessionResponseSchema": ListEnrollmentsInSessionResponseSchema,
    "ListInactiveUsersResponseSchema": ListInactiveUsersResponseSchema,
    "ListSessionsInCourseResponseSchema": ListSessionsInCourseResponseSchema,
    "ToggleUserActivationSchema": ToggleUserActivationSchema,
    "UplimitCourseInformationSchema": UplimitCourseInformationSchema,
    "UplimitSessionInformationSchema": UplimitSessionInformationSchema,
    "UplimitUserInformationSchema": UplimitUserInformationSchema,
    "UplimitUserInformationWithSessionCompletionStatusSchema": UplimitUserInformationWithSessionCompletionStatusSchema,
    "V1CreateUserPost200Response": V1CreateUserPost200Response,
    "V1CreateUserPost400Response": V1CreateUserPost400Response,
    "V1CreateUserPostRequest": V1CreateUserPostRequest,
    "V1EnrollUserIntoSessionPost200Response": V1EnrollUserIntoSessionPost200Response,
    "V1EnrollUserIntoSessionPostRequest": V1EnrollUserIntoSessionPostRequest,
    "V1GetUserInformationEmailAddressOrUserIdGet200Response": V1GetUserInformationEmailAddressOrUserIdGet200Response,
    "V1HealthcheckGet200Response": V1HealthcheckGet200Response,
    "V1HealthcheckGet403Response": V1HealthcheckGet403Response,
    "V1ListActiveUsersGet200Response": V1ListActiveUsersGet200Response,
    "V1ListCoursesGet200Response": V1ListCoursesGet200Response,
    "V1ListCoursesGet200ResponseCoursesInner": V1ListCoursesGet200ResponseCoursesInner,
    "V1ListEnrollmentsInSessionGet200Response": V1ListEnrollmentsInSessionGet200Response,
    "V1ListEnrollmentsInSessionGet200ResponseUsersInner": V1ListEnrollmentsInSessionGet200ResponseUsersInner,
    "V1ListSessionsInCourseGet200Response": V1ListSessionsInCourseGet200Response,
    "V1ListSessionsInCourseGet200ResponseSessionsInner": V1ListSessionsInCourseGet200ResponseSessionsInner,
    "V1ToggleUserActivationPostRequest": V1ToggleUserActivationPostRequest,
}

// Check if a string starts with another string without using es6 features
function startsWith(str: string, match: string): boolean {
    return str.substring(0, match.length) === match;
}

// Check if a string ends with another string without using es6 features
function endsWith(str: string, match: string): boolean {
    return str.length >= match.length && str.substring(str.length - match.length) === match;
}

const nullableSuffix = " | null";
const optionalSuffix = " | undefined";
const arrayPrefix = "Array<";
const arraySuffix = ">";
const mapPrefix = "{ [key: string]: ";
const mapSuffix = "; }";

export class ObjectSerializer {
    public static findCorrectType(data: any, expectedType: string) {
        if (data == undefined) {
            return expectedType;
        } else if (primitives.indexOf(expectedType.toLowerCase()) !== -1) {
            return expectedType;
        } else if (expectedType === "Date") {
            return expectedType;
        } else {
            if (enumsMap[expectedType]) {
                return expectedType;
            }

            if (!typeMap[expectedType]) {
                return expectedType; // w/e we don't know the type
            }

            // Check the discriminator
            let discriminatorProperty = typeMap[expectedType].discriminator;
            if (discriminatorProperty == null) {
                return expectedType; // the type does not have a discriminator. use it.
            } else {
                if (data[discriminatorProperty]) {
                    var discriminatorType = data[discriminatorProperty];
                    if(typeMap[discriminatorType]){
                        return discriminatorType; // use the type given in the discriminator
                    } else {
                        return expectedType; // discriminator did not map to a type
                    }
                } else {
                    return expectedType; // discriminator was not present (or an empty string)
                }
            }
        }
    }

    public static serialize(data: any, type: string): any {
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (endsWith(type, nullableSuffix)) {
            let subType: string = type.slice(0, -nullableSuffix.length); // Type | null => Type
            return ObjectSerializer.serialize(data, subType);
        } else if (endsWith(type, optionalSuffix)) {
            let subType: string = type.slice(0, -optionalSuffix.length); // Type | undefined => Type
            return ObjectSerializer.serialize(data, subType);
        } else if (startsWith(type, arrayPrefix)) {
            let subType: string = type.slice(arrayPrefix.length, -arraySuffix.length); // Array<Type> => Type
            let transformedData: any[] = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.serialize(datum, subType));
            }
            return transformedData;
        } else if (startsWith(type, mapPrefix)) {
            let subType: string = type.slice(mapPrefix.length, -mapSuffix.length); // { [key: string]: Type; } => Type
            let transformedData: { [key: string]: any } = {};
            for (let key in data) {
                transformedData[key] = ObjectSerializer.serialize(
                    data[key],
                    subType,
                );
            }
            return transformedData;
        } else if (type === "Date") {
            return data.toISOString();
        } else {
            if (enumsMap[type]) {
                return data;
            }
            if (!typeMap[type]) { // in case we dont know the type
                return data;
            }

            // Get the actual type of this object
            type = this.findCorrectType(data, type);

            // get the map for the correct type.
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            let instance: {[index: string]: any} = {};
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.baseName] = ObjectSerializer.serialize(data[attributeType.name], attributeType.type);
            }
            return instance;
        }
    }

    public static deserialize(data: any, type: string): any {
        // polymorphism may change the actual type.
        type = ObjectSerializer.findCorrectType(data, type);
        if (data == undefined) {
            return data;
        } else if (primitives.indexOf(type.toLowerCase()) !== -1) {
            return data;
        } else if (endsWith(type, nullableSuffix)) {
            let subType: string = type.slice(0, -nullableSuffix.length); // Type | null => Type
            return ObjectSerializer.deserialize(data, subType);
        } else if (endsWith(type, optionalSuffix)) {
            let subType: string = type.slice(0, -optionalSuffix.length); // Type | undefined => Type
            return ObjectSerializer.deserialize(data, subType);
        } else if (startsWith(type, arrayPrefix)) {
            let subType: string = type.slice(arrayPrefix.length, -arraySuffix.length); // Array<Type> => Type
            let transformedData: any[] = [];
            for (let index = 0; index < data.length; index++) {
                let datum = data[index];
                transformedData.push(ObjectSerializer.deserialize(datum, subType));
            }
            return transformedData;
        } else if (startsWith(type, mapPrefix)) {
            let subType: string = type.slice(mapPrefix.length, -mapSuffix.length); // { [key: string]: Type; } => Type
            let transformedData: { [key: string]: any } = {};
            for (let key in data) {
                transformedData[key] = ObjectSerializer.deserialize(
                    data[key],
                    subType,
                );
            }
            return transformedData;
        } else if (type === "Date") {
            return new Date(data);
        } else {
            if (enumsMap[type]) {// is Enum
                return data;
            }

            if (!typeMap[type]) { // dont know the type
                return data;
            }
            let instance = new typeMap[type]();
            let attributeTypes = typeMap[type].getAttributeTypeMap();
            for (let index = 0; index < attributeTypes.length; index++) {
                let attributeType = attributeTypes[index];
                instance[attributeType.name] = ObjectSerializer.deserialize(data[attributeType.baseName], attributeType.type);
            }
            return instance;
        }
    }
}

export interface Authentication {
    /**
    * Apply authentication settings to header and query params.
    */
    applyToRequest(requestOptions: localVarRequest.Options): Promise<void> | void;
}

export class HttpBasicAuth implements Authentication {
    public username: string = '';
    public password: string = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        requestOptions.auth = {
            username: this.username, password: this.password
        }
    }
}

export class HttpBearerAuth implements Authentication {
    public accessToken: string | (() => string) = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (requestOptions && requestOptions.headers) {
            const accessToken = typeof this.accessToken === 'function'
                            ? this.accessToken()
                            : this.accessToken;
            requestOptions.headers["Authorization"] = "Bearer " + accessToken;
        }
    }
}

export class ApiKeyAuth implements Authentication {
    public apiKey: string = '';

    constructor(private location: string, private paramName: string) {
    }

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (this.location == "query") {
            (<any>requestOptions.qs)[this.paramName] = this.apiKey;
        } else if (this.location == "header" && requestOptions && requestOptions.headers) {
            requestOptions.headers[this.paramName] = this.apiKey;
        } else if (this.location == 'cookie' && requestOptions && requestOptions.headers) {
            if (requestOptions.headers['Cookie']) {
                requestOptions.headers['Cookie'] += '; ' + this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
            else {
                requestOptions.headers['Cookie'] = this.paramName + '=' + encodeURIComponent(this.apiKey);
            }
        }
    }
}

export class OAuth implements Authentication {
    public accessToken: string = '';

    applyToRequest(requestOptions: localVarRequest.Options): void {
        if (requestOptions && requestOptions.headers) {
            requestOptions.headers["Authorization"] = "Bearer " + this.accessToken;
        }
    }
}

export class VoidAuth implements Authentication {
    public username: string = '';
    public password: string = '';

    applyToRequest(_: localVarRequest.Options): void {
        // Do nothing
    }
}

export type Interceptor = (requestOptions: localVarRequest.Options) => (Promise<void> | void);
