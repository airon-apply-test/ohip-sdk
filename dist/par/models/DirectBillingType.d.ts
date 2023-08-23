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
/**
 * This is the preconfigured routing instruction type.
 * @export
 * @interface DirectBillingType
 */
export interface DirectBillingType {
    /**
     * Hotel Code for which the details of direct billing is provided for a profile.
     * @type {string}
     * @memberof DirectBillingType
     */
    hotelId?: string;
    /**
     * Hotel Code for which the details of direct billing is provided for a profile.
     * @type {string}
     * @memberof DirectBillingType
     */
    aRNumber?: string;
}
/**
 * Check if a given object implements the DirectBillingType interface.
 */
export declare function instanceOfDirectBillingType(value: object): boolean;
export declare function DirectBillingTypeFromJSON(json: any): DirectBillingType;
export declare function DirectBillingTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): DirectBillingType;
export declare function DirectBillingTypeToJSON(value?: DirectBillingType | null): any;
