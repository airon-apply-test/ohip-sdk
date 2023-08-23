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
import type { UpsellInfoTypeOriginalInfo } from './UpsellInfoTypeOriginalInfo';
import type { UpsellInfoTypeUpsellInfo } from './UpsellInfoTypeUpsellInfo';
/**
 * Information regarding upsell for a reservation.
 * @export
 * @interface UpsellInfoType
 */
export interface UpsellInfoType {
    /**
     *
     * @type {UpsellInfoTypeOriginalInfo}
     * @memberof UpsellInfoType
     */
    originalInfo?: UpsellInfoTypeOriginalInfo;
    /**
     *
     * @type {UpsellInfoTypeUpsellInfo}
     * @memberof UpsellInfoType
     */
    upsellInfo?: UpsellInfoTypeUpsellInfo;
}
/**
 * Check if a given object implements the UpsellInfoType interface.
 */
export declare function instanceOfUpsellInfoType(value: object): boolean;
export declare function UpsellInfoTypeFromJSON(json: any): UpsellInfoType;
export declare function UpsellInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpsellInfoType;
export declare function UpsellInfoTypeToJSON(value?: UpsellInfoType | null): any;
