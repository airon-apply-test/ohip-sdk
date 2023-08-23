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
 * This element has revenue amount data for its revenue category such as Room Revenue, Food and Beverage Revenue.
 * @export
 * @interface RevenueCategorySummaryType
 */
export interface RevenueCategorySummaryType {
    /**
     * The representation of a revenue category.
     * @type {string}
     * @memberof RevenueCategorySummaryType
     */
    revenueCategoryCode?: string;
    /**
     * A monetary amount.
     * @type {number}
     * @memberof RevenueCategorySummaryType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof RevenueCategorySummaryType
     */
    currencyCode?: string;
}

/**
 * Check if a given object implements the RevenueCategorySummaryType interface.
 */
export function instanceOfRevenueCategorySummaryType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RevenueCategorySummaryTypeFromJSON(json: any): RevenueCategorySummaryType {
    return RevenueCategorySummaryTypeFromJSONTyped(json, false);
}

export function RevenueCategorySummaryTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RevenueCategorySummaryType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'revenueCategoryCode': !exists(json, 'revenueCategoryCode') ? undefined : json['revenueCategoryCode'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
    };
}

export function RevenueCategorySummaryTypeToJSON(value?: RevenueCategorySummaryType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'revenueCategoryCode': value.revenueCategoryCode,
        'amount': value.amount,
        'currencyCode': value.currencyCode,
    };
}

