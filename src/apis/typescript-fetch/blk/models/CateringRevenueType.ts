/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Block API
 * APIs to cater for Business Block functionality in OPERA Cloud. <br /><br /> A block is a group of rooms held for guests who are attending an event, meeting, or function. You can create blocks for family reunions, business conferences, weddings, and so on. You can also set aside rooms for the event (block).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';

/**
 * Information about Catering Revenue.
 * @export
 * @interface CateringRevenueType
 */
export interface CateringRevenueType {
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CateringRevenueType
     */
    serviceCharge?: CurrencyAmountType;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CateringRevenueType
     */
    taxAmount?: CurrencyAmountType;
    /**
     * Item discount for the block.
     * @type {number}
     * @memberof CateringRevenueType
     */
    itemDiscount?: number;
    /**
     * The discount percentage to be applied to resource items associated with catering events.
     * @type {number}
     * @memberof CateringRevenueType
     */
    itemDiscountPercentage?: number;
}

/**
 * Check if a given object implements the CateringRevenueType interface.
 */
export function instanceOfCateringRevenueType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CateringRevenueTypeFromJSON(json: any): CateringRevenueType {
    return CateringRevenueTypeFromJSONTyped(json, false);
}

export function CateringRevenueTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CateringRevenueType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'serviceCharge': !exists(json, 'serviceCharge') ? undefined : CurrencyAmountTypeFromJSON(json['serviceCharge']),
        'taxAmount': !exists(json, 'taxAmount') ? undefined : CurrencyAmountTypeFromJSON(json['taxAmount']),
        'itemDiscount': !exists(json, 'itemDiscount') ? undefined : json['itemDiscount'],
        'itemDiscountPercentage': !exists(json, 'itemDiscountPercentage') ? undefined : json['itemDiscountPercentage'],
    };
}

export function CateringRevenueTypeToJSON(value?: CateringRevenueType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'serviceCharge': CurrencyAmountTypeToJSON(value.serviceCharge),
        'taxAmount': CurrencyAmountTypeToJSON(value.taxAmount),
        'itemDiscount': value.itemDiscount,
        'itemDiscountPercentage': value.itemDiscountPercentage,
    };
}

