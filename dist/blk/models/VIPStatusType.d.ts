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
 * VIP status of the customer.
 * @export
 * @interface VIPStatusType
 */
export interface VIPStatusType {
    /**
     * Used for Character Strings, length 0 to 2000.
     * @type {string}
     * @memberof VIPStatusType
     */
    value?: string;
    /**
     * VIP status of the customer.
     * @type {string}
     * @memberof VIPStatusType
     */
    code?: string;
}
/**
 * Check if a given object implements the VIPStatusType interface.
 */
export declare function instanceOfVIPStatusType(value: object): boolean;
export declare function VIPStatusTypeFromJSON(json: any): VIPStatusType;
export declare function VIPStatusTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): VIPStatusType;
export declare function VIPStatusTypeToJSON(value?: VIPStatusType | null): any;
