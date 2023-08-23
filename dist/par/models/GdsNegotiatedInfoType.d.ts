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
 * This provides information for a channel negotiated rate.
 * @export
 * @interface GdsNegotiatedInfoType
 */
export interface GdsNegotiatedInfoType {
    /**
     * The GDS negotiated rate code.
     * @type {string}
     * @memberof GdsNegotiatedInfoType
     */
    accessCode?: string;
    /**
     * The sell order.
     * @type {number}
     * @memberof GdsNegotiatedInfoType
     */
    order?: number;
    /**
     * The GDS Negotiated Rate is inactive or not
     * @type {boolean}
     * @memberof GdsNegotiatedInfoType
     */
    inactive?: boolean;
    /**
     * The starting value of the date range.
     * @type {Date}
     * @memberof GdsNegotiatedInfoType
     */
    start?: Date;
    /**
     * The ending value of the date range.
     * @type {Date}
     * @memberof GdsNegotiatedInfoType
     */
    end?: Date;
}
/**
 * Check if a given object implements the GdsNegotiatedInfoType interface.
 */
export declare function instanceOfGdsNegotiatedInfoType(value: object): boolean;
export declare function GdsNegotiatedInfoTypeFromJSON(json: any): GdsNegotiatedInfoType;
export declare function GdsNegotiatedInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): GdsNegotiatedInfoType;
export declare function GdsNegotiatedInfoTypeToJSON(value?: GdsNegotiatedInfoType | null): any;
