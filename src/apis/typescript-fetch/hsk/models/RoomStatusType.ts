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
import type { FrontOfficeRoomStatusType } from './FrontOfficeRoomStatusType';
import {
    FrontOfficeRoomStatusTypeFromJSON,
    FrontOfficeRoomStatusTypeFromJSONTyped,
    FrontOfficeRoomStatusTypeToJSON,
} from './FrontOfficeRoomStatusType';
import type { HousekeepingRoomReservationStatusType } from './HousekeepingRoomReservationStatusType';
import {
    HousekeepingRoomReservationStatusTypeFromJSON,
    HousekeepingRoomReservationStatusTypeFromJSONTyped,
    HousekeepingRoomReservationStatusTypeToJSON,
} from './HousekeepingRoomReservationStatusType';
import type { HousekeepingRoomStatusType } from './HousekeepingRoomStatusType';
import {
    HousekeepingRoomStatusTypeFromJSON,
    HousekeepingRoomStatusTypeFromJSONTyped,
    HousekeepingRoomStatusTypeToJSON,
} from './HousekeepingRoomStatusType';

/**
 * 
 * @export
 * @interface RoomStatusType
 */
export interface RoomStatusType {
    /**
     * List of status of the reservation to which this Room is assigned..
     * @type {Array<HousekeepingRoomReservationStatusType>}
     * @memberof RoomStatusType
     */
    reservationStatusList?: Array<HousekeepingRoomReservationStatusType>;
    /**
     * 
     * @type {HousekeepingRoomStatusType}
     * @memberof RoomStatusType
     */
    housekeepingRoomStatus?: HousekeepingRoomStatusType;
    /**
     * 
     * @type {FrontOfficeRoomStatusType}
     * @memberof RoomStatusType
     */
    frontOfficeStatus?: FrontOfficeRoomStatusType;
    /**
     * 
     * @type {FrontOfficeRoomStatusType}
     * @memberof RoomStatusType
     */
    housekeepingStatus?: FrontOfficeRoomStatusType;
}

/**
 * Check if a given object implements the RoomStatusType interface.
 */
export function instanceOfRoomStatusType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomStatusTypeFromJSON(json: any): RoomStatusType {
    return RoomStatusTypeFromJSONTyped(json, false);
}

export function RoomStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomStatusType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationStatusList': !exists(json, 'reservationStatusList') ? undefined : ((json['reservationStatusList'] as Array<any>).map(HousekeepingRoomReservationStatusTypeFromJSON)),
        'housekeepingRoomStatus': !exists(json, 'housekeepingRoomStatus') ? undefined : HousekeepingRoomStatusTypeFromJSON(json['housekeepingRoomStatus']),
        'frontOfficeStatus': !exists(json, 'frontOfficeStatus') ? undefined : FrontOfficeRoomStatusTypeFromJSON(json['frontOfficeStatus']),
        'housekeepingStatus': !exists(json, 'housekeepingStatus') ? undefined : FrontOfficeRoomStatusTypeFromJSON(json['housekeepingStatus']),
    };
}

export function RoomStatusTypeToJSON(value?: RoomStatusType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationStatusList': value.reservationStatusList === undefined ? undefined : ((value.reservationStatusList as Array<any>).map(HousekeepingRoomReservationStatusTypeToJSON)),
        'housekeepingRoomStatus': HousekeepingRoomStatusTypeToJSON(value.housekeepingRoomStatus),
        'frontOfficeStatus': FrontOfficeRoomStatusTypeToJSON(value.frontOfficeStatus),
        'housekeepingStatus': FrontOfficeRoomStatusTypeToJSON(value.housekeepingStatus),
    };
}

