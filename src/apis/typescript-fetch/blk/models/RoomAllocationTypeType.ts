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
import type { AllocationType } from './AllocationType';
import {
    AllocationTypeFromJSON,
    AllocationTypeFromJSONTyped,
    AllocationTypeToJSON,
} from './AllocationType';
import type { RoomAllocationInfoType } from './RoomAllocationInfoType';
import {
    RoomAllocationInfoTypeFromJSON,
    RoomAllocationInfoTypeFromJSONTyped,
    RoomAllocationInfoTypeToJSON,
} from './RoomAllocationInfoType';

/**
 * A collection of Allocation objects for a block, such as Current Rooms, Original Rooms, Rate Amounts, etc.
 * @export
 * @interface RoomAllocationTypeType
 */
export interface RoomAllocationTypeType {
    /**
     * Allocation objects for a block.
     * @type {Array<RoomAllocationInfoType>}
     * @memberof RoomAllocationTypeType
     */
    roomAllocationInfo?: Array<RoomAllocationInfoType>;
    /**
     * 
     * @type {AllocationType}
     * @memberof RoomAllocationTypeType
     */
    allocation?: AllocationType;
}

/**
 * Check if a given object implements the RoomAllocationTypeType interface.
 */
export function instanceOfRoomAllocationTypeType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function RoomAllocationTypeTypeFromJSON(json: any): RoomAllocationTypeType {
    return RoomAllocationTypeTypeFromJSONTyped(json, false);
}

export function RoomAllocationTypeTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): RoomAllocationTypeType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'roomAllocationInfo': !exists(json, 'roomAllocationInfo') ? undefined : ((json['roomAllocationInfo'] as Array<any>).map(RoomAllocationInfoTypeFromJSON)),
        'allocation': !exists(json, 'allocation') ? undefined : AllocationTypeFromJSON(json['allocation']),
    };
}

export function RoomAllocationTypeTypeToJSON(value?: RoomAllocationTypeType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'roomAllocationInfo': value.roomAllocationInfo === undefined ? undefined : ((value.roomAllocationInfo as Array<any>).map(RoomAllocationInfoTypeToJSON)),
        'allocation': AllocationTypeToJSON(value.allocation),
    };
}

