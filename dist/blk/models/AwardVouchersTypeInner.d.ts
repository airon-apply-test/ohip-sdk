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
 *
 * @export
 * @interface AwardVouchersTypeInner
 */
export interface AwardVouchersTypeInner {
    /**
     * Membership Award code applied on the reservation.
     * @type {string}
     * @memberof AwardVouchersTypeInner
     */
    awardCode?: string;
    /**
     * Membership Award number.
     * @type {string}
     * @memberof AwardVouchersTypeInner
     */
    voucherNo?: string;
}
/**
 * Check if a given object implements the AwardVouchersTypeInner interface.
 */
export declare function instanceOfAwardVouchersTypeInner(value: object): boolean;
export declare function AwardVouchersTypeInnerFromJSON(json: any): AwardVouchersTypeInner;
export declare function AwardVouchersTypeInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): AwardVouchersTypeInner;
export declare function AwardVouchersTypeInnerToJSON(value?: AwardVouchersTypeInner | null): any;
