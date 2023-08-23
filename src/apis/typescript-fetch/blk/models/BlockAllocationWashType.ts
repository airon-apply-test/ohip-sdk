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
import type { BlockAllocationWashTypeOccPercentByDay } from './BlockAllocationWashTypeOccPercentByDay';
import {
    BlockAllocationWashTypeOccPercentByDayFromJSON,
    BlockAllocationWashTypeOccPercentByDayFromJSONTyped,
    BlockAllocationWashTypeOccPercentByDayToJSON,
} from './BlockAllocationWashTypeOccPercentByDay';
import type { BlockGridInvType } from './BlockGridInvType';
import {
    BlockGridInvTypeFromJSON,
    BlockGridInvTypeFromJSONTyped,
    BlockGridInvTypeToJSON,
} from './BlockGridInvType';
import type { BlockId } from './BlockId';
import {
    BlockIdFromJSON,
    BlockIdFromJSONTyped,
    BlockIdToJSON,
} from './BlockId';

/**
 * Holds criteria for performing a wash operation on a block.
 * @export
 * @interface BlockAllocationWashType
 */
export interface BlockAllocationWashType {
    /**
     * 
     * @type {BlockId}
     * @memberof BlockAllocationWashType
     */
    blockId?: BlockId;
    /**
     * Hotel Code to which the block belongs to.
     * @type {string}
     * @memberof BlockAllocationWashType
     */
    hotelId?: string;
    /**
     * Start date for the wash operation.
     * @type {Date}
     * @memberof BlockAllocationWashType
     */
    startDate?: Date;
    /**
     * End date for the wash operation.
     * @type {Date}
     * @memberof BlockAllocationWashType
     */
    endDate?: Date;
    /**
     * List of room types on which wash operation should be performed.
     * @type {Array<string>}
     * @memberof BlockAllocationWashType
     */
    roomTypes?: Array<string>;
    /**
     * String of length seven, each position representing various days of week from Sunday to Saturday, each character is either 0 or 1 indicating whether to apply wash for that day.
     * @type {string}
     * @memberof BlockAllocationWashType
     */
    includedDays?: string;
    /**
     * A true value indicates that absolute value is being passed for occupancy, else it indicates that a percentage value is passed.
     * @type {boolean}
     * @memberof BlockAllocationWashType
     */
    byValue?: boolean;
    /**
     * 
     * @type {BlockGridInvType}
     * @memberof BlockAllocationWashType
     */
    blockInventory?: BlockGridInvType;
    /**
     * A true value indicates that percentage values are being passed day-wise.
     * @type {boolean}
     * @memberof BlockAllocationWashType
     */
    percentByDay?: boolean;
    /**
     * 
     * @type {BlockAllocationWashTypeOccPercentByDay}
     * @memberof BlockAllocationWashType
     */
    occPercentByDay?: BlockAllocationWashTypeOccPercentByDay;
    /**
     * Indicates if the Allocation objects refer to Generic Room Types (Room Pools).
     * @type {boolean}
     * @memberof BlockAllocationWashType
     */
    genericRoomType?: boolean;
}

/**
 * Check if a given object implements the BlockAllocationWashType interface.
 */
export function instanceOfBlockAllocationWashType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockAllocationWashTypeFromJSON(json: any): BlockAllocationWashType {
    return BlockAllocationWashTypeFromJSONTyped(json, false);
}

export function BlockAllocationWashTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockAllocationWashType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'startDate': !exists(json, 'startDate') ? undefined : (new Date(json['startDate'])),
        'endDate': !exists(json, 'endDate') ? undefined : (new Date(json['endDate'])),
        'roomTypes': !exists(json, 'roomTypes') ? undefined : json['roomTypes'],
        'includedDays': !exists(json, 'includedDays') ? undefined : json['includedDays'],
        'byValue': !exists(json, 'byValue') ? undefined : json['byValue'],
        'blockInventory': !exists(json, 'blockInventory') ? undefined : BlockGridInvTypeFromJSON(json['blockInventory']),
        'percentByDay': !exists(json, 'percentByDay') ? undefined : json['percentByDay'],
        'occPercentByDay': !exists(json, 'occPercentByDay') ? undefined : BlockAllocationWashTypeOccPercentByDayFromJSON(json['occPercentByDay']),
        'genericRoomType': !exists(json, 'genericRoomType') ? undefined : json['genericRoomType'],
    };
}

export function BlockAllocationWashTypeToJSON(value?: BlockAllocationWashType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blockId': BlockIdToJSON(value.blockId),
        'hotelId': value.hotelId,
        'startDate': value.startDate === undefined ? undefined : (value.startDate.toISOString().substr(0,10)),
        'endDate': value.endDate === undefined ? undefined : (value.endDate.toISOString().substr(0,10)),
        'roomTypes': value.roomTypes,
        'includedDays': value.includedDays,
        'byValue': value.byValue,
        'blockInventory': BlockGridInvTypeToJSON(value.blockInventory),
        'percentByDay': value.percentByDay,
        'occPercentByDay': BlockAllocationWashTypeOccPercentByDayToJSON(value.occPercentByDay),
        'genericRoomType': value.genericRoomType,
    };
}

