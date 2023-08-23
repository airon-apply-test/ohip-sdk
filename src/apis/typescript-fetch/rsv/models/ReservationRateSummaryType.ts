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
import type { ReservationRateSummaryDetailType } from './ReservationRateSummaryDetailType';
import {
    ReservationRateSummaryDetailTypeFromJSON,
    ReservationRateSummaryDetailTypeFromJSONTyped,
    ReservationRateSummaryDetailTypeToJSON,
} from './ReservationRateSummaryDetailType';

/**
 * 
 * @export
 * @interface ReservationRateSummaryType
 */
export interface ReservationRateSummaryType {
    /**
     * 
     * @type {Array<ReservationRateSummaryDetailType>}
     * @memberof ReservationRateSummaryType
     */
    details?: Array<ReservationRateSummaryDetailType>;
    /**
     * Amount of revenue and package excluding tax.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    gross?: number;
    /**
     * Amount of revenue and package including tax.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    net?: number;
    /**
     * Amount of automatically posted charges along with room and taxes.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    fixedCharges?: number;
    /**
     * Amount paid prior to the stay.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    deposit?: number;
    /**
     * Amount of total cost of stay.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    totalCostOfStay?: number;
    /**
     * Remaining amount to be paid.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    outStandingCostOfStay?: number;
    /**
     * Amount to be paid by Guest.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    guestPay?: number;
    /**
     * Amount of automatically forwarded charges for specific transaction codes.
     * @type {number}
     * @memberof ReservationRateSummaryType
     */
    routing?: number;
    /**
     * Currency revenue calculation is based on.
     * @type {string}
     * @memberof ReservationRateSummaryType
     */
    currencyCode?: string;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof ReservationRateSummaryType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof ReservationRateSummaryType
     */
    end?: Date;
    /**
     * Whether suppressed rate is present or not.
     * @type {boolean}
     * @memberof ReservationRateSummaryType
     */
    hasSuppressedRate?: boolean;
}

/**
 * Check if a given object implements the ReservationRateSummaryType interface.
 */
export function instanceOfReservationRateSummaryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationRateSummaryTypeFromJSON(json: any): ReservationRateSummaryType {
    return ReservationRateSummaryTypeFromJSONTyped(json, false);
}

export function ReservationRateSummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationRateSummaryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'details': !exists(json, 'details') ? undefined : ((json['details'] as Array<any>).map(ReservationRateSummaryDetailTypeFromJSON)),
        'gross': !exists(json, 'gross') ? undefined : json['gross'],
        'net': !exists(json, 'net') ? undefined : json['net'],
        'fixedCharges': !exists(json, 'fixedCharges') ? undefined : json['fixedCharges'],
        'deposit': !exists(json, 'deposit') ? undefined : json['deposit'],
        'totalCostOfStay': !exists(json, 'totalCostOfStay') ? undefined : json['totalCostOfStay'],
        'outStandingCostOfStay': !exists(json, 'outStandingCostOfStay') ? undefined : json['outStandingCostOfStay'],
        'guestPay': !exists(json, 'guestPay') ? undefined : json['guestPay'],
        'routing': !exists(json, 'routing') ? undefined : json['routing'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'start': !exists(json, 'start') ? undefined : (new Date(json['start'])),
        'end': !exists(json, 'end') ? undefined : (new Date(json['end'])),
        'hasSuppressedRate': !exists(json, 'hasSuppressedRate') ? undefined : json['hasSuppressedRate'],
    };
}

export function ReservationRateSummaryTypeToJSON(value?: ReservationRateSummaryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'details': value.details === undefined ? undefined : ((value.details as Array<any>).map(ReservationRateSummaryDetailTypeToJSON)),
        'gross': value.gross,
        'net': value.net,
        'fixedCharges': value.fixedCharges,
        'deposit': value.deposit,
        'totalCostOfStay': value.totalCostOfStay,
        'outStandingCostOfStay': value.outStandingCostOfStay,
        'guestPay': value.guestPay,
        'routing': value.routing,
        'currencyCode': value.currencyCode,
        'start': value.start === undefined ? undefined : (value.start.toISOString().substr(0,10)),
        'end': value.end === undefined ? undefined : (value.end.toISOString().substr(0,10)),
        'hasSuppressedRate': value.hasSuppressedRate,
    };
}

