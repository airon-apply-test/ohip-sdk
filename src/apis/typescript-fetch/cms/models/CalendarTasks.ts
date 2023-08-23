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
import type { CalendarTaskCountType } from './CalendarTaskCountType';
import {
    CalendarTaskCountTypeFromJSON,
    CalendarTaskCountTypeFromJSONTyped,
    CalendarTaskCountTypeToJSON,
} from './CalendarTaskCountType';
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
 * Response for fetching calendar tasks.
 * @export
 * @interface CalendarTasks
 */
export interface CalendarTasks {
    /**
     * Detailed information regarding calendar task.
     * @type {Array<CalendarTaskType>}
     * @memberof CalendarTasks
     */
    taskSummary?: Array<CalendarTaskType>;
    /**
     * Defines summary count of calendar tasks that belongs to specific classification.
     * @type {Array<CalendarTaskCountType>}
     * @memberof CalendarTasks
     */
    countSummary?: Array<CalendarTaskCountType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof CalendarTasks
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof CalendarTasks
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the CalendarTasks interface.
 */
export function instanceOfCalendarTasks(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CalendarTasksFromJSON(json: any): CalendarTasks {
    return CalendarTasksFromJSONTyped(json, false);
}

export function CalendarTasksFromJSONTyped(json: any, ignoreDiscriminator: boolean): CalendarTasks {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'taskSummary': !exists(json, 'taskSummary') ? undefined : ((json['taskSummary'] as Array<any>).map(CalendarTaskTypeFromJSON)),
        'countSummary': !exists(json, 'countSummary') ? undefined : ((json['countSummary'] as Array<any>).map(CalendarTaskCountTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function CalendarTasksToJSON(value?: CalendarTasks | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'taskSummary': value.taskSummary === undefined ? undefined : ((value.taskSummary as Array<any>).map(CalendarTaskTypeToJSON)),
        'countSummary': value.countSummary === undefined ? undefined : ((value.countSummary as Array<any>).map(CalendarTaskCountTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

