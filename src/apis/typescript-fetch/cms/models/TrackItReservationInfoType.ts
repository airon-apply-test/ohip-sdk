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
import type { HousekeepingRoomStatusType } from './HousekeepingRoomStatusType';
import {
    HousekeepingRoomStatusTypeFromJSON,
    HousekeepingRoomStatusTypeFromJSONTyped,
    HousekeepingRoomStatusTypeToJSON,
} from './HousekeepingRoomStatusType';
import type { PMSResStatusType } from './PMSResStatusType';
import {
    PMSResStatusTypeFromJSON,
    PMSResStatusTypeFromJSONTyped,
    PMSResStatusTypeToJSON,
} from './PMSResStatusType';
import type { ResGuaranteeType } from './ResGuaranteeType';
import {
    ResGuaranteeTypeFromJSON,
    ResGuaranteeTypeFromJSONTyped,
    ResGuaranteeTypeToJSON,
} from './ResGuaranteeType';
import type { ResGuestInfoType } from './ResGuestInfoType';
import {
    ResGuestInfoTypeFromJSON,
    ResGuestInfoTypeFromJSONTyped,
    ResGuestInfoTypeToJSON,
} from './ResGuestInfoType';
import type { TimeSpanType } from './TimeSpanType';
import {
    TimeSpanTypeFromJSON,
    TimeSpanTypeFromJSONTyped,
    TimeSpanTypeToJSON,
} from './TimeSpanType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Reservation info associated with a Track It item.
 * @export
 * @interface TrackItReservationInfoType
 */
export interface TrackItReservationInfoType {
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof TrackItReservationInfoType
     */
    reservationIdList?: Array<UniqueIDType>;
    /**
     * 
     * @type {TimeSpanType}
     * @memberof TrackItReservationInfoType
     */
    timeSpan?: TimeSpanType;
    /**
     * 
     * @type {ResGuestInfoType}
     * @memberof TrackItReservationInfoType
     */
    guestInfo?: ResGuestInfoType;
    /**
     * 
     * @type {string}
     * @memberof TrackItReservationInfoType
     */
    roomType?: string;
    /**
     * 
     * @type {string}
     * @memberof TrackItReservationInfoType
     */
    roomId?: string;
    /**
     * 
     * @type {HousekeepingRoomStatusType}
     * @memberof TrackItReservationInfoType
     */
    roomStatus?: HousekeepingRoomStatusType;
    /**
     * 
     * @type {ResGuaranteeType}
     * @memberof TrackItReservationInfoType
     */
    guarantee?: ResGuaranteeType;
    /**
     * 
     * @type {PMSResStatusType}
     * @memberof TrackItReservationInfoType
     */
    reservationStatus?: PMSResStatusType;
    /**
     * 
     * @type {PMSResStatusType}
     * @memberof TrackItReservationInfoType
     */
    computedReservationStatus?: PMSResStatusType;
}

/**
 * Check if a given object implements the TrackItReservationInfoType interface.
 */
export function instanceOfTrackItReservationInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TrackItReservationInfoTypeFromJSON(json: any): TrackItReservationInfoType {
    return TrackItReservationInfoTypeFromJSONTyped(json, false);
}

export function TrackItReservationInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TrackItReservationInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationIdList': !exists(json, 'reservationIdList') ? undefined : ((json['reservationIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'timeSpan': !exists(json, 'timeSpan') ? undefined : TimeSpanTypeFromJSON(json['timeSpan']),
        'guestInfo': !exists(json, 'guestInfo') ? undefined : ResGuestInfoTypeFromJSON(json['guestInfo']),
        'roomType': !exists(json, 'roomType') ? undefined : json['roomType'],
        'roomId': !exists(json, 'roomId') ? undefined : json['roomId'],
        'roomStatus': !exists(json, 'roomStatus') ? undefined : HousekeepingRoomStatusTypeFromJSON(json['roomStatus']),
        'guarantee': !exists(json, 'guarantee') ? undefined : ResGuaranteeTypeFromJSON(json['guarantee']),
        'reservationStatus': !exists(json, 'reservationStatus') ? undefined : PMSResStatusTypeFromJSON(json['reservationStatus']),
        'computedReservationStatus': !exists(json, 'computedReservationStatus') ? undefined : PMSResStatusTypeFromJSON(json['computedReservationStatus']),
    };
}

export function TrackItReservationInfoTypeToJSON(value?: TrackItReservationInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationIdList': value.reservationIdList === undefined ? undefined : ((value.reservationIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'timeSpan': TimeSpanTypeToJSON(value.timeSpan),
        'guestInfo': ResGuestInfoTypeToJSON(value.guestInfo),
        'roomType': value.roomType,
        'roomId': value.roomId,
        'roomStatus': HousekeepingRoomStatusTypeToJSON(value.roomStatus),
        'guarantee': ResGuaranteeTypeToJSON(value.guarantee),
        'reservationStatus': PMSResStatusTypeToJSON(value.reservationStatus),
        'computedReservationStatus': PMSResStatusTypeToJSON(value.computedReservationStatus),
    };
}

