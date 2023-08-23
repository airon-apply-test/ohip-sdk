/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Price Availability Rate API
 * APIs to cater for Price and Rate Availability functionality in OPERA Cloud. <br /><br />Availability enables you to manage your room inventory by providing a detailed view of all available and sold rooms at a property. Some of the tasks you can perform include defining conditions for stay restrictions, setting room sell limits, and searching for and viewing room availability.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Currency Formatting information configuration of the hotel.
 * @export
 * @interface HotelInfoTypePropertyControlsCurrencyFormatting
 */
export interface HotelInfoTypePropertyControlsCurrencyFormatting {
    /**
     * The base currency code for this hotel..
     * @type {string}
     * @memberof HotelInfoTypePropertyControlsCurrencyFormatting
     */
    currencyCode?: string;
    /**
     * Format for the local currency.
     * @type {string}
     * @memberof HotelInfoTypePropertyControlsCurrencyFormatting
     */
    currencyFormat?: string;
    /**
     * Symbol to designate the default currency of the hotel.
     * @type {string}
     * @memberof HotelInfoTypePropertyControlsCurrencyFormatting
     */
    currencySymbol?: string;
    /**
     * Number of decimal positions used for this currency type.
     * @type {number}
     * @memberof HotelInfoTypePropertyControlsCurrencyFormatting
     */
    decimalPositions?: number;
}

/**
 * Check if a given object implements the HotelInfoTypePropertyControlsCurrencyFormatting interface.
 */
export function instanceOfHotelInfoTypePropertyControlsCurrencyFormatting(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HotelInfoTypePropertyControlsCurrencyFormattingFromJSON(json: any): HotelInfoTypePropertyControlsCurrencyFormatting {
    return HotelInfoTypePropertyControlsCurrencyFormattingFromJSONTyped(json, false);
}

export function HotelInfoTypePropertyControlsCurrencyFormattingFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInfoTypePropertyControlsCurrencyFormatting {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencyFormat': !exists(json, 'currencyFormat') ? undefined : json['currencyFormat'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPositions': !exists(json, 'decimalPositions') ? undefined : json['decimalPositions'],
    };
}

export function HotelInfoTypePropertyControlsCurrencyFormattingToJSON(value?: HotelInfoTypePropertyControlsCurrencyFormatting | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'currencyCode': value.currencyCode,
        'currencyFormat': value.currencyFormat,
        'currencySymbol': value.currencySymbol,
        'decimalPositions': value.decimalPositions,
    };
}

