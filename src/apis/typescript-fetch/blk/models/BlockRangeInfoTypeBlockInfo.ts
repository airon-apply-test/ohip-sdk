/* tslint:disable */
/* eslint-disable */
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

import { exists, mapValues } from '../runtime';
import type { BlockId } from './BlockId';
import {
    BlockIdFromJSON,
    BlockIdFromJSONTyped,
    BlockIdToJSON,
} from './BlockId';
import type { BlockInventoryControlType } from './BlockInventoryControlType';
import {
    BlockInventoryControlTypeFromJSON,
    BlockInventoryControlTypeFromJSONTyped,
    BlockInventoryControlTypeToJSON,
} from './BlockInventoryControlType';

/**
 * Block information needed to perform range operation on the block.
 * @export
 * @interface BlockRangeInfoTypeBlockInfo
 */
export interface BlockRangeInfoTypeBlockInfo {
    /**
     * 
     * @type {string}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    hotelId?: string;
    /**
     * 
     * @type {BlockId}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    blockId?: BlockId;
    /**
     * Block Code
     * @type {string}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    blockCode?: string;
    /**
     * Block Start Date
     * @type {Date}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    startDate?: Date;
    /**
     * Block End Date
     * @type {Date}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    endDate?: Date;
    /**
     * Block Shoulder Start Date.
     * @type {Date}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    shoulderStartDate?: Date;
    /**
     * Block Shoulder End Date.
     * @type {Date}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    shoulderEndDate?: Date;
    /**
     * Block Rate Code.
     * @type {string}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    rateCode?: string;
    /**
     * Block Status.
     * @type {string}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    blockStatus?: string;
    /**
     * 
     * @type {BlockInventoryControlType}
     * @memberof BlockRangeInfoTypeBlockInfo
     */
    inventoryControl?: BlockInventoryControlType;
}

/**
 * Check if a given object implements the BlockRangeInfoTypeBlockInfo interface.
 */
export function instanceOfBlockRangeInfoTypeBlockInfo(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockRangeInfoTypeBlockInfoFromJSON(json: any): BlockRangeInfoTypeBlockInfo {
    return BlockRangeInfoTypeBlockInfoFromJSONTyped(json, false);
}

export function BlockRangeInfoTypeBlockInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockRangeInfoTypeBlockInfo {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'blockCode': !exists(json, 'blockCode') ? undefined : json['blockCode'],
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'endDate': !exists(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'shoulderStartDate': !exists(json, 'shoulderStartDate') ? undefined : (new Date(json['shoulderStartDate'])),
        'shoulderEndDate': !exists(json, 'shoulderEndDate') ? undefined : (new Date(json['shoulderEndDate'])),
        'rateCode': !exists(json, 'rateCode') ? undefined : json['rateCode'],
        'blockStatus': !exists(json, 'blockStatus') ? undefined : json['blockStatus'],
        'inventoryControl': !exists(json, 'inventoryControl') ? undefined : BlockInventoryControlTypeFromJSON(json['inventoryControl']),
    };
}

export function BlockRangeInfoTypeBlockInfoToJSON(value?: BlockRangeInfoTypeBlockInfo | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'blockId': BlockIdToJSON(value.blockId),
        'blockCode': value.blockCode,
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0,10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0,10)),
        'shoulderStartDate': value.shoulderStartDate === undefined ? undefined : (value.shoulderStartDate.toISOString().substr(0,10)),
        'shoulderEndDate': value.shoulderEndDate === undefined ? undefined : (value.shoulderEndDate.toISOString().substr(0,10)),
        'rateCode': value.rateCode,
        'blockStatus': value.blockStatus,
        'inventoryControl': BlockInventoryControlTypeToJSON(value.inventoryControl),
    };
}

