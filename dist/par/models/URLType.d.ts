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
 * Web site address, in IETF(The Internet Engineering Task Force) specified format.
 * @export
 * @interface URLType
 */
export interface URLType {
    /**
     * Property Value
     * @type {string}
     * @memberof URLType
     */
    value?: string;
    /**
     * Defines the purpose of the URL address, such as personal, business, public, etc.
     * @type {string}
     * @memberof URLType
     */
    type?: string;
    /**
     * Describes the Type code
     * @type {string}
     * @memberof URLType
     */
    typeDescription?: string;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof URLType
     */
    primaryInd?: boolean;
    /**
     * Display Order sequence.
     * @type {number}
     * @memberof URLType
     */
    orderSequence?: number;
}
/**
 * Check if a given object implements the URLType interface.
 */
export declare function instanceOfURLType(value: object): boolean;
export declare function URLTypeFromJSON(json: any): URLType;
export declare function URLTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): URLType;
export declare function URLTypeToJSON(value?: URLType | null): any;
