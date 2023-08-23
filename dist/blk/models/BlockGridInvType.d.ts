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
 * Indicates the inventory values ( occupancy or availability ) per person occupancy.
 * @export
 * @interface BlockGridInvType
 */
export interface BlockGridInvType {
    /**
     * Indicates the one person inventory value ( occupancy or availability ).
     * @type {number}
     * @memberof BlockGridInvType
     */
    onePerson?: number;
    /**
     * Indicates the two person inventory value ( occupancy or availability ).
     * @type {number}
     * @memberof BlockGridInvType
     */
    twoPerson?: number;
    /**
     * Indicates the three person inventory value ( occupancy or availability ).
     * @type {number}
     * @memberof BlockGridInvType
     */
    threePerson?: number;
    /**
     * Indicates the four person inventory value ( occupancy or availability ).
     * @type {number}
     * @memberof BlockGridInvType
     */
    fourPerson?: number;
    /**
     * Indicates the sell limit
     * @type {number}
     * @memberof BlockGridInvType
     */
    sellLimit?: number;
    /**
     * Indicates the cutoff date.Date when inventory left in the block will be cut-off.
     * @type {Date}
     * @memberof BlockGridInvType
     */
    cutoffDate?: Date;
}
/**
 * Check if a given object implements the BlockGridInvType interface.
 */
export declare function instanceOfBlockGridInvType(value: object): boolean;
export declare function BlockGridInvTypeFromJSON(json: any): BlockGridInvType;
export declare function BlockGridInvTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockGridInvType;
export declare function BlockGridInvTypeToJSON(value?: BlockGridInvType | null): any;
