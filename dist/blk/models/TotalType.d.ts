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
import type { FeesType } from './FeesType';
import type { TaxesType } from './TaxesType';
/**
 * The total amount charged for the service including additional amounts and fees.
 * @export
 * @interface TotalType
 */
export interface TotalType {
    /**
     *
     * @type {TaxesType}
     * @memberof TotalType
     */
    taxes?: TaxesType;
    /**
     *
     * @type {FeesType}
     * @memberof TotalType
     */
    fees?: FeesType;
    /**
     *
     * @type {string}
     * @memberof TotalType
     */
    description?: string;
    /**
     * The total amount not including any associated tax (e.g., sales tax, VAT, GST or any associated tax).
     * @type {number}
     * @memberof TotalType
     */
    amountBeforeTax?: number;
    /**
     * The total amount including all associated taxes (e.g., sales tax, VAT, GST or any associated tax).
     * @type {number}
     * @memberof TotalType
     */
    amountAfterTax?: number;
    /**
     * The code specifying a monetary unit. Use ISO 4217, three alpha code.
     * @type {string}
     * @memberof TotalType
     */
    currencyCode?: string;
    /**
     * The symbol for the currency, e.g, for currencyCode USD the symbol is $.
     * @type {string}
     * @memberof TotalType
     */
    currencySymbol?: string;
    /**
     * Indicates the number of decimal places for a particular currency. This is equivalent to the ISO 4217 standard "minor unit". Typically used when the amount provided includes the minor unit of currency without a decimal point (e.g., USD 8500 needs DecimalPlaces="2" to represent $85).
     * @type {number}
     * @memberof TotalType
     */
    decimalPlaces?: number;
    /**
     * Type of charge.
     * @type {string}
     * @memberof TotalType
     */
    code?: string;
    /**
     * When true indicates that the rate amount has been overridden.
     * @type {boolean}
     * @memberof TotalType
     */
    rateOverride?: boolean;
}
/**
 * Check if a given object implements the TotalType interface.
 */
export declare function instanceOfTotalType(value: object): boolean;
export declare function TotalTypeFromJSON(json: any): TotalType;
export declare function TotalTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): TotalType;
export declare function TotalTypeToJSON(value?: TotalType | null): any;
