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
 * Identifies and provides details about the discount. This allows for both percentages and flat amounts. If one field is used, the other should be zero/not specified since logically.
 * @export
 * @interface DiscountType
 */
export interface DiscountType {
    /**
     * 
     * @type {string}
     * @memberof DiscountType
     */
    discountReason?: string;
    /**
     * Percentage discount.
     * @type {number}
     * @memberof DiscountType
     */
    percent?: number;
    /**
     * A monetary amount.
     * @type {number}
     * @memberof DiscountType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof DiscountType
     */
    currencyCode?: string;
    /**
     * Specifies the type of discount (e.g., No condition, LOS, Deposit or Total amount spent).
     * @type {string}
     * @memberof DiscountType
     */
    discountCode?: string;
}

/**
 * Check if a given object implements the DiscountType interface.
 */
export function instanceOfDiscountType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function DiscountTypeFromJSON(json: any): DiscountType {
    return DiscountTypeFromJSONTyped(json, false);
}

export function DiscountTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DiscountType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'discountReason': !exists(json, 'discountReason') ? undefined : json['discountReason'],
        'percent': !exists(json, 'percent') ? undefined : json['percent'],
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'discountCode': !exists(json, 'discountCode') ? undefined : json['discountCode'],
    };
}

export function DiscountTypeToJSON(value?: DiscountType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'discountReason': value.discountReason,
        'percent': value.percent,
        'amount': value.amount,
        'currencyCode': value.currencyCode,
        'discountCode': value.discountCode,
    };
}

