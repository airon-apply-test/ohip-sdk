/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { DateRangeType } from './DateRangeType';
import {
    DateRangeTypeFromJSON,
    DateRangeTypeFromJSONTyped,
    DateRangeTypeToJSON,
} from './DateRangeType';
import type { HousekeepingRoomStatusType } from './HousekeepingRoomStatusType';
import {
    HousekeepingRoomStatusTypeFromJSON,
    HousekeepingRoomStatusTypeFromJSONTyped,
    HousekeepingRoomStatusTypeToJSON,
} from './HousekeepingRoomStatusType';

/**
 * Out Of Order and Out Of Service Room information.
 * @export
 * @interface RoomOutOfOrderType
 */
export interface RoomOutOfOrderType {
    /**
     * Notes or Remarks on the OO/OS room
     * @type {string}
     * @memberof RoomOutOfOrderType
     */
    repairRemarks?: string;
    /**
     * 
     * @type {HousekeepingRoomStatusType}
     * @memberof RoomOutOfOrderType
     */
    returnStatus?: HousekeepingRoomStatusType;
    /**
     * Reason Code why the room is OO/OS.
     * @type {string}
     * @memberof RoomOutOfOrderType
     */
    reasonCode?: string;
    /**
     * Reason Description why the room is OO/OS.
     * @type {string}
     * @memberof RoomOutOfOrderType
     */
    reasonDesc?: string;
    /**
     * 
     * @type {DateRangeType}
     * @memberof RoomOutOfOrderType
     */
    newDateRange?: DateRangeType;
    /**
     * 
     * @type {HousekeepingRoomStatusType}
     * @memberof RoomOutOfOrderType
     */
    roomStatus?: HousekeepingRoomStatusType;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof RoomOutOfOrderType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof RoomOutOfOrderType
     */
    end?: Date;
}

/**
 * Check if a given object implements the RoomOutOfOrderType interface.
 */
export function instanceOfRoomOutOfOrderType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomOutOfOrderTypeFromJSON(json: any): RoomOutOfOrderType {
    return RoomOutOfOrderTypeFromJSONTyped(json, false);
}

export function RoomOutOfOrderTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomOutOfOrderType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'repairRemarks': !exists(json, 'repairRemarks') ? undefined : json['repairRemarks'],
        'returnStatus': !exists(json, 'returnStatus') ? undefined : HousekeepingRoomStatusTypeFromJSON(json['returnStatus']),
        'reasonCode': !exists(json, 'reasonCode') ? undefined : json['reasonCode'],
        'reasonDesc': !exists(json, 'reasonDesc') ? undefined : json['reasonDesc'],
        'newDateRange': !exists(json, 'newDateRange') ? undefined : DateRangeTypeFromJSON(json['newDateRange']),
        'roomStatus': !exists(json, 'roomStatus') ? undefined : HousekeepingRoomStatusTypeFromJSON(json['roomStatus']),
        'start': !exists(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !exists(json, 'end') ? undefined : (new Date(json['end'])),
    };
}

export function RoomOutOfOrderTypeToJSON(value?: RoomOutOfOrderType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'repairRemarks': value.repairRemarks,
        'returnStatus': HousekeepingRoomStatusTypeToJSON(value.returnStatus),
        'reasonCode': value.reasonCode,
        'reasonDesc': value.reasonDesc,
        'newDateRange': DateRangeTypeToJSON(value.newDateRange),
        'roomStatus': HousekeepingRoomStatusTypeToJSON(value.roomStatus),
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0,10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0,10)),
    };
}

