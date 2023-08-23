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

/**
 * 
 * @export
 * @interface ShiftBlockCriteriaType
 */
export interface ShiftBlockCriteriaType {
    /**
     * Hotel code of the business block.
     * @type {string}
     * @memberof ShiftBlockCriteriaType
     */
    hotelId?: string;
    /**
     * 
     * @type {BlockId}
     * @memberof ShiftBlockCriteriaType
     */
    blockId?: BlockId;
    /**
     * New start date of the business block.
     * @type {Date}
     * @memberof ShiftBlockCriteriaType
     */
    newStartDate?: Date;
    /**
     * When true, this will remove alternate dates for this block before shifting dates.
     * @type {boolean}
     * @memberof ShiftBlockCriteriaType
     */
    removeAlternateDates?: boolean;
    /**
     * When true, this will validate alternate dates for this block before shifting dates.
     * @type {boolean}
     * @memberof ShiftBlockCriteriaType
     */
    validateAlternateDates?: boolean;
    /**
     * When true, this will overbook allocated rooms if needed.
     * @type {boolean}
     * @memberof ShiftBlockCriteriaType
     */
    overbookAll?: boolean;
    /**
     * When true, this will shift the block even if there are rooms pre-allocated.
     * @type {boolean}
     * @memberof ShiftBlockCriteriaType
     */
    ignorePreAllocatedRooms?: boolean;
    /**
     * When true, this will shift the block even if there are active traces.
     * @type {boolean}
     * @memberof ShiftBlockCriteriaType
     */
    ignoreTraces?: boolean;
    /**
     * Indicates whether to ignore any warning during applying the changes to the events associated with the current block.
     * @type {boolean}
     * @memberof ShiftBlockCriteriaType
     */
    overrideEventsProcessingWarnings?: boolean;
}

/**
 * Check if a given object implements the ShiftBlockCriteriaType interface.
 */
export function instanceOfShiftBlockCriteriaType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ShiftBlockCriteriaTypeFromJSON(json: any): ShiftBlockCriteriaType {
    return ShiftBlockCriteriaTypeFromJSONTyped(json, false);
}

export function ShiftBlockCriteriaTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ShiftBlockCriteriaType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'newStartDate': !exists(json, 'newStartDate') ? undefined : (new Date(json['newStartDate'])),
        'removeAlternateDates': !exists(json, 'removeAlternateDates') ? undefined : json['removeAlternateDates'],
        'validateAlternateDates': !exists(json, 'validateAlternateDates') ? undefined : json['validateAlternateDates'],
        'overbookAll': !exists(json, 'overbookAll') ? undefined : json['overbookAll'],
        'ignorePreAllocatedRooms': !exists(json, 'ignorePreAllocatedRooms') ? undefined : json['ignorePreAllocatedRooms'],
        'ignoreTraces': !exists(json, 'ignoreTraces') ? undefined : json['ignoreTraces'],
        'overrideEventsProcessingWarnings': !exists(json, 'overrideEventsProcessingWarnings') ? undefined : json['overrideEventsProcessingWarnings'],
    };
}

export function ShiftBlockCriteriaTypeToJSON(value?: ShiftBlockCriteriaType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'blockId': BlockIdToJSON(value.blockId),
        'newStartDate': value.newStartDate === undefined ? undefined : (value.newStartDate.toISOString().substr(0,10)),
        'removeAlternateDates': value.removeAlternateDates,
        'validateAlternateDates': value.validateAlternateDates,
        'overbookAll': value.overbookAll,
        'ignorePreAllocatedRooms': value.ignorePreAllocatedRooms,
        'ignoreTraces': value.ignoreTraces,
        'overrideEventsProcessingWarnings': value.overrideEventsProcessingWarnings,
    };
}

