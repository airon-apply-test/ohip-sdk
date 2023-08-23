/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Customer Relationship Management Outbound API
 * APIs to cater for Customer Relationship Management external (outbound) functionality with OPERA. These APIs facilitate various operations related to getting data from an external system, and inserting it into OPERA.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { FacilityTaskType } from './FacilityTaskType';
import {
    FacilityTaskTypeFromJSON,
    FacilityTaskTypeFromJSONTyped,
    FacilityTaskTypeToJSON,
} from './FacilityTaskType';
import type { HousekeepingRoomStatusType } from './HousekeepingRoomStatusType';
import {
    HousekeepingRoomStatusTypeFromJSON,
    HousekeepingRoomStatusTypeFromJSONTyped,
    HousekeepingRoomStatusTypeToJSON,
} from './HousekeepingRoomStatusType';

/**
 * Holds housekeeping turndown service information for the room.
 * @export
 * @interface ResHousekeepingType
 */
export interface ResHousekeepingType {
    /**
     * Turndown instructions for the room.
     * @type {string}
     * @memberof ResHousekeepingType
     */
    instructions?: string;
    /**
     * 
     * @type {FacilityTaskType}
     * @memberof ResHousekeepingType
     */
    facilityTaskInfo?: FacilityTaskType;
    /**
     * Indicates if a linen change is necessary.
     * @type {boolean}
     * @memberof ResHousekeepingType
     */
    linenChange?: boolean;
    /**
     * Indicates whether guest wants turndown facility or not.
     * @type {boolean}
     * @memberof ResHousekeepingType
     */
    turndownRequested?: boolean;
    /**
     * This is the Turndown room service time.
     * @type {Date}
     * @memberof ResHousekeepingType
     */
    serviceTime?: Date;
    /**
     * Expected Start Time for housekeeping task(s).
     * @type {string}
     * @memberof ResHousekeepingType
     */
    expectedServiceTime?: string;
    /**
     * 
     * @type {HousekeepingRoomStatusType}
     * @memberof ResHousekeepingType
     */
    roomStatus?: HousekeepingRoomStatusType;
}

/**
 * Check if a given object implements the ResHousekeepingType interface.
 */
export function instanceOfResHousekeepingType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ResHousekeepingTypeFromJSON(json: any): ResHousekeepingType {
    return ResHousekeepingTypeFromJSONTyped(json, false);
}

export function ResHousekeepingTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResHousekeepingType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'instructions': !exists(json, 'instructions') ? undefined : json['instructions'],
        'facilityTaskInfo': !exists(json, 'facilityTaskInfo') ? undefined : FacilityTaskTypeFromJSON(json['facilityTaskInfo']),
        'linenChange': !exists(json, 'linenChange') ? undefined : json['linenChange'],
        'turndownRequested': !exists(json, 'turndownRequested') ? undefined : json['turndownRequested'],
        'serviceTime': !exists(json, 'serviceTime') ? undefined : (new Date(json['serviceTime'])),
        'expectedServiceTime': !exists(json, 'expectedServiceTime') ? undefined : json['expectedServiceTime'],
        'roomStatus': !exists(json, 'roomStatus') ? undefined : HousekeepingRoomStatusTypeFromJSON(json['roomStatus']),
    };
}

export function ResHousekeepingTypeToJSON(value?: ResHousekeepingType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'instructions': value.instructions,
        'facilityTaskInfo': FacilityTaskTypeToJSON(value.facilityTaskInfo),
        'linenChange': value.linenChange,
        'turndownRequested': value.turndownRequested,
        'serviceTime': value.serviceTime === undefined ? undefined : (value.serviceTime.toISOString().substr(0,10)),
        'expectedServiceTime': value.expectedServiceTime,
        'roomStatus': HousekeepingRoomStatusTypeToJSON(value.roomStatus),
    };
}

