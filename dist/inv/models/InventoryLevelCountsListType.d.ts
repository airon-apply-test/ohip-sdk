/**
 * OPERA Cloud Inventory API
 * APIs to cater for Inventory functionality in OPERA Cloud. This includes sell limits for date ranges, viewing and updating the property&apos;s inventory, as well as item inventory (such as rollaways, microwaves etc.).<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import type { InventoryCountsType } from './InventoryCountsType';
/**
 * Collection of Inventory counts for the date ranges.
 * @export
 * @interface InventoryLevelCountsListType
 */
export interface InventoryLevelCountsListType {
    /**
     * Collection of Inventory counts for the date ranges.
     * @type {Array<InventoryCountsType>}
     * @memberof InventoryLevelCountsListType
     */
    inventoryCounts?: Array<InventoryCountsType>;
    /**
     * Inventory Group/Level code.
     * @type {string}
     * @memberof InventoryLevelCountsListType
     */
    code?: string;
    /**
     * Integer Group/Level order sequence number.
     * @type {number}
     * @memberof InventoryLevelCountsListType
     */
    sequence?: number;
}
/**
 * Check if a given object implements the InventoryLevelCountsListType interface.
 */
export declare function instanceOfInventoryLevelCountsListType(value: object): boolean;
export declare function InventoryLevelCountsListTypeFromJSON(json: any): InventoryLevelCountsListType;
export declare function InventoryLevelCountsListTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): InventoryLevelCountsListType;
export declare function InventoryLevelCountsListTypeToJSON(value?: InventoryLevelCountsListType | null): any;
