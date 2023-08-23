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
import type { AllocationGridByRoomTypeType } from './AllocationGridByRoomTypeType';
import {
    AllocationGridByRoomTypeTypeFromJSON,
    AllocationGridByRoomTypeTypeFromJSONTyped,
    AllocationGridByRoomTypeTypeToJSON,
} from './AllocationGridByRoomTypeType';
import type { BlockId } from './BlockId';
import {
    BlockIdFromJSON,
    BlockIdFromJSONTyped,
    BlockIdToJSON,
} from './BlockId';

/**
 * A collection of Allocation objects for a block defined by room type.
 * @export
 * @interface AllocationGridByRoomTypesType
 */
export interface AllocationGridByRoomTypesType {
    /**
     * Opera Hotel code for the Set Allocation Grid request.
     * @type {string}
     * @memberof AllocationGridByRoomTypesType
     */
    hotelId?: string;
    /**
     * 
     * @type {BlockId}
     * @memberof AllocationGridByRoomTypesType
     */
    blockId?: BlockId;
    /**
     * Allocation objects of a block grouped by room types.
     * @type {Array<AllocationGridByRoomTypeType>}
     * @memberof AllocationGridByRoomTypesType
     */
    allocationRoomTypes?: Array<AllocationGridByRoomTypeType>;
    /**
     * Indicates if the Allocation objects refer to Generic Room Types (Room Pools).
     * @type {boolean}
     * @memberof AllocationGridByRoomTypesType
     */
    genericRoomType?: boolean;
}

/**
 * Check if a given object implements the AllocationGridByRoomTypesType interface.
 */
export function instanceOfAllocationGridByRoomTypesType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AllocationGridByRoomTypesTypeFromJSON(json: any): AllocationGridByRoomTypesType {
    return AllocationGridByRoomTypesTypeFromJSONTyped(json, false);
}

export function AllocationGridByRoomTypesTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): AllocationGridByRoomTypesType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'hotelId': !exists(json, 'hotelId') ? undefined : json['hotelId'],
        'blockId': !exists(json, 'blockId') ? undefined : BlockIdFromJSON(json['blockId']),
        'allocationRoomTypes': !exists(json, 'allocationRoomTypes') ? undefined : ((json['allocationRoomTypes'] as Array<any>).map(AllocationGridByRoomTypeTypeFromJSON)),
        'genericRoomType': !exists(json, 'genericRoomType') ? undefined : json['genericRoomType'],
    };
}

export function AllocationGridByRoomTypesTypeToJSON(value?: AllocationGridByRoomTypesType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'hotelId': value.hotelId,
        'blockId': BlockIdToJSON(value.blockId),
        'allocationRoomTypes': value.allocationRoomTypes === undefined ? undefined : ((value.allocationRoomTypes as Array<any>).map(AllocationGridByRoomTypeTypeToJSON)),
        'genericRoomType': value.genericRoomType,
    };
}

