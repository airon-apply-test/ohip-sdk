/**
 * OPERA Cloud Front Desk Operations Service
 * APIs to cater for Front Desk Operations and Front Desk Statistic functionality in OPERA Cloud. <br /><br /> Front Desk features some of the most commonly used operations in OPERA Cloud, such as managing guest arrivals, managing in-house guests, and managing guest departures. Some additional tasks you can complete from the Front Desk menu are room searches, room assignments, and quick check outs as well as opening folios, creating registration cards, setting wake up calls, and sending messages to guests.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { CodeDescriptionType } from './CodeDescriptionType';
import type { CurrencyAmountType } from './CurrencyAmountType';
/**
 * Original Values of the reservation prior to upgrade.
 * @export
 * @interface UpsellInfoTypeOriginalInfo
 */
export interface UpsellInfoTypeOriginalInfo {
    /**
     * The Rate Code of the Reservation before it was upgraded
     * @type {string}
     * @memberof UpsellInfoTypeOriginalInfo
     */
    rateCode?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof UpsellInfoTypeOriginalInfo
     */
    totalAmount?: CurrencyAmountType;
    /**
     *
     * @type {CodeDescriptionType}
     * @memberof UpsellInfoTypeOriginalInfo
     */
    roomType?: CodeDescriptionType;
    /**
     * The number of nights of the reservation before being upgraded.
     * @type {number}
     * @memberof UpsellInfoTypeOriginalInfo
     */
    nights?: number;
}
/**
 * Check if a given object implements the UpsellInfoTypeOriginalInfo interface.
 */
export declare function instanceOfUpsellInfoTypeOriginalInfo(value: object): boolean;
export declare function UpsellInfoTypeOriginalInfoFromJSON(json: any): UpsellInfoTypeOriginalInfo;
export declare function UpsellInfoTypeOriginalInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpsellInfoTypeOriginalInfo;
export declare function UpsellInfoTypeOriginalInfoToJSON(value?: UpsellInfoTypeOriginalInfo | null): any;
