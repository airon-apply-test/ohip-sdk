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
import type { BorrowInventoryType } from './BorrowInventoryType';
import {
    BorrowInventoryTypeFromJSON,
    BorrowInventoryTypeFromJSONTyped,
    BorrowInventoryTypeToJSON,
} from './BorrowInventoryType';
import type { InstanceLink } from './InstanceLink';
import {
    InstanceLinkFromJSON,
    InstanceLinkFromJSONTyped,
    InstanceLinkToJSON,
} from './InstanceLink';
import type { UniqueIDType } from './UniqueIDType';
import {
    UniqueIDTypeFromJSON,
    UniqueIDTypeFromJSONTyped,
    UniqueIDTypeToJSON,
} from './UniqueIDType';
import type { WarningType } from './WarningType';
import {
    WarningTypeFromJSON,
    WarningTypeFromJSONTyped,
    WarningTypeToJSON,
} from './WarningType';

/**
 * The standard optional Opera Context element is also included.
 * @export
 * @interface InventoryToBorrow
 */
export interface InventoryToBorrow {
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof InventoryToBorrow
     */
    hotelId?: string;
    /**
     * 
     * @type {BlockId}
     * @memberof InventoryToBorrow
     */
    blockId?: BlockId;
    /**
     * 
     * @type {UniqueIDType}
     * @memberof InventoryToBorrow
     */
    existingReservationId?: UniqueIDType;
    /**
     * Used for codes in the OPERA Code tables. Possible values of this pattern are 1, 101, 101.EQP, or 101.EQP.X.
     * @type {string}
     * @memberof InventoryToBorrow
     */
    roomType?: string;
    /**
     * The number of adults that are expected to stay in a room.
     * @type {number}
     * @memberof InventoryToBorrow
     */
    adultCount?: number;
    /**
     * Indicates whether to overbook the Sales Allowance in case there are no rooms left at the Generic Sales Allowance level.
     * @type {boolean}
     * @memberof InventoryToBorrow
     */
    overbookSalesAllowance?: boolean;
    /**
     * The date and number of rooms to borrow from either room type or House.
     * @type {Array<BorrowInventoryType>}
     * @memberof InventoryToBorrow
     */
    borrowInventoryList?: Array<BorrowInventoryType>;
    /**
     * 
     * @type {Array<InstanceLink>}
     * @memberof InventoryToBorrow
     */
    links?: Array<InstanceLink>;
    /**
     * Used in conjunction with the Success element to define a business error.
     * @type {Array<WarningType>}
     * @memberof InventoryToBorrow
     */
    warnings?: Array<WarningType>;
}

/**
 * Check if a given object implements the InventoryToBorrow interface.
 */
export function instanceOfInventoryToBorrow(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function InventoryToBorrowFromJSON(json: any): InventoryToBorrow {
    return InventoryToBorrowFromJSONTyped(json, false);
}

export function InventoryToBorrowFromJSONTyped(json: any, ignoreDiscriminator: boolean): InventoryToBorrow {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'existingReservationId': !exists(json, 'existingReservationId') ? undefined : UniqueIDTypeFromJSON(json['existingReservationId']),
        'roomType': !exists(json, 'roomType') ? undefined : json['roomType'],
        'adultCount': !exists(json, 'adultCount') ? undefined : json['adultCount'],
        'overbookSalesAllowance': !exists(json, 'overbookSalesAllowance') ? undefined : json['overbookSalesAllowance'],
        'borrowInventoryList': !exists(json, 'borrowInventoryList') ? undefined : ((json['borrowInventoryList'] as Array<any>).map(BorrowInventoryTypeFromJSON)),
        'links': !exists(json, 'links') ? undefined : ((json['links'] as Array<any>).map(InstanceLinkFromJSON)),
        'warnings': !exists(json, 'warnings') ? undefined : ((json['warnings'] as Array<any>).map(WarningTypeFromJSON)),
    };
}

export function InventoryToBorrowToJSON(value?: InventoryToBorrow | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'blockId': BlockIdToJSON(value.blockId),
        'existingReservationId': UniqueIDTypeToJSON(value.existingReservationId),
        'roomType': value.roomType,
        'adultCount': value.adultCount,
        'overbookSalesAllowance': value.overbookSalesAllowance,
        'borrowInventoryList': value.borrowInventoryList === undefined ? undefined : ((value.borrowInventoryList as Array<any>).map(BorrowInventoryTypeToJSON)),
        'links': value.links === undefined ? undefined : ((value.links as Array<any>).map(InstanceLinkToJSON)),
        'warnings': value.warnings === undefined ? undefined : ((value.warnings as Array<any>).map(WarningTypeToJSON)),
    };
}

