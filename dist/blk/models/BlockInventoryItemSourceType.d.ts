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
import type { BlockId } from './BlockId';
/**
 * Defines whether the item is setup due to a Rate Plan, Package or a Block.
 * @export
 * @interface BlockInventoryItemSourceType
 */
export interface BlockInventoryItemSourceType {
    /**
     * Rate Plan Code, If populated specifies that the item is setup due to a Rate Plan.
     * @type {string}
     * @memberof BlockInventoryItemSourceType
     */
    ratePlanCode?: string;
    /**
     * Package Code, If populated specifies that the item is setup due to a Package.
     * @type {string}
     * @memberof BlockInventoryItemSourceType
     */
    packageCode?: string;
    /**
     *
     * @type {BlockId}
     * @memberof BlockInventoryItemSourceType
     */
    blockId?: BlockId;
}
/**
 * Check if a given object implements the BlockInventoryItemSourceType interface.
 */
export declare function instanceOfBlockInventoryItemSourceType(value: object): boolean;
export declare function BlockInventoryItemSourceTypeFromJSON(json: any): BlockInventoryItemSourceType;
export declare function BlockInventoryItemSourceTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockInventoryItemSourceType;
export declare function BlockInventoryItemSourceTypeToJSON(value?: BlockInventoryItemSourceType | null): any;
