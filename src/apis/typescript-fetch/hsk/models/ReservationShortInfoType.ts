/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Housekeeping Service API
 * APIs to cater for Housekeeping functionality in OPERA Cloud. <br /><br />Housekeeping enables you to schedule daily room cleaning, maintenance, and housekeeping staff activities. It provides information on room status, out of order/out of service rooms, and forecasting.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { HousekeepingRoomReservationStatusType } from './HousekeepingRoomReservationStatusType';
import {
    HousekeepingRoomReservationStatusTypeFromJSON,
    HousekeepingRoomReservationStatusTypeFromJSONTyped,
    HousekeepingRoomReservationStatusTypeToJSON,
} from './HousekeepingRoomReservationStatusType';
import type { ReservationId } from './ReservationId';
import {
    ReservationIdFromJSON,
    ReservationIdFromJSONTyped,
    ReservationIdToJSON,
} from './ReservationId';
import type { StayDateInfoType } from './StayDateInfoType';
import {
    StayDateInfoTypeFromJSON,
    StayDateInfoTypeFromJSONTyped,
    StayDateInfoTypeToJSON,
} from './StayDateInfoType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';
import type { VIPStatusType } from './VIPStatusType';
import {
    VIPStatusTypeFromJSON,
    VIPStatusTypeFromJSONTyped,
    VIPStatusTypeToJSON,
} from './VIPStatusType';

/**
 * 
 * @export
 * @interface ReservationShortInfoType
 */
export interface ReservationShortInfoType {
    /**
     * Guest name that is registered for the reservation.
     * @type {string}
     * @memberof ReservationShortInfoType
     */
    guestName?: string;
    /**
     * 
     * @type {ReservationId}
     * @memberof ReservationShortInfoType
     */
    reservationId?: ReservationId;
    /**
     * 
     * @type {HousekeepingRoomReservationStatusType}
     * @memberof ReservationShortInfoType
     */
    reservationStatusInfo?: HousekeepingRoomReservationStatusType;
    /**
     * 
     * @type {StayDateInfoType}
     * @memberof ReservationShortInfoType
     */
    stayDateInfo?: StayDateInfoType;
    /**
     * 
     * @type {VIPStatusType}
     * @memberof ReservationShortInfoType
     */
    vipStatus?: VIPStatusType;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ReservationShortInfoType
     */
    profileId?: UniqueIDType;
}

/**
 * Check if a given object implements the ReservationShortInfoType interface.
 */
export function instanceOfReservationShortInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationShortInfoTypeFromJSON(json: any): ReservationShortInfoType {
    return ReservationShortInfoTypeFromJSONTyped(json, false);
}

export function ReservationShortInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationShortInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'guestName': !exists(json, 'guestName') ? undefined : json['guestName'],
        'reservationId': !exists(json, 'reservationId') ? undefined : ReservationIdFromJSON(json['reservationId']),
        'reservationStatusInfo': !exists(json, 'reservationStatusInfo') ? undefined : HousekeepingRoomReservationStatusTypeFromJSON(json['reservationStatusInfo']),
        'stayDateInfo': !exists(json, 'stayDateInfo') ? undefined : StayDateInfoTypeFromJSON(json['stayDateInfo']),
        'vipStatus': !exists(json, 'vipStatus') ? undefined : VIPStatusTypeFromJSON(json['vipStatus']),
        'profileId': !exists(json, 'profileId') ? undefined : UniqueIDTypeFromJSON(json['profileId']),
    };
}

export function ReservationShortInfoTypeToJSON(value?: ReservationShortInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'guestName': value.guestName,
        'reservationId': ReservationIdToJSON(value.reservationId),
        'reservationStatusInfo': HousekeepingRoomReservationStatusTypeToJSON(value.reservationStatusInfo),
        'stayDateInfo': StayDateInfoTypeToJSON(value.stayDateInfo),
        'vipStatus': VIPStatusTypeToJSON(value.vipStatus),
        'profileId': UniqueIDTypeToJSON(value.profileId),
    };
}

