export * from './courseApi';
import { CourseApi } from './courseApi';
export * from './enrollmentApi';
import { EnrollmentApi } from './enrollmentApi';
export * from './exportApi';
import { ExportApi } from './exportApi';
export * from './platformApi';
import { PlatformApi } from './platformApi';
export * from './sessionApi';
import { SessionApi } from './sessionApi';
export * from './userApi';
import { UserApi } from './userApi';
import * as http from 'http';

export class HttpError extends Error {
    constructor (public response: http.IncomingMessage, public body: any, public statusCode?: number) {
        super('HTTP request failed');
        this.name = 'HttpError';
    }
}

export { RequestFile } from '../model/models';

export const APIS = [CourseApi, EnrollmentApi, ExportApi, PlatformApi, SessionApi, UserApi];
