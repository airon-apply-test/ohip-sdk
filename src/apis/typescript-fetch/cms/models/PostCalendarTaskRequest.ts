/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud API for Customer Management Service
 * This API deals with the different aspect of the CustomerManagement.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CalendarTaskType } from './CalendarTaskType';
import {
    CalendarTaskTypeFromJSON,
    CalendarTaskTypeFromJSONTyped,
    CalendarTaskTypeToJSON,
} from './CalendarTaskType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * 
 * @export
 * @interface PostCalendarTaskRequest
 */
export interface PostCalendarTaskRequest {
    /**
     * 
     * @type {CalendarTaskType}
     * @memberof PostCalendarTaskRequest
     */
    calendarTaskDetails?: CalendarTaskType;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof PostCalendarTaskRequest
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof PostCalendarTaskRequest
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the PostCalendarTaskRequest interface.
 */
export function instanceOfPostCalendarTaskRequest(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function PostCalendarTaskRequestFromJSON(json: any): PostCalendarTaskRequest {
    return PostCalendarTaskRequestFromJSONTyped(json, false);
}

export function PostCalendarTaskRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostCalendarTaskRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'calendarTaskDetails': !exists(json, 'calendarTaskDetails') ? undefined : CalendarTaskTypeFromJSON(json['calendarTaskDetails']),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function PostCalendarTaskRequestToJSON(value?: PostCalendarTaskRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'calendarTaskDetails': CalendarTaskTypeToJSON(value.calendarTaskDetails),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

