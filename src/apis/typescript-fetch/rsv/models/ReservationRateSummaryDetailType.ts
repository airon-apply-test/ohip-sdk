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
/**
 * 
 * @export
 * @interface ReservationRateSummaryDetailType
 */
export interface ReservationRateSummaryDetailType {
    /**
     * Stay date for which revenue calculation is done.
     * @type {Date}
     * @memberof ReservationRateSummaryDetailType
     */
    summaryDate?: Date;
    /**
     * Base Amount of rate.
     * @type {number}
     * @memberof ReservationRateSummaryDetailType
     */
    revenue?: number;
    /**
     * Amount of an additional product or service that is sold as part of the rate or in addition to the rate.
     * @type {number}
     * @memberof ReservationRateSummaryDetailType
     */
    _package?: number;
    /**
     * Amount of tax generated separately on the revenue and package.
     * @type {number}
     * @memberof ReservationRateSummaryDetailType
     */
    tax?: number;
    /**
     * Amount of revenue and package excluding tax.
     * @type {number}
     * @memberof ReservationRateSummaryDetailType
     */
    gross?: number;
    /**
     * Amount of revenue and package including tax.
     * @type {number}
     * @memberof ReservationRateSummaryDetailType
     */
    net?: number;
    /**
     * Rate code calculation is based on.
     * @type {string}
     * @memberof ReservationRateSummaryDetailType
     */
    ratePlanCode?: string;
    /**
     * Date revenue calculation is based on.
     * @type {Date}
     * @memberof ReservationRateSummaryDetailType
     */
    revenueSimulationDate?: Date;
    /**
     * Currency revenue calculation is based on.
     * @type {string}
     * @memberof ReservationRateSummaryDetailType
     */
    currencyCode?: string;
    /**
     * Indicates if the rate is to be hidden.
     * @type {boolean}
     * @memberof ReservationRateSummaryDetailType
     */
    rateSuppressed?: boolean;
}

/**
 * Check if a given object implements the ReservationRateSummaryDetailType interface.
 */
export function instanceOfReservationRateSummaryDetailType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ReservationRateSummaryDetailTypeFromJSON(json: any): ReservationRateSummaryDetailType {
    return ReservationRateSummaryDetailTypeFromJSONTyped(json, false);
}

export function ReservationRateSummaryDetailTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReservationRateSummaryDetailType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'summaryDate': !exists(json, 'summaryDate') ? undefined : (new Date(json['summaryDate'])),
        'revenue': !exists(json, 'revenue') ? undefined : json['revenue'],
        '_package': !exists(json, 'package') ? undefined : json['package'],
        'tax': !exists(json, 'tax') ? undefined : json['tax'],
        'gross': !exists(json, 'gross') ? undefined : json['gross'],
        'net': !exists(json, 'net') ? undefined : json['net'],
        'ratePlanCode': !exists(json, 'ratePlanCode') ? undefined : json['ratePlanCode'],
        'revenueSimulationDate': !exists(json, 'revenueSimulationDate') ? undefined : (new Date(json['revenueSimulationDate'])),
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'rateSuppressed': !exists(json, 'rateSuppressed') ? undefined : json['rateSuppressed'],
    };
}

export function ReservationRateSummaryDetailTypeToJSON(value?: ReservationRateSummaryDetailType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'summaryDate': value.summaryDate === undefined ? undefined : (value.summaryDate.toISOString().substr(0,10)),
        'revenue': value.revenue,
        'package': value._package,
        'tax': value.tax,
        'gross': value.gross,
        'net': value.net,
        'ratePlanCode': value.ratePlanCode,
        'revenueSimulationDate': value.revenueSimulationDate === undefined ? undefined : (value.revenueSimulationDate.toISOString().substr(0,10)),
        'currencyCode': value.currencyCode,
        'rateSuppressed': value.rateSuppressed,
    };
}

