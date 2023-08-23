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
/**
 * A monetary value expressed with a currency code.
 * @export
 * @interface CurrencyAmountType
 */
export interface CurrencyAmountType {
    /**
     * A monetary amount.
     * @type {number}
     * @memberof CurrencyAmountType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof CurrencyAmountType
     */
    currencyCode?: string;
}
/**
 * Check if a given object implements the CurrencyAmountType interface.
 */
export declare function instanceOfCurrencyAmountType(value: object): boolean;
export declare function CurrencyAmountTypeFromJSON(json: any): CurrencyAmountType;
export declare function CurrencyAmountTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CurrencyAmountType;
export declare function CurrencyAmountTypeToJSON(value?: CurrencyAmountType | null): any;
