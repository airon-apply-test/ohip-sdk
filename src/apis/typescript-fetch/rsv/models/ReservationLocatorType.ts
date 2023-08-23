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
import type { DateRangeType } from './DateRangeType';
import {
    DateRangeTypeFromJSON,
    DateRangeTypeFromJSONTyped,
    DateRangeTypeToJSON,
} from './DateRangeType';
import type { DateTimeSpanType } from './DateTimeSpanType';
import {
    DateTimeSpanTypeFromJSON,
    DateTimeSpanTypeFromJSONTyped,
    DateTimeSpanTypeToJSON,
} from './DateTimeSpanType';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';

/**
 * Holds the information for a Reservation Guest Locator
 * @export
 * @interface ReservationLocatorType
 */
export interface ReservationLocatorType {
    /**
     * 
     * @type {DateRangeType}
     * @memberof ReservationLocatorType
     */
    dateSpan?: DateRangeType;
    /**
     * 
     * @type {DateTimeSpanType}
     * @memberof ReservationLocatorType
     */
    timeSpan?: DateTimeSpanType;
    /**
     * The Locator Text for the guest.
     * @type {string}
     * @memberof ReservationLocatorType
     */
    locatorText?: string;
    /**
     * Date and time of the Guest Locator.
     * @type {Date}
     * @memberof ReservationLocatorType
     */
    locatorOn?: Date;
    /**
     * User that entered this Guest Locator.
     * @type {string}
     * @memberof ReservationLocatorType
     */
    locatorBy?: string;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof ReservationLocatorType
     */
    locatorId?: UniqueIDType;
}

/**
 * Check if a given object implements the ReservationLocatorType interface.
 */
export function instanceOfReservationLocatorType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationLocatorTypeFromJSON(json: any): ReservationLocatorType {
    return ReservationLocatorTypeFromJSONTyped(json, false);
}

export function ReservationLocatorTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationLocatorType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dateSpan': !exists(json, 'dateSpan') ? undefined : DateRangeTypeFromJSON(json['dateSpan']),
        'timeSpan': !exists(json, 'timeSpan') ? undefined : DateTimeSpanTypeFromJSON(json['timeSpan']),
        'locatorText': !exists(json, 'locatorText') ? undefined : json['locatorText'],
        'locatorOn': !exists(json, 'locatorOn') ? undefined : (new Date(json['locatorOn'])),
        'locatorBy': !exists(json, 'locatorBy') ? undefined : json['locatorBy'],
        'locatorId': !exists(json, 'locatorId') ? undefined : UniqueIDTypeFromJSON(json['locatorId']),
    };
}

export function ReservationLocatorTypeToJSON(value?: ReservationLocatorType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dateSpan': DateRangeTypeToJSON(value.dateSpan),
        'timeSpan': DateTimeSpanTypeToJSON(value.timeSpan),
        'locatorText': value.locatorText,
        'locatorOn': value.locatorOn === undefined ? undefined : (value.locatorOn.toISOString()),
        'locatorBy': value.locatorBy,
        'locatorId': UniqueIDTypeToJSON(value.locatorId),
    };
}

