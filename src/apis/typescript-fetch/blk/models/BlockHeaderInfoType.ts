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
import type { BlockSourceOfSaleType } from './BlockSourceOfSaleType';
import {
    BlockSourceOfSaleTypeFromJSON,
    BlockSourceOfSaleTypeFromJSONTyped,
    BlockSourceOfSaleTypeToJSON,
} from './BlockSourceOfSaleType';
import type { CodeDescriptionType } from './CodeDescriptionType';
import {
    CodeDescriptionTypeFromJSON,
    CodeDescriptionTypeFromJSONTyped,
    CodeDescriptionTypeToJSON,
} from './CodeDescriptionType';

/**
 * Holds block header details.
 * @export
 * @interface BlockHeaderInfoType
 */
export interface BlockHeaderInfoType {
    /**
     * Hotel code to which the block belongs to.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    hotelId?: string;
    /**
     * Name of the hotel.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    hotelName?: string;
    /**
     * Block description.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    blockName?: string;
    /**
     * Unique code for the block.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    blockCode?: string;
    /**
     * 
     * @type {BlockId}
     * @memberof BlockHeaderInfoType
     */
    blockId?: BlockId;
    /**
     * Default currency code of the block.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    currencyCode?: string;
    /**
     * Primary rate code of the block.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    rateCode?: string;
    /**
     * Shoulder start date of the block, applicable if the block is elastic.
     * @type {Date}
     * @memberof BlockHeaderInfoType
     */
    shoulderStartDate?: Date;
    /**
     * Shoulder end date of the block, applicable if the block is elastic.
     * @type {Date}
     * @memberof BlockHeaderInfoType
     */
    shoulderEndDate?: Date;
    /**
     * Start date of the block.
     * @type {Date}
     * @memberof BlockHeaderInfoType
     */
    startDate?: Date;
    /**
     * End date of the block.
     * @type {Date}
     * @memberof BlockHeaderInfoType
     */
    endDate?: Date;
    /**
     * 
     * @type {BlockInventoryControlType}
     * @memberof BlockHeaderInfoType
     */
    inventoryControl?: BlockInventoryControlType;
    /**
     * Market code for the block.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    marketCode?: string;
    /**
     * 
     * @type {BlockSourceOfSaleType}
     * @memberof BlockHeaderInfoType
     */
    sourceOfSale?: BlockSourceOfSaleType;
    /**
     * Guarantee Code for the block.
     * @type {string}
     * @memberof BlockHeaderInfoType
     */
    guaranteeCode?: string;
    /**
     * 
     * @type {CodeDescriptionType}
     * @memberof BlockHeaderInfoType
     */
    blockStatus?: CodeDescriptionType;
    /**
     * Indicates whether rates are suppressed.
     * @type {boolean}
     * @memberof BlockHeaderInfoType
     */
    suppressRate?: boolean;
}

/**
 * Check if a given object implements the BlockHeaderInfoType interface.
 */
export function instanceOfBlockHeaderInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockHeaderInfoTypeFromJSON(json: any): BlockHeaderInfoType {
    return BlockHeaderInfoTypeFromJSONTyped(json, false);
}

export function BlockHeaderInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockHeaderInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'hotelName': !exists(json, 'hotelName') ? undefined : json['hotelName'],
        'blockName': !exists(json, 'blockName') ? undefined : json['blockName'],
        'blockCode': !exists(json, 'blockCode') ? undefined : json['blockCode'],
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'currencyCode': !exists(json, 'currencyCode') ? undefined : json['currencyCode'],
        'rateCode': !exists(json, 'rateCode') ? undefined : json['rateCode'],
        'shoulderStartDate': !exists(json, 'shoulderStartDate') ? undefined : (new Date(json['shoulderStartDate'])),
        'shoulderEndDate': !exists(json, 'shoulderEndDate') ? undefined : (new Date(json['shoulderEndDate'])),
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'endDate': !exists(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'inventoryControl': !exists(json, 'inventoryControl') ? undefined : BlockInventoryControlTypeFromJSON(json['inventoryControl']),
        'marketCode': !exists(json, 'marketCode') ? undefined : json['marketCode'],
        'sourceOfSale': !exists(json, 'sourceOfSale') ? undefined : BlockSourceOfSaleTypeFromJSON(json['sourceOfSale']),
        'guaranteeCode': !exists(json, 'guaranteeCode') ? undefined : json['guaranteeCode'],
        'blockStatus': !exists(json, 'blockStatus') ? undefined : CodeDescriptionTypeFromJSON(json['blockStatus']),
        'suppressRate': !exists(json, 'suppressRate') ? undefined : json['suppressRate'],
    };
}

export function BlockHeaderInfoTypeToJSON(value?: BlockHeaderInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'hotelName': value.hotelName,
        'blockName': value.blockName,
        'blockCode': value.blockCode,
        'blockId': BlockIdToJSON(value.blockId),
        'currencyCode': value.currencyCode,
        'rateCode': value.rateCode,
        'shoulderStartDate': value.shoulderStartDate === undefined ? undefined : (value.shoulderStartDate.toISOString().substr(0,10)),
        'shoulderEndDate': value.shoulderEndDate === undefined ? undefined : (value.shoulderEndDate.toISOString().substr(0,10)),
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0,10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0,10)),
        'inventoryControl': BlockInventoryControlTypeToJSON(value.inventoryControl),
        'marketCode': value.marketCode,
        'sourceOfSale': BlockSourceOfSaleTypeToJSON(value.sourceOfSale),
        'guaranteeCode': value.guaranteeCode,
        'blockStatus': CodeDescriptionTypeToJSON(value.blockStatus),
        'suppressRate': value.suppressRate,
    };
}

