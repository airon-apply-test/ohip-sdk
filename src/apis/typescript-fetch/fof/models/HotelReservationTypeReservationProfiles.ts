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
import type { CommissionPayoutToType } from './CommissionPayoutToType';
import {
    CommissionPayoutToTypeFromJSON,
    CommissionPayoutToTypeFromJSONTyped,
    CommissionPayoutToTypeToJSON,
} from './CommissionPayoutToType';
import type { ReservationProfileType } from './ReservationProfileType';
import {
    ReservationProfileTypeFromJSON,
    ReservationProfileTypeFromJSONTyped,
    ReservationProfileTypeToJSON,
} from './ReservationProfileType';

/**
 * Collection of guests associated with the reservation.
 * @export
 * @interface HotelReservationTypeReservationProfiles
 */
export interface HotelReservationTypeReservationProfiles {
    /**
     * 
     * @type {Array<ReservationProfileType>}
     * @memberof HotelReservationTypeReservationProfiles
     */
    reservationProfile?: Array<ReservationProfileType>;
    /**
     * 
     * @type {CommissionPayoutToType}
     * @memberof HotelReservationTypeReservationProfiles
     */
    commissionPayoutTo?: CommissionPayoutToType;
}

/**
 * Check if a given object implements the HotelReservationTypeReservationProfiles interface.
 */
export function instanceOfHotelReservationTypeReservationProfiles(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HotelReservationTypeReservationProfilesFromJSON(json: any): HotelReservationTypeReservationProfiles {
    return HotelReservationTypeReservationProfilesFromJSONTyped(json, false);
}

export function HotelReservationTypeReservationProfilesFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelReservationTypeReservationProfiles {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationProfile': !exists(json, 'reservationProfile') ? undefined : ((json['reservationProfile'] as Array<any>).map(ReservationProfileTypeFromJSON)),
        'commissionPayoutTo': !exists(json, 'commissionPayoutTo') ? undefined : CommissionPayoutToTypeFromJSON(json['commissionPayoutTo']),
    };
}

export function HotelReservationTypeReservationProfilesToJSON(value?: HotelReservationTypeReservationProfiles | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationProfile': value.reservationProfile === undefined ? undefined : ((value.reservationProfile as Array<any>).map(ReservationProfileTypeToJSON)),
        'commissionPayoutTo': CommissionPayoutToTypeToJSON(value.commissionPayoutTo),
    };
}

