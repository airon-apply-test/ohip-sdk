/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Sales Event Management API
 * APIs to cater for Event Management functionality in OPERA Cloud. <br /><br />The Events feature in OPERA Cloud is designed to manage any kind of catering activity. Events can be as simple as a one-hour reception or more complex, such as a three-day business meeting with meals, breaks, and specific meeting functionSpaceDetails with setupCode and resource requirements. Any group function can be an Event.<br /><br /> Compatible with OPERA Cloud release 21.5.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 21.5.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { BlockGridInvType } from './BlockGridInvType';
import {
    BlockGridInvTypeFromJSON,
    BlockGridInvTypeFromJSONTyped,
    BlockGridInvTypeToJSON,
} from './BlockGridInvType';

/**
 * Allows to select type of wash schedule: by number of rooms or by percentage.
 * @export
 * @interface BlockWashScheduleType
 */
export interface BlockWashScheduleType {
    /**
     * Date on which the block wash operation will be performed.
     * @type {Date}
     * @memberof BlockWashScheduleType
     */
    washDate?: Date;
    /**
     * A single functionSpaceDetails type on which wash operation should be performed.
     * @type {Array<string>}
     * @memberof BlockWashScheduleType
     */
    roomTypes?: Array<string>;
    /**
     * 
     * @type {BlockGridInvType}
     * @memberof BlockWashScheduleType
     */
    washByRooms?: BlockGridInvType;
    /**
     * When using the Wash by % option, indicate a wash percentage that will be applied it to the whole block.
     * @type {number}
     * @memberof BlockWashScheduleType
     */
    washByPercent?: number;
}

/**
 * Check if a given object implements the BlockWashScheduleType interface.
 */
export function instanceOfBlockWashScheduleType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function BlockWashScheduleTypeFromJSON(json: any): BlockWashScheduleType {
    return BlockWashScheduleTypeFromJSONTyped(json, false);
}

export function BlockWashScheduleTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockWashScheduleType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'washDate': !exists(json, 'washDate') ? undefined : (new Date(json['washDate'])),
        'roomTypes': !exists(json, 'roomTypes') ? undefined : json['roomTypes'],
        'washByRooms': !exists(json, 'washByRooms') ? undefined : BlockGridInvTypeFromJSON(json['washByRooms']),
        'washByPercent': !exists(json, 'washByPercent') ? undefined : json['washByPercent'],
    };
}

export function BlockWashScheduleTypeToJSON(value?: BlockWashScheduleType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'washDate': value.washDate === undefined ? undefined : (value.washDate.toISOString().substr(0,10)),
        'roomTypes': value.roomTypes,
        'washByRooms': BlockGridInvTypeToJSON(value.washByRooms),
        'washByPercent': value.washByPercent,
    };
}

