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
import type { TaxType } from './TaxType';
import {
    TaxTypeFromJSON,
    TaxTypeFromJSONTyped,
    TaxTypeToJSON,
} from './TaxType';

/**
 * A collection of taxes.
 * @export
 * @interface TaxesType
 */
export interface TaxesType {
    /**
     * An individual tax.
     * @type {Array<TaxType>}
     * @memberof TaxesType
     */
    tax?: Array<TaxType>;
    /**
     * A monetary amount.
     * @type {number}
     * @memberof TaxesType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof TaxesType
     */
    currencyCode?: string;
}

/**
 * Check if a given object implements the TaxesType interface.
 */
export function instanceOfTaxesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function TaxesTypeFromJSON(json: any): TaxesType {
    return TaxesTypeFromJSONTyped(json, false);
}

export function TaxesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TaxesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'tax': !exists(json, 'tax') ? undefined : ((json['tax'] as Array<any>).map(TaxTypeFromJSON)),
        'amount': !exists(json, 'amount') ? undefined : json['amount'],
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
    };
}

export function TaxesTypeToJSON(value?: TaxesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'tax': value.tax === undefined ? undefined : ((value.tax as Array<any>).map(TaxTypeToJSON)),
        'amount': value.amount,
        'currencyCode': value.currencyCode,
    };
}

