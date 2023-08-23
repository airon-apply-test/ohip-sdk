/**
 * OPERA Cloud Accounts Receivables API
 * APIs to cater for Accounts Receivables functionality in OPERA Cloud. <br /><br The OPERA Cloud Accounts Receivable module enables you to manage debtors’ accounts, invoices, and remittance.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Lamp indicator Type.
 * @export
 * @interface IndicatorType
 */
export interface IndicatorType {
    /**
     * Name of the indicator.
     * @type {string}
     * @memberof IndicatorType
     */
    indicatorName?: string;
    /**
     * Indicates number of occurrences of the indicator.
     * @type {number}
     * @memberof IndicatorType
     */
    count?: number;
}
/**
 * Check if a given object implements the IndicatorType interface.
 */
export declare function instanceOfIndicatorType(value: object): boolean;
export declare function IndicatorTypeFromJSON(json: any): IndicatorType;
export declare function IndicatorTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): IndicatorType;
export declare function IndicatorTypeToJSON(value?: IndicatorType | null): any;
