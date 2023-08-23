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
import type { ActivityId } from './ActivityId';
import {
    ActivityIdFromJSON,
    ActivityIdFromJSONTyped,
    ActivityIdToJSON,
} from './ActivityId';
import type { ActivityInfoType } from './ActivityInfoType';
import {
    ActivityInfoTypeFromJSON,
    ActivityInfoTypeFromJSONTyped,
    ActivityInfoTypeToJSON,
} from './ActivityInfoType';

/**
 * Linked Activity Related Information.
 * @export
 * @interface LinkedActivityDetailsType
 */
export interface LinkedActivityDetailsType {
    /**
     * 
     * @type {ActivityId}
     * @memberof LinkedActivityDetailsType
     */
    activityId?: ActivityId;
    /**
     * 
     * @type {ActivityInfoType}
     * @memberof LinkedActivityDetailsType
     */
    linkedActivityDetail?: ActivityInfoType;
}

/**
 * Check if a given object implements the LinkedActivityDetailsType interface.
 */
export function instanceOfLinkedActivityDetailsType(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function LinkedActivityDetailsTypeFromJSON(json: any): LinkedActivityDetailsType {
    return LinkedActivityDetailsTypeFromJSONTyped(json, false);
}

export function LinkedActivityDetailsTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): LinkedActivityDetailsType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'activityId': !exists(json, 'activityId') ? undefined : ActivityIdFromJSON(json['activityId']),
        'linkedActivityDetail': !exists(json, 'linkedActivityDetail') ? undefined : ActivityInfoTypeFromJSON(json['linkedActivityDetail']),
    };
}

export function LinkedActivityDetailsTypeToJSON(value?: LinkedActivityDetailsType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'activityId': ActivityIdToJSON(value.activityId),
        'linkedActivityDetail': ActivityInfoTypeToJSON(value.linkedActivityDetail),
    };
}

