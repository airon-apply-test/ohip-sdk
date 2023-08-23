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
import type { CurrencyAmountType } from './CurrencyAmountType';
import {
    CurrencyAmountTypeFromJSON,
    CurrencyAmountTypeFromJSONTyped,
    CurrencyAmountTypeToJSON,
} from './CurrencyAmountType';

/**
 * Currency code configuration.
 * @export
 * @interface CurrencyExchangeRateType
 */
export interface CurrencyExchangeRateType {
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof CurrencyExchangeRateType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof CurrencyExchangeRateType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof CurrencyExchangeRateType
     */
    decimalPlaces?: number;
    /**
     * Description of the currency code.
     * @type {string}
     * @memberof CurrencyExchangeRateType
     */
    description?: string;
    /**
     * 
     * @type {CurrencyAmountType}
     * @memberof CurrencyExchangeRateType
     */
    exchangeRate?: CurrencyAmountType;
}

/**
 * Check if a given object implements the CurrencyExchangeRateType interface.
 */
export function instanceOfCurrencyExchangeRateType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function CurrencyExchangeRateTypeFromJSON(json: any): CurrencyExchangeRateType {
    return CurrencyExchangeRateTypeFromJSONTyped(json, false);
}

export function CurrencyExchangeRateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CurrencyExchangeRateType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'currencySymbol': !exists(json, 'currencySymbol') ? undefined : json['currencySymbol'],
        'decimalPlaces': !exists(json, 'decimalPlaces') ? undefined : json['decimalPlaces'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'exchangeRate': !exists(json, 'exchangeRate') ? undefined : CurrencyAmountTypeFromJSON(json['exchangeRate']),
    };
}

export function CurrencyExchangeRateTypeToJSON(value?: CurrencyExchangeRateType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'currencyCode': value.currencyCode,
        'currencySymbol': value.currencySymbol,
        'decimalPlaces': value.decimalPlaces,
        'description': value.description,
        'exchangeRate': CurrencyAmountTypeToJSON(value.exchangeRate),
    };
}

