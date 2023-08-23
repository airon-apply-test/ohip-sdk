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
import type { AmountDeterminationType } from './AmountDeterminationType';
/**
 * Applicable tax element. This element allows for both percentages and flat amounts. If one field is used, the other should be zero since logically, taxes should be calculated in only one of the two ways.
 * @export
 * @interface TaxType
 */
export interface TaxType {
    /**
     *
     * @type {string}
     * @memberof TaxType
     */
    description?: string;
    /**
     *
     * @type {AmountDeterminationType}
     * @memberof TaxType
     */
    type?: AmountDeterminationType;
    /**
     * Code identifying the fee (e.g.,agency fee, municipality fee).
     * @type {string}
     * @memberof TaxType
     */
    code?: string;
    /**
     * A monetary amount.
     * @type {number}
     * @memberof TaxType
     */
    amount?: number;
    /**
     * Provides a currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof TaxType
     */
    currencyCode?: string;
}
/**
 * Check if a given object implements the TaxType interface.
 */
export declare function instanceOfTaxType(value: object): boolean;
export declare function TaxTypeFromJSON(json: any): TaxType;
export declare function TaxTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TaxType;
export declare function TaxTypeToJSON(value?: TaxType | null): any;
