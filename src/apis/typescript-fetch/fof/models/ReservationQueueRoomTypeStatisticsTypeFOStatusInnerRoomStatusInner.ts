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
import type { HousekeepingRoomStatusType } from './HousekeepingRoomStatusType';
import {
    HousekeepingRoomStatusTypeFromJSON,
    HousekeepingRoomStatusTypeFromJSONTyped,
    HousekeepingRoomStatusTypeToJSON,
} from './HousekeepingRoomStatusType';

/**
 * 
 * @export
 * @interface ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner
 */
export interface ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner {
    /**
     * 
     * @type {HousekeepingRoomStatusType}
     * @memberof ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner
     */
    housekeepingRoomStatus?: HousekeepingRoomStatusType;
    /**
     * The total number of rooms for the room status.
     * @type {number}
     * @memberof ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner
     */
    totalRooms?: number;
}

/**
 * Check if a given object implements the ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner interface.
 */
export function instanceOfReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInnerFromJSON(json: any): ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner {
    return ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInnerFromJSONTyped(json, false);
}

export function ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'housekeepingRoomStatus': !exists(json, 'housekeepingRoomStatus') ? undefined : HousekeepingRoomStatusTypeFromJSON(json['housekeepingRoomStatus']),
        'totalRooms': !exists(json, 'totalRooms') ? undefined : json['totalRooms'],
    };
}

export function ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInnerToJSON(value?: ReservationQueueRoomTypeStatisticsTypeFOStatusInnerRoomStatusInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'housekeepingRoomStatus': HousekeepingRoomStatusTypeToJSON(value.housekeepingRoomStatus),
        'totalRooms': value.totalRooms,
    };
}

