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
import type { CurrencyAmountType } from './CurrencyAmountType';
/**
 * Contains last stay related details for the profile
 * @export
 * @interface LastStayInfoType
 */
export interface LastStayInfoType {
    /**
     * Used to hold last stay information for the profile.
     * @type {Date}
     * @memberof LastStayInfoType
     */
    lastVisit?: Date;
    /**
     * Used to hold last room information for the profile.
     * @type {string}
     * @memberof LastStayInfoType
     */
    lastRoom?: string;
    /**
     *
     * @type {CurrencyAmountType}
     * @memberof LastStayInfoType
     */
    lastRate?: CurrencyAmountType;
    /**
     * The total number of previous stay of the profile.
     * @type {number}
     * @memberof LastStayInfoType
     */
    totalStay?: number;
}
/**
 * Check if a given object implements the LastStayInfoType interface.
 */
export declare function instanceOfLastStayInfoType(value: object): boolean;
export declare function LastStayInfoTypeFromJSON(json: any): LastStayInfoType;
export declare function LastStayInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): LastStayInfoType;
export declare function LastStayInfoTypeToJSON(value?: LastStayInfoType | null): any;
