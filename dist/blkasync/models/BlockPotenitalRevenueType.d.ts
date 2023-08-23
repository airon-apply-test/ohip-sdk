/**
 * OPERA Cloud Block Reservation Asynchronous API
 * APIs to cater Block Reservation related asynchronous functionality in OPERA.<br /><br /> Compatible with OPERA Cloud release 22.4.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.4.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
/**
 * Indicates the sum of revenue breakup per room type per block per allocation date for the reservations. Applicable for business date or future dates.
 * @export
 * @interface BlockPotenitalRevenueType
 */
export interface BlockPotenitalRevenueType {
    /**
     * Indicates room revenue amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    roomRevenue?: number;
    /**
     * Indicates food revenue amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    foodRevenue?: number;
    /**
     * Indicates other revenue amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    otherRevenue?: number;
    /**
     * Indicates non revenue amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    nonRevenue?: number;
    /**
     * Indicates total revenue amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    totalRevenue?: number;
    /**
     * Indicates room revenue tax amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    roomRevenueTax?: number;
    /**
     * Indicates food revenue tax amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    foodRevenueTax?: number;
    /**
     * Indicates other revenue tax amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    otherRevenueTax?: number;
    /**
     * Indicates non revenue tax amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    nonRevenueTax?: number;
    /**
     * Indicates total revenue tax amount.
     * @type {number}
     * @memberof BlockPotenitalRevenueType
     */
    totalRevenueTax?: number;
    /**
     * Currency code to reflect the currency in which an amount may be expressed.
     * @type {string}
     * @memberof BlockPotenitalRevenueType
     */
    currency?: string;
}
/**
 * Check if a given object implements the BlockPotenitalRevenueType interface.
 */
export declare function instanceOfBlockPotenitalRevenueType(value: object): boolean;
export declare function BlockPotenitalRevenueTypeFromJSON(json: any): BlockPotenitalRevenueType;
export declare function BlockPotenitalRevenueTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockPotenitalRevenueType;
export declare function BlockPotenitalRevenueTypeToJSON(value?: BlockPotenitalRevenueType | null): any;
