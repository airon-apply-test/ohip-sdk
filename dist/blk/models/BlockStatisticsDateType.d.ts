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
import type { BlockRoomStatisticsType } from './BlockRoomStatisticsType';
/**
 * Holds the room type level statistics and total statistics for an allotment date.
 * @export
 * @interface BlockStatisticsDateType
 */
export interface BlockStatisticsDateType {
    /**
     *
     * @type {number}
     * @memberof BlockStatisticsDateType
     */
    total?: number;
    /**
     * Collection of room type level statistics.
     * @type {Array<BlockRoomStatisticsType>}
     * @memberof BlockStatisticsDateType
     */
    roomStatisticsList?: Array<BlockRoomStatisticsType>;
    /**
     *
     * @type {Date}
     * @memberof BlockStatisticsDateType
     */
    date?: Date;
}
/**
 * Check if a given object implements the BlockStatisticsDateType interface.
 */
export declare function instanceOfBlockStatisticsDateType(value: object): boolean;
export declare function BlockStatisticsDateTypeFromJSON(json: any): BlockStatisticsDateType;
export declare function BlockStatisticsDateTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockStatisticsDateType;
export declare function BlockStatisticsDateTypeToJSON(value?: BlockStatisticsDateType | null): any;
