/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Enterprise Configuration API
 * APIs to cater for Enterprise Configuration functionality in OPERA Cloud. <br /><br In this module, you can configure a variety of options related to your properties such as their locations, facilities, and local attractions. The available options are dependant on the active controls at your property.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
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
 * Catering Currency Formatting information configuration of the hotel.
 * @export
 * @interface HotelInfoTypePropertyControlsCateringCurrencyFormatting
 */
export interface HotelInfoTypePropertyControlsCateringCurrencyFormatting {
    /**
     * The base currency code for this hotel..
     * @type {string}
     * @memberof HotelInfoTypePropertyControlsCateringCurrencyFormatting
     */
    currencyCode?: string;
    /**
     * Format for the local currency.
     * @type {string}
     * @memberof HotelInfoTypePropertyControlsCateringCurrencyFormatting
     */
    currencyFormat?: string;
}

/**
 * Check if a given object implements the HotelInfoTypePropertyControlsCateringCurrencyFormatting interface.
 */
export function instanceOfHotelInfoTypePropertyControlsCateringCurrencyFormatting(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function HotelInfoTypePropertyControlsCateringCurrencyFormattingFromJSON(json: any): HotelInfoTypePropertyControlsCateringCurrencyFormatting {
    return HotelInfoTypePropertyControlsCateringCurrencyFormattingFromJSONTyped(json, false);
}

export function HotelInfoTypePropertyControlsCateringCurrencyFormattingFromJSONTyped(json: any, ignoreDiscriminator: boolean): HotelInfoTypePropertyControlsCateringCurrencyFormatting {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencyFormat': !exists(json, 'currencyFormat') ? undefined : json['currencyFormat'],
    };
}

export function HotelInfoTypePropertyControlsCateringCurrencyFormattingToJSON(value?: HotelInfoTypePropertyControlsCateringCurrencyFormatting | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'currencyCode': value.currencyCode,
        'currencyFormat': value.currencyFormat,
    };
}

