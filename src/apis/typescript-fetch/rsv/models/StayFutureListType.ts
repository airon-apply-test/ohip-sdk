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
import type { StayReservationInfoType } from './StayReservationInfoType';
import {
    StayReservationInfoTypeFromJSON,
    StayReservationInfoTypeFromJSONTyped,
    StayReservationInfoTypeToJSON,
} from './StayReservationInfoType';

/**
 * A collection of reservation history details attached to Profiles.
 * @export
 * @interface StayFutureListType
 */
export interface StayFutureListType {
    /**
     * Additional reservation information attached to the profile . Eg : Future reservation details
     * @type {Array<StayReservationInfoType>}
     * @memberof StayFutureListType
     */
    reservationInfo?: Array<StayReservationInfoType>;
    /**
     * Indicates whether all the records are included in the response or not. Absence of the attribute values should be consider as all rows fetched in the response.
     * @type {boolean}
     * @memberof StayFutureListType
     */
    hasMore?: boolean;
    /**
     * Total number of rows queried
     * @type {number}
     * @memberof StayFutureListType
     */
    totalResults?: number;
    /**
     * Total number of rows returned
     * @type {number}
     * @memberof StayFutureListType
     */
    count?: number;
}

/**
 * Check if a given object implements the StayFutureListType interface.
 */
export function instanceOfStayFutureListType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function StayFutureListTypeFromJSON(json: any): StayFutureListType {
    return StayFutureListTypeFromJSONTyped(json, false);
}

export function StayFutureListTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): StayFutureListType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'reservationInfo': !exists(json, 'reservationInfo') ? undefined : ((json['reservationInfo'] as Array<any>).map(StayReservationInfoTypeFromJSON)),
        'hasMore': !exists(json, 'hasMore') ? undefined : json['hasMore'],
        'totalResults': !exists(json, 'totalResults') ? undefined : json['totalResults'],
        'count': !exists(json, 'count') ? undefined : json['count'],
    };
}

export function StayFutureListTypeToJSON(value?: StayFutureListType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'reservationInfo': value.reservationInfo === undefined ? undefined : ((value.reservationInfo as Array<any>).map(StayReservationInfoTypeToJSON)),
        'hasMore': value.hasMore,
        'totalResults': value.totalResults,
        'count': value.count,
    };
}

