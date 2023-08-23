/* tslint:disable */
/* eslint-disable */
/**
 * OPERA Cloud Activity API
 * APIs to cater for Sales Activity functionality in OPERA Cloud. <br /><br /> Activities provide you with an account management tool for managing daily tasks such as appointments, sales calls, contact follow-up, and so on.<br /><br /> Compatible with OPERA Cloud release 22.3.0.0.<br /><br /><p> This document and all content within is available under the Universal Permissive License v 1.0 (https://oss.oracle.com/licenses/upl). Copyright (c) 2020, 2022 Oracle and/or its affiliates.</p>
 *
 * The version of the OpenAPI document: 22.3.0.0
 * Contact: hospitality_apis_ww_grp@oracle.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { TimeSpanType } from './TimeSpanType';
import {
    TimeSpanTypeFromJSON,
    TimeSpanTypeFromJSONTyped,
    TimeSpanTypeToJSON,
} from './TimeSpanType';
import type { UniqueNameIDType } from './UniqueNameIDType';
import {
    UniqueNameIDTypeFromJSON,
    UniqueNameIDTypeFromJSONTyped,
    UniqueNameIDTypeToJSON,
} from './UniqueNameIDType';

/**
 * 
 * @export
 * @interface ActivityBlockInfoType
 */
export interface ActivityBlockInfoType {
    /**
     * Defines descriptive name and unique identification combination.
     * @type {Array<UniqueNameIDType>}
     * @memberof ActivityBlockInfoType
     */
    blockIdList?: Array<UniqueNameIDType>;
    /**
     * Block code for the block.
     * @type {string}
     * @memberof ActivityBlockInfoType
     */
    blockCode?: string;
    /**
     * Name of the block.
     * @type {string}
     * @memberof ActivityBlockInfoType
     */
    blockName?: string;
    /**
     * 
     * @type {TimeSpanType}
     * @memberof ActivityBlockInfoType
     */
    timeSpan?: TimeSpanType;
    /**
     * When true, indicates a primary information.
     * @type {boolean}
     * @memberof ActivityBlockInfoType
     */
    primary?: boolean;
}

/**
 * Check if a given object implements the ActivityBlockInfoType interface.
 */
export function instanceOfActivityBlockInfoType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ActivityBlockInfoTypeFromJSON(json: any): ActivityBlockInfoType {
    return ActivityBlockInfoTypeFromJSONTyped(json, false);
}

export function ActivityBlockInfoTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ActivityBlockInfoType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blockIdList': !exists(json, 'blockIdList') ? undefined : ((json['blockIdList'] as Array<any>).map(UniqueNameIDTypeFromJSON)),
        'blockCode': !exists(json, 'blockCode') ? undefined : json['blockCode'],
        'blockName': !exists(json, 'blockName') ? undefined : json['blockName'],
        'timeSpan': !exists(json, 'timeSpan') ? undefined : TimeSpanTypeFromJSON(json['timeSpan']),
        'primary': !exists(json, 'primary') ? undefined : json['primary'],
    };
}

export function ActivityBlockInfoTypeToJSON(value?: ActivityBlockInfoType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blockIdList': value.blockIdList === undefined ? undefined : ((value.blockIdList as Array<any>).map(UniqueNameIDTypeToJSON)),
        'blockCode': value.blockCode,
        'blockName': value.blockName,
        'timeSpan': TimeSpanTypeToJSON(value.timeSpan),
        'primary': value.primary,
    };
}

