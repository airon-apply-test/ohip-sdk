/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Reservation API
 * APIs to cater for Reservation functionality in OPERA Cloud. <br /><br />OPERA Cloud Reservations provides a complete set of capabilities for creating and updating reservations. Reservations are a central feature of OPERA Cloud. As a key source of information, the reservation specifies a guest&apos;s arrival date, departure date, room type, rate, packages, and many other details. It is also a gateway to dozens of other functions that contribute to the guest&apos;s experience.  All reservations in OPERA Cloud require a guest profile.<br /><br /> You can create profiles while booking a reservation. If a profile already exists, you can look it up (using getProfiles in the Customer Relationship Management module) and attach it to the reservation during the reservation booking process using the profileId.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { EffectiveRateType } from './EffectiveRateType';
import {
    EffectiveRateTypeFromJSON,
    EffectiveRateTypeFromJSONTyped,
    EffectiveRateTypeToJSON,
} from './EffectiveRateType';
import type { HotelReservationType } from './HotelReservationType';
import {
    HotelReservationTypeFromJSON,
    HotelReservationTypeFromJSONTyped,
    HotelReservationTypeToJSON,
} from './HotelReservationType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * Overriding the validation check will log the supposed errors in the warnings log if they would fail.
 * @export
 * @interface ReinstateReservation
 */
export interface ReinstateReservation {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof ReinstateReservation
     */
    hotelId?: string;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof ReinstateReservation
     */
    reservationIdList?: Array<UniqueIDType>;
    /**
     * 
     * @type {number}
     * @memberof ReinstateReservation
     */
    reservationLockHandle?: number;
    /**
     * 
     * @type {HotelReservationType}
     * @memberof ReinstateReservation
     */
    reservation?: HotelReservationType;
    /**
     * Flag that indicates if room inventory check should be skipped when the reservation is being reinstated.
     * @type {boolean}
     * @memberof ReinstateReservation
     */
    overrideInventory?: boolean;
    /**
     * Flag that indicates if rate code inventory check should be skipped when the reservation is being reinstated.
     * @type {boolean}
     * @memberof ReinstateReservation
     */
    overrideRates?: boolean;
    /**
     * Flag that indicates if the check on the housekeeping status for out of service should be skipped.
     * @type {boolean}
     * @memberof ReinstateReservation
     */
    overrideRoomOutOfService?: boolean;
    /**
     * Flag that indicates if the check on room allocation should be skipped.
     * @type {boolean}
     * @memberof ReinstateReservation
     */
    overrideRoomAllocation?: boolean;
    /**
     * Unique Id that references an object uniquely in the system.
     * @type {Array<UniqueIDType>}
     * @memberof ReinstateReservation
     */
    additionalReservationIdList?: Array<UniqueIDType>;
    /**
     * Collection of effective rate amount per guest on specific dates.
     * @type {Array<EffectiveRateType>}
     * @memberof ReinstateReservation
     */
    effectiveRates?: Array<EffectiveRateType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof ReinstateReservation
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof ReinstateReservation
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the ReinstateReservation interface.
 */
export function instanceOfReinstateReservation(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReinstateReservationFromJSON(json: any): ReinstateReservation {
    return ReinstateReservationFromJSONTyped(json, false);
}

export function ReinstateReservationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReinstateReservation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'reservationIdList': !exists(json, 'reservationIdList') ? undefined : ((json['reservationIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'reservationLockHandle': !exists(json, 'reservationLockHandle') ? undefined : json['reservationLockHandle'],
        'reservation': !exists(json, 'reservation') ? undefined : HotelReservationTypeFromJSON(json['reservation']),
        'overrideInventory': !exists(json, 'overrideInventory') ? undefined : json['overrideInventory'],
        'overrideRates': !exists(json, 'overrideRates') ? undefined : json['overrideRates'],
        'overrideRoomOutOfService': !exists(json, 'overrideRoomOutOfService') ? undefined : json['overrideRoomOutOfService'],
        'overrideRoomAllocation': !exists(json, 'overrideRoomAllocation') ? undefined : json['overrideRoomAllocation'],
        'additionalReservationIdList': !exists(json, 'additionalReservationIdList') ? undefined : ((json['additionalReservationIdList'] as Array<any>).map(UniqueIDTypeFromJSON)),
        'effectiveRates': !exists(json, 'effectiveRates') ? undefined : ((json['effectiveRates'] as Array<any>).map(EffectiveRateTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function ReinstateReservationToJSON(value?: ReinstateReservation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'reservationIdList': value.reservationIdList === undefined ? undefined : ((value.reservationIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'reservationLockHandle': value.reservationLockHandle,
        'reservation': HotelReservationTypeToJSON(value.reservation),
        'overrideInventory': value.overrideInventory,
        'overrideRates': value.overrideRates,
        'overrideRoomOutOfService': value.overrideRoomOutOfService,
        'overrideRoomAllocation': value.overrideRoomAllocation,
        'additionalReservationIdList': value.additionalReservationIdList === undefined ? undefined : ((value.additionalReservationIdList as Array<any>).map(UniqueIDTypeToJSON)),
        'effectiveRates': value.effectiveRates === undefined ? undefined : ((value.effectiveRates as Array<any>).map(EffectiveRateTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

